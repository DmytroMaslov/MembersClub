package domain

import (
	"time"
)

type Member struct {
	Name            string    `json:"name" validate:"alpha|containsany= ."`
	Email           string    `json:"email" validate:"email"`
	RegistationDate time.Time `json:"registration_date" validate:"lte"`
}
