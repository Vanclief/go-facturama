package multiemissor

import (
	"context"
	"net/url"
	"strconv"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// GetCfdiStatusRequest represents a request to get the SAT status of a CFDI
type GetCfdiStatusRequest struct {
	UUID        string
	IssuerRFC   string
	ReceiverRFC string
	Total       float64 // Total including taxes, exactly as stamped. SAT does not find the CFDI if it is rounded
}

// Validate validates the request to get the SAT status of a CFDI
func (request *GetCfdiStatusRequest) Validate() error {
	if request.UUID == "" {
		return ez.New(ez.EINVALID, "UUID is required", nil)
	}
	if request.IssuerRFC == "" {
		return ez.New(ez.EINVALID, "IssuerRFC is required", nil)
	}
	if request.ReceiverRFC == "" {
		return ez.New(ez.EINVALID, "ReceiverRFC is required", nil)
	}

	return nil
}

// GetCfdiStatus retrieves the status of a CFDI with the SAT and whether it can
// be canceled. It works for any CFDI, even if it was not issued with Facturama.
// The sandbox cannot reach the SAT, so there it only returns simulated or error responses
// Endpoint: GET /cfdi/status?uuid={uuid}&issuerRfc={issuerRfc}&receiverRfc={receiverRfc}&total={total}
func (c *Client) GetCfdiStatus(ctx context.Context, request GetCfdiStatusRequest) (*models.CfdiStatusModel, error) {
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(err)
	}

	query := url.Values{}
	query.Set("uuid", request.UUID)
	query.Set("issuerRfc", request.IssuerRFC)
	query.Set("receiverRfc", request.ReceiverRFC)
	query.Set("total", strconv.FormatFloat(request.Total, 'f', -1, 64))

	path := "/cfdi/status?" + query.Encode()
	var result models.CfdiStatusModel

	err = c.Get(ctx, path, &result)
	if err != nil {
		return nil, ez.Wrap(err)
	}

	return &result, nil
}
