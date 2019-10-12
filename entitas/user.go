package entitas

import (
	"time"

	"github.com/jinzhu/gorm"
)

type (
	Login struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	User struct {
		gorm.Model
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserTamp struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
