package domain

type Member struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	RegistationDate string `json:"registration_date"`
}
