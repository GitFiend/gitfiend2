package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"os"
)

// const port = ":0"
const port = ":29998"

func StartServer() {
	http.HandleFunc(
		"/f/{funcName}", func(writer http.ResponseWriter, req *http.Request) {
			body, err := io.ReadAll(req.Body)
			if err != nil {
				slog.Error(err.Error())
			}

			funcName := req.PathValue("funcName")
			if funcName == "" {
				slog.Error("funcName is empty")
				return
			}
			result, ok := handleFuncRequest(funcName, body)
			if ok {
				_, err = writer.Write(result)
				if err != nil {
					slog.Error(err.Error())
				}
			}
		},
	)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// Needs to print to stdout.
	fmt.Printf("PORT:%d\n", listener.Addr().(*net.TCPAddr).Port)

	err = http.Serve(listener, nil)

	if errors.Is(err, http.ErrServerClosed) {
		slog.Info("server closed")
	} else if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func callFunc[P, R any](f func(p P) R, data []byte) (R, bool) {
	var result P
	err := json.Unmarshal(data, &result)
	if err == nil {
		return f(result), true
	}
	slog.Error(err.Error())
	return *new(R), false
}
