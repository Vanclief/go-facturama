package multiemissor

import (
	"github.com/vanclief/go-facturama/api/models"
)

func (s *APIClientSuite) TestCreateCfdiV4() {
	// Step 1: Create a CFDI para PUBLIC EN GENERAL
	request := CreateCfdiV4Request{
		NameID:          1,
		ExpeditionPlace: "78116", // Valid 5-digit ZIP code
		Folio:           "Test-001",
		CfdiType:        "I",   // Ingreso
		PaymentForm:     "01",  // Efectivo
		PaymentMethod:   "PUE", // Pago en una sola exhibición
		Currency:        "MXN",
		GlobalInformation: &models.GlobalInformationV4Model{
			Periodicity: "01", // Mensual
			Months:      "05", // January
			Year:        2025,
		},
		Issuer: models.IssuerV4BindingModel{
			Name:         "FRANCO VALENCIA ADAN",
			Rfc:          s.RFC, // Test RFC for sandbox
			FiscalRegime: "612", // Persona Física con actividad empresarial
		},
		Receiver: models.ReceiverV4BindingModel{
			Rfc:          "XAXX010101000", // RFC genérico nacional
			Name:         "PUBLICO EN GENERAL",
			CfdiUse:      "S01", // Sin efectos fiscales
			FiscalRegime: "616", // RFC genérico nacional, // Sin obligaciones fiscales
			TaxZipCode:   "78116",
		},
		Items: []models.ItemFullBindingModel{
			{
				ProductCode: "01010101", // Standard code for products and services
				Description: "Test product",
				Unit:        "PIECE",
				UnitCode:    "H87", // Pieza
				UnitPrice:   100.0,
				Quantity:    1.0,
				Subtotal:    100.0,
				Total:       116.0, // Price with tax
				TaxObject:   "02",  // Sí objeto de impuesto
				Taxes: []models.TaxBindingModel{
					{
						Name:        "IVA",
						Base:        100.0,
						Rate:        0.16,
						Total:       16.0,
						IsRetention: false,
						IsQuota:     false,
					},
				},
			},
		},
	}

	cfdi, err := s.Client.CreateCfdiV4(s.Context, request)

	// Validate results
	s.Nil(err, "Error creating CFDI")
	s.NotNil(cfdi, "Expected non-nil result")

	// Verify basic structure of the response
	s.NotEmpty(cfdi.ID, "ID should not be empty")
	s.Equal("ingreso", cfdi.CfdiType, "Expected CFDI type to be 'ingreso'")
	s.Equal("Test-001", cfdi.Folio, "Expected CFDI folio to match request")

	// Step 2: Cancel this new CFDI
	cancelRequest := CancelCfdiRequest{
		ID:     cfdi.ID,
		Motive: "02",
	}

	// Call the method being tested
	response, err := s.Client.CancelCfdi(s.Context, cancelRequest)

	// Validate results
	s.Nil(err, "Error deleting CFDI file")
	s.NotNil(response, "Expected non-nil result")
	s.Equal("Ok", response.Status, "Wrong cancelation response status")
}

func (s *APIClientSuite) TestGetCfdi() {
	// Create the request
	request := GetCfdiByIdRequest{
		ID: "VRHXPSsy-Xx0i0LyHNziJA2",
	}

	// Call the method being tested
	cfdi, err := s.Client.GetCfdiById(s.Context, request)

	// Validate results
	s.Nil(err, "Error getting CFDI by ID")
	s.NotNil(cfdi, "Expected non-nil result")

	// Verify we got the requested CFDI
	s.Equal(request.ID, cfdi.ID, "Expected CFDI ID to match requested ID")

	// Verify basic structure of the response
	s.NotEmpty(cfdi.CfdiType, "CfdiType should not be empty")
	s.NotEmpty(cfdi.Date, "Date should not be empty")
}

func (s *APIClientSuite) TestGetCfdiFile() {
	// Create the request
	request := GetCfdiFileRequest{
		ID:       "VRHXPSsy-Xx0i0LyHNziJA2",
		Format:   "pdf",
		CfdiType: "issuedLite",
	}

	// Call the method being tested
	file, err := s.Client.GetCfdiFile(s.Context, request)

	// Validate results
	s.Nil(err, "Error getting CFDI file?")
	s.NotNil(file, "Expected non-nil result")

	// Verify basic structure of the response
	s.Equal("base64", file.ContentEncoding, "Expected base64 encoding")
	s.Equal("pdf", file.ContentType, "Expected content type to match request format")
	s.Greater(file.ContentLength, 0, "Expected content length > 0")
	s.NotEmpty(file.Content, "Content should not be empty")
}
