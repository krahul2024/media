package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/krahul2024/go-multer/routes"
    "github.com/krahul2024/go-multer/utils"
)

func main() {
    utils.SetEnv()
    PORT := os.Getenv("PORT")

    mux := http.NewServeMux()

    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("This is index Route"));
    }))

    // Auth Routes
    mux.Handle("/auth/", routes.AuthMux())

    fmt.Println(fmt.Sprintf("Server runinng on the port: %v", PORT))
    http.ListenAndServe(fmt.Sprintf(":%v", PORT), mux)
}


/*
I have following routes,
1. Auth Route
2. ....more

*/
