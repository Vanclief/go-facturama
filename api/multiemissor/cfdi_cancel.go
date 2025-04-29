package multiemissor

import (
	"context"
	"fmt"
	"net/http"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// CancelCfdiRequest represents the request to cancel a CFDI
type CancelCfdiRequest struct {
	ID              string
	Motive          string
	UUIDReplacement string
}

// Validate validates the request to cancel a CFDI
func (request *CancelCfdiRequest) Validate() error {
	const op = "CancelCfdiRequest.Validate"

	// Validate required parameters
	if request.ID == "" {
		return ez.New(op, ez.EINVALID, "CFDI ID is required", nil)
	}

	// Validate motive parameter (if provided)
	if request.Motive != "" {
		validMotives := []string{"01", "02", "03", "04"}
		isValid := false
		for _, v := range validMotives {
			if request.Motive == v {
				isValid = true
				break
			}
		}
		if !isValid {
			return ez.New(op, ez.EINVALID, "Motive must be one of: 01, 02, 03, 04", nil)
		}
	}

	// Validate uuidReplacement if motive is "01" (CFDI issued with errors with relation)
	if request.Motive == "01" && request.UUIDReplacement == "" {
		return ez.New(op, ez.EINVALID, "UUID replacement is required when motive is 01", nil)
	}

	return nil
}

// CancelCfdi cancels a CFDI (Version 2018)
// Endpoint: DELETE /api-lite/cfdis/{id}?motive={motive}&uuidReplacement={uuidReplacement}
func (c *Client) CancelCfdi(ctx context.Context, request CancelCfdiRequest) (*models.CancelationStatusLite, error) {
	const op = "multiemissor.CancelCfdi"

	// Validate request
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	// Build the URL with query parameters
	path := fmt.Sprintf("/api-lite/cfdis/%s", request.ID)
	if request.Motive != "" {
		path = fmt.Sprintf("%s?motive=%s", path, request.Motive)
		if request.UUIDReplacement != "" {
			path = fmt.Sprintf("%s&uuidReplacement=%s", path, request.UUIDReplacement)
		}
	}

	var result models.CancelationStatusLite

	// Use Request method with DELETE HTTP method
	err = c.Request(ctx, http.MethodDelete, path, nil, &result)
	if err != nil {
		return nil, ez.Wrap(op, err)
	}

	return &result, nil
}
