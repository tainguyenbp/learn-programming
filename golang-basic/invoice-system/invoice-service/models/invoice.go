package models

import (
    "time"
)

type Item struct {
    Description string  `json:"description" mongo:"description"`
    Quantity    int     `json:"quantity" mongo:"quantity"`
    Price       float64 `json:"price" mongo:"price"`
}

type Invoice struct {
    ID       string    `json:"id" mongo:"_id"`
    Customer string    `json:"customer" mongo:"customer"`
    Date     time.Time `json:"date" mongo:"date"`
    Items    []Item    `json:"items" mongo:"items"`
    Total    float64   `json:"total" mongo:"total"`
}
