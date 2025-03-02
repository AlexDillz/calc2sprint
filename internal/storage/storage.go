package storage

import (
	"errors"
	"sync"
)

type Expression struct {
	ID         int
	Expression string
	Status     string
	Result     *float64
}

type Task struct {
	ID            int
	Arg1          string
	Arg2          string
	Operation     string
	OperationTime int
}

type Storage interface {
	SaveExpression(expr Expression) error
	GetExpression(id int) (Expression, error)
	ListExpressions() ([]Expression, error)
	UpdateExpression(expr Expression) error

	EnqueueTask(task Task) error
	DequeueTask() (Task, error)
	ListTasks() ([]Task, error)
}

type InMemoryStorage struct {
	exprMu  sync.RWMutex
	exprMap map[int]Expression

	taskMu    sync.Mutex
	taskQueue []Task
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		exprMap:   make(map[int]Expression),
		taskQueue: make([]Task, 0),
	}
}

func (s *InMemoryStorage) SaveExpression(expr Expression) error {
	s.exprMu.Lock()
	defer s.exprMu.Unlock()
	if _, exists := s.exprMap[expr.ID]; exists {
		return errors.New("expression already exists")
	}
	s.exprMap[expr.ID] = expr
	return nil
}

func (s *InMemoryStorage) GetExpression(id int) (Expression, error) {
	s.exprMu.RLock()
	defer s.exprMu.RUnlock()
	expr, exists := s.exprMap[id]
	if !exists {
		return Expression{}, errors.New("expression not found")
	}
	return expr, nil
}

func (s *InMemoryStorage) ListExpressions() ([]Expression, error) {
	s.exprMu.RLock()
	defer s.exprMu.RUnlock()
	list := make([]Expression, 0, len(s.exprMap))
	for _, expr := range s.exprMap {
		list = append(list, expr)
	}
	return list, nil
}

func (s *InMemoryStorage) UpdateExpression(expr Expression) error {
	s.exprMu.Lock()
	defer s.exprMu.Unlock()
	if _, exists := s.exprMap[expr.ID]; !exists {
		return errors.New("expression not found")
	}
	s.exprMap[expr.ID] = expr
	return nil
}

func (s *InMemoryStorage) EnqueueTask(task Task) error {
	s.taskMu.Lock()
	defer s.taskMu.Unlock()
	s.taskQueue = append(s.taskQueue, task)
	return nil
}

func (s *InMemoryStorage) DequeueTask() (Task, error) {
	s.taskMu.Lock()
	defer s.taskMu.Unlock()
	if len(s.taskQueue) == 0 {
		return Task{}, errors.New("no tasks available")
	}
	task := s.taskQueue[0]
	s.taskQueue = s.taskQueue[1:]
	return task, nil
}

func (s *InMemoryStorage) ListTasks() ([]Task, error) {
	s.taskMu.Lock()
	defer s.taskMu.Unlock()
	tasks := make([]Task, len(s.taskQueue))
	copy(tasks, s.taskQueue)
	return tasks, nil
}
