package asserting_test

import (
	"testing"

	. "github.com/thehungry-dev/asserting"
	"github.com/thehungry-dev/asserting/controls"
)

func TestAssertPanicMsg(t *testing.T) {
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
			t.Run("Error message matches", func(t *testing.T) {
				controller := controls.NewTestController()

				Assert.PanicMsg(controller, controls.WillPanic, controls.MatchesMsg)

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

			t.Run("Error message doesn't match", func(t *testing.T) {
				controller := controls.NewTestController()

				Assert.PanicMsg(controller, controls.WillPanic, controls.RejectMsg)

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
		})
	})
}
