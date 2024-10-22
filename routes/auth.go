package routes

import (
	"net/http"

	"github.com/krahul2024/go-multer/handlers"
)

func AuthMux() http.Handler {
    authMux := http.NewServeMux()

    authMux.Handle("/register", Post(http.HandlerFunc(handlers.Register)))
    authMux.Handle("/login", Post(http.HandlerFunc(handlers.Login)))
    authMux.Handle("/logout", Get(http.HandlerFunc(handlers.Logout)))

    return http.StripPrefix("/auth", authMux)
}
