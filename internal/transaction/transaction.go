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

func (t Transaction) Print() { /* The Transaction have a method called Print*/
	/*t Transaction is called Receiver, he says that This Method belongs to Transation Type""*/ 
	fmt.Println("========== TRANSACTION ==========")
	fmt.Println("ID:", t.ID)
	fmt.Println("Descrição:", t.Description)
	fmt.Printf("Valor: R$ %.2f\n", t.Amount)
	fmt.Println("Categoria:", t.Category)
	fmt.Println("Loja:", t.Merchant)
	fmt.Println("Data:", t.Date.Format("02/01/2006 15:04"))
	fmt.Println("=================================")
}