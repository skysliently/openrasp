package test_api

import (
	"net/http"
	"rasp-cloud/controllers"
	"rasp-cloud/conf"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"rasp-cloud/models"
	"rasp-cloud/models/waf"
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

// @router /testMgoFunc [get]
func (o *TestController) TestMgoFunc()  {
	err := models.TestMgo()
	if err != nil {
		o.ServeError(http.StatusBadRequest, err.Error())
	}
	o.Serve(map[string]interface{}{
		"testresult": "success",
	})
	o.ServeWithEmptyData()
}

// @router /testWafFunc [get]
func (o *TestController) TestWafFunc()  {
	total, wafs, err := waf.GetAllWafIdSort()
	if err != nil {
		// o.ServeError(http.StatusBadRequest, err.Error())
		o.ServeError(http.StatusBadRequest, "failed to get apps", err)
	}
	if wafs == nil {
		wafs = make([]*waf.Waf, 0)
		// o.ServeError(http.StatusBadRequest, "wafs is nil")/
		var result = make(map[string]interface{})
		result["total"] = total
		result["wafs"] = wafs
		result["wafs nil"] = "wafs nil"
		o.Serve(result)
	}else {
		var result = make(map[string]interface{})
		result["total"] = total
		result["wafs"] = wafs
		o.Serve(result)
	}
	
}

// @router /getWafStatus [post]
func (o *TestController) GetWafStatus()  {
	var param struct {
		Id string `json:"id"`
	}
	o.UnmarshalJson(&param)
	if param.Id == ""{
		o.ServeError(http.StatusBadRequest, "Param can not be empty")
	}
	waf, err := waf.GetWafById(param.Id)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to get waf", err)
	}
	resp, err := http.Get(waf.Addr+"/api/stat")
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to Request",err)
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

// @router /configNewWaf [post]
func (o *TestController) ConfigNewWaf()  {
	var param struct {
		WafName string `json:"wafName"`
		WafAddr string `json:"wafAddr"`
	}
	o.UnmarshalJson(&param)
	if param.WafName == "" || param.WafName == "" {
		o.ServeError(http.StatusBadRequest, "WafName and WafAddr can not be empty")
	}
	resp, err := http.Get(param.WafAddr+"/api/stat")
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to Request",err)
	}
	if resp.StatusCode != 200 {
		o.ServeError(http.StatusBadRequest, "Response code not 200")
	}
	waf, err := waf.ConfigNewWaf(param.WafName,param.WafAddr)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to ConfigNewWaf", err)
	}
	var result = make(map[string]interface{})
	result["waf"] = waf
	o.Serve(result)
	// o.Ctx.WriteString("test")
}