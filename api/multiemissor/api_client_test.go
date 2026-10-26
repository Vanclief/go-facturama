package multiemissor

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"github.com/vanclief/go-facturama/api/common"
	"github.com/vanclief/go-facturama/utils"
)

// APIClientSuite is a test suite for Facturama API client tests
type APIClientSuite struct {
	suite.Suite
	Client             *Client
	Context            context.Context
	Cancel             context.CancelFunc
	RFC                string
	Certificate        string
	PrivateKey         string
	PrivateKeyPassword string
}

// SetupSuite initializes the test suite before running tests
func (s *APIClientSuite) SetupSuite() {
	username := os.Getenv("FACTURAMA_USERNAME")
	password := os.Getenv("FACTURAMA_PASSWORD")
	s.RFC = os.Getenv("FACTURAMA_RFC")
	certificatePath := os.Getenv("FACTURAMA_CERT_PATH")
	privateKeyPath := os.Getenv("FACTURAMA_PK_PATH")
	privateKeyPassword := os.Getenv("FACTURAMA_PK_PASSWORD")

	// Without sandbox credentials only the offline tests run
	if username == "" || password == "" || s.RFC == "" {
		s.T().Skip("Set FACTURAMA_USERNAME, FACTURAMA_PASSWORD and FACTURAMA_RFC to run the sandbox tests")
	}

	// The CSD files are only needed by the CSD tests
	s.PrivateKeyPassword = privateKeyPassword

	var err error

	if certificatePath != "" {
		s.Certificate, err = utils.FileToBase64String(certificatePath)
		if err != nil {
			s.T().Fatalf("Error reading certificate file: %v", err)
		}
	}

	if privateKeyPath != "" {
		s.PrivateKey, err = utils.FileToBase64String(privateKeyPath)
		if err != nil {
			s.T().Fatalf("Error reading private key file: %v", err)
		}
	}

	// Create client with sandbox environment
	s.Client = NewClient(username, password, common.WithEnvironment(common.Sandbox))

	// Create a context with timeout for all tests
	s.Context, s.Cancel = context.WithTimeout(context.Background(), 30*time.Second)
}

// TearDownSuite cleans up after all tests have run
func (s *APIClientSuite) TearDownSuite() {
	s.Cancel()
}

func TestAPIClientSuite(t *testing.T) {
	suite.Run(t, new(APIClientSuite))
}
