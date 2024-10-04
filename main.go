package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/krahul2024/go-multer/routes"
	"github.com/krahul2024/go-multer/utils"
)


func main(){
    utils.SetEnv()
    PORT := os.Getenv("PORT")

    routes.Routes()
    
    fmt.Println(fmt.Sprintf("Server runinng on the port: %v", PORT))
    http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
} 

