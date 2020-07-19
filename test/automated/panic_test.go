package asserting_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"
	"github.com/thehungry-dev/asserting/controls"
)

func TestAssertPanic(t *testing.T) {
	t.Run("Assert Panic", func(t *testing.T) {
		t.Run("Function asserted doesn't panic", func(t *testing.T) {
			controller := controls.NewTestController()

			Assert.Panic(controller, controls.NotPanic)

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

		t.Run("Function asserted panicks", func(t *testing.T) {
			controller := controls.NewTestController()

			Assert.Panic(controller, controls.WillPanic)

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
