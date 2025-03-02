package calculation

import (
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
		hasError bool
	}{
		// Базовые примеры
		{"2+2", 4, false},
		{"2+2*2", 6, false},
		{"(2+2)*2", 8, false},
		{"10/0", 0, true},
		{"2+3-1", 4, false},
		{"3*4/2", 6, false},
		{"2-2-2", -2, false},

		// Тесты с вложенными скобками
		{"((2+3)*4)", 20, false},         // Два уровня: (2+3) умножить на 4
		{"(2+(3*4))", 14, false},         // Скобки вокруг (3*4) внутри внешних скобок
		{"((1+2)*(3+4))", 21, false},     // Два выражения в отдельных скобках, перемноженные друг на друга
		{"((2+3)*((4-2)+1))", 15, false}, // Вложенные скобки на втором операнде
		// Ошибочные случаи с вложенными скобками
		{"(2+(3*4)", 0, true}, // Пропущена закрывающая скобка
		{"(2+3))*4", 0, true}, // Лишняя закрывающая скобка
	}

	for _, test := range tests {
		result, err := Calc(test.expr)
		if test.hasError {
			if err == nil {
				t.Errorf("expected error for expression %q, got result %v", test.expr, result)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for expression %q: %v", test.expr, err)
			} else if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("expected result %v for expression %q, got %v", test.expected, test.expr, result)
			}
		}
	}
}
