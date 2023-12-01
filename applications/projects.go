package applications

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ai-coder/utils"
	"github.com/charmbracelet/bubbles/list"
)

func LoadProjectFrameworks() []utils.Framework {
  configurations := utils.LoadContext().Frameworks
  return configurations
}

func initializeProject(name string) {
  fmt.Println(fmt.Sprintf("Initializing Project: %s", name))

  projects := []list.Item{}

  for _, v := range LoadProjectFrameworks() {
    projects = append(projects, utils.Item(v.Name))
  }

  utils.DynamicMenu(projects, "Select Framework")

  selected := utils.StartMenuChoice

  var command string
  
  for _, v := range LoadProjectFrameworks() {
    if selected == v.Name {
      command = strings.Replace(v.Command, "{project_name}", name, 1)
    }
  }

  projectDirectory := utils.LoadContext().PROJECT_DIRECTORY

  fmt.Println(">", command)
  utils.RunCmd(command, projectDirectory)
}

func createProject() {
  fmt.Print("Project Name: ")
  scanner := bufio.NewScanner(os.Stdin)
  var projectName string
  for scanner.Scan() {
    projectName = scanner.Text()
    break
  }

  projectPath := filepath.Join(utils.LoadContext().PROJECT_DIRECTORY, projectName)
  
  fmt.Println("Project Path:", projectPath)

  initializeProject(projectName)
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
    fmt.Println("Path:", filepath.Join(utils.LoadContext().PROJECT_DIRECTORY, selectedProjectPath))
    utils.RunCmd(fmt.Sprintf("nvim %s", selectedProjectPath), filepath.Join(utils.LoadContext().PROJECT_DIRECTORY, selectedProjectPath))
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

  selected := utils.StartMenuChoice

  if selected == "New Project" {
    createProject()
  } else
  if selected == "List Projects" {
    ListProjects()
  }
}
