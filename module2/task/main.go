package main

import (
    "fmt"
    "net/http"
    "runtime"
)

func taskHandler(w http.ResponseWriter, r *http.Request) {
    for name, values := range r.Header {
        for _, value := range values {
            w.Header().Add(name, value)
        }
    }
    w.Header().Add("go_version", runtime.Version())

    w.Write([]byte("success"))

    //w.WriteHeader(http.StatusOK)

    fmt.Printf("remoteAddr: %s status: %d \n", r.RemoteAddr, http.StatusOK)
}


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/healthz", taskHandler)
    server := &http.Server{
        Handler: mux,
        Addr:    ":80",
    }
    server.ListenAndServe()
}
