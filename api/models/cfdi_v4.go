package models

// GlobalInformationV4Model represents global information for a CFDI v4
type GlobalInformationV4Model struct {
	Periodicity string `json:"Periodicity"`
	Months      string `json:"Months"`
	Year        int    `json:"Year"`
}

// Cfdiv4Relations represents CFDI v4 relations
type Cfdiv4Relations struct {
	Type  string       `json:"Type"`
	Cfdis []CfdiUuidID `json:"Cfdis"`
}

// CfdiUuidID represents a CFDI UUID ID
type CfdiUuidID struct {
	Uuid string `json:"Uuid"`
}

// IssuerV4BindingModel represents the issuer for a CFDI v4
type IssuerV4BindingModel struct {
	FiscalRegime   string               `json:"FiscalRegime"`
	Rfc            string               `json:"Rfc"`
	Name           string               `json:"Name,omitempty"`
	FacAtrAcquirer string               `json:"FacAtrAcquirer,omitempty"`
	Address        *AddressBindingModel `json:"Address,omitempty"`
}

// ReceiverV4BindingModel represents the receiver for a CFDI v4
type ReceiverV4BindingModel struct {
	ID                    string               `json:"Id,omitempty"`
	Rfc                   string               `json:"Rfc"`
	Name                  string               `json:"Name"`
	CfdiUse               string               `json:"CfdiUse"`
	FiscalRegime          string               `json:"FiscalRegime"`
	TaxZipCode            string               `json:"TaxZipCode"`
	TaxResidence          string               `json:"TaxResidence,omitempty"`
	TaxRegistrationNumber string               `json:"TaxRegistrationNumber,omitempty"`
	Address               *AddressBindingModel `json:"Address,omitempty"`
}

// AddressBindingModel represents an address
type AddressBindingModel struct {
	Street         string `json:"Street"`
	ExteriorNumber string `json:"ExteriorNumber"`
	InteriorNumber string `json:"InteriorNumber,omitempty"`
	Neighborhood   string `json:"Neighborhood"`
	ZipCode        string `json:"ZipCode"`
	Locality       string `json:"Locality,omitempty"`
	Municipality   string `json:"Municipality"`
	State          string `json:"State"`
	Country        string `json:"Country"`
	ID             string `json:"Id,omitempty"`
}

// ItemFullBindingModel represents an item in a CFDI v4
type ItemFullBindingModel struct {
	IDProduct            string                  `json:"IdProduct,omitempty"`
	ProductCode          string                  `json:"ProductCode"`
	IdentificationNumber string                  `json:"IdentificationNumber,omitempty"`
	SKU                  string                  `json:"SKU,omitempty"`
	Description          string                  `json:"Description"`
	Unit                 string                  `json:"Unit"`
	UnitCode             string                  `json:"UnitCode"`
	UnitPrice            float64                 `json:"UnitPrice"`
	Quantity             float64                 `json:"Quantity"`
	Subtotal             float64                 `json:"Subtotal"`
	Discount             float64                 `json:"Discount,omitempty"`
	TaxObject            string                  `json:"TaxObject"`
	Taxes                []TaxBindingModel       `json:"Taxes,omitempty"`
	ThirdPartyAccount    *ThirdPartyAccountModel `json:"ThirdPartyAccount,omitempty"`
	PropertyTaxIDNumber  []string                `json:"PropertyTaxIDNumber,omitempty"`
	NumerosPedimento     []string                `json:"NumerosPedimento,omitempty"`
	Parts                []ItemPartBindingModel  `json:"Parts,omitempty"`
	Total                float64                 `json:"Total"`
	Complement           *ItemComplementModel    `json:"Complement,omitempty"`
}

// TaxBindingModel represents a tax in a CFDI v4
type TaxBindingModel struct {
	Total       float64 `json:"Total"`
	Name        string  `json:"Name"`
	Base        float64 `json:"Base"`
	Rate        float64 `json:"Rate"`
	IsRetention bool    `json:"IsRetention"`
	IsQuota     bool    `json:"IsQuota"`
	TaxObject   string  `json:"TaxObject,omitempty"`
}

// ThirdPartyAccountModel represents a third party account in a CFDI v4
type ThirdPartyAccountModel struct {
	Rfc          string `json:"Rfc"`
	Name         string `json:"Name"`
	FiscalRegime string `json:"FiscalRegime"`
	TaxZipCode   string `json:"TaxZipCode"`
}

