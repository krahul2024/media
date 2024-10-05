package routes

import (
	"log"
	"net/http"
)

func postRoutes(router *Router){
    router.HandleFunc("", listPosts)    
    router.HandleFunc("/create", createPost)
}

func listPosts(w http.ResponseWriter, r *http.Request){
    log.Println("from the route /post")
}

func createPost(w http.ResponseWriter, r *http.Request){
    log.Println("from the route /post/create")
}

