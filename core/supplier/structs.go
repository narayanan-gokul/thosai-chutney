package supplier

type Supplier struct {
	SuppId int `json:"supp_id"`
	Postcode int `json:"postcode"`
	Name string `json:"name"`
	Password string `json:"password"`
}
