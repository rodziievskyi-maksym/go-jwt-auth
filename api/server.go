package api

import "github.com/labstack/echo/v4"

type HTTPServer struct {
	cfg string
}

func (s *HTTPServer) Start() error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {

		return c.String(200, "Hello, World!")
	})

	address := s.cfg
	return e.Start(address)
}

func NewHTTPServer(cfg string) *HTTPServer {
	return &HTTPServer{cfg: cfg}
}
