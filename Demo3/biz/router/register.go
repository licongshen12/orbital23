// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	demo3 "Demo3/biz/router/demo3"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	demo3.Register(r)
}
