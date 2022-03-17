package testutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func AssertProtoEqual(t testing.TB, expected, actual proto.Message) {
	assert.True(
		t,
		proto.Equal(expected, actual),
		fmt.Sprintf("Protobuf messages are not equal:\nexpected: %v\nactual:  %v", expected, actual),
	)
}
