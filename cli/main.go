package main

import "fmt"
import "github.com/avatar29A/lambdacube/untyped"

const Version = "0.0.1v"

func main() {
	identity := untyped.NewFunc("x", untyped.NewName("x"))
	printExpression("identity", identity)

	z := untyped.Substitution(identity.Var, identity.Body, untyped.NewName("z"))
	untyped.PrintLambda(z)
	fmt.Println()

	apply := untyped.NewFunc("x",
		untyped.NewFunc("y",
			untyped.NewApplication(
				untyped.NewName("x"),
				untyped.NewName("y"))))

	printExpression("apply", apply)

	selfApply := untyped.NewFunc("s", untyped.NewApplication(untyped.NewName("s"), untyped.NewName("s")))
	printExpression("selfApply", selfApply)

	selfApply2 := untyped.Substitution(selfApply.Var, selfApply.Body, selfApply)
	printExpression("selfApply2", selfApply2)

	y := untyped.NewFunc("y", untyped.NewFunc("y", untyped.NewName("y")))
	printExpression("y", y)

	y2 := untyped.Substitution(y.Var, y.Body, untyped.NewName("b"))
	printExpression("y2", y2)

	xxx := untyped.NewFunc("x", untyped.NewFunc("x", untyped.NewApplication(untyped.NewName("x"), untyped.NewFunc("x", untyped.NewName("x")))))

	printExpression("xxx", xxx)

	xxx2 := untyped.Substitution(xxx.Var, xxx.Body, untyped.NewName("x"))
	printExpression("xxx2", xxx2)
}

func printExpression(name string, expr untyped.LambdaTerm) {
	fmt.Printf("(defun %s ::=\n", name)
	untyped.PrintLambda(expr)
	fmt.Println(")")
	fmt.Println()
}
