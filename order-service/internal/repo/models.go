package repo

type OrderStatus string

const (
	StatusConfirmed OrderStatus = "CONFIRMED"
	StatusRejected  OrderStatus = "REJECTED"
)

type Order struct {
	ID     string
	UserID string
	Status OrderStatus
}

type OrderItem struct {
	OrderID  string
	SKU      string
	Quantity int32
}
