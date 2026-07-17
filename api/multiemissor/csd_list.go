package multiemissor

import (
	"context"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// ListCSDs retrieves all CSDs (Certificados de Sello Digital) available
// Endpoint: GET /api-lite/csds
func (c *Client) ListCSDs(ctx context.Context) ([]models.TaxEntityCSD, error) {
	path := "/api-lite/csds"
	var response []models.TaxEntityCSD

	err := c.Get(ctx, path, &response)
	if err != nil {
		return nil, ez.Wrap(err)
	}

	return response, nil
}
