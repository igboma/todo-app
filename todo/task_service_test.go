package todo_test

import (
	"testing"

	"github.com/igboma/todo-app/mocks"
	"github.com/igboma/todo-app/todo"
	"github.com/stretchr/testify/assert"
)

func TestAddTask(t *testing.T) {
	// Initialize the mock TaskRepository
	mockRepo := new(mocks.TaskRepository)

	// Create the TaskService with the mocked repository
	service := todo.NewTaskService(mockRepo)

	// Initialize the mock Task
	mockTask := new(mocks.Task)

	// Expect the Save method to be called with the mockTask
	mockRepo.On("Save", mockTask).Return(nil)

	// Invoke the AddTask method with the mockTask
	err := service.AddTask(mockTask)

	// Assert that no error was returned
	assert.NoError(t, err)

	// Verify that all expectations were met
	mockRepo.AssertExpectations(t)
}

func TestListTasks(t *testing.T) {
	// Initialize the mock TaskRepository
	mockRepo := new(mocks.TaskRepository)

	// Create the TaskService with the mocked repository
	service := todo.NewTaskService(mockRepo)

	// Create a list of mock tasks to be returned by the repository
	mockTask1 := new(mocks.Task)
	mockTask2 := new(mocks.Task)
	tasks := []todo.Task{mockTask1, mockTask2}

	// Expect the List method to be called and return the list of tasks
	mockRepo.On("List").Return(tasks, nil)

	// Invoke the ListTasks method
	result, err := service.ListTasks()

	// Assert that no error was returned
	assert.NoError(t, err)

	// Assert that the result matches the expected list of tasks
	assert.Equal(t, tasks, result)

	// Verify that all expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	// Initialize the mock TaskRepository
	mockRepo := new(mocks.TaskRepository)

	// Create the TaskService with the mocked repository
	service := todo.NewTaskService(mockRepo)

	// Define the task ID to be deleted
	taskID := "1"

	// Expect the Delete method to be called with the task ID
	mockRepo.On("Delete", taskID).Return(nil)

	// Invoke the DeleteTask method
	err := service.DeleteTask(taskID)

	// Assert that no error was returned
	assert.NoError(t, err)

	// Verify that all expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_InvalidID(t *testing.T) {
	// Initialize the mock TaskRepository
	mockRepo := new(mocks.TaskRepository)

	// Create the TaskService with the mocked repository
	service := todo.NewTaskService(mockRepo)

	// Define invalid task IDs
	invalidIDs := []string{"", "0"}

	for _, taskID := range invalidIDs {
		// Expect the Delete method NOT to be called for invalid IDs
		mockRepo.On("Delete", taskID).Return(nil).Maybe()

		// Invoke the DeleteTask method with invalid IDs
		err := service.DeleteTask(taskID)

		// Assert that an error is returned
		assert.Error(t, err, "Expected an error when deleting task with ID: %s", taskID)

		// Verify that Delete was not called
		mockRepo.AssertNotCalled(t, "Delete", taskID)
	}

	// Verify that all expectations were met
	mockRepo.AssertExpectations(t)
}
