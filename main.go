package main

import (
	"fmt"
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

	fmt.Printf(
		"ID: %d\nDescrição: %s\nValor: R$ %.2f\nCategoria: %s\nLoja: %s\nData: %s\n",
		t.ID,
		t.Description,
		t.Amount,
		t.Category,
		t.Merchant,
		t.Date.Format("02/01/2006 15:04"),
	)
}