package applications

import (
	"ai-coder/utils"
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	gl "cloud.google.com/go/ai/generativelanguage/apiv1beta2"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/list"
	"github.com/eiannone/keyboard"
	"google.golang.org/api/option"
)

func getCode(s string, language string) string {
  r := regexp.MustCompile("`([^`]+)`")
  match := r.FindStringSubmatch(s)
  if match != nil {
    if strings.HasPrefix(match[1], language) {
      return strings.Replace(match[1], language, "", 1)
    }
    return match[1]
  }
  return ""
}

func generate(prompt string) string {

  ctx := context.Background()
  client, err := gl.NewTextRESTClient(ctx, option.WithAPIKey("AIzaSyBi2LGmRB-ZOCNlZ13URits-GTm1EdCC_E"))
  
  if err != nil {
    panic(err)
  }

  defer client.Close()
  req := &pb.GenerateTextRequest{
    Model: "models/text-bison-001",
    Prompt: &pb.TextPrompt{
      Text: prompt,
    },
  }

  resp, err := client.GenerateText(ctx, req)
  if err != nil {
    panic(err)
  }

  return resp.Candidates[0].Output
}

func formatResponse(response string) (string, string) {
  language := "Unknown"

  // TODO: Use Regex Instead, for easier and more accurate approach
  if strings.Contains(response, "```") {
    result := strings.Split(response, "\n")
    language := strings.Replace(result[0], "```", "", 99999999)
    return strings.Replace(strings.Join(result[1:], "\n"), "```", "", 999999999), language
  }
  return "", language
}

func runCodeInTerminal(scanner bufio.Scanner, code string) {
  fmt.Print("Do you want to execute the commands? y (Yes) / n (No): ")
  char, _, err := keyboard.GetSingleKey()
  if (err != nil) {
      panic(err)
  }
  if char == 'y' {
    fmt.Println("Yes")
    for _, value := range strings.Split(code, "\n") {
      if value == "" {
        continue
      }
      fmt.Println("Agent >", value)
      utils.RunCmd(value)
    }
  }
}

func RunAiApplication() {
  scanner := bufio.NewScanner(os.Stdin)

  var prompt string
  fmt.Print("Prompt > ") 

  var language, code string

	for scanner.Scan() {
    prompt = scanner.Text()
    fmt.Println("Prompt:", prompt)
    break
	}
  
  output := generate(prompt)

  fmt.Println(output)
  code, language = formatResponse(output)

  if code != "" {
    fmt.Println("Language:", language)
    fmt.Println("Code ------------------------")
    fmt.Println(code)

    fmt.Print("Do you want to execute this code or save it? r (Run) / s (Save) / c (Copy) / q (quit): ")
    char, _, err := keyboard.GetSingleKey()
    if (err != nil) {
        panic(err)
    }

    if char == 'q' {
      fmt.Println("Quitting...")
      os.Exit(1)
      return
    }

    if char == 's' {
      fmt.Println("Save")
      var fileName string
      
      items := []list.Item{
        utils.Item("Yes"),
        utils.Item("No"),
      }
      utils.DynamicMenu(items, "Do you want to save it to an existing project?")

      if utils.StartMenuChoice == "Yes" {      
        items := []list.Item{}
        entries, _ := os.ReadDir(utils.LoadContext().PROJECT_DIRECTORY)

        for _, e := range entries {
          items = append(items, utils.Item(e.Name()))
        }
        utils.DynamicMenu(items, "Select Project")
        projectName := utils.StartMenuChoice
        fileName = filepath.Join(utils.LoadContext().PROJECT_DIRECTORY, projectName)
        fmt.Println("FilePath:", fileName)

        fmt.Print("Filename: ")
        for scanner.Scan() {
          fileName = filepath.Join(fileName, scanner.Text())
          break
        }
      } else {
        fmt.Print("Filename: ")
        for scanner.Scan() {
          fileName = scanner.Text()
          break
        }
      }

      utils.SaveFile(fileName, "code", code)
      
    } else if char == 'c' {
      fmt.Println("Copy")
      
      err := clipboard.WriteAll(strings.Replace(code, "\n", "", 99))
      if err != nil {
        fmt.Println("Copied Code:", code)
      }        
    
    } else if char == 'r' {
      fmt.Println("Run")

      fmt.Print("Are you sure? y (Yes) / n (No): ")
      char, _, err := keyboard.GetSingleKey()
      if (err != nil) {
          panic(err)
      }
      
      if char == 'n' {
        os.Exit(1)
      }

      if language != "bash" {
        fmt.Print("Do you want to consider this language as bash? y (Yes) / n (No): ")
        char, _, err := keyboard.GetSingleKey()
        if (err != nil) {
            panic(err)
        }
        if char == 'y' {
          fmt.Println("Yes")
          language = "bash"
        }
      }

      if language == "bash" {
        runCodeInTerminal(*scanner, code)
      }
    } else {
      fmt.Println("Invalid Input!")
    }
  }
}
