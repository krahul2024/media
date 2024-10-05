package routes

import "net/http"

func postRoutes(router *Router){
    router.HandleFunc("", listPosts)    
    router.HandleFunc("/create", createPost)
}

func listPosts(w http.ResponseWriter, r *http.Request){

}

func createPost(w http.ResponseWriter, r *http.Request){

}

