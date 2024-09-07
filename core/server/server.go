package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitfiend2/core/git"
	"io"
	"net"
	"net/http"
	"os"
)

// const port = ":0"
const port = ":29998"

var Store = git.NewStore()

// StartServer
// This is a server for testing against the current Rust implementation.
func StartServer() {
	http.HandleFunc(
		"/f/{funcName}", func(writer http.ResponseWriter, req *http.Request) {
			funcName := req.PathValue("funcName")

			body, err := io.ReadAll(req.Body)
			if err != nil {
				fmt.Println(err)
			}

			result, ok := handleFuncRequest(funcName, body)

			if ok {
				_, err := writer.Write(result)

				if err != nil {
					fmt.Println(err)
				}
			}
		},
	)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	fmt.Printf("PORT:%d\n", listener.Addr().(*net.TCPAddr).Port)

	err = http.Serve(listener, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func callFunc[P, R any](f func(p P) R, data []byte) (R, bool) {
	var result P
	err := json.Unmarshal(data, &result)

	if err == nil {
		return f(result), true
	}

	return *new(R), false
}
