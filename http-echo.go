package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
    "os"
)

const defaultPort = "8000"

func log(logLevel string, message string) {
    const dateLayout = "2006-01-02 15:04:05"
    curTime := time.Now()
    fmt.Printf("%s [%-5s] %s\n", curTime.Format(dateLayout), logLevel, message)
}

// EchoHandler echos back the request in the response body.
func EchoHandler(writer http.ResponseWriter, request *http.Request) {
    log("INFO", fmt.Sprintf(
        "%s - %s %s %s",
        request.RemoteAddr, request.Method, request.URL.Path, request.Proto,
    ))

    // permissive CORS configuration
    writer.Header().Set("Access-Control-Allow-Origin", "*")
    writer.Header().Set("Access-Control-Allow-Methods", "*")
    writer.Header().Set("Access-Control-Allow-Headers", "*")

    request.Write(writer)
}

func main() {
    port := defaultPort
    args := os.Args[1:]
    if len(args) != 0 {
        port = args[0]
        if _, err := strconv.Atoi(port); err != nil {
            log("ERROR", "invalid port \"" + port + "\"")
            os.Exit(1)
        }
    }

    log("INFO", "Listening on port " + port)

    http.HandleFunc("/", EchoHandler)

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        log("ERROR", err.Error())
        os.Exit(1)
    }
}
