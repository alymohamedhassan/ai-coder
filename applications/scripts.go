package applications

import (
	"ai-coder/utils"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
)



func RunScriptsApplication(args []string) {
  fmt.Println("Init Scripts App")
  scripts := utils.LoadContext().Scripts

  if len(scripts) == 0 {
    fmt.Println("No Scripts added yet!")
    os.Exit(1)
  }

  items := []list.Item{}

  for _, v := range scripts {
    items = append(items, utils.Item(v.Name))
  }

  utils.DynamicMenu(items, "Select script to run")

  selected := utils.StartMenuChoice

  var command string
  for _, v := range scripts {
    if selected == v.Name {
      command = v.Command
    }
  }
  utils.RunCmd(command, "./")
}
