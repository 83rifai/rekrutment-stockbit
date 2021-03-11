package presenter

import (
	"net/http"
	"stockbit/cnf"
	"stockbit/module/anagram/model"
	"stockbit/module/anagram/usecase"
	"stockbit/utl"

	"github.com/labstack/echo"
)

type HTTPAnagramHandler struct {
	Config cnf.Config
}

func NewHTTPHandler(conf cnf.Config) *HTTPAnagramHandler {
	return &HTTPAnagramHandler{
		Config: conf,
	}
}

func (h *HTTPAnagramHandler) MountAdmin(group *echo.Group) {
	// config
	group.POST("/anagram", h.AnagramHandle)

}

func (h *HTTPAnagramHandler) AnagramHandle(c echo.Context) error {

	// get body
	body := new(model.Anagram)
	err := c.Bind(body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utl.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "error bad request",
			Error:      err,
			Data:       nil,
		})
	}

	resp, err := usecase.UseCaseAnagram(body)

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
		Data:       resp,
	})
}
