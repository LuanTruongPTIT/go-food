package environment

import (
	"errors"
	"fmt"

	"log"
	"os"
	"path/filepath"

	"github.com/LuanTruongPTIT/go-food/internal/pkg/constants"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

var (
	Development = Environment(constants.Dev)
	Production  = Environment(constants.Production)
	Test        = Environment(constants.Test)
)

func ConfigAppEnv(environments ...Environment) Environment {
	environment := Environment("")
	if len(environments) > 0 {
		environment = environments[0]
	} else {
		environment = Production
	}

	viper.AutomaticEnv()
	err := loadEnvFilesRecursively()
	if err != nil {
		log.Printf(".env file cannot be found, err: %v", err)
	}
	setRootWorkingDirectoryEnvironment()
	FixProjectRootWorkingDirectory()

	manualEnv := os.Getenv(constants.AppEnv)

	if manualEnv != "" {
		environment = Environment(manualEnv)
	}
	return environment
}

func loadEnvFilesRecursively() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		envFilePath := filepath.Join(dir, ".env")
		err := godotenv.Load(envFilePath)

		if err == nil {
			return nil
		}
		parentDir := filepath.Dir(dir)
		fmt.Println("file: environment.go", parentDir)
		if parentDir == dir {
			break
		}
		dir = parentDir
	}
	return errors.New("no .env file found")
}

func setRootWorkingDirectoryEnvironment() {
	absoluteRootWorkingDirectory := GetProjectRootWorkingDirectory()
	viper.Set(constants.AppRootPath, absoluteRootWorkingDirectory)
}
