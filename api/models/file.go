package models

// FileViewModel represents a file returned by the Facturama API
type FileViewModel struct {
	ContentEncoding string `json:"ContentEncoding"`
	ContentType     string `json:"ContentType"`
	ContentLength   int    `json:"ContentLength"`
	Content         string `json:"Content"`
}