package multiemissor

import (
	"slices"

	"github.com/vanclief/go-facturama/api/models"
)

// incomeRequest returns a valid income CFDI from the test issuer for 100 + IVA
func (s *APIClientSuite) incomeRequest(folio string) CreateCfdiV4Request {
	return CreateCfdiV4Request{
		ExpeditionPlace: "78116",
		Folio:           folio,
		CfdiType:        "I",
		PaymentForm:     "03",
		PaymentMethod:   "PUE",
		Currency:        "MXN",
		Exportation:     "01",
		Issuer: models.IssuerV4BindingModel{
			Rfc:          s.RFC,
			Name:         "FRANCO VALENCIA ADAN",
			FiscalRegime: "612",
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
				ProductCode: "01010101",
				Description: "Test product",
				Unit:        "PIEZA",
				UnitCode:    "H87",
				UnitPrice:   100,
				Quantity:    1,
				Subtotal:    100,
				TaxObject:   "02",
				Taxes: []models.TaxBindingModel{
					{Name: "IVA", Base: 100, Rate: 0.16, Total: 16},
				},
				Total: 116,
			},
		},
	}
}

func (s *APIClientSuite) TestCreateCfdiV4XMLAndList() {
	request := s.incomeRequest("Test-v4")
	request.Serie = "LIST"

	cfdi, err := s.Client.CreateCfdiV4XML(s.Context, request)
	s.Require().NoError(err)
	s.NotEmpty(cfdi.ID)
	s.NotEmpty(cfdi.XMLBase64)
	s.NotEmpty(cfdi.TaxStamp.UUID)

	// Other CFDIs may be issued meanwhile, so look it up by ID instead of position
	cfdis, err := s.Client.ListCfdis(s.Context, ListCfdisRequest{IssuerRFC: s.RFC, Serie: request.Serie, Status: "active"})
	s.Require().NoError(err)

	index := slices.IndexFunc(cfdis, func(listed models.CfdiSearchResult) bool {
		return listed.ID == cfdi.ID
	})
	s.Require().NotEqual(-1, index, "created CFDI not found in the list")
	s.Equal(cfdi.TaxStamp.UUID, cfdis[index].UUID)
	s.Equal(116.0, cfdis[index].Total)
	s.Equal(request.Serie, cfdis[index].Serie)
}

func (s *APIClientSuite) TestCreatePaymentComplement() {
	// Step 1: Create the invoice to be paid later (PPD)
	invoiceRequest := s.incomeRequest("Test-PPD")
	invoiceRequest.PaymentForm = "99"
	invoiceRequest.PaymentMethod = "PPD"

	invoice, err := s.Client.CreateCfdiV4(s.Context, invoiceRequest)
	s.Require().NoError(err)

	// Step 2: Register its payment. Payment CFDIs have no items
	paymentRequest := s.incomeRequest("Test-Pago")
	paymentRequest.CfdiType = "P"
	paymentRequest.PaymentForm = ""
	paymentRequest.PaymentMethod = ""
	paymentRequest.Currency = ""
	paymentRequest.Receiver.CfdiUse = "CP01"
	paymentRequest.Items = nil
	paymentRequest.Complemento = &models.Complementv4{
		Payments: []models.PaymentModel{
			{
				Date:        invoice.Date,
				PaymentForm: "03",
				Currency:    "MXN",
				Amount:      116,
				RelatedDocuments: []models.RelatedDocumentModel{
					{
						Uuid:                  invoice.Complement.TaxStamp.UUID,
						Folio:                 invoice.Folio,
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
	}

	payment, err := s.Client.CreateCfdiV4XML(s.Context, paymentRequest)
	s.Require().NoError(err)
	s.NotEmpty(payment.TaxStamp.UUID)
}

func (s *APIClientSuite) TestCreateDonationComplement() {
	request := s.incomeRequest("Test-Donativo")
	request.PaymentConditions = "DONATIVO"
	request.Receiver = models.ReceiverV4BindingModel{
		Rfc:          "CACX7605101P8",
		Name:         "XOCHILT CASAS CHAVEZ",
		CfdiUse:      "D04",
		FiscalRegime: "605",
		TaxZipCode:   "36257",
	}
	request.Items[0].TaxObject = "01"
	request.Items[0].Taxes = nil
	request.Items[0].Total = 100
	request.Complemento = &models.Complementv4{
		Donation: &models.DonationModel{
			AuthorizationDate:   "30/01/2019",
			AuthorizationNumber: "B400-05-08-2014-005",
			Legend:              "El comprobante es un donativo",
		},
	}

	cfdi, err := s.Client.CreateCfdiV4XML(s.Context, request)
	s.Require().NoError(err)
	s.NotEmpty(cfdi.TaxStamp.UUID)
}

func (s *APIClientSuite) TestCreateThirdPartyAccountComplement() {
	request := s.incomeRequest("Test-Terceros")
	request.Items[0].Unit = ""
	request.Items[0].ThirdPartyAccount = &models.ThirdPartyAccountModel{
		Rfc:          "CACX7605101P8",
		Name:         "XOCHILT CASAS CHAVEZ",
		FiscalRegime: "616",
		TaxZipCode:   "36257",
	}

	cfdi, err := s.Client.CreateCfdiV4XML(s.Context, request)
	s.Require().NoError(err)
	s.NotEmpty(cfdi.TaxStamp.UUID)
}
