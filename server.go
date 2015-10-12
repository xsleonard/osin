package osin

import "time"

// Server is an OAuth2 implementation
type Server struct {
	Config                *ServerConfig
	Storage               Storage
	AuthorizeTokenGen     AuthorizeTokenGen
	AccessTokenGen        AccessTokenGen
	Now                   func() time.Time
	AccessRequestHandlers map[AccessRequestType]AccessRequestHandler
}

// NewServer creates a new server instance
func NewServer(config *ServerConfig, storage Storage) *Server {
	s := &Server{
		Config:            config,
		Storage:           storage,
		AuthorizeTokenGen: &AuthorizeTokenGenDefault{},
		AccessTokenGen:    &AccessTokenGenDefault{},
		Now:               time.Now,
		AccessRequestHandlers: make(map[AccessRequestType]AccessRequestHandler),
	}
	s.registerDefaultAccessRequestHandlers()
	return s
}

// NewResponse creates a new response for the server
func (s *Server) NewResponse() *Response {
	r := NewResponse(s.Storage)
	r.ErrorStatusCode = s.Config.ErrorStatusCode
	return r
}
