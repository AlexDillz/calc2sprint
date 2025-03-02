package server

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/AlexDillz/calc2sprint/pkg/calculation"
)

type DecomposedTask struct {
	ID            int    `json:"id"`
	Arg1          string `json:"arg1"`
	Arg2          string `json:"arg2"`
	Operation     string `json:"operation"`
	OperationTime int    `json:"operation_time"`
}

var (
	taskQueue   []DecomposedTask
	taskQueueMu sync.Mutex
	nextTaskID  = 1
)

func DecomposeExpression(expression string) ([]DecomposedTask, error) {
	tokens, err := calculation.Tokenize(expression)
	if err != nil {
		return nil, err
	}

	var tasks []DecomposedTask
	if len(tokens) == 3 {
		op := tokens[1]
		task := DecomposedTask{
			ID:            nextTaskID,
			Arg1:          tokens[0],
			Arg2:          tokens[2],
			Operation:     op,
			OperationTime: getOperationTime(op),
		}
		nextTaskID++
		tasks = append(tasks, task)
		EnqueueTask(task)
		return tasks, nil
	}

	task := DecomposedTask{
		ID:            nextTaskID,
		Arg1:          expression,
		Arg2:          "",
		Operation:     "",
		OperationTime: 0,
	}
	nextTaskID++
	tasks = append(tasks, task)
	EnqueueTask(task)
	return tasks, nil
}

func EnqueueTask(task DecomposedTask) {
	taskQueueMu.Lock()
	defer taskQueueMu.Unlock()
	taskQueue = append(taskQueue, task)
	log.Printf("Task enqueued: %+v", task)
}

func GetNextTask() (DecomposedTask, bool) {
	taskQueueMu.Lock()
	defer taskQueueMu.Unlock()
	if len(taskQueue) == 0 {
		return DecomposedTask{}, false
	}
	task := taskQueue[0]
	taskQueue = taskQueue[1:]
	return task, true
}

func getOperationTime(op string) int {
	switch op {
	case "+":
		return getEnvInt("TIME_ADDITION_MS", 2000)
	case "-":
		return getEnvInt("TIME_SUBTRACTION_MS", 2000)
	case "*":
		return getEnvInt("TIME_MULTIPLICATION_MS", 3000)
	case "/":
		return getEnvInt("TIME_DIVISIONS_MS", 3000)
	default:
		return 1000
	}
}

func getEnvInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr != "" {
		if val, err := strconv.Atoi(valStr); err == nil {
			return val
		}
	}
	return defaultVal
}
