package domain

// ItemInput ..
type ItemInput struct {
	Items []Item
}

// Item ..
type Item struct {
	ID       int
	ItemName string
	MenuID   int
}
