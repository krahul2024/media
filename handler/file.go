package handler

import (
    "encoding/json"
    "media/service"
    "net/http"
)

type FileHandler struct {
    FileService service.FileService
}

func (h *FileHandler) GetAllFiles(w http.ResponseWriter, r *http.Request) {
    files, err := h.FileService.GetAllFiles()
    if err != nil {
        http.Error(w, "error get files", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(files)
}

func (h *FileHandler) CreateNewFile(w http.ResponseWriter, r *http.Request) {
    var file struct {
        Name string `json:"name"`
    }

    err := json.NewDecoder(r.Body).Decode(&file);
    if err != nil {
        http.Error(w, "invalid body/payload", http.StatusBadRequest)
        return
    }
    if file.Name == "" {
        http.Error(w, "incomplete information", http.StatusBadRequest)
        return
    }

    err = h.FileService.CreateFile(file.Name)
    if err != nil {
        http.Error(w, "file create error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func (h *FileHandler) GetFileInfo(w http.ResponseWriter, r *http.Request) {
}

func (h *FileHandler) LoadFile(w http.ResponseWriter, r *http.Request) {

}

func (h *FileHandler) UpdateFile(w http.ResponseWriter, r *http.Request) {

}

func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request) {

}
