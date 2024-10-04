package routes

import (
	"fmt"
	"net/http"
)

func textRoutes(){
    http.HandleFunc("/text", textIndexHandler)
    http.HandleFunc("/text/upload", textUploadHandler)
}

func textIndexHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println("This is the index handler for text")
}

func textUploadHandler(w http.ResponseWriter, r *http.Request){
    fmt.Println("This is the upload handler for text")
}
