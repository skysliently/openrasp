package test_api

import (
//	"net/http"
	"rasp-cloud/controllers"
)

type TestController struct {
	controllers.BaseController
}

// @router /test [get]
func (o *TestController) Test(){
	o.Serve(map[string]interface{}{
		"gugu": "gu",
	})
}