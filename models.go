package models

import "time"

type Product struct {
	ID          string  `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Type        string  `db:"type" json:"type"`
	Gender      string  `db:"gender" json:"gender"`
	Size        string  `db:"size" json:"size"`
	Color       string  `db:"color" json:"color"`
	Description string  `db:"description" json:"description"`
	Brand       string  `db:"brand" json:"brand"`
	Season      string  `db:"season" json:"season"`
	Price       float64 `db:"price" json:"price"`
}

type Photo struct {
	ID        string `db:"id" json:"id"`
	Filename  string `db:"filename" json:"filename"`
	ProductID string `db:"product_id" json:"product_id"`
}

type Role struct {
	ID   string `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type User struct {
	ID               string    `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	Email            string    `db:"email" json:"email"`
	Phone            string    `db:"phone" json:"phone"`
	RoleID           string    `db:"role_id" json:"role_id"`
	RegistrationDate time.Time `db:"registration_date" json:"registration_date"`
}

type Cart struct {
	ID        string `db:"id" json:"id"`
	UserID    string `db:"user_id" json:"user_id"`
	ProductID string `db:"product_id" json:"product_id"`
	Count     int    `db:"count" json:"count"`
}

type Order struct {
	ID       string    `db:"id" json:"id"`
	UserID   string    `db:"user_id" json:"user_id"`
	DateTime time.Time `db:"date_time" json:"date_time"`
	Status   string    `db:"status" json:"status"`
}

type OrderItem struct {
	ID        string  `db:"id" json:"id"`
	OrderID   string  `db:"order_id" json:"order_id"`
	ProductID string  `db:"product_id" json:"product_id"`
	Price     float64 `db:"price" json:"price"`
}