package entities

// Menu menu
type Menu struct {
	ID         uint    `json:"id"`
	ParentID   uint    `json:"parent_id"`
	Title      string  `json:"title"`
	Name       string  `json:"name"`
	Route      string  `json:"route"`
	IsChildren bool    `json:"is_children"`
	Menu       []*Menu `json:"children" gorm:"foreignKey:parent_id;references:id"`
}
