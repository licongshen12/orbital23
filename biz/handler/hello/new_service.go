// Code generated by hertz generator.

package hello

import (
	"context"

	hello "demosiu/biz/model/hello"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// NewMethod .
// @router /new [GET]
func NewMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(hello.HelloResp)

	c.JSON(consts.StatusOK, resp)
}
