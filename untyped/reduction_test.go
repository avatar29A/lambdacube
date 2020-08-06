package untyped

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name     Name
	expr     LambdaTerm
	apply    LambdaTerm
	expected LambdaTerm
}

func TestSubstitution(t *testing.T) {

	cases := []testCase{
		{
			name:     NewName("x"),
			expr:     NewFunc("x", NewName("x")), // Identity: (lambda x.x z) => z,
			apply:    NewName("z"),
			expected: NewName("z"),
		},

		{
			name:     NewName("selfApply"),
			expr:     NewFunc("s", NewApplication(NewName("s"), NewName("s"))), // Self apply: ((lambda s . (s s))(lambda s . (s s)))
			apply:    NewFunc("s", NewApplication(NewName("s"), NewName("s"))),
			expected: NewFunc("s", NewApplication(NewName("s"), NewName("s"))),
		},
	}

	for _, test := range cases {
		result := Substitution(test.name, test.expr, test.apply)

		assert.ObjectsAreEqual(test.expected, result)
	}
}
