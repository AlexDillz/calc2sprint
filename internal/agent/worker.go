package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AlexDillz/calc2sprint/pkg/calculation"
)

type Task struct {
	ID            int    `json:"id"`
	Arg1          string `json:"arg1"`
	Arg2          string `json:"arg2"`
	Operation     string `json:"operation"`
	OperationTime int    `json:"operation_time"`
}

type TaskResponse struct {
	Task Task `json:"task"`
}

type TaskResult struct {
	ID     int     `json:"id"`
	Result float64 `json:"result"`
}

func fetchTask() (*Task, error) {
	resp, err := http.Get("http://localhost:8080/internal/task")
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе задачи: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный статус запроса задачи: %d", resp.StatusCode)
	}

	var taskResp TaskResponse
	if err := json.NewDecoder(resp.Body).Decode(&taskResp); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа задачи: %v", err)
	}
	return &taskResp.Task, nil
}

func postResult(result TaskResult) error {
	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("ошибка при кодировании результата в JSON: %v", err)
	}
	resp, err := http.Post("http://localhost:8080/internal/task", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("ошибка при отправке результата: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось отправить результат, статус: %d", resp.StatusCode)
	}
	return nil
}

func worker() {
	for {
		task, err := fetchTask()
		if err != nil {
			log.Printf("Ошибка получения задачи: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}
		if task == nil {

			time.Sleep(1 * time.Second)
			continue
		}

		time.Sleep(time.Duration(task.OperationTime) * time.Millisecond)

		var result float64
		if task.Operation == "" {
			result, err = calculation.Calc(task.Arg1)
			if err != nil {
				log.Printf("Ошибка вычисления выражения '%s': %v", task.Arg1, err)
				continue
			}
		} else {
			var num1, num2 float64
			num1, err = strconv.ParseFloat(task.Arg1, 64)
			if err != nil {
				log.Printf("Ошибка преобразования Arg1 '%s': %v", task.Arg1, err)
				continue
			}
			num2, err = strconv.ParseFloat(task.Arg2, 64)
			if err != nil {
				log.Printf("Ошибка преобразования Arg2 '%s': %v", task.Arg2, err)
				continue
			}
			switch task.Operation {
			case "+":
				result = num1 + num2
			case "-":
				result = num1 - num2
			case "*":
				result = num1 * num2
			case "/":
				if num2 == 0 {
					log.Printf("Задача %d: деление на ноль", task.ID)
					continue
				}
				result = num1 / num2
			default:
				log.Printf("Задача %d: неизвестная операция '%s'", task.ID, task.Operation)
				continue
			}
		}

		err = postResult(TaskResult{ID: task.ID, Result: result})
		if err != nil {
			log.Printf("Задача %d: ошибка отправки результата: %v", task.ID, err)
		} else {
			log.Printf("Задача %d успешно обработана, результат: %f", task.ID, result)
		}
	}
}
