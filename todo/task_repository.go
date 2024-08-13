package todo

type TaskRepository interface {
	Save(task Task) error
	List() ([]Task, error)
	Delete(id string) error
}
