package handlers

import (
	"fmt"
	"net/http"
)

type AuthHandler struct {
   AuthService service.AuthService 
}
func Register(w http.ResponseWriter, r *http.Request) {

}

func Login (w http.ResponseWriter, r *http.Request) {
   fmt.Println("this is login function"); 
    return;
}

func Logout (w http.ResponseWriter, r *http.Request) {
   fmt.Println("this is logout function"); 
    return;
}
