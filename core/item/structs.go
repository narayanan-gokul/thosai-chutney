package item

type Item struct {
	Name string `json:"name"`
	ItemId string `json:"item_id"`
	MaxCap string `json:"max_cap"`
	Price float64 `json:"price"`
}
