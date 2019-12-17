package main

import (
	"os"

	"github.com/noah-blockchain/explorer-genesis-uploader/internal/core"
	"github.com/noah-blockchain/explorer-genesis-uploader/internal/env"
)

func main() {
	envData := env.New()
	uploader := core.New(envData)
	err := uploader.Do()
	if err != nil {
		panic(err)
	}

	os.Exit(0)
}
