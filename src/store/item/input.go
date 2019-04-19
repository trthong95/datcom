package item

// Items ..
type Items struct {
	ItemNames []Item
}

// Item ..
type Item struct {
	ID       int
	ItemName string
	MenuID   int
}
