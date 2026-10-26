package multiemissor

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// GetCancellationReceiptRequest represents a request to get the cancellation receipt (acuse) of a CFDI
type GetCancellationReceiptRequest struct {
	ID     string
	Format string // pdf or html
}

// Validate validates the request to get a cancellation receipt
func (request *GetCancellationReceiptRequest) Validate() error {
	if request.ID == "" {
		return ez.New(ez.EINVALID, "CFDI ID is required", nil)
	}

	request.Format = strings.ToLower(request.Format)
	if request.Format != "pdf" && request.Format != "html" {
		return ez.New(ez.EINVALID, "Format must be one of: pdf, html", nil)
	}

	return nil
}

// GetCancellationReceipt retrieves the cancellation receipt (acuse) of a CFDI
// The XML acuse comes in CancelationStatusLite.AcuseXmlBase64 when canceling
// Endpoint: GET /acuse/{format}/issuedLite/{id}
func (c *Client) GetCancellationReceipt(ctx context.Context, request GetCancellationReceiptRequest) (*models.FileViewModel, error) {
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(err)
	}

	path := fmt.Sprintf("/acuse/%s/issuedLite/%s", request.Format, url.PathEscape(request.ID))
	var result models.FileViewModel

	err = c.Get(ctx, path, &result)
	if err != nil {
		return nil, ez.Wrap(err)
	}

	return &result, nil
}
