package main

import (
	"bufio"
	"fmt"
	"log"
	"media/route"
	"net/http"
	"os"
	"strings"
)

func main() {
    setEnvs()
    PORT := os.Getenv("PORT")
    if len(PORT) == 0 || PORT == "" {
        log.Fatalln("PORT Not found, exiting...")
        return
    }

    serveMux := http.NewServeMux()

    fileRoute := route.FileRoutes()
    serveMux.Handle("/api/file/", http.StripPrefix("/api/file", fileRoute))

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