// ItemPartBindingModel represents a part of an item in a CFDI v4
type ItemPartBindingModel struct {
	Quantity             float64                   `json:"Quantity"`
	UnitCode             string                    `json:"UnitCode"`
	ProductCode          string                    `json:"ProductCode,omitempty"`
	IdentificationNumber string                    `json:"IdentificationNumber,omitempty"`
	Description          string                    `json:"Description"`
	UnitPrice            float64                   `json:"UnitPrice"`
	Amount               float64                   `json:"Amount"`
	CustomsInformation   []CustomsInformationModel `json:"CustomsInformation,omitempty"`
}

// CustomsInformationModel represents customs information in a CFDI v4
type CustomsInformationModel struct {
	Number  string `json:"Number"`
	Date    string `json:"Date"`
	Customs string `json:"Customs"`
}

// ItemComplementModel represents item complement in a CFDI v4
type ItemComplementModel struct {
	EducationalInstitution *EducationalInstitutionModel `json:"EducationalInstitution,omitempty"`
	ThirdPartyAccount      *ThirdPartyAccountFullModel  `json:"ThirdPartyAccount,omitempty"`
}

// EducationalInstitutionModel represents educational institution in a CFDI v4
type EducationalInstitutionModel struct {
	StudentsName   string `json:"StudentsName"`
	Curp           string `json:"Curp"`
	EducationLevel string `json:"EducationLevel"`
	AutRvoe        string `json:"AutRvoe"`
	PaymentRfc     string `json:"PaymentRfc"`
}

// ThirdPartyAccountFullModel represents third party account in a CFDI v4
type ThirdPartyAccountFullModel struct {
	Rfc                 string                    `json:"Rfc"`
	Name                string                    `json:"Name"`
	FiscalRegime        string                    `json:"FiscalRegime"`
	TaxZipCode          string                    `json:"TaxZipCode"`
	ThirdTaxInformation *ThirdTaxInformationModel `json:"ThirdTaxInformation,omitempty"`
	CustomsInformation  *CustomsInformationModel  `json:"CustomsInformation,omitempty"`
	Parts               []PartModel               `json:"Parts,omitempty"`
	PropertyTaxNumber   string                    `json:"PropertyTaxNumber,omitempty"`
	Taxes               []ThirdPartyTaxModel      `json:"Taxes,omitempty"`
}

// ThirdTaxInformationModel represents third tax information in a CFDI v4
type ThirdTaxInformationModel struct {
	Street         string `json:"Street"`
	ExteriorNumber string `json:"ExteriorNumber"`
	InteriorNumber string `json:"InteriorNumber,omitempty"`
	Neighborhood   string `json:"Neighborhood"`
	Locality       string `json:"Locality,omitempty"`
	Reference      string `json:"Reference,omitempty"`
	Municipality   string `json:"Municipality"`
	State          string `json:"State"`
	Country        string `json:"Country"`
	PostalCode     string `json:"PostalCode"`
	ZipCode        string `json:"ZipCode"`
}

// PartModel represents a part in a CFDI v4
type PartModel struct {
	Quantity             float64                   `json:"Quantity"`
	Unit                 string                    `json:"Unit"`
	IdentificationNumber string                    `json:"IdentificationNumber,omitempty"`
	Description          string                    `json:"Description"`
	UnitPrce             float64                   `json:"UnitPrce"`
	Amount               float64                   `json:"Amount"`
	CustomsInformation   []CustomsInformationModel `json:"CustomsInformation,omitempty"`
}

// ThirdPartyTaxModel represents third party tax in a CFDI v4
type ThirdPartyTaxModel struct {
	Name   string  `json:"Name"`
	Rate   float64 `json:"Rate"`
	Amount float64 `json:"Amount"`
}

