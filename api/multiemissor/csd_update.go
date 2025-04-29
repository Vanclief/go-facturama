package multiemissor

import (
	"context"
	"fmt"

	"github.com/vanclief/ez"
)

// UpdateCSD updates an existing CSD (Certificado de Sello Digital)
// Endpoint: PUT /api-lite/csds/{rfc}
func (c *Client) UpdateCSD(ctx context.Context, request CreateCSDRequest) error {
	const op = "multiemissor.UpdateCSD"

	err := request.Validate()
	if err != nil {
		return ez.Wrap(op, err)
	}

	path := fmt.Sprintf("/api-lite/csds/%s", request.RFC)

	// The API doesn't provide any specific response for this endpoint
	err = c.Put(ctx, path, request, nil)
	if err != nil {
		return ez.Wrap(op, err)
	}

	return nil
}
