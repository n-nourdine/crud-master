package main

import (
	"encoding/json"
	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"log"
)

type BillingHandler struct {
	db   *sqlx.DB
	conn *amqp.Connection
}

func (h *BillingHandler) consumeMessages() {
	ch, err := h.conn.Channel()
	if err != nil {
		log.Fatal("Failed to open channel: ", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"billing_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to declare queue: ", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to register consumer: ", err)
	}

	for msg := range msgs {
		var order Order
		if err := json.Unmarshal(msg.Body, &order); err != nil {
			log.Println("Failed to parse message: ", err)
			continue
		}

		_, err = h.db.Exec("INSERT INTO orders (user_id, number_of_items, total_amount) VALUES ($1, $2, $3)",
			order.UserID, order.NumberOfItems, order.TotalAmount)
		if err != nil {
			log.Println("Failed to insert order: ", err)
			continue
		}

		msg.Ack(false)
		log.Println("Processed order: ", order)
	}
}