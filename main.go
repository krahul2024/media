package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"media/db"
	"media/handler"
	"media/route"
	"media/service"
	"net/http"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

type server struct {
    FileRouter route.FileRouter
}


func main() {
    setEnvs()
    PORT := os.Getenv("PORT")
    DB_NAME := os.Getenv("DB_NAME")
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_HOST := os.Getenv("DB_HOST")
    DB_PORT := os.Getenv("DB_PORT")

    dbURL := fmt.Sprintf(
        "postgres://%v:%v@%v:%v/%v?sslmode=disable",
        DB_USER, DB_PASSWORD,
        DB_HOST, DB_PORT,
        DB_NAME,
        )

    dbConn, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("db connection error, exiting...")
        return
    }

    fileModel := &db.SQLFileModel{DB: dbConn}
    fileService := service.FileService{FileModel: fileModel}
    fileHandler := handler.FileHandler{ FileService: fileService }

    if len(PORT) == 0 || PORT == "" {
        log.Fatalln("PORT Not found, exiting...")
        return
    }

    serveMux := http.NewServeMux()


    log.Println("Starting the server at PORT: ", PORT)
    if err := http.ListenAndServe(
        fmt.Sprintf(":%v", PORT), serveMux,
        ); err != nil {
        log.Fatalln("Error starting the server, exiting...")
        return
    }

}


func setEnvs() {
    fmt.Println("Setting the Environment Variables...")
    file, err := os.Open(".env")
    if err != nil {
        log.Fatal(err)
        return
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lineBytes := scanner.Bytes()
        lineString := strings.TrimSpace(string(lineBytes))
        if len(lineString) == 0 || strings.HasPrefix(lineString, "#") {
            continue
        }
        line := strings.SplitN(lineString, "=", 2)
        if len(line) == 1 {
            line = append(line, "")
        }
        key, value := strings.TrimSpace(line[0]), strings.TrimSpace(line[1]);
        value = strings.Trim(value, `"`)
        value = strings.Trim(value, `'`)
        os.Setenv(key, value)
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading the file contents")
    }
}
