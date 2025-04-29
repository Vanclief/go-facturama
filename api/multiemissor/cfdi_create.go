package multiemissor

import (
	"context"
	"fmt"
	"regexp"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// CreateCfdiV4Request represents a request to create a CFDI v4
type CreateCfdiV4Request struct {
	NameID               int                       `json:"NameId,omitempty"`
	LogoURL              string                    `json:"LogoUrl,omitempty"`
	Date                 string                    `json:"Date,omitempty"`
	Serie                string                    `json:"Serie,omitempty"`
	PaymentAccountNumber string                    `json:"PaymentAccountNumber,omitempty"`
	CurrencyExchangeRate float64                   `json:"CurrencyExchangeRate,omitempty"`
	Currency             string                    `json:"Currency,omitempty"`
	ExpeditionPlace      string                    `json:"ExpeditionPlace"`
	Exportation          string                    `json:"Exportation,omitempty"`
	PaymentConditions    string                    `json:"PaymentConditions,omitempty"`
	GlobalInformation    *models.GlobalInformationV4Model `json:"GlobalInformation,omitempty"`
	Relations            *models.Cfdiv4Relations          `json:"Relations,omitempty"`
	Folio                string                    `json:"Folio"`
	CfdiType             string                    `json:"CfdiType"`
	PaymentForm          string                    `json:"PaymentForm,omitempty"`
	PaymentMethod        string                    `json:"PaymentMethod,omitempty"`
	Issuer               models.IssuerV4BindingModel      `json:"Issuer"`
	Receiver             models.ReceiverV4BindingModel    `json:"Receiver"`
	Items                []models.ItemFullBindingModel    `json:"Items"`
	Complemento          *models.Complementv4             `json:"Complemento,omitempty"`
	Observations         string                    `json:"Observations,omitempty"`
	OrderNumber          string                    `json:"OrderNumber,omitempty"`
	PaymentBankName      string                    `json:"PaymentBankName,omitempty"`
}

// Validate validates the request to create a CFDI v4
func (request *CreateCfdiV4Request) Validate() error {
	const op = "CreateCfdiV4Request.Validate"

	// Required fields
	if request.ExpeditionPlace == "" {
		return ez.New(op, ez.EINVALID, "ExpeditionPlace is required", nil)
	}
	if request.Folio == "" {
		return ez.New(op, ez.EINVALID, "Folio is required", nil)
	}
	if request.CfdiType == "" {
		return ez.New(op, ez.EINVALID, "CfdiType is required", nil)
	}

	// Validate CfdiType (I|E|T|N|P)
	if !regexp.MustCompile(`^[IETNP]$`).MatchString(request.CfdiType) {
		return ez.New(op, ez.EINVALID, "CfdiType must be one of: I, E, T, N, P", nil)
	}

	// Validate ExpeditionPlace (5 digit zip code)
	if !regexp.MustCompile(`^[0-9]{5}$`).MatchString(request.ExpeditionPlace) {
		return ez.New(op, ez.EINVALID, "ExpeditionPlace must be a 5-digit zip code", nil)
	}

	// Validate Folio (1-40 chars)
	if len(request.Folio) < 1 || len(request.Folio) > 40 {
		return ez.New(op, ez.EINVALID, "Folio must be between 1 and 40 characters", nil)
	}

	// Validate conditional fields
	if request.PaymentForm != "" {
		validPaymentForms := []string{"01", "02", "03", "04", "05", "06", "08", "12", "13", "14", "15", "17", "23", "24", "25", "26", "27", "28", "29", "30", "31", "99"}
		isValid := false
		for _, v := range validPaymentForms {
			if request.PaymentForm == v {
				isValid = true
				break
			}
		}
		if !isValid {
			return ez.New(op, ez.EINVALID, "Invalid PaymentForm value", nil)
		}
	}

	if request.PaymentMethod != "" {
		if request.PaymentMethod != "PUE" && request.PaymentMethod != "PPD" {
			return ez.New(op, ez.EINVALID, "PaymentMethod must be either PUE or PPD", nil)
		}
	}

	// Validate Issuer
	if request.Issuer.Rfc == "" {
		return ez.New(op, ez.EINVALID, "Issuer.Rfc is required", nil)
	}
	if request.Issuer.FiscalRegime == "" {
		return ez.New(op, ez.EINVALID, "Issuer.FiscalRegime is required", nil)
	}

	// Validate Receiver
	if request.Receiver.Rfc == "" {
		return ez.New(op, ez.EINVALID, "Receiver.Rfc is required", nil)
	}
	if request.Receiver.Name == "" {
		return ez.New(op, ez.EINVALID, "Receiver.Name is required", nil)
	}
	if request.Receiver.CfdiUse == "" {
		return ez.New(op, ez.EINVALID, "Receiver.CfdiUse is required", nil)
	}
	if request.Receiver.FiscalRegime == "" {
		return ez.New(op, ez.EINVALID, "Receiver.FiscalRegime is required", nil)
	}
	if request.Receiver.TaxZipCode == "" {
		return ez.New(op, ez.EINVALID, "Receiver.TaxZipCode is required", nil)
	}

	// Validate Items
	if len(request.Items) == 0 {
		return ez.New(op, ez.EINVALID, "At least one item is required", nil)
	}

	for i, item := range request.Items {
		if item.ProductCode == "" {
			return ez.New(op, ez.EINVALID, fmt.Sprintf("Items[%d].ProductCode is required", i), nil)
		}
		if item.Description == "" {
			return ez.New(op, ez.EINVALID, fmt.Sprintf("Items[%d].Description is required", i), nil)
		}
		if item.Unit == "" {
			return ez.New(op, ez.EINVALID, fmt.Sprintf("Items[%d].Unit is required", i), nil)
		}
		if item.UnitCode == "" {
			return ez.New(op, ez.EINVALID, fmt.Sprintf("Items[%d].UnitCode is required", i), nil)
		}
		if item.TaxObject == "" {
			return ez.New(op, ez.EINVALID, fmt.Sprintf("Items[%d].TaxObject is required", i), nil)
		}
	}

	return nil
}

// CreateCfdiV4 creates a new CFDI v4 (Mexican digital invoice)
// Endpoint: POST /api-lite/3/cfdis
func (c *Client) CreateCfdiV4(ctx context.Context, request CreateCfdiV4Request) (*models.CfdiInfoModel, error) {
	const op = "multiemissor.CreateCfdiV4"

	// Validate required fields
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	// Using /api-lite/3/cfdis as the endpoint for CFDI v4 creation
	path := "/api-lite/3/cfdis"
	var result models.CfdiInfoModel

	err = c.Post(ctx, path, request, &result)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return &result, nil
}
