package distributor

type Distributor struct {
	DistId int `json:"dist_id"`
	Postcode int `json:"postcode"`
	Name string `json:"name"`
	Password string `json:"password"`
}
