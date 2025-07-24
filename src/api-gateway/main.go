package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

func main() {
	err := godotenv.Load("/home/vagrant/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/movies", proxyInventory).Methods("GET", "POST", "DELETE")
	r.HandleFunc("/api/movies/{id}", proxyInventory).Methods("GET", "PUT", "DELETE")
	r.HandleFunc("/api/billing", sendToBillingQueue).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+os.Getenv("GATEWAY_PORT"), r))
}

func proxyInventory(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:" + os.Getenv("INVENTORY_PORT") + r.URL.Path
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	for k, v := range resp.Header {
		for _, vv := range v {
			w.Header().Add(k, vv)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func sendToBillingQueue(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	conn, err := amqp.Dial("amqp://" + os.Getenv("RABBITMQ_USER") + ":" + os.Getenv("RABBITMQ_PASSWORD") + "@" + os.Getenv("RABBITMQ_HOST") + ":" + os.Getenv("RABBITMQ_PORT") + "/")
	if err != nil {
		http.Error(w, "Failed to connect to RabbitMQ", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		http.Error(w, "Failed to open channel", http.StatusInternalServerError)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		os.Getenv("RABBITMQ_QUEUE"),
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		http.Error(w, "Failed to declare queue", http.StatusInternalServerError)
		return
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Message posted"))
}
