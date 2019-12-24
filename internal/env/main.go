package env

import (
	"flag"
	"os"

	"github.com/noah-blockchain/coinExplorer-tools/models"
)

func New() *models.ExtenderEnvironment {
	appName := flag.String("app_name", "Coin Extender", "App name")
	flag.Parse()

	envData := new(models.ExtenderEnvironment)
	envData.DbUser = os.Getenv("DB_USER")
	envData.DbName = os.Getenv("DB_NAME")
	envData.DbPassword = os.Getenv("DB_PASSWORD")
	envData.DbHost = os.Getenv("DB_HOST")
	envData.DbPort = getEnvAsInt("DB_PORT", 5432)
	envData.NodeApi = os.Getenv("NOAH_API_NODE")
	envData.ApiHost = os.Getenv("COIN_EXTENDER_API_HOST")
	envData.ApiPort = getEnvAsInt("COIN_EXTENDER_API_PORT", 10000)
	envData.Debug = getEnvAsBool("DEBUG", true)

	envData.AppName = *appName
	envData.DbMinIdleConns = 10
	envData.DbPoolSize = 20
	return envData
}
