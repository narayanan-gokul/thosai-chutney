package shipment

type Shipment struct {
	ShipId int `json:"ship_id"`
	ItemId int `json:"item_id"`
	Quantity int `json:"quantity"`
	SuppId int `json:"supp_id"`
}
