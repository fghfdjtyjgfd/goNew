package repo

// PageInformation page information
type PageInformation struct {
	Page                  int   `json:"page,omitempty"`
	Size                  int   `json:"size,omitempty"`
	TotalNumberOfEntities int64 `json:"total_number_of_entities,omitempty"`
	TotalNumberOfPages    int   `json:"total_number_of_pages,omitempty"`
}
