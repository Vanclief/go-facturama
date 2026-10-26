package models

// CfdiInfoModel represents the information of a CFDI (Mexican digital invoice)
type CfdiInfoModel struct {
	ID                   string                 `json:"Id"`
	CfdiType             string                 `json:"CfdiType"`
	Type                 string                 `json:"Type"`
	Serie                string                 `json:"Serie"`
	Folio                string                 `json:"Folio"`
	Date                 string                 `json:"Date"`
	CertNumber           string                 `json:"CertNumber"`
	PaymentTerms         string                 `json:"PaymentTerms"`
	PaymentConditions    string                 `json:"PaymentConditions"`
	PaymentMethod        string                 `json:"PaymentMethod"`
	PaymentAccountNumber string                 `json:"PaymentAccountNumber"`
	PaymentBankName      string                 `json:"PaymentBankName"`
	ExpeditionPlace      string                 `json:"ExpeditionPlace"`
	ExchangeRate         float64                `json:"ExchangeRate"`
	Currency             string                 `json:"Currency"`
	Subtotal             float64                `json:"Subtotal"`
	Discount             float64                `json:"Discount"`
	Total                float64                `json:"Total"`
	Observations         string                 `json:"Observations"`
	OrderNumber          string                 `json:"OrderNumber"`
	Issuer               TaxEntityInfoViewModel `json:"Issuer"`
	Receiver             ReceiverViewModel      `json:"Receiver"`
	Items                []ItemInfoModel        `json:"Items"`
	Taxes                []TaxInfoModel         `json:"Taxes"`
	Complement           CfdiComplement         `json:"Complement"`
}

// TaxEntityInfoViewModel represents tax entity information for the CFDI issuer
type TaxEntityInfoViewModel struct {
	FiscalRegime string `json:"FiscalRegime"`
	Rfc          string `json:"Rfc"`
	TaxName      string `json:"TaxName"`
}

// ReceiverViewModel represents information for the CFDI receiver
type ReceiverViewModel struct {
	Rfc  string `json:"Rfc"`
	Name string `json:"Name"`
}

// ItemInfoModel represents information for a CFDI item
type ItemInfoModel struct {
	Discount    float64 `json:"Discount"`
	Quantity    float64 `json:"Quantity"`
	Unit        string  `json:"Unit"`
	Description string  `json:"Description"`
	UnitValue   float64 `json:"UnitValue"`
	Total       float64 `json:"Total"`
}

// TaxInfoModel represents information for a CFDI tax
type TaxInfoModel struct {
	Total float64 `json:"Total"`
	Name  string  `json:"Name"`
	Rate  float64 `json:"Rate"`
	Type  string  `json:"Type"`
}

// CfdiComplement represents complementary information for a CFDI
type CfdiComplement struct {
	TaxStamp CfdiTaxStamp `json:"TaxStamp"`
}

// CfdiTaxStamp represents a CFDI tax stamp
type CfdiTaxStamp struct {
	Version                string `json:"Version"`
	TaxStampOriginalString string `json:"TaxStampOriginalString"`
	UUID                   string `json:"Uuid"`
	Date                   string `json:"Date"`
	CfdiSign               string `json:"CfdiSign"`
	SatCertNumber          string `json:"SatCertNumber"`
	SatSign                string `json:"SatSign"`
	RfcProvCertif          string `json:"RfcProvCertif"`
}

// CfdiStampModel is the response of creating a CFDI through /api-lite/4/cfdis
type CfdiStampModel struct {
	ID             string       `json:"Id"`
	XMLBase64      string       `json:"XmlBase64"`
	OriginalString string       `json:"OriginalString"`
	TaxStamp       CfdiTaxStamp `json:"TaxStamp"`
}

// CfdiSearchResult is a CFDI as returned when listing CFDIs
type CfdiSearchResult struct {
	ID        string  `json:"Id"`
	Type      string  `json:"Type"`
	Serie     string  `json:"Serie"`
	Folio     string  `json:"Folio"`
	TaxName   string  `json:"TaxName"`
	Rfc       string  `json:"Rfc"`
	RfcIssuer string  `json:"RfcIssuer"`
	Date      string  `json:"Date"`
	Subtotal  float64 `json:"Subtotal"`
	Total     float64 `json:"Total"`
	UUID      string  `json:"Uuid"`
	IsActive  bool    `json:"IsActive"`
	Status    string  `json:"Status"`
}

// CfdiStatusModel is the status of a CFDI with the SAT
type CfdiStatusModel struct {
	// Status is one of "Vigente", "Cancelado" or "No encontrado"
	Status string `json:"Status"`
	// IsCancelable is one of "No cancelable", "Cancelable con aceptación" or "Cancelable sin aceptación"
	IsCancelable string `json:"IsCancelable"`
	UUID         string `json:"Uuid"`
}