// Complementv4 represents complement in a CFDI v4
type Complementv4 struct {
	NotariosPublicos *NotariosPublicosModel `json:"NotariosPublicos,omitempty"`
	Ine              *IneModel              `json:"Ine,omitempty"`
	Detallista       *DetallistaModel       `json:"Detallista,omitempty"`
	Payments         []PaymentModel         `json:"Payments,omitempty"`
	Donation         *DonationModel         `json:"Donation,omitempty"`
	ForeignTrade     *ForeignTradeModel     `json:"ForeignTrade,omitempty"`
	Payroll          *PayrollModel          `json:"Payroll,omitempty"`
	TaxLegends       *TaxLegendsModel       `json:"TaxLegends,omitempty"`
	CartaPorte31     *CartaPorte31Model     `json:"CartaPorte31,omitempty"`
	ValesDeDespensa  *ValesDeDespensaModel  `json:"ValesDeDespensa,omitempty"`
}

// NotariosPublicosModel represents notarios públicos in a CFDI v4
type NotariosPublicosModel struct {
	DescInmuebles    []DescInmueblesModel   `json:"DescInmuebles,omitempty"`
	DatosOperacion   *DatosOperacionModel   `json:"DatosOperacion,omitempty"`
	DatosNotario     *DatosNotarioModel     `json:"DatosNotario,omitempty"`
	DatosEnajenante  *DatosEnajenanteModel  `json:"DatosEnajenante,omitempty"`
	DatosAdquiriente *DatosAdquirienteModel `json:"DatosAdquiriente,omitempty"`
}

// PlaceholderModels to satisfy compilation - implement as needed
type (
	DescInmueblesModel    struct{}
	DatosOperacionModel   struct{}
	DatosNotarioModel     struct{}
	DatosEnajenanteModel  struct{}
	DatosAdquirienteModel struct{}
	IneModel              struct{}
	DetallistaModel       struct{}
)

// PaymentModel represents payment in a CFDI v4
type PaymentModel struct {
	SignPayment                   string                 `json:"SignPayment,omitempty"`
	CertPayment                   string                 `json:"CertPayment,omitempty"`
	OriginalString                string                 `json:"OriginalString,omitempty"`
	StringTypePayment             string                 `json:"StringTypePayment,omitempty"`
	RelatedDocuments              []RelatedDocumentModel `json:"RelatedDocuments,omitempty"`
	Taxes                         []TaxBindingModel      `json:"Taxes,omitempty"`
	Date                          string                 `json:"Date"`
	PaymentForm                   string                 `json:"PaymentForm"`
	Currency                      string                 `json:"Currency"`
	ExchangeRate                  float64                `json:"ExchangeRate,omitempty"`
	Amount                        float64                `json:"Amount"`
	OperationNumber               string                 `json:"OperationNumber,omitempty"`
	RfcIssuerPayerAccount         string                 `json:"RfcIssuerPayerAccount,omitempty"`
	ForeignAccountNamePayer       string                 `json:"ForeignAccountNamePayer,omitempty"`
	PayerAccount                  string                 `json:"PayerAccount,omitempty"`
	RfcReceiverBeneficiaryAccount string                 `json:"RfcReceiverBeneficiaryAccount,omitempty"`
	BeneficiaryAccount            string                 `json:"BeneficiaryAccount,omitempty"`
	ExpectedPaid                  float64                `json:"ExpectedPaid,omitempty"`
}

// RelatedDocumentModel represents related document in a payment
type RelatedDocumentModel struct {
	Uuid                  string            `json:"Uuid"`
	Serie                 string            `json:"Serie,omitempty"`
	Folio                 string            `json:"Folio,omitempty"`
	Currency              string            `json:"Currency,omitempty"`
	EquivalenceDocRel     float64           `json:"EquivalenceDocRel,omitempty"`
	ExchangeRate          float64           `json:"ExchangeRate,omitempty"`
	PaymentMethod         string            `json:"PaymentMethod,omitempty"`
	PartialityNumber      int               `json:"PartialityNumber,omitempty"`
	PreviousBalanceAmount float64           `json:"PreviousBalanceAmount,omitempty"`
	AmountPaid            float64           `json:"AmountPaid"`
	TaxObject             string            `json:"TaxObject,omitempty"`
	Taxes                 []TaxBindingModel `json:"Taxes,omitempty"`
}

// Placeholder models for other complementos - implement as needed
type (
	DonationModel        struct{}
	ForeignTradeModel    struct{}
	PayrollModel         struct{}
	TaxLegendsModel      struct{}
	CartaPorte31Model    struct{}
	ValesDeDespensaModel struct{}
)
