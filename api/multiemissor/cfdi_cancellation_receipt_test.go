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

func TestGetCancellationReceipt(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The ID is escaped so it can't change the path
		assert.Equal(t, "/acuse/html/issuedLite/abc%2F..%3Fx", r.URL.EscapedPath())
		assert.Empty(t, r.URL.RawQuery)

		w.Write([]byte(`{"ContentEncoding":"base64","ContentType":"html","ContentLength":4,"Content":"PGI+"}`))
	}))
	defer server.Close()

	client := NewClient("user", "pass", common.WithBaseURL(server.URL))

	receipt, err := client.GetCancellationReceipt(context.Background(), GetCancellationReceiptRequest{
		ID:     "abc/..?x",
		Format: "HTML",
	})
	require.NoError(t, err)

	assert.Equal(t, "html", receipt.ContentType)
	assert.Equal(t, "PGI+", receipt.Content)
}
