package utils

import (
	"encoding/json"
	"net/http"
)

func HttpRes(w http.ResponseWriter, data any, statusCode int){
    dataBytes, err := json.Marshal(data);
    if err != nil {
        // this should be redirected to the HttpErr function
    }

    w.Write(dataBytes);
    w.WriteHeader(statusCode)
    return
}


func HttpErr(w http.ResponseWriter, errMessage string, statusCode int){
    http.Error(w, errMessage, statusCode)
    // here we can try doing some special kind of logging for errors
    return
}


