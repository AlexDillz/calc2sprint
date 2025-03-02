package server

import (
	"log"
	"sync"
)

type Expression struct {
	ID         int      `json:"id"`
	Expression string   `json:"expression"`
	Status     string   `json:"status"`
	Result     *float64 `json:"result"`
}

var (
	expressions   = make(map[int]*Expression)
	expressionsMu sync.RWMutex
	nextExprID    = 1
)

func AddExpression(expr string) *Expression {
	expressionsMu.Lock()
	defer expressionsMu.Unlock()
	e := &Expression{
		ID:         nextExprID,
		Expression: expr,
		Status:     "pending",
		Result:     nil,
	}
	nextExprID++
	expressions[e.ID] = e
	log.Printf("Expression added: %+v", e)
	return e
}

func GetExpressions() []*Expression {
	expressionsMu.RLock()
	defer expressionsMu.RUnlock()
	var list []*Expression
	for _, e := range expressions {
		list = append(list, e)
	}
	return list
}

func GetExpression(id int) (*Expression, bool) {
	expressionsMu.RLock()
	defer expressionsMu.RUnlock()
	e, ok := expressions[id]
	return e, ok
}
