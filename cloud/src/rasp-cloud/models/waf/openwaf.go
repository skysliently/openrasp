package waf

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"rasp-cloud/mongo"
	"rasp-cloud/tools"
	"github.com/astaxie/beego"
	"errors"
	"strconv"
	"time"
	"fmt"
	"crypto/sha1"
	"math/rand"
)

type Waf struct {
	Id               string                 `json:"id" bson:"_id"`
	Addr             string                 `json:"addr" bson:"addr"`
	Name             string                 `json:"name" bson:"name"`
	CreateTime       int64                  `json:"create_time"  bson:"create_time"`
	WafStatusResp    WafStatusResp          `json:"wafStatusResp"  bson:"waf_status_resp"`
}

const (
	wafCollectionName = "waf"
	// statusAddr = "/api/stat"
)

type WafStatusResp struct {
	Success          int                    `json:"success" bson:"success"`
	Result           Result                 `json:"result" bson:"result"`
}

type Result struct {
	Main             Mian                   `json:"mian" bson:"mian"`
}

type Mian struct {
	Resetsec         int                    `json:"resetsec" bson:"resetsec"`
	Safe             Safe                   `json:"safe" bson:"safe"`
	Access           Access                 `json:"access" bson:"access"`
	Nginx_version    string                 `json:"nginx_version" bson:"nginx_version"`
	Connection       Connection             `json:"connection" bson:"connection"`
	Upstream         Upstream               `json:"upstream" bson:"upstream"`
}

type Safe struct {
	// Main Safe `json:"mian" bson:""`
}

type Access struct {
	Bytes_out        int                    `json:"bytes_out" bson:"bytes_out"`
	Req_total        int                    `json:"req_total" bson:"req_total"`
	Attack_total     int                    `json:"attack_total" bson:"attack_total"`
	StatCode1xx      int                    `json:"1xx" bson:"statCode1xx"`
	StatCode2xx      int                    `json:"2xx" bson:"statCode2xx"`
	StatCode3xx      int                    `json:"3xx" bson:"statCode3xx"`
	StatCode4xx      int                    `json:"4xx" bson:"statCode4xx"`
	StatCode5xx      int                    `json:"5xx" bson:"statCode5xx"`
	Addr_total       int                    `json:"addr_total" bson:"addr_total"`
	Bytes_in         int                    `json:"bytes_in" bson:"bytes_in"`
	Conn_total       int                    `json:"conn_total" bson:"conn_total"`
}

type Connection struct {
	Writing          int                    `json:"writing" bson:"writing"`
	Waiting          int                    `json:"waiting" bson:"waiting"`
	Accepted         int                    `json:"accepted" bson:"accepted"`
	Active           int                    `json:"active" bson:"active"`
	Requests         int                    `json:"requests" bson:"requests"`
	Handled          int                    `json:"handled" bson:"handled"`
	Reading          int                    `json:"reading" bson:"reading"`
}

type Upstream struct {
	StatCode400      int                    `json:"400" bson:"400"`
	StatCode401      int                    `json:"401" bson:"401"`
	StatCode402      int                    `json:"402" bson:"402"`
	StatCode403      int                    `json:"403" bson:"403"`
	StatCode404      int                    `json:"404" bson:"404"`
	StatCode405      int                    `json:"405" bson:"405"`
	StatCode406      int                    `json:"406" bson:"406"`
	StatCode407      int                    `json:"407" bson:"407"`
	StatCode408      int                    `json:"408" bson:"408"`
	StatCode409      int                    `json:"409" bson:"409"`
	StatCode410      int                    `json:"410" bson:"410"`
	StatCode411      int                    `json:"411" bson:"411"`
	StatCode412      int                    `json:"412" bson:"412"`
	StatCode413      int                    `json:"413" bson:"413"`
	StatCode414      int                    `json:"414" bson:"414"`
	StatCode415      int                    `json:"415" bson:"415"`
	StatCode416      int                    `json:"416" bson:"416"`
	StatCode417      int                    `json:"417" bson:"417"`
	StatCode500      int                    `json:"500" bson:"500"`
	StatCode501      int                    `json:"501" bson:"501"`
	StatCode502      int                    `json:"502" bson:"502"`
	StatCode503      int                    `json:"503" bson:"503"`
	StatCode504      int                    `json:"504" bson:"504"`
	StatCode505      int                    `json:"505" bson:"505"`
	StatCode507      int                    `json:"507" bson:"507"`
	StatCode1xx      int                    `json:"1xx" bson:"1xx"`
	StatCode2xx      int                    `json:"2xx" bson:"2xx"`
	StatCode3xx      int                    `json:"3xx" bson:"3xx"`
	StatCode4xx      int                    `json:"4xx" bson:"4xx"`
	StatCode5xx      int                    `json:"5xx" bson:"5xx"`
	Req_total        int                    `json:"req_total" bson:"req_total"`
	Bytes_in         int                    `json:"bytes_in" bson:"bytes_in"`
	Bytes_out        int                    `json:"bytes_out" bson:"bytes_out"`
}

