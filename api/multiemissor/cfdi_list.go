package multiemissor

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/vanclief/ez"
	"github.com/vanclief/go-facturama/api/models"
)

// ListCfdisRequest represents the filters to list CFDIs. Zero values are not sent
type ListCfdisRequest struct {
	Status        string // all (default), active or canceled
	ReceiverRFC   string
	IssuerRFC     string
	TaxEntityName string
	Serie         string
	DateStart     time.Time // Sent as the clock time in its own location, pass it in the issuer's local time zone
	DateEnd       time.Time
	Page          int // Pages start at 0 and hold up to 10 CFDIs
}

// Validate validates the request to list CFDIs
func (request *ListCfdisRequest) Validate() error {
	if request.Status != "" && request.Status != "all" && request.Status != "active" && request.Status != "canceled" {
		return ez.New(ez.EINVALID, "Status must be one of: all, active, canceled", nil)
	}

	if request.Page < 0 {
		return ez.New(ez.EINVALID, "Page must be 0 or greater", nil)
	}

	return nil
}

// ListCfdis lists the CFDIs issued through the Multiemissor API, newest first
// Endpoint: GET /cfdi?type=issuedLite
func (c *Client) ListCfdis(ctx context.Context, request ListCfdisRequest) ([]models.CfdiSearchResult, error) {
	err := request.Validate()
	if err != nil {
		return nil, ez.Wrap(err)
	}

	status := request.Status
	if status == "" {
		status = "all"
	}

	query := url.Values{}
	query.Set("type", "issuedLite")
	query.Set("status", status)
	query.Set("page", strconv.Itoa(request.Page))

	optional := map[string]string{
		"rfc":           request.ReceiverRFC,
		"rfcIssuer":     request.IssuerRFC,
		"taxEntityName": request.TaxEntityName,
		"serie":         request.Serie,
	}
	for key, value := range optional {
		if value != "" {
			query.Set(key, value)
		}
	}

	if !request.DateStart.IsZero() {
		query.Set("dateStart", request.DateStart.Format("2006-01-02T15:04:05"))
	}
	if !request.DateEnd.IsZero() {
		query.Set("dateEnd", request.DateEnd.Format("2006-01-02T15:04:05"))
	}

	path := "/cfdi?" + query.Encode()
	var result []models.CfdiSearchResult

	err = c.Get(ctx, path, &result)
	if err != nil {
		return nil, ez.Wrap(err)
	}

	return result, nil
}
