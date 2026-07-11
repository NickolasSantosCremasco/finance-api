package transaction

import (
	"fmt"
	"time"
)


type Transaction struct {
	ID 			int
	Description string
	Amount 		float64
	Category	string
	Merchant	string
	Date		time.Time
}

func (t Transaction) Print() { /* The Transaction has a method called Print*/
	/*t Transaction. It tells Go that this method belongs to Transaction*/ 
	fmt.Println("========== TRANSACTION ==========")
	fmt.Println("ID:", t.ID)
	fmt.Println("Descrição:", t.Description)
	fmt.Printf("Valor: R$ %.2f\n", t.Amount)
	fmt.Println("Categoria:", t.Category)
	fmt.Println("Loja:", t.Merchant)
	fmt.Println("Data:", t.Date.Format("02/01/2006 15:04"))
	fmt.Println("=================================")
}

func (t Transaction) IsExpense() bool {
	return t.Amount < 0
}

func (t Transaction) IsIncome() bool {
	return t.Amount > 0
}

func (t Transaction) Summary() string {
	return fmt.Sprintf("%s - R$ %.2f", t.Description, t.Amount)
}