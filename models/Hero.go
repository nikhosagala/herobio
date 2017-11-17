package models

// Hero model
type Hero struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Born        string `json:"born"`
	Death       string `json:"death"`
	Province    string `json:"province"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}
