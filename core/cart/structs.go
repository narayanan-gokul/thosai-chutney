package cart

type CreateCartItem struct {
	ItemId int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type CreateCartRequest struct {
	DistId int
	Items []CreateCartItem
}
