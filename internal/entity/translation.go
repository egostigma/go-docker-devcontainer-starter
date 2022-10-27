// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Response -.
type Response struct {
	Status  string `json:"status" example:"Success"`
	Message string `json:"message" example:"Successfully Saved"`
}

// Translation -.
type Translation struct {
	ID          uint   `gorm:"primaryKey" json:"id" example:"1"`
	Source      string `json:"source" example:"auto"`
	Destination string `json:"destination" example:"en"`
	Original    string `json:"original" example:"текст для перевода"`
	Translation string `json:"translation" example:"text for translation"`
}
