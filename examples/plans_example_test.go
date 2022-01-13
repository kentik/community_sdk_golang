//go:build examples
// +build examples

//nolint:testpackage,forbidigo
package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlansAPIExample(t *testing.T) {
	assert.NoError(t, runGetAllPlans())
}

func runGetAllPlans() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	plans, err := client.Plans.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(plans)
	fmt.Println()

	return nil
}
