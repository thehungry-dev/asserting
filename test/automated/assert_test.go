package asserting_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"
	"github.com/thehungry-dev/asserting/controls"
)

func TestAssert(t *testing.T) {
	t.Run("Assert", func(t *testing.T) {
		t.Run("Value asserted is false", func(t *testing.T) {
			controller := controls.NewTestController()

			Assert(controller, false)

			t.Run("Helper is invoked", func(t *testing.T) {
				if !controller.IsHelper() {
					t.Fail()
				}
			})

			t.Run("Errors", func(t *testing.T) {
				if !controller.HasError() {
					t.Fail()
				}
			})
		})

		t.Run("Value asserted is true", func(t *testing.T) {
			controller := controls.NewTestController()

			Assert(controller, true)

			t.Run("Helper is invoked", func(t *testing.T) {
				if !controller.IsHelper() {
					t.Fail()
				}
			})

			t.Run("Successful", func(t *testing.T) {
				if controller.HasError() {
					t.Fail()
				}
			})
		})
	})
}
