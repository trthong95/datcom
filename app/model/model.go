package model

// Response ..
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Item ...
type Item struct {
	ItemNames []string `json:"item_names"`
}
