package utils

import (
	"strconv"
	"strings"
	"testing"

	"github.com/antchfx/jsonquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// For testing if json contains required fields set to required values
type xpath struct {
	assert  *assert.Assertions
	require *require.Assertions
	doc     *jsonquery.Node
}

func NewJSONPayloadXPath(t *testing.T, jsonString string) *xpath {
	doc, err := jsonquery.Parse(strings.NewReader(jsonString))
	if err != nil {
		t.Fatalf("%v\n%s\n", err, jsonString)
	}
	return &xpath{assert: assert.New(t), require: require.New(t), doc: doc}
}

func (v *xpath) Get(path string) *jsonquery.Node {
	return jsonquery.FindOne(v.doc, path)
}

func (v *xpath) Count(path string) int {
	return len(jsonquery.Find(v.doc, path))
}

func (v *xpath) String(path string) string {
	doc := v.Get(path)
	v.require.NotNil(doc)
	return doc.InnerText()
}

func (v *xpath) Int(path string) int {
	doc := v.Get(path)
	v.require.NotNil(doc)
	i, err := strconv.Atoi(doc.InnerText())
	v.assert.NoError(err)
	return i
}

func (v *xpath) Bool(path string) bool {
	doc := v.Get(path)
	v.require.NotNil(doc)
	b, err := strconv.ParseBool(doc.InnerText())
	v.assert.NoError(err)
	return b
}
