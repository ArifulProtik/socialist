package main

func (s *Server) Routes() {
	s.e.GET("/", s.c.Signup)
	s.e.GET("/signin", s.c.SignInShow)
}
