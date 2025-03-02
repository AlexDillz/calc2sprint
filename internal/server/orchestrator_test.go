package server

import "testing"

func TestAddAndGetExpression(t *testing.T) {
	exprText := "2+2*2"
	expr := AddExpression(exprText)
	if expr.ID == 0 {
		t.Error("Expected non-zero expression ID")
	}
	if expr.Expression != exprText {
		t.Errorf("Expected expression %s, got %s", exprText, expr.Expression)
	}
	if expr.Status != "pending" {
		t.Errorf("Expected status 'pending', got %s", expr.Status)
	}
	retExpr, found := GetExpression(expr.ID)
	if !found {
		t.Errorf("Expression with ID %d not found", expr.ID)
	}
	if retExpr.Expression != exprText {
		t.Errorf("Expected expression %s, got %s", exprText, retExpr.Expression)
	}
}
