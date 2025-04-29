package multiemissor

import (
	"context"
	"fmt"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// GetCSDByRFCRequest represents a request to get a CSD by RFC
type GetCSDByRFCRequest struct {
	RFC string
}

// Validate validates the request to get a CSD by RFC
func (request *GetCSDByRFCRequest) Validate() error {
	const op = "GetCSDByRFCRequest.Validate"

	if request.RFC == "" {
		return ez.New(op, ez.EINVALID, "RFC is required", nil)
	}

	return nil
}

// GetCSDByRFC retrieves a CSD (Certificado de Sello Digital) by RFC (Registro Federal de Contribuyentes)
// Endpoint: GET /api-lite/csds/{rfc}
func (c *Client) GetCSDByRFC(ctx context.Context, request GetCSDByRFCRequest) (*models.TaxEntityCSD, error) {
	const op = "multiemissor.GetCSDByRFC"

	// Validate request
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	path := fmt.Sprintf("/api-lite/csds/%s", request.RFC)
	response := &models.TaxEntityCSD{}

	err = c.Get(ctx, path, response)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return response, nil
}
