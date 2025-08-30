package distributor

type Distributor struct {
	DistId int `json:"dist_id"`
	Postcode int `json:"postcode"`
	Name string `json:"name"`
	Password string `json:"password"`
}

type ReadDistributorResponse struct {
	DistId int `json:"dist_id"`
	Name string `json:"name"`
	Postcode int `json:"postcode"`
}
