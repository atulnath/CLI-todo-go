package main

// import "testing"
import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// Task represents a single task with an ID and content.
// It is used to store the task information in a slice.
func resettingTasks() {
	tasks = []Task{}
	nextID = 1
}

// Test functions to test the functionality of the task manager

func TestAddTask(t *testing.T) {
	resettingTasks()
	addTask("Test for addTask run 1")
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for addTask run 1" {
		t.Errorf("Expected 'Test for addTask run 1', got '%s'", tasks[0].Content)
	}
}

// Test for adding multiple tasks
func TestAddMultipleTasks(t *testing.T) {
	resettingTasks()
	addTask("Test for addTask run 1")
	addTask("Test for addTask run 2")
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for addTask run 1" {
		t.Errorf("Expected 'Test for addTask run 1', got '%s'", tasks[0].Content)
	}
	if tasks[1].Content != "Test for addTask run 2" {
		t.Errorf("Expected 'Test for addTask run 2', got '%s'", tasks[1].Content)
	}
}

// Test for deleting a task
func TestDeleteTask(t *testing.T) {
	resettingTasks()
	addTask("Test for deleteTask run 1")
	addTask("Test for deleteTask run 2")
	deleteTasks(1)
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for deleteTask run 2" {
		t.Errorf("Expected 'Test for deleteTask run 2', got '%s'", tasks[0].Content)
	}
}

// Test Listing tasks
// This test checks if the tasks are listed correctly
func TestListTasks(t *testing.T) {
	resettingTasks()
	addTask("Test for listTasks run 1")
	addTask("Test for listTasks run 2")

	// Capture the output
	var output strings.Builder
	writer := bufio.NewWriter(&output)
	listTasks()
	writer.Flush()

	// Check if the output contains the expected tasks
	if !strings.Contains(output.String(), "Test for listTasks run 1") {
		t.Errorf("Expected 'Test for listTasks run 1' in output, got '%s'", output.String())
	}
	if !strings.Contains(output.String(), "Test for listTasks run 2") {
		t.Errorf("Expected 'Test for listTasks run 2' in output, got '%s'", output.String())
	}
}

// Test for deleting an invalid task
// This test checks if the function handles invalid task IDs correctly
func TestDeleteInvalidTask(t *testing.T) {
	resettingTasks()
	addTask("Test for delete invalid task run 1")
	addTask("Test for delete invalid task run 2")
	deleteTasks(3) // Invalid ID
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for delete invalid task run 1" {
		t.Errorf("Expected 'Test for delete invalid task run 1', got '%s'", tasks[0].Content)
	}
	if tasks[1].Content != "Test for delete invalid task run 2" {
		t.Errorf("Expected 'Test for delete invalid task run 2', got '%s'", tasks[1].Content)
	}
}

// Test for creating a task
// This test checks if the function creates a task correctly
func TestCreateOperation(t *testing.T) {
	resettingTasks()
	scanner := bufio.NewScanner(strings.NewReader("Test for create operation run 1\n"))
	createOperation(scanner)
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for create operation run 1" {
		t.Errorf("Expected 'Test for create operation run 1', got '%s'", tasks[0].Content)
	}
}

// Test for creating a task with empty content
// This test checks if the function handles empty content correctly
func TestDeleteOperation(t *testing.T) {
	resettingTasks()
	addTask("Test for delete operation run 1")
	addTask("Test for delete operation run 2")
	scanner := bufio.NewScanner(strings.NewReader("1\n"))
	deleteOperation(scanner)
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for delete operation run 2" {
		t.Errorf("Expected 'Test for delete operation run 2', got '%s'", tasks[0].Content)
	}
}

// Test for deleting a task with an invalid ID
// This test checks if the function handles invalid task IDs correctly
func TestDeleteOperationInvalidID(t *testing.T) {
	resettingTasks()
	addTask("Test for delete operation invalid ID run 1")
	addTask("Test for delete operation invalid ID run 2")
	scanner := bufio.NewScanner(strings.NewReader("3\n"))
	deleteOperation(scanner)
	if len(tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(tasks))
	}
	if tasks[0].Content != "Test for delete operation invalid ID run 1" {
		t.Errorf("Expected 'Test for delete operation invalid ID run 1', got '%s'", tasks[0].Content)
	}
	if tasks[1].Content != "Test for delete operation invalid ID run 2" {
		t.Errorf("Expected 'Test for delete operation invalid ID run 2', got '%s'", tasks[1].Content)
	}
}

// TestMain is the entry point for running tests.
// It resets the tasks before running the tests and exits with the code from the tests.
func TestMain(m *testing.M) {
	// Run the tests
	code := m.Run()

	// Reset the tasks after running tests
	resettingTasks()

	// Exit with the code from the tests
	os.Exit(code)
}
