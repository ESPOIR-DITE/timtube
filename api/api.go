package api

import (
	"gopkg.in/resty.v1"
	"timtube/config"
)

//const BASE_URL string = "http://localhost:8081/"

//const BASE_URL string = "http://172.17.0.4:8081/"

//const BASE_URL string = "http://172.18.0.3:8081/"

const BASE_URL string = "http://74.208.50.103:8081/"

func Rest() *resty.Request {
	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
