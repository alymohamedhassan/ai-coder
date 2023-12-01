package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
  "bufio"
  "encoding/json"
)

func LoadContext() ConfigType {
  path, filename := GetConfigFilePath()
  configurations, _ := os.ReadFile(filepath.Join(path, filename))
  
  var object ConfigType
  err := json.Unmarshal(configurations, &object)

  if err != nil {
    fmt.Println("Error:", err)
  }

  return ConfigType{
    API_KEY: object.API_KEY,
    PROJECT_DIRECTORY: object.PROJECT_DIRECTORY,
    Frameworks: object.Frameworks,
    Scripts: object.Scripts,
  }
}

func setupProjectDirectory(scanner *bufio.Scanner) string {
  var path string
  fmt.Print("Enter your projects directory path: ")
  for scanner.Scan() {
    path = scanner.Text()
    break
  }
  return path
}

func IsSetup(path string) bool {
  _, err := os.Open(path)
  return err == nil
}

func GetConfigFilePath() (string, string) {
  var path string
  fileName := "c.conf"

  if runtime.GOOS == "windows" {
    path = "C:\\ai-coder\\"
  } else {
    path = "/usr/ai-coder"
  }

  return path, fileName
}

func Setup(scanner *bufio.Scanner) {
  configDirectoryPath, filename := GetConfigFilePath()
  isSetup := IsSetup(filepath.Join(configDirectoryPath, filename))
  
  if !isSetup {
    projectDir := setupProjectDirectory(scanner)

    var apiKey string
    fmt.Print("Palm 2 Api Key: ")
    for scanner.Scan() {
      apiKey = scanner.Text()
      break
    }
    configurations := ConfigType{
      API_KEY: apiKey,
      PROJECT_DIRECTORY: projectDir,
      Frameworks: []Framework{
        {Name: "NestJs", Command: "nest new {project_name}"},
        {Name: "Django", Command: "django-admin startproject {project_name}"},
      },
    }
    fmt.Println("configurations:", configurations)
    CreateConfigFile(configDirectoryPath, "c.conf", configurations)
  } else {
    fmt.Println("You already setup your application")
  }
}
