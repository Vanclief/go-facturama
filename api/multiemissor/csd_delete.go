package multiemissor

import (
	"context"
	"fmt"

	"github.com/vanclief/ez"
)

// DeleteCSDRequest represents a request to delete a CSD by RFC
type DeleteCSDRequest struct {
	RFC string
}

// Validate validates the request to delete a CSD
func (request *DeleteCSDRequest) Validate() error {
	// Validate RFC parameter
	if request.RFC == "" {
		return ez.New(ez.EINVALID, "RFC is required", nil)
	}

	return nil
}

// DeleteCSD deletes a CSD (Certificado de Sello Digital) by RFC
// Endpoint: DELETE /api-lite/csds/{rfc}
func (c *Client) DeleteCSD(ctx context.Context, request DeleteCSDRequest) error {
	// Validate request
	err := request.Validate()
	if err != nil {
		return ez.Wrap(err)
	}

	path := fmt.Sprintf("/api-lite/csds/%s", request.RFC)

	// The API doesn't provide any specific response for this endpoint
	err = c.Delete(ctx, path)
	if err != nil {
		return ez.Wrap(err)
	}

	return nil
}
