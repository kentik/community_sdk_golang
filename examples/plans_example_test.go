//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlansAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runGetAllPlans())
}

func runGetAllPlans() error {
	client := NewClient()

	fmt.Println("### GET ALL")
	plans, err := client.Plans.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(plans)
	fmt.Println()

	return nil
}
