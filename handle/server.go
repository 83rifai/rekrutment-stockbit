package handle

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stockbit/cnf/env"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	echoMid "github.com/labstack/echo/middleware"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http
func (s *Service) HTTPServerMain() *echo.Echo {
	e := echo.New()

	e.Use(echoMid.Recover())
	e.Use(echoMid.CORS())

	e.Validator = &CustomValidator{V: validator.New()}
	custErr := &customErrHandler{e: e}
	e.HTTPErrorHandler = custErr.handler
	e.Binder = &CustomBinder{b: &echo.DefaultBinder{}}

	// external endpoints
	api := e.Group("/api")
	s.AnagramHandle.MountAdmin(api)
	s.ImdbHandle.MountAdmin(api)

	return e
}

func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", env.Conf.HttpPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

// NewBinder initializes custom server binder
func NewBinder() *CustomBinder {
	return &CustomBinder{b: &echo.DefaultBinder{}}
}

// CustomBinder struct
type CustomBinder struct {
	b echo.Binder
}

// Bind tries to bind request into interface, and if it does then validate it
func (cb *CustomBinder) Bind(i interface{}, c echo.Context) error {
	if err := cb.b.Bind(i, c); err != nil && err != echo.ErrUnsupportedMediaType {
		return err
	}
	return c.Validate(i)
}

// CustomValidator holds custom validator
type CustomValidator struct {
	V *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.V.Struct(i)
}

type customErrHandler struct {
	e *echo.Echo
}

func (ce *customErrHandler) handler(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  interface{}
	)

	type resp struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message interface{} `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}

	switch e := err.(type) {
	case *echo.HTTPError:
		code = e.Code
		msg = e.Message
		if e.Internal != nil {
			msg = fmt.Sprintf("%v, %v", err, e.Internal)
		}
	default:
		msg = http.StatusText(code)
	}
	if _, ok := msg.(string); ok {
		msg = resp{Message: msg}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == "HEAD" {
			err = c.NoContent(code)
		} else {
			err = c.JSON(code, msg)
		}
		if err != nil {
			ce.e.Logger.Error(err)
		}
	}
}
