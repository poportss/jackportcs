package dto

type PagarmeCreateCustomerRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	Type         string `json:"type"`

	Address PagarmeAddress `json:"address"`
	Phones  PagarmePhones  `json:"phones"`
}

type PagarmeAddress struct {
	ID        string `json:"id,omitempty"`
	Line1     string `json:"line_1"`
	Line2     string `json:"line_2"`
	ZipCode   string `json:"zip_code"`
	City      string `json:"city"`
	State     string `json:"state"`
	Country   string `json:"country"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type PagarmePhones struct {
	HomePhone PagarmeHomePhone `json:"home_phone"`
}

type PagarmeHomePhone struct {
	CountryCode string `json:"country_code"`
	AreaCode    string `json:"area_code"`
	Number      string `json:"number"`
}

type PagarmeCreateCustomerResponse struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Document     string `json:"document"`
	DocumentType string `json:"document_type"`
	Type         string `json:"type"`
}
