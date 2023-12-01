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


func addNewScript() utils.Script {
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

  return utils.Script{
    Name: name,
    Command: command,
  }
}

func RunConfigurationApplication(args []string) {
  fmt.Println("Configuration Application")

  items := []list.Item{
    utils.Item("Add new Framework"),
    utils.Item("Remove Framework"),
    utils.Item("Add new Script"),
    utils.Item("Remove Script"),
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
  } else 

  if utils.StartMenuChoice == "Remove Framework" {
    Frameworks := utils.LoadContext().Frameworks

    var items []list.Item
    for _, v := range Frameworks {
      items = append(items, utils.Item(v.Name))
    }

    utils.DynamicMenu(items, "Frameworks")

    var newFrameworks []utils.Framework
    for _, v := range Frameworks {
      if v.Name == utils.StartMenuChoice {
        continue
      }
      newFrameworks = append(newFrameworks, utils.Framework{Name: v.Name, Command: v.Command})
    }

    configDirectorypath, _ := utils.GetConfigFilePath()
    oldConfigurations := utils.LoadContext()
    configurations := utils.ConfigType{
      API_KEY: oldConfigurations.API_KEY,
      PROJECT_DIRECTORY: oldConfigurations.PROJECT_DIRECTORY,
      Frameworks: newFrameworks,
    }

    utils.CreateConfigFile(configDirectorypath, "c.conf", configurations)
  } else 

  if utils.StartMenuChoice == "Remove Script" {
    Scripts := utils.LoadContext().Scripts

    var items []list.Item
    for _, v := range Scripts {
      items = append(items, utils.Item(v.Name))
    }

    utils.DynamicMenu(items, "Scripts")

    var newScripts []utils.Script
    for _, v := range Scripts {
      if v.Name == utils.StartMenuChoice {
        continue
      }
      newScripts = append(newScripts, utils.Script{Name: v.Name, Command: v.Command})
    }

    configDirectorypath, _ := utils.GetConfigFilePath()
    oldConfigurations := utils.LoadContext()
    configurations := utils.ConfigType{
      API_KEY: oldConfigurations.API_KEY,
      PROJECT_DIRECTORY: oldConfigurations.PROJECT_DIRECTORY,
      Frameworks: oldConfigurations.Frameworks,
      Scripts: newScripts,
    }

    utils.CreateConfigFile(configDirectorypath, "c.conf", configurations)
  } else 

  if utils.StartMenuChoice == "Add new Script" {
    script := addNewScript()
    Scripts := utils.LoadContext().Scripts
    newScripts := append(Scripts, script)

    configDirectorypath, _ := utils.GetConfigFilePath()
    oldConfigurations := utils.LoadContext()
    configurations := utils.ConfigType{
      API_KEY: oldConfigurations.API_KEY,
      PROJECT_DIRECTORY: oldConfigurations.PROJECT_DIRECTORY,
      Scripts: newScripts,
    }

    utils.CreateConfigFile(configDirectorypath, "c.conf", configurations)
  }
}
