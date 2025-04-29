package multiemissor

import (
	"context"

	"github.com/vanclief/ez"
)

// CreateCSDRequest represents a request to create a new CSD
type CreateCSDRequest struct {
	RFC                string `json:"Rfc"`
	Certificate        string `json:"Certificate"`
	PrivateKey         string `json:"PrivateKey"`
	PrivateKeyPassword string `json:"PrivateKeyPassword"`
}

func (request *CreateCSDRequest) Validate() error {
	const op = "CreateCSDRequest.Validate"

	// Validate required fields
	if request.RFC == "" {
		return ez.New(op, ez.EINVALID, "RFC is required", nil)
	}
	if request.Certificate == "" {
		return ez.New(op, ez.EINVALID, "Certificate is required", nil)
	}
	if request.PrivateKey == "" {
		return ez.New(op, ez.EINVALID, "PrivateKey is required", nil)
	}
	if request.PrivateKeyPassword == "" {
		return ez.New(op, ez.EINVALID, "PrivateKeyPassword is required", nil)
	}

	return nil
}

// CreateCSD uploads a new CSD (Certificado de Sello Digital)
// Endpoint: POST /api-lite/csds
func (c *Client) CreateCSD(ctx context.Context, request CreateCSDRequest) error {
	const op = "multiemissor.CreateCSD"

	err := request.Validate()
	if err != nil {
		return ez.Wrap(op, err)
	}

	path := "/api-lite/csds"

	// The API doesn't provide any specific response for this endpoint
	err = c.Post(ctx, path, request, nil)
	if err != nil {
		return ez.Wrap(op, err)
	}

	return nil
}
