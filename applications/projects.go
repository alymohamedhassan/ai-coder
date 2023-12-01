package applications

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"ai-coder/utils"
	"github.com/charmbracelet/bubbles/list"
)

func LoadProjectFrameworks() []utils.Framework {
  configurations := utils.LoadContext().Frameworks
  fmt.Println("Configs:", configurations)
  return configurations
}

func initializeProject() {
  fmt.Println("Initializing Project")

  projects := []list.Item{}

  for _, v := range LoadProjectFrameworks() {
    projects = append(projects, utils.Item(v.Name))
  }

  utils.DynamicMenu(projects, "Select Framework")

  selected := utils.StartMenuChoice

  var command string
  
  for _, v := range LoadProjectFrameworks() {
    if selected == v.Name {
      command = v.Command
    }
  }

  utils.RunCmd(command)
}

func createProject() {
  fmt.Print("Project Name: ")
  scanner := bufio.NewScanner(os.Stdin)
  var projectName string
  for scanner.Scan() {
    projectName = scanner.Text()
    break
  }

  var environment string

  // items := []list.Item{
  //   utils.Item("New Environment"),
  // }
  // entries, _ := os.ReadDir(utils.LoadContext().PROJECT_DIRECTORY)

  // for _, e := range entries {
  //   items = append(items, utils.Item(e.Name()))
  // }
  // utils.DynamicMenu(items)

  // if utils.StartMenuChoice != "New Environment" {
  //   environment = utils.StartMenuChoice
  // } else {
  //   fmt.Print("Enter environment name: ")
  //   for scanner.Scan() {
  //     environment = scanner.Text()
  //     break
  //   }
  // }

  projectPath := filepath.Join(utils.LoadContext().PROJECT_DIRECTORY, environment, projectName)
  
  fmt.Println("Project Path:", projectPath)

  initializeProject()
}

func ListProjects() {
  items := []list.Item{}
  entries, _ := os.ReadDir(utils.LoadContext().PROJECT_DIRECTORY)

  for _, e := range entries {
    items = append(items, utils.Item(e.Name()))
  }
  utils.DynamicMenu(items, "Select Project")

  selectedProjectPath := utils.StartMenuChoice

  fmt.Print("Do you want to open it using nvim?")
  options := []list.Item{
    utils.Item("Yes"),
    utils.Item("No"),
  }
  utils.DynamicMenu(options, "Confirm")

  if utils.StartMenuChoice == "Yes" {
    utils.RunCmd(fmt.Sprintf("nvim %s", selectedProjectPath))
  }
}

func RunProjectsApplication(args []string) {
  fmt.Println("Run Projects Application")

  if len(args) == 1 {
    items := []list.Item{
      utils.Item("List Projects"),
      utils.Item("New Project"),
    }
    utils.DynamicMenu(items, "Select Option")
  }
  fmt.Println("Selected:", utils.StartMenuChoice)

  selected := utils.StartMenuChoice

  if selected == "New Project" {
    createProject()
  } else
  if selected == "List Projects" {
    ListProjects()
  }
}
