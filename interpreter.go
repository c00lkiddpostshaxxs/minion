package main

import (
	"fmt"
	"strconv"
)

type Interpreter struct {
	vars      map[string]interface{}
	functions map[string]*FunctionDef
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		vars:      make(map[string]interface{}),
		functions: make(map[string]*FunctionDef),
	}
}

func (i *Interpreter) Interpret(program *Program) {
	for _, stmt := range program.Statements {
		i.evalStatement(stmt)
	}
}

func (i *Interpreter) evalStatement(stmt Node) interface{} {
	switch s := stmt.(type) {
	case *VarDecl:
		val := i.evalExpr(s.Value)
		i.vars[s.Name] = val
		return nil
	case *FuncCall:
		return i.evalFuncCall(s)
	case *WhileLoop:
		wl := stmt.(*WhileLoop)
		for i.isTruthy(i.evalExpr(wl.Condition)) {
			for _, s := range wl.Body {
				result := i.evalStatement(s)
				if _, ok := result.(*ReturnValue); ok {
					return result
				}
			}
		}
		return nil
	case *FunctionDef:
		i.functions[s.Name] = s
		return nil
	case *ReturnStmt:
		return &ReturnValue{Value: i.evalExpr(s.Value)}
	case *IfStatement:
		ifstmt := stmt.(*IfStatement)
		cond := i.evalExpr(ifstmt.Condition)
		if i.isTruthy(cond) {
			for _, s := range ifstmt.Body {
				result := i.evalStatement(s)
				if _, ok := result.(*BreakSignal); ok {
					return result
				}
				if _, ok := result.(*ContinueSignal); ok {
					return result
				}
				if _, ok := result.(*ReturnValue); ok {
					return result
				}
			}
		} else if len(ifstmt.Else) > 0 {
			for _, s := range ifstmt.Else {
				result := i.evalStatement(s)
				if _, ok := result.(*BreakSignal); ok {
					return result
				}
				if _, ok := result.(*ContinueSignal); ok {
					return result
				}
				if _, ok := result.(*ReturnValue); ok {
					return result
				}
			}
		}
		return nil
	case *ForLoop:
		fl := stmt.(*ForLoop)
		start := int(i.toFloat(i.evalExpr(fl.Start)))
		end := int(i.toFloat(i.evalExpr(fl.End)))

		shouldBreak := false
		for j := start; j < end; j++ {
			i.vars[fl.Variable] = float64(j)
			for _, s := range fl.Body {
				result := i.evalStatement(s)
				if _, ok := result.(*BreakSignal); ok {
					shouldBreak = true
					break
				}
				if _, ok := result.(*ContinueSignal); ok {
					break
				}
				if _, ok := result.(*ReturnValue); ok {
					return result
				}
			}
			if shouldBreak {
				break
			}
		}
		return nil
	case *BreakStmt:
		return &BreakSignal{}
	case *ContinueStmt:
		return &ContinueSignal{}
	default:
		return nil
	}
	return nil
}

func (i *Interpreter) evalExpr(expr Node) interface{} {
	switch e := expr.(type) {
	case *Number:
		num, _ := strconv.ParseFloat(e.Value, 64)
		return num
	case *Identifier:
		return i.vars[e.Name]
	case *BinaryOp:
		left := i.evalExpr(e.Left)
		right := i.evalExpr(e.Right)
		return i.applyOp(left, e.Op, right)
	case *BoolLit:
		return e.Value
	case *ListLit:
		var result []interface{}
		for _, elem := range e.Elements {
			result = append(result, i.evalExpr(elem))
		}
		return result
	case *FuncCall:
		return i.evalFuncCall(e)
	case *StringLit:
		return e.Value
	case *IndexAccess:
		list := i.evalExpr(e.List)
		index := int(i.toFloat(i.evalExpr(e.Index)))
		if arr, ok := list.([]interface{}); ok {
			if index >= 0 && index < len(arr) {
				return arr[index]
			}
		}
	case *UnaryOp:
		right := i.evalExpr(e.Right)
		if e.Op == "nod" {
			return !i.isTruthy(right)
		}
		return nil
	default:
		return nil
	}
	return nil
}

func (i *Interpreter) applyOp(left interface{}, op string, right interface{}) interface{} {
	l := i.toFloat(left)
	r := i.toFloat(right)

	switch op {
	case "+":
		if ls, ok := left.(string); ok {
			return ls + fmt.Sprintf("%v", right)
		}
		if rs, ok := right.(string); ok {
			return fmt.Sprintf("%v", left) + rs
		}
		l := i.toFloat(left)
		r := i.toFloat(right)
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	case ">":
		return l > r
	case "<":
		return l < r
	case "e":
		return i.isTruthy(left) && i.isTruthy(right)
	case "ro":
		return i.isTruthy(left) || i.isTruthy(right)
	case "==":
		return l == r
	default:
		return 0
	}
}

func (i *Interpreter) toFloat(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return f
	default:
		return 0
	}
}

func (i *Interpreter) isTruthy(val interface{}) bool {
	switch v := val.(type) {
	case float64:
		return v != 0
	case bool:
		return v
	default:
		return false
	}
}

type ReturnValue struct {
	Value interface{}
}

type BreakSignal struct{}
type ContinueSignal struct{}

func (i *Interpreter) evalFuncCall(call *FuncCall) interface{} {
	if fn, exists := i.functions[call.Name]; exists {
		// Save current vars
		oldVars := i.vars
		i.vars = make(map[string]interface{})
		// Copy global vars
		for k, v := range oldVars {
			i.vars[k] = v
		}
		// Set parameters
		for idx, param := range fn.Params {
			if idx < len(call.Args) {
				i.vars[param] = i.evalExpr(call.Args[idx])
			}
		}
		// Execute function body
		var result interface{}
		for _, stmt := range fn.Body {
			result = i.evalStatement(stmt)
			if retVal, ok := result.(*ReturnValue); ok {
				result = retVal.Value
				break
			}
		}
		// Restore vars
		i.vars = oldVars
		return result
	}

	// Built-in functions
	switch call.Name {
	case "stampa":
		for _, arg := range call.Args {
			val := i.evalExpr(arg)
			switch v := val.(type) {
			case []interface{}:
				fmt.Println(v)
			default:
				fmt.Println(val)
			}
		}
	case "numero":
		if len(call.Args) > 0 {
			val := i.evalExpr(call.Args[0])
			return i.toFloat(val)
		}
		return 0
	case "stringo":
		if len(call.Args) > 0 {
			val := i.evalExpr(call.Args[0])
			return fmt.Sprintf("%v", val)
		}
		return ""
	case "lungo":
		if len(call.Args) > 0 {
			val := i.evalExpr(call.Args[0])
			if arr, ok := val.([]interface{}); ok {
				return float64(len(arr))
			}
			if str, ok := val.(string); ok {
				return float64(len(str))
			}
		}
		return 0
	case "suma":
		total := 0.0
		if len(call.Args) > 0 {
			val := i.evalExpr(call.Args[0])
			if arr, ok := val.([]interface{}); ok {
				for _, elem := range arr {
					total += i.toFloat(elem)
				}
			}
		}
		return total
	case "maximo":
		if len(call.Args) > 0 {
			val := i.evalExpr(call.Args[0])
			if arr, ok := val.([]interface{}); ok {
				max := i.toFloat(arr[0])
				for _, elem := range arr {
					if i.toFloat(elem) > max {
						max = i.toFloat(elem)
					}
				}
				return max
			}
		}
		return 0
	}
	return nil
}
