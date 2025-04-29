package multiemissor

// TODO: Uncomment once everything is set

// func (s *APIClientSuite) TestCSDCreationAndDeletion() {
// 	// Step 1: Check if we already have a CSD for this RFC
// 	csd, _ := s.Client.GetCSDByRFC(s.Context, GetCSDByRFCRequest{
// 		RFC: s.RFC,
// 	})
//
// 	// Step 2: If we already have a CSD for this RFC, delete it
// 	if csd != nil {
// 		// If we have a CSD, delete it
// 		err := s.Client.DeleteCSD(s.Context, DeleteCSDRequest{
// 			RFC: s.RFC,
// 		})
//
// 		s.Nil(err, "Should not have error deleting CSD")
// 	}
//
// 	// Step 3: Create a new CSD
// 	err := s.Client.CreateCSD(s.Context, CreateCSDRequest{
// 		RFC:                s.RFC,
// 		Certificate:        s.Certificate,
// 		PrivateKey:         s.PrivateKey,
// 		PrivateKeyPassword: s.PrivateKeyPassword,
// 	})
// 	s.Nil(err, "Should not have error creating CSD")
//
// 	// Step 4: We should be able to get the CSD we just created
// 	csd, err = s.Client.GetCSDByRFC(s.Context, GetCSDByRFCRequest{
// 		RFC: s.RFC,
// 	})
// 	s.Nil(err, "Should not have error getting CSD")
// 	s.Equal(s.RFC, csd.RFC, "Should have the same RFC")
// 	s.Equal(s.Certificate, csd.Certificate, "Should have the same Certificate")
// 	s.Equal(s.PrivateKey, csd.PrivateKey, "Should have the same PrivateKey")
// 	s.Equal(s.PrivateKeyPassword, csd.PrivateKeyPassword, "Should have the same PrivateKeyPassword")
//
// 	// Step 5: We should be able to update the CSD
// 	err = s.Client.UpdateCSD(s.Context, CreateCSDRequest{
// 		RFC:                s.RFC,
// 		Certificate:        s.Certificate,
// 		PrivateKey:         s.PrivateKey,
// 		PrivateKeyPassword: s.PrivateKeyPassword,
// 	})
// 	s.Nil(err, "Should not have error updating the CSD")
//
// 	// Step 5: We should be able to get a list with the CSD we just created as well
// 	csds, err := s.Client.ListCSDs(s.Context)
// 	s.Nil(err, "Should not have error getting CSD")
// 	s.Len(csds, 1, "Should have one CSD")
// 	s.Equal(s.RFC, csds[0].RFC, "Should have the same RFC")
//
// 	// Step 5: Delete the CSD
// 	err = s.Client.DeleteCSD(s.Context, DeleteCSDRequest{
// 		RFC: s.RFC,
// 	})
//
// 	s.Nil(err, "Should not have error deleting CSD")
// }
//
// func (s *APIClientSuite) TestCSDCreateValidation() {
// 	// Test missing RFC
// 	err := s.Client.CreateCSD(s.Context, CreateCSDRequest{
// 		Certificate:        "cert-data",
// 		PrivateKey:         "key-data",
// 		PrivateKeyPassword: "password",
// 	})
// 	s.Error(err, "Should error with missing RFC")
// 	s.Contains(err.Error(), "RFC is required")
//
// 	// Test missing Certificate
// 	err = s.Client.CreateCSD(s.Context, CreateCSDRequest{
// 		RFC:                "TEST010101AAA",
// 		PrivateKey:         "key-data",
// 		PrivateKeyPassword: "password",
// 	})
// 	s.Error(err, "Should error with missing Certificate")
// 	s.Contains(err.Error(), "Certificate is required")
//
// 	// Test missing PrivateKey
// 	err = s.Client.CreateCSD(s.Context, CreateCSDRequest{
// 		RFC:                "TEST010101AAA",
// 		Certificate:        "cert-data",
// 		PrivateKeyPassword: "password",
// 	})
// 	s.Error(err, "Should error with missing PrivateKey")
// 	s.Contains(err.Error(), "PrivateKey is required")
//
// 	// Test missing PrivateKeyPassword
// 	err = s.Client.CreateCSD(s.Context, CreateCSDRequest{
// 		RFC:         "TEST010101AAA",
// 		Certificate: "cert-data",
// 		PrivateKey:  "key-data",
// 	})
// 	s.Error(err, "Should error with missing PrivateKeyPassword")
// 	s.Contains(err.Error(), "PrivateKeyPassword is required")
// }
//
// func (s *APIClientSuite) TestCSDGetValidation() {
// 	// Test missing RFC
// 	csd, err := s.Client.GetCSDByRFC(s.Context, GetCSDByRFCRequest{})
// 	s.Error(err, "Should error with missing RFC")
// 	s.Nil(csd, "CSD should be nil")
// 	s.Contains(err.Error(), "RFC is required")
//
// 	// Test random RFC
// 	csd, err = s.Client.GetCSDByRFC(s.Context, GetCSDByRFCRequest{RFC: "NOSOYUNRFC"})
// 	s.Error(err, "Should error with random RFC")
// 	s.Nil(csd, "CSD should be nil")
// }
