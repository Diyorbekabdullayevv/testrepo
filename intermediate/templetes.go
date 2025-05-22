package main

// import (
// 	"bufio"
// 	"fmt"
// 	"html/template"

// 	// "html/template"
// 	"os"
// 	"strings"
// )

// func main() {

// 	// tmpl := template
// 	// tmpl, err := template.New("example").Parse("Welcome {{.name}}! How are you doing?")
// 	// if err != nil{
// 	// 	fmt.Println("Syntax error")
// 	// 	panic(err)
// 	// }

// 	// data := map[string]any{
// 	// 	"name": "John",
// 	// }

// 	// err = tmpl.Execute(os.Stdout, data)
// 	// if err != nil{
// 	// 	fmt.Println("EXECUTION ERROR")
// 	// 	panic(err)
// 	// }

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter your name: ")
// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	// fmt.Printf("Hello %s!", name)

// 	templates := map[string]string{
// 		"welcome":      "Welcome {{.name}}! We are glad you joined!",
// 		"notification": "{{.name}}, you have a new notification: {{.nft}}",
// 		"error":        "Oops! Error occured: {{.errorMessage}}",
// 	}

// 	parsedTemplates := make(map[string]*template.Template)

// 	for name, tmpl := range templates {
// 		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
// 	}

// 	for {

// 		fmt.Println("\nMenu")
// 		fmt.Println("1. Join")
// 		fmt.Println("2. Get notification")
// 		fmt.Println("3. Get error")
// 		fmt.Println("4. Exit")
// 		fmt.Print("Choose an option: ")

// 		choice, _ := reader.ReadString('\n')
// 		choice = strings.TrimSpace(choice)

// 		var data map[string]any
// 		var tmpl *template.Template

// 		switch choice {
// 		case "1":
// 			tmpl = parsedTemplates["welcome"]
// 			data = map[string]any{"name": name}
// 		case "2":
// 			fmt.Print("Enter your notification: ")
// 			notification, _ := reader.ReadString('\n')
// 			notification = strings.TrimSpace(notification)
// 			tmpl = parsedTemplates["notification"]
// 			data = map[string]any{"name": name, "nft": notification}
// 		case "3":
// 			fmt.Print("Enter your error message: ")
// 			errorMessage, _ := reader.ReadString('\n')
// 			errorMessage = strings.TrimSpace(errorMessage)
// 			tmpl = parsedTemplates["error"]
// 			data = map[string]any{"name": name, "errorMessage": errorMessage}
// 		case "4":
// 			fmt.Print("You exited successfully!")
// 			return
// 		default:
// 			fmt.Print("Invalid input! Try again!")
// 			continue
// 		}

// 		err := tmpl.Execute(os.Stdout, data)
// 		if err != nil {
// 			fmt.Println("Error executing template:", err)
// 			return
// 		}
// 	}
// }
