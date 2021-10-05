//+build examples

package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

func TestSavedFiltersAPIExample(t *testing.T) {
	assert := assert.New(t)
	assert.NoError(runGetAllSavedFilters())
	assert.NoError(runCRUDSavedFilters())
}

func runGetAllSavedFilters() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### GET ALL")
	savedFilters, err := client.SavedFilters.GetAll(context.Background())
	if err != nil {
		return err
	}
	PrettyPrint(savedFilters)
	fmt.Println()

	return nil
}

func runCRUDSavedFilters() error {
	client, err := NewClient()
	if err != nil {
		return err
	}

	fmt.Println("### CREATE")
	newSavedFilter := models.SavedFilter{
		FilterName:        "New_Filter_test",
		FilterDescription: "description of freshly created saved filter",
		Filters: models.Filters{
			Connector: "Any",
			FilterGroups: []models.FilterGroups{
				{
					Connector: "Any",
					Not:       false,
					Filters: []models.Filter{
						{
							FilterField: "dst_as",
							FilterValue: "82",
							Operator:    "=",
						},
					},
				},
			},
		},
	}
	savedFilter, err := client.SavedFilters.Create(context.Background(), newSavedFilter)
	if err != nil {
		return err
	}
	PrettyPrint(savedFilter)
	fmt.Println()

	fmt.Println("### GET")
	savedFilter, err = client.SavedFilters.Get(context.Background(), savedFilter.ID)
	if err != nil {
		return err
	}
	PrettyPrint(savedFilter)
	fmt.Println()

	fmt.Println("### UPDATE")
	savedFilter.FilterDescription = "This description was updated just now."
	savedFilter.Filters.Connector = "All"
	savedFilter, err = client.SavedFilters.Update(context.Background(), *savedFilter)
	if err != nil {
		return err
	}
	PrettyPrint(savedFilter)
	fmt.Println()

	fmt.Println("## DELETE")
	err = client.SavedFilters.Detete(context.Background(), savedFilter.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Succesfully deleted Saved Filter: %v\n", savedFilter.ID)
	fmt.Println()

	return nil
}
