package routes

import "net/http"


func Post(next http.Handler) http.Handler {
    return http.HandlerFunc( func(w http.ResponseWriter, r *http.Request){
            if r.Method != http.MethodPost {
                http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
                return
            }

            next.ServeHTTP(w, r)
        })
}

func Get(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if r.Method != http.MethodGet {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func Put(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if r.Method != http.MethodPut {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        next.ServeHTTP(w, r)
    })
}

func Delete(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        if r.Method != http.MethodDelete {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }

        next.ServeHTTP(w, r)
    })
}
