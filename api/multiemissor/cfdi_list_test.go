package multiemissor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanclief/go-facturama/api/common"
)

// listQuery returns the query ListCfdis sends for the request
func listQuery(t *testing.T, request ListCfdisRequest) url.Values {
	var query url.Values
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/cfdi", r.URL.Path)
		query = r.URL.Query()
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	client := NewClient("user", "pass", common.WithBaseURL(server.URL))

	_, err := client.ListCfdis(context.Background(), request)
	require.NoError(t, err)

	return query
}

func TestListCfdisSendsOnlyRequiredParamsByDefault(t *testing.T) {
	query := listQuery(t, ListCfdisRequest{})

	assert.Equal(t, url.Values{
		"type":   {"issuedLite"},
		"status": {"all"},
		"page":   {"0"},
	}, query)
}

func TestListCfdisSendsFilters(t *testing.T) {
	mexicoCity := time.FixedZone("CST", -6*60*60)

	query := listQuery(t, ListCfdisRequest{
		Status:        "canceled",
		ReceiverRFC:   "URE180429TM6",
		IssuerRFC:     "EKU9003173C9",
		TaxEntityName: "UNIVERSIDAD",
		Serie:         "FA",
		// Same-day interval, the time of day must survive
		DateStart: time.Date(2026, 10, 26, 14, 0, 0, 0, mexicoCity),
		DateEnd:   time.Date(2026, 10, 26, 16, 30, 59, 0, mexicoCity),
		Page:      2,
	})

	assert.Equal(t, url.Values{
		"type":          {"issuedLite"},
		"status":        {"canceled"},
		"page":          {"2"},
		"rfc":           {"URE180429TM6"},
		"rfcIssuer":     {"EKU9003173C9"},
		"taxEntityName": {"UNIVERSIDAD"},
		"serie":         {"FA"},
		"dateStart":     {"2026-10-26T14:00:00"},
		"dateEnd":       {"2026-10-26T16:30:59"},
	}, query)
}
