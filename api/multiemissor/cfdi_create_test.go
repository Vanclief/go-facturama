package multiemissor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vanclief/go-facturama/api/models"
)

func TestCreateCfdiV4RequestValidateItems(t *testing.T) {
	request := CreateCfdiV4Request{
		ExpeditionPlace: "78116",
		Folio:           "1",
		Issuer:          models.IssuerV4BindingModel{Rfc: "EKU9003173C9", FiscalRegime: "601"},
		Receiver: models.ReceiverV4BindingModel{
			Rfc:          "URE180429TM6",
			Name:         "UNIVERSIDAD ROBOTICA ESPAÑOLA",
			CfdiUse:      "CP01",
			FiscalRegime: "601",
			TaxZipCode:   "86991",
		},
	}

	// Payment CFDIs carry their amounts in the payments complement, not in items
	request.CfdiType = "P"
	assert.NoError(t, request.Validate())

	request.CfdiType = "I"
	assert.ErrorContains(t, request.Validate(), "At least one item is required")

	// Unit is optional, the API only needs UnitCode
	request.Items = []models.ItemFullBindingModel{
		{ProductCode: "01010101", Description: "Test", UnitCode: "H87", TaxObject: "01"},
	}
	assert.NoError(t, request.Validate())
}
