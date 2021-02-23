//+build examples

package examples

import (
	"context"
	"fmt"
	"runtime/debug"
	"testing"
)

func TestPlansAPIExample(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Fatal(err)
		}
	}()

	runGetAllPlans()
}

func runGetAllPlans() {
	client := NewClient()

	fmt.Println("### GET ALL")
	plans, err := client.Plans.GetAll(context.Background())
	PanicOnError(err)
	PrettyPrint(plans)
	fmt.Println()
}
