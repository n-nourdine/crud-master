package main

type Order struct {
	ID            int     `json:"id"`
	UserID        string  `json:"user_id"`
	NumberOfItems int     `json:"number_of_items"`
	TotalAmount   float64 `json:"total_amount"`
}