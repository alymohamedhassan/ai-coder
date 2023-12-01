package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"


)

func isPath(path string) bool {
  return filepath.IsAbs(path) || filepath.Dir(path) != ""
}

func SaveFile(name string, fileType string, content string) {
  var filePath string
  if isPath(name) {
    filePath = name
  } else{
    filePath = filepath.Join("/usr/ai-coder/", name)
  }

  f, err := os.Create(filePath)
  
  defer f.Close()

  _, err = f.WriteString(content)

  if err != nil {
    fmt.Println("Error writing to file:", err)
    return
  }
}

func RunCmd(command string, directory string) {
  cmd := exec.Command("bash", "-c", command)
  cmd.Dir = directory
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  _ = cmd.Run() // add error checking
}

type Framework struct {
  Name string
  Command string
}

type Script struct {
  Name string
  Command string
}

type ConfigType struct {
  API_KEY string
  PROJECT_DIRECTORY string 
  Frameworks []Framework
  Scripts []Script
}

func CreateConfigFile(path string, fileName string, configurations ConfigType) {
  _, err := os.Stat(path)
  if err != nil {
    os.Mkdir(path, 0755)
  }

  filePath := filepath.Join(path, fileName)
  f, err := os.Create(filePath)
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
    return
  }

  // Marshal the struct to JSON
  jsonBytes, err := json.Marshal(configurations)
  if err != nil {
    fmt.Println("Error marshaling JSON:", err)
    return
  }

	fmt.Println(configurations)
	fmt.Println(string(jsonBytes))

  defer f.Close()

  _, err = f.WriteString(string(jsonBytes))
  if err != nil {
    fmt.Println("Error writing to file:", err)
    return
  }

  fmt.Println("File created successfully")
}

