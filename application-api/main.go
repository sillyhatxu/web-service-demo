package main

import "github.com/sillyhatxu/web-service-demo/application-api/api"

func main() {
	//os.Setenv("TEST","hello api")
	//os.Setenv("INTERNAL_HOST","http://localhost:8082")
	//api.InitialAPI(8081)
	api.InitialAPI(8080)
}
