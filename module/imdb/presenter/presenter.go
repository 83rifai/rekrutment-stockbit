package presenter

import (
	"errors"
	"net/http"
	"stockbit/cnf"
	"stockbit/module/imdb/usecase"
	"stockbit/utl"

	"github.com/labstack/echo"
)

type HTTPImdbHandler struct {
	Config cnf.Config
}

func NewHTTPHandler(conf cnf.Config) *HTTPImdbHandler {
	return &HTTPImdbHandler{
		Config: conf,
	}
}

func (h *HTTPImdbHandler) MountAdmin(group *echo.Group) {
	// config
	group.GET("/imdb/list", h.GetList)
	group.GET("/imdb/detail/:id", h.GetDetail)

}

func (h *HTTPImdbHandler) GetList(c echo.Context) error {

	return c.JSON(http.StatusOK, utl.Response{
		StatusCode: http.StatusOK,
		Message:    nil,
		Data:       c.QueryParams(),
	})
}

func (h *HTTPImdbHandler) GetDetail(c echo.Context) error {
	if c.Param("id") == "" {
		return c.JSON(http.StatusBadRequest, utl.Response{
			StatusCode: http.StatusBadRequest,
			Message:    nil,
			Error:      errors.New("error body request"),
			Data:       nil,
		})
	}

	detail, err := usecase.GetDetail(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, utl.Response{
			StatusCode: http.StatusNotFound,
			Message:    nil,
			Error:      err,
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, utl.Response{
		StatusCode: http.StatusOK,
		Message:    nil,
		Data:       detail,
		Param:      c.QueryParams(),
	})
}
