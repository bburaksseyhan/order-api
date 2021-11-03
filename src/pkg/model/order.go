package model

type Order struct {
	Id             int        `json:"id"`
	TrackingNumber string     `json:"trackingnumber"`
	Customer       Customer   `json:"customer"`
	Product        []Products `json:"products"`
}

type Customer struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

type Products struct {
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float32 `json:"price"`
}
