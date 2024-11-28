package server

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Post("/verify-auth-key", s.VerifyAuthKeyHandler)
	s.App.Post("/regenerate-sequence/:id", s.RegenerateSequence)
	s.App.Post("/validate-sequence/:id", s.ValidateSequence)

	s.App.Get("/existing-gotchi/:id", s.ExistingGotchiHandler)
}
