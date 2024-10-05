package routes

import (
	"log"
	"net/http"
)

func userRoutes(router *Router){
    router.HandleFunc("", listUsers)
    router.HandleFunc("/create", createUser)
}

func listUsers(w http.ResponseWriter, r *http.Request){
    log.Println("from the route /user")
}

func createUser(w http.ResponseWriter, r *http.Request){
    log.Println("from the route /user/create")
}
