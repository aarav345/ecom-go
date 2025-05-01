package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // don't return password
	CreatedAt time.Time `json:"created_at"`
}

type ProductStore interface {
	GetProducts() ([]ProductWithInventory, error)
	GetProductsByID(ps []int) ([]ProductWithInventory, error)
	UpdateProduct(Product ProductWithInventory, updateProductFields bool) error
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type Inventory struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductWithInventory struct {
	Product   `json:"product"`
	Inventory `json:"inventory"`
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type History struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	ProductName  string `json:"product_name"`
	ProductImage string `json:"product_image"`
	Quantity     int    `json:"quantity"`
	Price        int    `json:"price"`
	Status       string `json:"status"`
	OrderDate    string `json:"order_date"`
}

type HistoryStore interface {
	CreateHistory(history History) error
}

type CartItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}
