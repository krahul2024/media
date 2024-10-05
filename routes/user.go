package routes

import "net/http"

func userRoutes(router *Router){
    router.HandleFunc("/", listUsers)
    router.HandleFunc("/create", createUser)
}

func listUsers(w http.ResponseWriter, r *http.Request){

}

func createUser(w http.ResponseWriter, r *http.Request){

}
