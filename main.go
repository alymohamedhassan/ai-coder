package main

import (
	"fmt"
	"os"
  	"bufio"
	"ai-coder/utils"
	"ai-coder/applications"
)

func main() {
  var args []string

  if len(os.Args) < 2 {
    utils.StartMenu()
  } else {
    argument := os.Args[1]
    if argument == "setup" {
      utils.StartMenuChoice = "Setup"
    } else 
    if argument == "s" || argument == "scripts" {
      utils.StartMenuChoice = "Scripts"
      fmt.Println("Module", utils.StartMenuChoice)
    } else 
    if argument == "ai" {
      utils.StartMenuChoice = "AI Assistant"
      fmt.Println("Module", utils.StartMenuChoice)
    } else 
    if argument == "p" || argument == "project" {
      args = os.Args[1:]
      utils.StartMenuChoice = "Projects"
      fmt.Println("Module", utils.StartMenuChoice)
    } else
    if argument == "c" {
      utils.StartMenuChoice = "Configurations"
      fmt.Println("Module", utils.StartMenuChoice)
    }  
  }

  // Run the selected application
  if utils.StartMenuChoice == "AI Assistant" {
    applications.RunAiApplication()
  } else

  if utils.StartMenuChoice == "Projects" {
    if len(args) == 0 { args = append(args, "p")}
    applications.RunProjectsApplication(args)
  } else

  if utils.StartMenuChoice == "Configurations" {
    if len(args) == 0 { args = append(args, "c")}
    applications.RunConfigurationApplication(args)
  } else

  if utils.StartMenuChoice == "Scripts" {
    if len(args) == 0 { args = append(args, "s")}
    applications.RunScriptsApplication(args)
  } else

  if utils.StartMenuChoice == "Setup" {
    scanner := bufio.NewScanner(os.Stdin)
    utils.Setup(scanner)
  }
}

