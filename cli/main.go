package main

import "fmt"
import "github.com/avatar29A/lambdacube/untyped"

const Version = "0.0.1v"

func main() {
	identity := untyped.NewExpression(untyped.NewFunc("x", untyped.NewExpression(untyped.NewName("x"))))

	printExpression("identity", identity)

	apply := untyped.NewExpression(
		untyped.NewFunc("x",
			untyped.NewExpression(
				untyped.NewFunc("y",
					untyped.NewExpression(
						untyped.NewApplication(
							untyped.NewName("x"),
							untyped.NewName("y")))))))

	printExpression("apply", apply)

}

func printExpression(name string, expr untyped.LambdaTerm) {
	fmt.Printf("(defun %s ::=\n", name)
	expr.Print()
	fmt.Println(")")
	fmt.Println()

}
