package multiemissor_test

import (
	"context"
	"fmt"

	"github.com/vanclief/go-facturama/api/common"
	"github.com/vanclief/go-facturama/api/models"
	"github.com/vanclief/go-facturama/api/multiemissor"
)

// Registers the payment of a PPD invoice with the pagos complement.
// Payment CFDIs (CfdiType "P") have no items and no top level payment form
func ExampleClient_CreateCfdiV4XML_payment() {
	client := multiemissor.NewClient("username", "password", common.WithEnvironment(common.Sandbox))

	cfdi, err := client.CreateCfdiV4XML(context.Background(), multiemissor.CreateCfdiV4Request{
		Folio:           "100",
		Serie:           "COP",
		CfdiType:        "P",
		ExpeditionPlace: "78000",
		Exportation:     "01",
		Issuer: models.IssuerV4BindingModel{
			Rfc:          "EKU9003173C9",
			Name:         "ESCUELA KEMPER URGATE",
			FiscalRegime: "601",
		},
		Receiver: models.ReceiverV4BindingModel{
			Rfc:          "URE180429TM6",
			Name:         "UNIVERSIDAD ROBOTICA ESPAÑOLA",
			CfdiUse:      "CP01",
			FiscalRegime: "601",
			TaxZipCode:   "86991",
		},
		Complemento: &models.Complementv4{
			Payments: []models.PaymentModel{
				{
					Date:            "2026-10-26T12:00:00",
					PaymentForm:     "03",
					Currency:        "MXN",
					Amount:          116,
					OperationNumber: "126456",
					RelatedDocuments: []models.RelatedDocumentModel{
						{
							Uuid:                  "11c9536f-6bd0-4c32-9337-f8a73a15b775", // UUID of the PPD invoice being paid
							Folio:                 "45",
							Currency:              "MXN",
							PaymentMethod:         "PPD",
							PartialityNumber:      1,
							PreviousBalanceAmount: 116,
							AmountPaid:            116,
							TaxObject:             "02",
							Taxes: []models.TaxBindingModel{
								{Name: "IVA", Base: 100, Rate: 0.16, Total: 16},
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfdi.TaxStamp.UUID)
}

// Issues a donation receipt with the donatarias complement. The issuer must be
// an authorized donataria and the items are not subject to taxes
func ExampleClient_CreateCfdiV4XML_donation() {
	client := multiemissor.NewClient("username", "password", common.WithEnvironment(common.Sandbox))

	cfdi, err := client.CreateCfdiV4XML(context.Background(), multiemissor.CreateCfdiV4Request{
		Folio:             "100",
		Serie:             "DONST",
		CfdiType:          "I",
		Currency:          "MXN",
		PaymentForm:       "03",
		PaymentMethod:     "PUE",
		PaymentConditions: "DONATIVO",
		ExpeditionPlace:   "78000",
		Exportation:       "01",
		Issuer: models.IssuerV4BindingModel{
			Rfc:          "EKU9003173C9",
			Name:         "ESCUELA KEMPER URGATE",
			FiscalRegime: "601",
		},
		Receiver: models.ReceiverV4BindingModel{
			Rfc:          "CACX7605101P8",
			Name:         "XOCHILT CASAS CHAVEZ",
			CfdiUse:      "D04",
			FiscalRegime: "605",
			TaxZipCode:   "36257",
		},
		Items: []models.ItemFullBindingModel{
			{
				ProductCode: "20102000",
				Description: "Cobija de lana y algodon",
				Unit:        "NO APLICA",
				UnitCode:    "EA",
				UnitPrice:   1000,
				Quantity:    3,
				Subtotal:    3000,
				TaxObject:   "01",
				Total:       3000,
			},
		},
		Complemento: &models.Complementv4{
			Donation: &models.DonationModel{
				AuthorizationDate:   "30/01/2019",
				AuthorizationNumber: "B400-05-08-2014-005",
				Legend:              "El comprobante es un donativo",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfdi.TaxStamp.UUID)
}

// Bills an item on behalf of a third party with the a cuenta de terceros complement
func ExampleClient_CreateCfdiV4XML_thirdPartyAccount() {
	client := multiemissor.NewClient("username", "password", common.WithEnvironment(common.Sandbox))

	cfdi, err := client.CreateCfdiV4XML(context.Background(), multiemissor.CreateCfdiV4Request{
		Folio:             "100",
		Serie:             "FA",
		CfdiType:          "I",
		Currency:          "MXN",
		PaymentForm:       "03",
		PaymentMethod:     "PUE",
		PaymentConditions: "Por cuenta de terceros",
		ExpeditionPlace:   "78000",
		Exportation:       "01",
		Issuer: models.IssuerV4BindingModel{
			Rfc:          "EKU9003173C9",
			Name:         "ESCUELA KEMPER URGATE",
			FiscalRegime: "601",
		},
		Receiver: models.ReceiverV4BindingModel{
			Rfc:          "URE180429TM6",
			Name:         "UNIVERSIDAD ROBOTICA ESPAÑOLA",
			CfdiUse:      "G03",
			FiscalRegime: "601",
			TaxZipCode:   "86991",
		},
		Items: []models.ItemFullBindingModel{
			{
				ProductCode: "25173108",
				Description: "GPS estandar pruebas",
				UnitCode:    "E48",
				UnitPrice:   100,
				Quantity:    1,
				Subtotal:    100,
				TaxObject:   "02",
				Taxes: []models.TaxBindingModel{
					{Name: "IVA", Base: 100, Rate: 0.16, Total: 16},
				},
				Total: 116,
				ThirdPartyAccount: &models.ThirdPartyAccountModel{
					Rfc:          "CACX7605101P8",
					Name:         "XOCHILT CASAS CHAVEZ",
					FiscalRegime: "616",
					TaxZipCode:   "36257",
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cfdi.TaxStamp.UUID)
}
