package utils_test

import (
	"strconv"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestConvertOrNoneReturnsValue(t *testing.T) {
	t.Parallel()

	// given
	input := new(string)
	*input = "42"
	var output *int

	// when
	err := utils.ConvertOrNone(input, strconv.Atoi, &output)

	// then
	assert := assert.New(t)
	assert.NoError(err)
	assert.NotNil(output)
	assert.Equal(*output, 42)
}

func TestConvertOrNoneReturnsNil(t *testing.T) {
	t.Parallel()

	// given
	var input *string
	var output *int

	// when
	err := utils.ConvertOrNone(input, strconv.Atoi, &output)

	// then
	assert := assert.New(t)
	assert.NoError(err)
	assert.Nil(output)
}

func TestConvertOrNoneReturnsError(t *testing.T) {
	t.Parallel()

	// given
	input := new(string)
	*input = "0xFF"
	var output *int

	// when
	err := utils.ConvertOrNone(input, strconv.Atoi, &output)

	// then
	assert := assert.New(t)
	assert.Error(err)
}

func TestConvertListSuccess(t *testing.T) {
	t.Parallel()

	// given
	input := [...]string{"-13", "22", "742"}
	var output []int

	// when
	err := utils.ConvertList(input, strconv.Atoi, &output)

	// then
	assert := assert.New(t)
	assert.NoError(err)
	assert.Equal(output[0], -13)
	assert.Equal(output[1], 22)
	assert.Equal(output[2], 742)
}

func TestConvertListError(t *testing.T) {
	t.Parallel()

	// given
	input := []string{"42", "0xFF"}
	var output []int

	// when
	err := utils.ConvertList(input, strconv.Atoi, &output)

	// then
	assert := assert.New(t)
	assert.Error(err)
}
