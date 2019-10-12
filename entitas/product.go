package entitas

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Product struct {
		gorm.Model
		Name   string `json:"name"`
		Price  string `json:"price"`
		ImgUrl string `json:"imageurl"`
	}
	ProductTamp struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Price     string    `json:"price"`
		ImgUrl    string    `json:"imageurl"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
