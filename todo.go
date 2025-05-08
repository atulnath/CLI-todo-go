// This program is a simple command-line To-Do list application.
// It allows users to add, delete, and view tasks.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a single task with an ID and content.
// It is used to store the task information in a slice.
type Task struct {
	ID      int
	Content string
}

// tasks is a slice that holds all the tasks created by the user.
// It is used to manage the tasks in memory.
var tasks []Task
var nextID = 1

// addTask adds a new task to the tasks slice with a unique ID.
// It increments the nextID variable to ensure each task has a unique ID.
func addTask(content string) {
	task := Task{ID: nextID, Content: content}
	tasks = append(tasks, task)
	nextID++
	fmt.Println("Task is added:", content)
}

// deleteTasks removes a task from the tasks slice based on its ID.
// It searches for the task with the given ID and removes it from the slice.
func deleteTasks(id int) {

	if len(tasks) < id {
		fmt.Println("Invalid task ID.")
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted:", task.Content)
			return
		}
	}
}

// listTasks prints all the tasks in the tasks slice to the console.
// It iterates through the slice and prints each task's ID and content.
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tasks {
		fmt.Printf("ID: %d, Content: %s\n", task.ID, task.Content)
	}
}

// createOperation prompts the user to enter a new task content and adds it to the tasks slice.
// It uses a scanner to read user input from the console.
func createOperation(scanner *bufio.Scanner) {
	fmt.Print("Enter task content: ")
	scanner.Scan()
	content := strings.TrimSpace(scanner.Text())
	if content != "" {
		addTask(content)
	} else {
		fmt.Println("Task content cannot be empty.")
	}
}

// deleteOperation prompts the user to enter a task ID to delete and removes it from the tasks slice.
// It uses a scanner to read user input from the console.
func deleteOperation(scanner *bufio.Scanner) {
	fmt.Print("Enter the task ID to delete: ")
	scanner.Scan()
	var deleteID int
	// Use fmt.Sscanf to parse the input string into an integer
	// This is a more efficient way to convert the string to an integer.
	_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d", &deleteID)
	if err == nil {
		deleteTasks(deleteID)
	} else {
		fmt.Println("Invalid task ID.")
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("1. Enter a new task: ")
		fmt.Print("2. See all tasks: ")
		fmt.Print("3. delete a task: ")
		fmt.Print("Choose an option from 1-3: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		if choice == "1" {
			createOperation(scanner)
		} else if choice == "2" {
			listTasks()
		} else if choice == "3" {
			deleteOperation(scanner)
		} else if choice == "4" {
			fmt.Println("Exiting the program.")
			break
		}
	}
}
