package main

import (
	"finance-api/internal/transaction"
	"fmt"
	"time"
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

	if t.IsExpense() {
		fmt.Println("É uma despesa")
	} else {
		fmt.Println("É uma receita")
	}
 }