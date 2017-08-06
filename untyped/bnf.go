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
}

// Term ...
type Term struct {
	kind TermKind
}

func (t Term) Kind() TermKind {
	return t.kind
}

// Func correspond to func from BNF
type Func struct {
	Term
	Var  Name
	Body LambdaTerm
}

func NewFunc(varName string, body LambdaTerm) Func {
	f := Func{Term{FuncKind}, NewName(varName), body}

	return f
}

// Name corresponds to name from BNF
type Name struct {
	Term
	Name string "contains variable's name"
}

func NewName(name string) Name {
	n := Name{Term{NameKind}, name}

	return n
}

type Application struct {
	Term
	Left  LambdaTerm
	Right LambdaTerm
}

func NewApplication(left, right LambdaTerm) Application {
	app := Application{Term{ApplicationKind}, left, right}

	return app
}

func PrintLambda(t LambdaTerm) {
	recursionTermPrint(t, 0)
}

func recursionTermPrint(t LambdaTerm, level int) {
	tab := spaces(level)
	switch t.Kind() {
	default:
		fmt.Printf("%s%s | %v | level: %d\n", tab, "unknown term", t, level)
	case NameKind:
		fmt.Println(tab, t.(Name).Name)
	case ApplicationKind:
		app := t.(Application)
		fmt.Print("(")
		recursionTermPrint(app.Left, level)
		recursionTermPrint(app.Right, level)
		fmt.Print(")")
	case FuncKind:
		f := t.(Func)
		fmt.Println(tab, "(lambda", f.Var.Name, ".")
		recursionTermPrint(f.Body, level+1)
		fmt.Print(")")
	}
}

func spaces(level int) string {
	s := ""
	for i := 0; i < level; i++ {
		s += " "
	}

	return s
}
