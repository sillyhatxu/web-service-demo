package main

import (
	"github.com/sillyhatxu/web-service-demo/application-internal-api/api"
)

func main() {
	//os.Setenv("TEST","hello internal")
	//os.Setenv("MOCK_DETAIL",`{"id":1,"value":"hello world","name":"test","time":"2020-09-27"}`)
	//api.InitialAPI(8082)
	api.InitialAPI(8080)
}
