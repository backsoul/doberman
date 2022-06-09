package services

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "doberman" // change to relevant project name
// use os package to get the env variable which is already set
func Get(key string) string {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/configs/dev.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	// return the env variable using os package
	return os.Getenv(key)
}
