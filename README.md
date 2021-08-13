# http-echo

A simple HTTP echo server written in Go.

## How to use

    go run http-echo.go

This will compile and start the HTTP server. By default, it listens on port 8000.
Once running, you can make requests from any HTTP client to `http://localhost:8000/`
and the server will echo the request back in the response body.

To use a different port, specify the desired port as a command-line argument:

    go run http-echo.go 12345
