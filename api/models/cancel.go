package models

// CancelationStatusLite represents the response for a CFDI cancellation
type CancelationStatusLite struct {
	Status           string `json:"Status"`
	Message          string `json:"Message"`
	UUID             string `json:"Uuid,omitempty"`
	RequestDate      string `json:"RequestDate,omitempty"`
	AcuseXmlBase64   string `json:"AcuseXmlBase64,omitempty"`
	CancelationDate  string `json:"CancelationDate,omitempty"`
}