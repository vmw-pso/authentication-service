package main

func (s *server) routes() {
	s.mux.Post("/signup", s.handleSignup())
	s.mux.Post("/signin", s.handleSignin())
}
