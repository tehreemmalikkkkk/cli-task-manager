package main

import (
	"fmt"
)

var tasks []Task
var idCounter int = 1

func main() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			var desc string
			fmt.Scanln(&desc)
			tasks = append(tasks, Task{ID: idCounter, Description: desc})
			fmt.Println("Task added!")
			idCounter++
		case 2:
			fmt.Println("Tasks:")
			for _, task := range tasks {
				fmt.Printf("%d. %s\n", task.ID, task.Description)
			}
		case 3:
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
