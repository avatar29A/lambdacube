package untyped

import "fmt"

/*

Lambda Calculus BNF

expression ::= name | func | application
func ::= lambda <name>.<body>
body ::= expression

*/

// TermKind is inner identity of kind of terms. Use to analyze syntax tree.
type TermKind int

// Predifned kind types
const (
	Undefined TermKind = iota
	ExpressionKind
	NameKind
	FuncKind
	BodyKind
	ApplicationKind
)

type LambdaTerm interface {
	Kind() TermKind
	Print()
}

// Term ...
type Term struct {
	kind TermKind
}

func (t Term) Kind() TermKind {
	return t.kind
}

func (t Term) Print() {
	panic("Not implemented")
}

// Expression correspond to expression from BNF
type Expression struct {
	Term
	Terms []LambdaTerm
}

func (e Expression) Print() {
	recursionTermPrint(e, 0)
}

func NewExpression(terms ...LambdaTerm) LambdaTerm {
	expr := Expression{Term{kind: ExpressionKind}, terms}

	return expr
}

// Func correspond to func from BNF
type Func struct {
	Term
	Var  Name
	Body Expression
}

func NewFunc(varName string, body LambdaTerm) LambdaTerm {
	f := Func{Term{FuncKind}, NewName(varName).(Name), body.(Expression)}

	return f
}

// Name corresponds to name from BNF
type Name struct {
	Term
	Name string "contains variable's name"
}

func NewName(name string) LambdaTerm {
	n := Name{Term{NameKind}, name}

	return n
}

type Application struct {
	Term
	Left  LambdaTerm
	Right LambdaTerm
}

func NewApplication(left, right LambdaTerm) LambdaTerm {
	app := Application{Term{ApplicationKind}, left, right}

	return app
}

func recursionTermPrint(t LambdaTerm, level int) {
	tab := spaces(level)
	switch t.Kind() {
	default:
		fmt.Printf("%s%s | %v | level: %d\n", tab, "unknown term", t, level)
	case ExpressionKind:
		for _, expr := range t.(Expression).Terms {
			recursionTermPrint(expr, level+1)
		}
	case NameKind:
		fmt.Println(tab, t.(Name).Name)
	case ApplicationKind:
		app := t.(Application)
		recursionTermPrint(app.Left, level)
		recursionTermPrint(app.Right, level)
	case FuncKind:
		f := t.(Func)
		fmt.Println(tab, "lambda", f.Var.Name, ".")
		recursionTermPrint(f.Body, level+1)
	}
}

func spaces(level int) string {
	s := ""
	for i := 0; i < level; i++ {
		s += " "
	}

	return s
}
