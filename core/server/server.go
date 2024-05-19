package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitfiend2/core/git"
	"io"
	"net/http"
	"os"
)

const port = 29998

// StartServer
// This is a server for testing against the current Rust implementation.
func StartServer() {
	http.HandleFunc("/f/{funcName}", func(writer http.ResponseWriter, req *http.Request) {
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
	})

	addr := fmt.Sprint(":", port)
	err := http.ListenAndServe(addr, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Started server on port ", port)
	}
}

type ReqOptions struct {
	RepoPath string `json:"repoPath"`
}

func ReqGitVersion(_ ReqOptions) git.VersionInfo {
	git.LoadGitVersion()
	return git.Version
}

func handleFuncRequest(name string, reqData []byte) ([]byte, bool) {
	if name == "git_version" {
		res, ok := CallFunc(ReqGitVersion, reqData)

		if ok {
			fmt.Println("Func Result: ", res)
			resBytes, err := json.Marshal(res)

			if err == nil {
				return resBytes, true
			}
		}
	}

	return []byte{}, false
}

func CallFunc[P, R any](f func(p P) R, data []byte) (R, bool) {
	var result P
	err := json.Unmarshal(data, &result)

	if err == nil {
		return f(result), true
	}

	return *new(R), false
}
