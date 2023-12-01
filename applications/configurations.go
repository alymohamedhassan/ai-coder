package applications

import (
	"ai-coder/utils"
	"bufio"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
)

func addNewFramework() utils.Framework {
  fmt.Print("Name: ")
  scanner := bufio.NewScanner(os.Stdin)
  var name string
  for scanner.Scan() {
    name = scanner.Text()
    break
  }

  fmt.Print("Command: ")
  var command string
  for scanner.Scan() {
    command = scanner.Text()
    break
  }

  return utils.Framework{
    Name: name,
    Command: command,
  }
}

func RunConfigurationApplication(args []string) {
  fmt.Println("Configuration Application")

  items := []list.Item{
    utils.Item("Add new Framework"),
  }
  utils.DynamicMenu(items, "Select Option")

  if utils.StartMenuChoice == "Add new Framework" {
    framework := addNewFramework()
    Frameworks := utils.LoadContext().Frameworks
    newFramework := append(Frameworks, framework)
    fmt.Println(newFramework)

    configDirectorypath, _ := utils.GetConfigFilePath()
    oldConfigurations := utils.LoadContext()
    configurations := utils.ConfigType{
      API_KEY: oldConfigurations.API_KEY,
      PROJECT_DIRECTORY: oldConfigurations.PROJECT_DIRECTORY,
      Frameworks: newFramework,
    }

    utils.CreateConfigFile(configDirectorypath, "c.conf", configurations)
  }
}
