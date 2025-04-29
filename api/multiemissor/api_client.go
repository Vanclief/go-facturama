package multiemissor

import (
	"github.com/vanclief/go-facturama/api/common"
)

// Client represents a client for the Facturama Multiemissor API
type Client struct {
	*common.Client
}

// NewClient creates a new Multiemissor API client
func NewClient(username, password string, options ...common.Option) *Client {
	return &Client{
		Client: common.NewClient(username, password, options...),
	}
}
