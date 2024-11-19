package route

import (
    "media/handler"

    "net/http"
)

type FileRouter struct {
    FileHandler handler.FileHandler
}

func (r *FileRouter) Routes() *http.ServeMux {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /all", r.FileHandler.GetAllFiles)
    mux.HandleFunc("POST /new", r.FileHandler.CreateNewFile)
    mux.HandleFunc("GET /{id}/info", r.FileHandler.GetFileInfo)
    mux.HandleFunc("GET /{id}/load", r.FileHandler.LoadFile)
    mux.HandleFunc("GET /{id}/update", r.FileHandler.UpdateFile)
    mux.HandleFunc("GET /{id}/delete", r.FileHandler.DeleteFile)

    return mux
}

