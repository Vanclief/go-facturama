package multiemissor

import (
	"context"
	"fmt"
	"strings"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// GetCfdiFileRequest represents a request to get a CFDI file
type GetCfdiFileRequest struct {
	Format   string
	CfdiType string
	ID       string
}

// Validate validates the request to get a CFDI file
func (request *GetCfdiFileRequest) Validate() error {
	// Validate required parameters
	if request.ID == "" {
		return ez.New(ez.EINVALID, "CFDI ID is required", nil)
	}

	// Validate format parameter
	request.Format = strings.ToLower(request.Format)
	if request.Format != "pdf" && request.Format != "html" && request.Format != "xml" {
		return ez.New(ez.EINVALID, "Format must be one of: pdf, html, xml", nil)
	}

	// Validate cfdiType parameter
	if request.CfdiType != "payroll" && request.CfdiType != "received" && request.CfdiType != "issued" && request.CfdiType != "issuedLite" {
		return ez.New(ez.EINVALID, "Type must be one of: payroll, received, issued, issuedLite", nil)
	}

	return nil
}

// GetCfdiFile retrieves a CFDI file in the specified format
// Endpoint: GET /cfdi/{format}/{type}/{id}
func (c *Client) GetCfdiFile(ctx context.Context, request GetCfdiFileRequest) (*models.FileViewModel, error) {
	// Validate request
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(err)
	}

	path := fmt.Sprintf("/cfdi/%s/%s/%s", request.Format, request.CfdiType, request.ID)
	var result models.FileViewModel

	err = c.Get(ctx, path, &result)
	if err != nil {
		return nil, ez.Wrap(err)
	}

	return &result, nil
}
