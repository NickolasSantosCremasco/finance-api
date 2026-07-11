package main

import (

	"time"
	"finance-api/internal/transaction"
)

func main() {
	t := transaction.Transaction{
		ID: 	1,
		Description: "Mercado",
		Amount: 120.50,
		Category: "Alimentação",
		Merchant: "Supermercado ABC",
		Date: time.Now(),
	}

	t.Print()
}