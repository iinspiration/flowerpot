package responses

type Member struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Lastname     string `json:"lastname"`
	MobileNumber string `json:"mobile_number"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	TaxDetail    string `json:"tax_detail"`
	Level        string `json:"level"`
	Status       string `json:"status"`
	RegisteredAt string `json:"registered_at"`
}
