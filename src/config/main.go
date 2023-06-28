// The purpose of the main.go file is to house the ENV load-up function, which exports all the env in the specified file so that it can be assessed using the os module
package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.otrex/blog/src/utils"
)

func SetupConfig(path string) error {
	err := godotenv.Load(utils.AbsPath(path))
	os.Setenv("TEMPLATE_PATH", utils.AbsPath("templates"))
	os.Setenv("ASSETS_PATH", utils.AbsPath("assets"))
	os.Setenv("BUTTERCMS_BASE_URL", "https://api.buttercms.com/v2/")
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", "8000")
	}
	return err
}
