package main

type Node interface{}

type Program struct {
	Statements []Node
}

type VarDecl struct {
	Name  string
	Value Node
}

type FuncCall struct {
	Name string
	Args []Node
}

type Number struct {
	Value string
}

type Identifier struct {
	Name string
}

type BinaryOp struct {
	Left  Node
	Op    string
	Right Node
}

type IfStatement struct {
	Condition Node
	Body      []Node
	Else      []Node
}

type WhileLoop struct {
	Condition Node
	Body      []Node
}

type BoolLit struct {
	Value bool
}

type ListLit struct {
	Elements []Node
}

type IndexAccess struct {
	List  Node
	Index Node
}

type FunctionDef struct {
	Name   string
	Params []string
	Body   []Node
}

type ReturnStmt struct {
	Value Node
}

type StringLit struct {
	Value string
}

type UnaryOp struct {
	Op    string
	Right Node
}

type ForLoop struct {
	Variable string
	Start    Node
	End      Node
	Body     []Node
}

type BreakStmt struct{}

type ContinueStmt struct{}
