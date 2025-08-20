package order

type Order struct {
	ID     int
	UserID int
	Amount float64
}

func New(id int, userID int, amount float64) Order {
	return Order{
		ID:     id,
		UserID: userID,
		Amount: amount,
	}
}
