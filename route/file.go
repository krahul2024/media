package route

import (
	"media/handler"

	"net/http"
)

func FileRoutes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /all", handler.GetAllFiles)
    mux.HandleFunc("POST /new", handler.CreateNewFile)
    mux.HandleFunc("GET /{id}/info", handler.GetFileInfo)
    mux.HandleFunc("GET /{id}/load", handler.LoadFile)
    mux.HandleFunc("GET /{id}/update", handler.UpdateFile)
    mux.HandleFunc("GET /{id}/delete", handler.DeleteFile)

    return mux
}
