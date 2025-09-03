package server

import "github.com/labstack/echo/v4"

func (s *Server) RegisterUserRoutes(e *echo.Echo, m ...echo.MiddlewareFunc) {

	group := e.Group("/user", m...)
	// group.GET(":id")            // TODO: add find by id handler
	// group.POST("/confirmation") // TODO: add user confirmation handler
	group.POST("", s.handleCreateUser) // TODO: add create user handler
}
