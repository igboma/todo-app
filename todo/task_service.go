package todo

import "errors"

type Task interface {
	ID() string
	Name() string
	PerformTask() error
}

type TaskService interface {
	AddTask(task Task) error
	ListTasks() ([]Task, error)
	DeleteTask(id string) error
}

type DatabaseTask struct {
	TaskID   string
	TaskName string
}

func (t *DatabaseTask) ID() string {
	return t.TaskID
}

func (t *DatabaseTask) Name() string {
	return t.TaskName
}

func (t *DatabaseTask) PerformTask() error {
	// Implementation for database task
	return nil
}

type TaskServiceImpl struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &TaskServiceImpl{repo: repo}
}

func (s *TaskServiceImpl) AddTask(task Task) error {
	return s.repo.Save(task)
}

func (s *TaskServiceImpl) ListTasks() ([]Task, error) {
	return s.repo.List()
}

func (s *TaskServiceImpl) DeleteTask(id string) error {
	if id == "" || id == "0" {
		return errors.New("invalid task ID")
	}
	return s.repo.Delete(id)
}
