package environment

import (
	"errors"

	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/LuanTruongPTIT/go-food/internal/pkg/constants"
	"github.com/spf13/viper"
)

func GetProjectRootWorkingDirectory() string {
	var rootWorkingDirectory string

	pn := viper.GetString(constants.PROJECT_NAME_ENV)

	if pn != "" {
		rootWorkingDirectory = getProjectRootDirectoryFromProjectName(pn)
	} else {
		wd, _ := os.Getwd()
		dir, err := searchRootDirectory(wd)

		if err != nil {
			log.Fatal(err)
		}
		rootWorkingDirectory = dir
	}
	absoluteRootWorkingDirectory, _ := filepath.Abs(rootWorkingDirectory)
	return absoluteRootWorkingDirectory
}

func getProjectRootDirectoryFromProjectName(pn string) string {
	wd, _ := os.Getwd()

	for !strings.HasSuffix(wd, pn) {
		wd = filepath.Dir(wd)
	}

	return wd
}
func searchRootDirectory(dir string) (string, error) {
	files, err := os.ReadDir(dir)

	if err != nil {
		return "", errors.New("go.mod file not found")
	}
	for _, file := range files {
		if !file.IsDir() {
			filerName := file.Name()
			if strings.EqualFold(filerName, "go.mod") {
				return dir, nil
			}
		}
	}
	// check parent directory
	parentDir := filepath.Dir(dir)

	if parentDir == dir {
		return "", errors.New("go.mod file not found")
	}
	return searchRootDirectory(parentDir)
}

func FixProjectRootWorkingDirectory() {
	currentWD, _ := os.Getwd()
	log.Printf("Current working directory: %s", currentWD)

	rootDir := GetProjectRootWorkingDirectory()
	_ = os.Chdir(rootDir)
	newWD, _ := os.Getwd()

	log.Printf("New working directory: %s", newWD)
}
