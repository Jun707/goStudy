package main

import (
    "log/slog"
    "net/http"
    "flag"
    "os"
)

func main() {
    addr := flag.String("addr", ":8000", "Http network-address")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    mux := http.NewServeMux()
    
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
    mux.HandleFunc("GET /{$}", home)
    mux.HandleFunc("GET /snippet/view/{id}", snippetView)
    mux.HandleFunc("GET /snippet/create", snippetCreate)
    mux.HandleFunc("POST /snippet/create", snippetCreatePost)

    logger.Info("starting server", slog.Any("addr", ":8000"))  
    
    err := http.ListenAndServe(*addr, mux)
    logger.Error(err.Error())
    os.Exit(1)   
}