func init()  {
	_, err := mongo.Count(wafCollectionName)
	if err != nil {
		tools.Panic(tools.ErrCodeMongoInitFailed, "failed to get waf collection count", err)
	}
}

func GetAllWafNameSort() (count int, result []*Waf, err error) {
	// count, err = mongo.FindAll(appCollectionName, nil, &result, perpage*(page-1), perpage, "name")
	count, err = mongo.FindAll(wafCollectionName, nil, &result, 0, 10, "name")
	// if err == nil && result != nil {
		
	// }
	return
}

func GetAllWafIdSort() (count int, result []*Waf, err error) {
	// count, err = mongo.FindAll(appCollectionName, nil, &result, perpage*(page-1), perpage, "name")
	count, err = mongo.FindAll(wafCollectionName, nil, &result, 0, 10, "_id")
	// if err == nil && result != nil {
		
	// }
	return
}

func GetWafById( id string ) (result *Waf, err error) {
	err =  mongo.FindId(wafCollectionName, id, &result)
	return
}

func ConfigNewWaf(wafName string, wafAddr string) (result *Waf, err error) {
	waf, err := AddWaf(&Waf{
		Name:        wafName,
		Addr:        wafAddr,
	})
	if err != nil {
		// beego.Error("failed to config new waf :" + err.Error())
		// tools.Panic(tools.ErrCodeInitDefaultAppFailed, "failed to config new waf", err)
		return nil,errors.New("failed to config new waf: " + err.Error())
		// return nil,err
	}
	// result = waf
	result, err = GetWafById(waf.Id)
	if err != nil {
		return nil,errors.New("failed to GetWafById: " + err.Error())
	}
	return
}

func ConfigNewWafWithStatus(wafName string, wafAddr string, wafStatusResp *WafStatusResp) (result *Waf, err error) {
	wafInput := Waf{
		Name:              wafName,
		Addr:              wafAddr,
	}
	wafInput.WafStatusResp = wafStatusResp
	waf, err := AddWaf(wafInput)
	if err != nil {
		return nil,errors.New("failed to config new waf: " + err.Error())
	}
	result, err = GetWafById(waf.Id)
	if err != nil {
		return nil,errors.New("failed to GetWafById: " + err.Error())
	}
	return
}

func AddWaf(waf *Waf) (result *Waf, err error) {
	waf.Id = GenerateWafId(waf)
	waf.CreateTime = time.Now().Unix()
	if err := mongo.FindOne(wafCollectionName, bson.M{"name": waf.Name}, &Waf{}); err != mgo.ErrNotFound {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("duplicate waf name")
	}
	err = mongo.Insert(wafCollectionName, waf)
	if err != nil {
		return nil, errors.New("failed to insert waf to db: " + err.Error())
	}
	result = waf
	beego.Info("Succeed to create app, name: " + waf.Name)
	// selectDefaultPlugin(app)
	return
}

func GenerateWafId(waf *Waf) string {
	random := "openwaf" + waf.Name + strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(rand.Intn(10000))
	return fmt.Sprintf("%x", sha1.Sum([]byte(random)))
}