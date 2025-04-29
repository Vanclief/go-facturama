package models

import (
	"strings"
	"time"
)

// TaxEntityCSD represents a CSD (Certificado de Sello Digital) entity
type TaxEntityCSD struct {
	RFC                string        `json:"Rfc"`
	Certificate        string        `json:"Certificate"`
	PrivateKey         string        `json:"PrivateKey"`
	PrivateKeyPassword string        `json:"PrivateKeyPassword"`
	CsdExpirationDate  FacturamaTime `json:"CsdExpirationDate,omitempty"`
	UploadDate         FacturamaTime `json:"UploadDate,omitempty"`
}

// FacturamaTime is a wrapper around time.Time that handles the custom date format
type FacturamaTime struct {
	time.Time
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ct *FacturamaTime) UnmarshalJSON(data []byte) error {
	// Remove quotes
	s := strings.Trim(string(data), "\"")
	if s == "null" || s == "" {
		ct.Time = time.Time{}
		return nil
	}

	// Try first format without milliseconds
	t, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		// Try with milliseconds
		t, err = time.Parse("2006-01-02T15:04:05.99", s)
		if err != nil {
			return err
		}
	}

	ct.Time = t
	return nil
}

