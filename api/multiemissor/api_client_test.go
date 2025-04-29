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

	// Require credentials to be set
	if username == "" {
		s.T().Fatal("FACTURAMA_USERNAME environment variable must be set to run tests")
	}

	if password == "" {
		s.T().Fatal("FACTURAMA_PASSWORD environment variable must be set to run tests")
	}

	if s.RFC == "" {
		s.T().Fatal("FACTURAMA_RFC environment variable must be set to run tests")
	}

	if privateKeyPath == "" {
		s.FailNow("FACTURAMA_PK_PATH is not set")
	}

	if certificatePath == "" {
		s.FailNow("FACTURAMA_CERT_PATH is not set")
	}

	if privateKeyPassword == "" {
		s.FailNow("FACTURAMA_PK_PASSWORD is not set")
	}

	s.PrivateKeyPassword = privateKeyPassword

	var err error

	s.Certificate, err = utils.FileToBase64String(certificatePath)
	if err != nil {
		s.T().Fatalf("Error reading certificate file: %v", err)
	}

	s.PrivateKey, err = utils.FileToBase64String(privateKeyPath)
	if err != nil {
		s.T().Fatalf("Error reading private key file: %v", err)
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
