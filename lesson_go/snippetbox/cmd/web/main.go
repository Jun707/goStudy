package main

import (
    "log/slog"
    "net/http"
    "flag"
    "os"
)

type application struct {
    logger *slog.Logger
}

func main() {
    addr := flag.String("addr", ":8000", "Http network-address")
    flag.Parse()

    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    app := &application {
        logger: logger,
    }

    logger.Info("starting server", slog.Any("addr", ":8000"))  
    
    err := http.ListenAndServe(*addr, app.routes())
    logger.Error(err.Error())
    os.Exit(1)   
}