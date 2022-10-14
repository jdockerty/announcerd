package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jdockerty/announcerd/pkg/config"
	"github.com/rs/zerolog"
)

const (
    ghAppIdEnv = "ANNOUNCERD_GH_APP_ID"
    ghAppKeyFileEnv = "ANNOUNCERD_GH_APP_KEY_FILE"
)

func main() {

    appIdEnv, ok := os.LookupEnv(ghAppIdEnv)
    if !ok {
        fmt.Printf("%s is required to be set", ghAppIdEnv)
        return
    }

    appId, err := strconv.ParseInt(appIdEnv, 10, 64)
    if err != nil {
        fmt.Printf("could not parse %s value to int64: %s\n", ghAppIdEnv, err)
    }

    ghAppKeyFilePath, ok := os.LookupEnv(ghAppKeyFileEnv)
    if !ok {
        fmt.Printf("%s is required to be set", ghAppKeyFileEnv)
        return
    }

    c, err := config.New(appId, 1, ghAppKeyFilePath)
	if err != nil {
		panic(err)
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &logger

	if err != nil {
		panic(err)
	}

    http.HandleFunc("/pulls", func(w http.ResponseWriter, req *http.Request) {
        PullRequestEventHandler(w, req, c.Client)
    })

    addr := fmt.Sprintf("%s:%d", "localhost", 6000)
	logger.Info().Msgf("Starting server on %s...", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
