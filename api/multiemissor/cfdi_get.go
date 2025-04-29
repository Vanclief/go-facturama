package multiemissor

import (
	"context"
	"fmt"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// GetCfdiByIdRequest represents a request to get a CFDI by ID
type GetCfdiByIdRequest struct {
	ID string
}

// Validate validates the request to get a CFDI by ID
func (request *GetCfdiByIdRequest) Validate() error {
	const op = "GetCfdiByIdRequest.Validate"

	// Validate required parameters
	if request.ID == "" {
		return ez.New(op, ez.EINVALID, "CFDI ID is required", nil)
	}

	return nil
}

// GetCfdiById retrieves the details of a CFDI (Mexican digital invoice) by its ID
// Endpoint: GET /api-lite/cfdis/{id}
func (c *Client) GetCfdiById(ctx context.Context, request GetCfdiByIdRequest) (*models.CfdiInfoModel, error) {
	const op = "multiemissor.GetCfdiById"

	// Validate request
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	path := fmt.Sprintf("/api-lite/cfdis/%s", request.ID)
	var result models.CfdiInfoModel

	err = c.Get(ctx, path, &result)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return &result, nil
}
