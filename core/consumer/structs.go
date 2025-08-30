package consumer

type Consumer struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Postcode int `json:"postcode"`
	ConsId int `json:"cons_id"`
	Password string
}

