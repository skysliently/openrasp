package models

import (
	
)

type TestJson struct {
	Success int `json:"success" `
	Result Result `json:"result" `
}

type Result struct {
	Main Mian `json:"mian" `
}

type Mian struct {
	Resetsec int  `json:"resetsec" `
	Safe Safe  `json:"safe" `
	Access Access `json:"access" `
	Nginx_version string `json:"nginx_version" `
	Connection Connection `json:"connection" `
	Upstream Upstream `json:"upstream" `
}

type Safe struct {
	// Main Safe `json:"mian" `
}

type Access struct {
	Bytes_out int `json:"bytes_out" `
	Req_total int `json:"req_total" `
	Attack_total int `json:"attack_total" `
	StatCode1xx int `json:"1xx" `
	StatCode2xx int `json:"2xx" `
	StatCode3xx int `json:"3xx" `
	StatCode4xx int `json:"4xx" `
	StatCode5xx int `json:"5xx" `
	Addr_total int `json:"addr_total" `
	Bytes_in int `json:"bytes_in" `
	Conn_total int `json:"conn_total" `
}

type Connection struct {
	Writing int `json:"writing" `
	Waiting int `json:"waiting" `
	Accepted int `json:"accepted" `
	Active int `json:"active" `
	Requests int `json:"requests" `
	Handled int `json:"handled" `
	Eeading int `json:"reading" `
}

type Upstream struct {
	StatCode400 int `json:"400" `
	StatCode401 int `json:"401" `
	StatCode402 int `json:"402" `
	StatCode403 int `json:"403" `
	StatCode404 int `json:"404" `
	StatCode405 int `json:"405" `
	StatCode406 int `json:"406" `
	StatCode407 int `json:"407" `
	StatCode408 int `json:"408" `
	StatCode409 int `json:"409" `
	StatCode410 int `json:"410" `
	StatCode411 int `json:"411" `
	StatCode412 int `json:"412" `
	StatCode413 int `json:"413" `
	StatCode414 int `json:"414" `
	StatCode415 int `json:"415" `
	StatCode416 int `json:"416" `
	StatCode417 int `json:"417" `
	StatCode500 int `json:"500" `
	StatCode501 int `json:"501" `
	StatCode502 int `json:"502" `
	StatCode503 int `json:"503" `
	StatCode504 int `json:"504" `
	StatCode505 int `json:"505" `
	StatCode507 int `json:"507" `
	StatCode1xx int `json:"1xx" `
	StatCode2xx int `json:"2xx" `
	StatCode3xx int `json:"3xx" `
	StatCode4xx int `json:"4xx" `
	StatCode5xx int `json:"5xx" `
	Req_total int `json:"req_total" `
	Bytes_in int `json:"bytes_in" `
	Bytes_out int `json:"bytes_out" `
}