package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ApiRequests struct {
	Foo string `json:"sample"`
	Bar []int  `json:"request"`
}

type ApiResponse struct {
	Baz struct {
		Prop string `json:"prop"`
	} `json:"baz"`
}


func ApiHandler(ctx echo.Context) error {
	req := ApiRequests{}
	if err := ctx.Bind(&req); err != nil {
		return echo.ErrBadRequest
	}
	resp := doSthWithRequest(req)
	return ctx.JSON(http.StatusOK, resp)
}

func doSthWithRequest(req ApiRequests) ApiResponse {
	return ApiResponse{}
}