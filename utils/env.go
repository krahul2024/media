package utils

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func SetEnv() error {
    file, err := os.Open(".env")
    defer file.Close()

    if err != nil {
        return errors.New("no env file found")
    }

    reader := bufio.NewReader(file)
    for {
        lineBytes, _, err := reader.ReadLine()
        if len(lineBytes) > 0 {
            lineParts := strings.Split(strings.TrimSpace(string(lineBytes)), "=")
            if len(lineParts) == 1 {
                lineParts = append(lineParts, "")
            }
            if lineParts[0] == "" {
                continue
            }        
            key, value := strings.TrimSpace(lineParts[0]), strings.TrimSpace(lineParts[1])
            os.Setenv(key, value)
        } else if err == io.EOF {
            break
        } else if err != nil {
            return errors.New("error reading environment variables")
        }
    }

    return nil
}
