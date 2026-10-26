package multiemissor

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanclief/go-facturama/api/common"
)

// The sandbox cannot reach the SAT, so this runs against a local server
func TestGetCfdiStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/cfdi/status", r.URL.Path)
		assert.Equal(t, "27568D31-7E57-442F-BA77-798CBF30BD7D", r.URL.Query().Get("uuid"))
		assert.Equal(t, "CACX7605101P8", r.URL.Query().Get("issuerRfc"))
		assert.Equal(t, "FUNK671228PH6", r.URL.Query().Get("receiverRfc"))
		assert.Equal(t, "11000.456", r.URL.Query().Get("total"))

		w.Write([]byte(`{"Status":"Vigente","IsCancelable":"Cancelable sin aceptación","Uuid":"27568D31-7E57-442F-BA77-798CBF30BD7D"}`))
	}))
	defer server.Close()

	client := NewClient("user", "pass", common.WithBaseURL(server.URL))

	status, err := client.GetCfdiStatus(context.Background(), GetCfdiStatusRequest{
		UUID:        "27568D31-7E57-442F-BA77-798CBF30BD7D",
		IssuerRFC:   "CACX7605101P8",
		ReceiverRFC: "FUNK671228PH6",
		Total:       11000.456,
	})
	require.NoError(t, err)

	assert.Equal(t, "Vigente", status.Status)
	assert.Equal(t, "Cancelable sin aceptación", status.IsCancelable)
	assert.Equal(t, "27568D31-7E57-442F-BA77-798CBF30BD7D", status.UUID)
}
