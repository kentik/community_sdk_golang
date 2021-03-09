package testutil

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ParseISO8601Timestamp(t testing.TB, timestamp string) *time.Time {
	const iso8601Layout = "2006-01-02T15:04:05Z0700"
	ts, err := time.Parse(iso8601Layout, timestamp)
	assert.NoError(t, err)

	return &ts
}

func UnmarshalJSONToIf(t testing.TB, jsonString string) interface{} {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	assert.NoError(t, err)
	return data
}

func StringPointer(s string) *string {
	return &s
}
