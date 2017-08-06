package untyped

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func SubstitutionTest(t *testing.T) {

	Convey("Identity: (lambda x.x z) => z", t, func() {
		id := NewFunc("x", NewName("x"))

		z := NewName("z")
		body := Substitution(id.Var, id.Body, z)

		So(z, ShouldResemble, body)
	})

	Convey("Self apply: ((lambda s . (s s))(lambda s . (s s))) ", t, func() {
		selfApply := NewFunc("s", NewApplication(NewName("s"), NewName("s")))

		_ = Substitution(selfApply.Var, selfApply.Body, selfApply)

		So(true, ShouldBeFalse)
		So(selfApply, ShouldResemble, "a")
	})
}
