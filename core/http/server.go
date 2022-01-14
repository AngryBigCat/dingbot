package http

import (
	"ding/core"
	"ding/core/http/routers"
)

func Run() {
	router := routers.Register()

	router.Run(core.HTTP_HOST + ":" + core.HTTP_PORT)
}
