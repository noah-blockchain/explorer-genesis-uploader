package explorer_genesis_uploader

import (
	"github.com/noah-blockchain/explorer-genesis-uploader/core"
	"github.com/noah-blockchain/noah-explorer-extender/env"
	"os"
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
