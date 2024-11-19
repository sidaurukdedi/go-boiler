package domain

type User struct {
	Name    string `json:"name"`
	Address string `json:"address" validate:"required"`
}
