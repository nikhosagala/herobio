package models

// Hero model
type Hero struct {
	ID          int    `json:"id" gorm:"AUTO_INCREMENT"`
	Name        string `json:"name"`
	Born        string `json:"born"`
	Death       string `json:"death"`
	Province    string `json:"province"`
	ImageURL    string `json:"imageUrl"`
	Description string `json:"description"`
}
