package untyped

import "fmt"

type Scope struct {
	Vars map[string]int
}

func NewScope() Scope {
	return Scope{make(map[string]int)}
}

// Next returns a new name for variable. For example: x => x0, if x0 has been already taken, then will return x1 and so on.
func (s Scope) Next(name string) string {
	if counter, ok := s.Vars[name]; ok {
		newCounter := counter + 1
		s.Vars[name] = newCounter
		return fmt.Sprintf("%s%d", name, newCounter)
	} else {
		s.Vars[name] = 0
		return name + "0"
	}
}

// Reduct does one step a beta-reduction. Returns false if has wasn't changed.
func Reduct(app Application) (LambdaTerm, bool) {
	switch app.Left.Kind() {
	default:
		panic(fmt.Sprintf("Unknown kind of lambda term.  Kind = %v", app.Left.Kind()))
	case NameKind:
		return app, false
	case ApplicationKind:
		t, changed := Reduct(app.Left.(Application))
		if !changed {
			return app, false
		}

		return Reduct(NewApplication(t, app.Right))
	case FuncKind:
		f := app.Left.(Func)
		return Substitution(f.Var, f.Body, app.Right), true
	}
}

// Substitution replaces func's variable to target lambda term
func Substitution(v Name, s LambdaTerm, t LambdaTerm) LambdaTerm {
	scope := NewScope()

	return substitutionWithScope(v, s, t, scope)
}

func substitutionWithScope(v Name, s LambdaTerm, t LambdaTerm, scope Scope) LambdaTerm {
	switch s.Kind() {
	default:
		panic(fmt.Sprintf("Unsupported kind of term: %v", s.Kind()))
	case NameKind:
		name := s.(Name)
		if v.Name == name.Name {
			return t
		}

		return s

	case FuncKind:
		f := s.(Func)

		if v.Name == f.Var.Name {
			newName := NewName(scope.Next(f.Var.Name))

			fmt.Println("alfa-conversion: ", v.Name, " => ", newName.Name)
			f.Body = substitutionWithScope(f.Var, f.Body, newName, scope)
			f.Var = newName
		}

		f.Body = substitutionWithScope(v, f.Body, t, scope)

		return f

	case ApplicationKind:
		app := s.(Application)
		app.Left = substitutionWithScope(v, app.Left, t, scope)
		app.Right = substitutionWithScope(v, app.Right, t, scope)

		return app
	}
}
