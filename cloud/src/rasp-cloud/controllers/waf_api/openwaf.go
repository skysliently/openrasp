package waf_api

import (
	"net/http"
	"rasp-cloud/controllers"
	// "rasp-cloud/conf"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"rasp-cloud/models"
	"rasp-cloud/models/waf"
)

type WafController struct {
	controllers.BaseController
}

const (
	StatusAddr = "/api/stat"
)

// @router /getWafStatusByIdRefresh [post]
func (o *WafController) GetWafStatusByIdRefresh()  {
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
	resp, err := http.Get(waf.Addr+StatusAddr)
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
func (o *WafController) ConfigNewWaf()  {
	var param struct {
		WafName string `json:"wafName"`
		WafAddr string `json:"wafAddr"`
	}
	o.UnmarshalJson(&param)
	if param.WafName == "" || param.WafName == "" {
		o.ServeError(http.StatusBadRequest, "WafName and WafAddr can not be empty")
	}
	resp, err := http.Get(param.WafAddr+StatusAddr)
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

// @router /configNewWafWithStatus [post]
func (o *WafController) ConfigNewWafWithStatus()  {
	var param struct {
		WafName string `json:"wafName"`
		WafAddr string `json:"wafAddr"`
	}
	o.UnmarshalJson(&param)
	if param.WafName == "" || param.WafName == "" {
		o.ServeError(http.StatusBadRequest, "WafName and WafAddr can not be empty")
	}
	resp, err := http.Get(param.WafAddr+StatusAddr)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to Request",err)
	}
	if resp.StatusCode != 200 {
		o.ServeError(http.StatusBadRequest, "Response code not 200")
	}
	body, err := ioutil.ReadAll(resp.Body)
	var wafStatusResp *waf.WafStatusResp
	json.Unmarshal(body,&wafStatusResp)
	waf, err := waf.ConfigNewWafWithStatus(param.WafName,param.WafAddr,wafStatusResp)
	// waf, err := waf.ConfigNewWaf(param.WafName,param.WafAddr)
	if err != nil {
		o.ServeError(http.StatusBadRequest, "failed to ConfigNewWaf", err)
	}
	var result = make(map[string]interface{})
	result["waf"] = waf
	o.Serve(result)
	// o.Ctx.WriteString("test")
}