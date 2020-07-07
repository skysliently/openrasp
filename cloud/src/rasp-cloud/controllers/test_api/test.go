package test_api

import (
	"net/http"
	"rasp-cloud/controllers"
	"rasp-cloud/conf"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"rasp-cloud/models"
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


// @router /testFun [post]
func (o *TestController) TestFun()  {
	// var query map[string]interface{}
	// o.UnmarshalJson(&query)
	// testString = 

	var param struct {
		TestPostParam1 string `json:"testPostParam1"`
		TestPostParam2 string `json:"testPostParam2"`
		TestPostParam3 string `json:"testPostParam3"`
	}
	o.UnmarshalJson(&param)
	if param.TestPostParam1 == "" {
		o.ServeError(http.StatusBadRequest, "testPostParam1 can not be empty")
	}
	if param.TestPostParam2 == "" {
		o.ServeError(http.StatusBadRequest, "testPostParam2 can not be empty")
	}
	if param.TestPostParam3 == "" {
		o.ServeError(http.StatusBadRequest, "testPostParam3 can not be empty")
	}
	if param.TestPostParam3 =="TestPostParam3" {
		param.TestPostParam3 = "testchange"
	}
	o.Serve(map[string]interface{}{
		"testPostParam1": param.TestPostParam1,
		"testPostParam2": param.TestPostParam2,
		"testPostParam3": param.TestPostParam3,
	})
}

// @router /openWafConfig [get]
func (o *TestController) OpenWafConfig()  {
	//conf.AppConfig.OwAddr
	// o.Serve(map[string]interface{}{
	// 	"gugu": conf.AppConfigWaf.OwAddr,
	// })
	// var result = make(map[string]interface{})
	// result["total"] = total
	// result["total_page"] = math.Ceil(float64(total) / float64(param.Perpage))
	// result["page"] = param.Page
	// result["perpage"] = param.Perpage
	// result["data"] = records
	// o.Serve(result)
	var result = make(map[string]interface{})
	for i,r := range  conf.AppConfigWaf.OwAddr{
		result["url"+strconv.Itoa(i)] = r
	}
	o.Serve(result)
}

// @router /urlTest [get]
func (o *TestController) UrlTest()  {
	resp, err := http.Get("http://172.20.27.60:61111/api/stat")
    if err != nil {
        // fmt.Println(err)
		// return
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		// fmt.Println("ok")
		o.Ctx.WriteString(string(body))
	}else {
		o.ServeError(http.StatusBadRequest, strconv.Itoa(resp.StatusCode))
	}
}

// @router /urlJsonTest [post]
func (o *TestController) UrlJsonTest()  {
	var param struct {
		TestPostParam1 string `json:"testPostParam1"`
		TestPostParam2 string `json:"testPostParam2"`
		TestPostParam3 string `json:"testPostParam3"`
	}
	o.UnmarshalJson(&param)
	if param.TestPostParam1 == "" || param.TestPostParam2 == "" || param.TestPostParam3 == "" {
		o.ServeError(http.StatusBadRequest, "Param can not be empty e.g:{			'testPostParam1':'1',			'testPostParam2':'2', 			'testPostParam3':'3'}")
	}
	resp, err := http.Get("http://172.20.27.60:61111/api/stat")
    if err != nil {
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		var res *models.TestJson
		json.Unmarshal(body,&res)
		var result = make(map[string]interface{})
		result["res"] = res
		o.Serve(result)
	}else {
		o.ServeError(http.StatusBadRequest, strconv.Itoa(resp.StatusCode))
	}
	
}