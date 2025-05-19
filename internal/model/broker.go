package model

type BrokerRequest struct {
	Code string `json:"code,omitempty" validate:"omitempty,len=2"`
}

type BrokerResponse struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	License string `json:"license"`
}
