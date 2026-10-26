package common

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanclief/ez"
)

func requestError(t *testing.T, status int, body string) error {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	defer server.Close()

	client := NewClient("user", "pass", WithBaseURL(server.URL))

	return client.Get(context.Background(), "/", nil)
}

func asAPIError(t *testing.T, err error) *APIError {
	var apiErr *APIError
	require.True(t, errors.As(err, &apiErr), "expected an *APIError in the chain, got %v", err)

	return apiErr
}

func TestAPIErrorIncludesModelStateMessages(t *testing.T) {
	// Body captured from the sandbox when creating a CFDI with an unknown product code
	body := `{"Message":"La solicitud no es válida.","ModelState":{"cfdiToCreate.items[0].ProductCode":["El campo clave del producto o del servicio no se encuentra en los catalogos."],"cfdiToCreate.Receiver.TaxZipCode":["El código postal no es válido.","Debe tener 5 dígitos."]}}`

	err := requestError(t, http.StatusBadRequest, body)

	// Apps show ez.ErrorMessage to users, so SAT's reason has to reach it
	assert.Equal(t, ez.EINVALID, ez.ErrorCode(err))
	assert.Equal(t,
		"Facturama API error 400: La solicitud no es válida."+
			" cfdiToCreate.Receiver.TaxZipCode: El código postal no es válido. Debe tener 5 dígitos."+
			" cfdiToCreate.items[0].ProductCode: El campo clave del producto o del servicio no se encuentra en los catalogos.",
		ez.ErrorMessage(err))

	apiErr := asAPIError(t, err)
	assert.Equal(t, http.StatusBadRequest, apiErr.StatusCode)
	assert.Equal(t, []string{"El campo clave del producto o del servicio no se encuentra en los catalogos."}, apiErr.ModelState["cfdiToCreate.items[0].ProductCode"])
}

func TestAPIErrorWithoutModelState(t *testing.T) {
	err := requestError(t, http.StatusInternalServerError, `{"Message":"Sucedió un error interno y no fue posible verificar el estado del CFDi"}`)

	assert.Equal(t, ez.EINTERNAL, ez.ErrorCode(err))
	assert.Equal(t, "Facturama API error 500: Sucedió un error interno y no fue posible verificar el estado del CFDi", ez.ErrorMessage(err))
}

func TestAPIErrorFallsBackToRawBody(t *testing.T) {
	err := requestError(t, http.StatusNotFound, "<html>Not Found</html>")

	// Users only see the status, the body stays in the error text for logs
	assert.Equal(t, ez.ENOTFOUND, ez.ErrorCode(err))
	assert.Equal(t, "Facturama API error 404", ez.ErrorMessage(err))
	assert.Equal(t, "Facturama API error 404: <html>Not Found</html>", asAPIError(t, err).Error())
	assert.Equal(t, "<html>Not Found</html>", asAPIError(t, err).RawBody)
}

func TestAPIErrorCapsLongRawBody(t *testing.T) {
	page := "<html>" + strings.Repeat("ñ", 1000) + "</html>"

	err := requestError(t, http.StatusBadGateway, page)

	assert.Equal(t, "Facturama API error 502: <html>"+strings.Repeat("ñ", maxRawBodyInError-6)+"…", asAPIError(t, err).Error())
	assert.Equal(t, page, asAPIError(t, err).RawBody)
}

func TestAPIErrorReportsFacturamaAuthStatus(t *testing.T) {
	// Facturama answers wrong credentials with an empty 401. The code says what
	// Facturama answered, callers decide what it means for their own users
	err := requestError(t, http.StatusUnauthorized, "")

	assert.Equal(t, ez.ENOTAUTHENTICATED, ez.ErrorCode(err))
	assert.Equal(t, "Facturama API error 401", ez.ErrorMessage(err))
	assert.Equal(t, http.StatusUnauthorized, asAPIError(t, err).StatusCode)

	err = requestError(t, http.StatusForbidden, `{"Message":"Se ha denegado la autorización para esta solicitud."}`)

	assert.Equal(t, ez.ENOTAUTHORIZED, ez.ErrorCode(err))
	assert.Equal(t, http.StatusForbidden, asAPIError(t, err).StatusCode)
}

func TestAPIErrorSkipsEmptyParts(t *testing.T) {
	err := requestError(t, http.StatusBadRequest, `{"ModelState":{"":["El cuerpo de la solicitud no es válido."],"Rfc":["El RFC no es válido."]}}`)

	assert.Equal(t, "Facturama API error 400: El cuerpo de la solicitud no es válido. Rfc: El RFC no es válido.", ez.ErrorMessage(err))

	err = requestError(t, http.StatusBadRequest, `{"ModelState":{"Rfc":[],"":null}}`)

	assert.Equal(t, "Facturama API error 400", ez.ErrorMessage(err))

	err = requestError(t, http.StatusUnauthorized, "")

	assert.Equal(t, "Facturama API error 401", asAPIError(t, err).Error())
}
