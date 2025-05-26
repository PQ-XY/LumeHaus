package main

import (
    "fmt"
    "log"
    "net/http" 

	"estore/internal"
    "estore/handler"   
	"estore/util"
)
func main() {
    fmt.Println("started-service")

	config, err := util.LoadApplicationConfig("conf", "deploy.yml")
    if err != nil {
        panic(err)
    }

	internal.InitElasticsearchBackend(config.ElasticsearchConfig)
	internal.InitGCSBackend(config.GCSConfig)

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter(config.TokenConfig)))
}