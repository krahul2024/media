package routes

import "net/http"

func mediaRoutes(){
    http.HandleFunc("/media", mediaIndexHandler)
    http.HandleFunc("/media/upload", mediaUploadHandler)
}

func mediaIndexHandler(w http.ResponseWriter, r *http.Request){

}

func mediaUploadHandler(w http.ResponseWriter, r *http.Request){

}
