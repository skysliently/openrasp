//Copyright 2017-2019 Baidu Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http: //www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package routers

import (
	"github.com/astaxie/beego"
	"rasp-cloud/conf"
	"rasp-cloud/controllers"
	"rasp-cloud/controllers/agent"
	"rasp-cloud/controllers/agent/agent_logs"
	"rasp-cloud/controllers/api"
	"rasp-cloud/controllers/api/fore_logs"
	"rasp-cloud/controllers/iast"
	"rasp-cloud/tools"
	"rasp-cloud/controllers/test_api"
	"rasp-cloud/controllers/waf_api"
)

func InitRouter() {
	agentNS := beego.NewNamespace("/agent",
		beego.NSNamespace("/heartbeat",
			beego.NSInclude(
				&agent.HeartbeatController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSNamespace("/attack",
				beego.NSInclude(
					&agent_logs.AttackAlarmController{},
				),
			),
			beego.NSNamespace("/policy",
				beego.NSInclude(
					&agent_logs.PolicyAlarmController{},
				),
			),
			beego.NSNamespace("/error",
				beego.NSInclude(
					&agent_logs.ErrorController{},
				),
			),
		),
		beego.NSNamespace("/rasp",
			beego.NSInclude(
				&agent.RaspController{},
			),
		),
		beego.NSNamespace("/report",
			beego.NSInclude(
				&agent.ReportController{},
			),
		),
		beego.NSNamespace("/dependency",
			beego.NSInclude(
				&agent.DependencyController{},
			),
		),
		beego.NSNamespace("/crash",
			beego.NSInclude(
				&agent.CrashController{},
			),
		),
	)
	foregroudNS := beego.NewNamespace("/api",

		beego.NSNamespace("/plugin",
			beego.NSInclude(
				&api.PluginController{},
			),
		),
		beego.NSNamespace("/log",
			beego.NSNamespace("/attack",
				beego.NSInclude(
					&fore_logs.AttackAlarmController{},
				),
			),
			beego.NSNamespace("/policy",
				beego.NSInclude(
					&fore_logs.PolicyAlarmController{},
				),
			),
			beego.NSNamespace("/error",
				beego.NSInclude(
					&fore_logs.ErrorController{},
				),
			),
		),
		beego.NSNamespace("/app",
			beego.NSInclude(
				&api.AppController{},
			),
		),
		beego.NSNamespace("/rasp",
			beego.NSInclude(
				&api.RaspController{},
			),
		),
		beego.NSNamespace("/strategy",
			beego.NSInclude(
				&api.StrategyController{},
			),
		),
		beego.NSNamespace("/token",
			beego.NSInclude(
				&api.TokenController{},
			),
		),
		beego.NSNamespace("/report",
			beego.NSInclude(
				&api.ReportController{},
			),
		),
		beego.NSNamespace("/operation",
			beego.NSInclude(
				&api.OperationController{},
			),
		),
		beego.NSNamespace("/server",
			beego.NSInclude(
				&api.ServerController{},
			),
		),
		beego.NSNamespace("/dependency",
			beego.NSInclude(
				&api.DependencyController{},
			),
		),
	)
	iastNS := beego.NewNamespace("/iast",
		beego.NSInclude(
			&iast.WebsocketController{},
		),
		beego.NSInclude(
			&iast.IastController{},
		),
	)
	userNS := beego.NewNamespace("/user", beego.NSInclude(&api.UserController{}))
	pingNS := beego.NewNamespace("/ping", beego.NSInclude(&controllers.PingController{}))
	versionNS := beego.NewNamespace("/version", beego.NSInclude(&controllers.GeneralController{}))
	//test router DELETE when dev over
	testNS := beego.NewNamespace("/test_router", beego.NSInclude(&test_api.TestController{}))
	//waf router
	wafNS := beego.NewNamespace("/waf", beego.NSInclude(&waf_api.WafController{}))
	ns := beego.NewNamespace("/v1")
	//test router DELETE when dev over
	ns.Namespace(testNS)
	//waf router
	ns.Namespace(wafNS)
	ns.Namespace(pingNS)
	ns.Namespace(versionNS)
	startType := *conf.AppConfig.Flag.StartType
	if startType == conf.StartTypeForeground {
		ns.Namespace(foregroudNS, agentNS, userNS, iastNS)
	} else if startType == conf.StartTypeAgent {
		ns.Namespace(agentNS)
	} else if startType == conf.StartTypeDefault {
		ns.Namespace(foregroudNS, agentNS, userNS, iastNS)
	} else {
		tools.Panic(tools.ErrCodeStartTypeNotSupport, "Unknown -type parameter provided: "+startType, nil)
	}
	if startType == conf.StartTypeForeground || startType == conf.StartTypeDefault {
		beego.SetStaticPath("//", "dist")
	}
	beego.AddNamespace(ns)
}
