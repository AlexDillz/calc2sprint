package storage

import (
	"testing"
)

func TestInMemoryStorageExpressions(t *testing.T) {
	s := NewInMemoryStorage()

	expr := Expression{
		ID:         1,
		Expression: "2+2",
		Status:     "pending",
	}

	if err := s.SaveExpression(expr); err != nil {
		t.Fatalf("failed to save expression: %v", err)
	}

	if err := s.SaveExpression(expr); err == nil {
		t.Error("expected error when saving duplicate expression, got nil")
	}

	got, err := s.GetExpression(1)
	if err != nil {
		t.Fatalf("failed to get expression: %v", err)
	}
	if got.Expression != "2+2" {
		t.Errorf("expected expression '2+2', got %s", got.Expression)
	}

	newResult := 4.0
	expr.Status = "done"
	expr.Result = &newResult
	if err := s.UpdateExpression(expr); err != nil {
		t.Fatalf("failed to update expression: %v", err)
	}

	updated, err := s.GetExpression(1)
	if err != nil {
		t.Fatalf("failed to get expression: %v", err)
	}
	if updated.Status != "done" {
		t.Errorf("expected status 'done', got %s", updated.Status)
	}
	if updated.Result == nil || *updated.Result != 4.0 {
		t.Errorf("expected result 4.0, got %v", updated.Result)
	}
}

func TestInMemoryStorageTasks(t *testing.T) {
	s := NewInMemoryStorage()

	task1 := Task{
		ID:            1,
		Arg1:          "2",
		Arg2:          "2",
		Operation:     "+",
		OperationTime: 1000,
	}

	task2 := Task{
		ID:            2,
		Arg1:          "4",
		Arg2:          "2",
		Operation:     "/",
		OperationTime: 1000,
	}

	if err := s.EnqueueTask(task1); err != nil {
		t.Fatalf("failed to enqueue task1: %v", err)
	}
	if err := s.EnqueueTask(task2); err != nil {
		t.Fatalf("failed to enqueue task2: %v", err)
	}

	tasks, err := s.ListTasks()
	if err != nil {
		t.Fatalf("failed to list tasks: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("expected 2 tasks, got %d", len(tasks))
	}

	dTask, err := s.DequeueTask()
	if err != nil {
		t.Fatalf("failed to dequeue task: %v", err)
	}
	if dTask.ID != 1 {
		t.Errorf("expected task ID 1, got %d", dTask.ID)
	}

	dTask2, err := s.DequeueTask()
	if err != nil {
		t.Fatalf("failed to dequeue task: %v", err)
	}
	if dTask2.ID != 2 {
		t.Errorf("expected task ID 2, got %d", dTask2.ID)
	}

	_, err = s.DequeueTask()
	if err == nil {
		t.Error("expected error when dequeuing from empty queue, got nil")
	}
}
