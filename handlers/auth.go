package handlers

import (
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
   fmt.Println("this is register function"); 
    return;
}

func Login (w http.ResponseWriter, r *http.Request) {
   fmt.Println("this is login function"); 
    return;
}

func Logout (w http.ResponseWriter, r *http.Request) {
   fmt.Println("this is logout function"); 
    return;
}
