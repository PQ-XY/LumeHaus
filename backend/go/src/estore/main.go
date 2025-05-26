package main

import (
    "fmt"
    "log"
    "net/http" 

	"estore/internal"
    "estore/handler"   
)
func main() {
    fmt.Println("started-service")

	internal.InitElasticsearchBackend()

    log.Fatal(http.ListenAndServe(":8080", handler.InitRouter()))
}