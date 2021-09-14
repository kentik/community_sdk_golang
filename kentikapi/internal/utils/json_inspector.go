package utils

import (
	"strconv"
	"strings"
	"testing"

	"github.com/antchfx/jsonquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type JSONPayloadInspector struct {
	assert  *assert.Assertions
	require *require.Assertions
	doc     *jsonquery.Node
}

// NewJSONPayloadInspector creates json inspector object for evaluating correctness of json document
// It uses XPath for addressing JSON fields.
func NewJSONPayloadInspector(t *testing.T, jsonString string) *JSONPayloadInspector {
	doc, err := jsonquery.Parse(strings.NewReader(jsonString))
	if err != nil {
		t.Fatalf("%v\n%s\n", err, jsonString)
	}
	return &JSONPayloadInspector{assert: assert.New(t), require: require.New(t), doc: doc}
}

// Exists checks if field at given path exists.
func (i *JSONPayloadInspector) Exists(path string) bool {
	return jsonquery.Find(i.doc, path) != nil
}

// Get returns whatever can be found at given path or nil if nothing is there.
func (i *JSONPayloadInspector) Get(path string) *jsonquery.Node {
	return jsonquery.FindOne(i.doc, path)
}

// GetAll returns list of whatever can be found at given path or nil if nothing is there.
func (i *JSONPayloadInspector) GetAll(path string) []*jsonquery.Node {
	return jsonquery.Find(i.doc, path)
}

// Count returns number of array elements at given path.
func (i *JSONPayloadInspector) Count(path string) int {
	return len(jsonquery.Find(i.doc, path))
}

// String returns text found at given path.
func (i *JSONPayloadInspector) String(path string) string {
	doc := i.Get(path)
	i.require.NotNil(doc)
	return doc.InnerText()
}

// Int returns integer found at given path.
func (i *JSONPayloadInspector) Int(path string) int {
	doc := i.Get(path)
	i.require.NotNil(doc)
	result, err := strconv.Atoi(doc.InnerText())
	i.assert.NoError(err)
	return result
}

// Float returns floating point number found at given path.
//nolint:gomnd
func (i *JSONPayloadInspector) Float(path string) float64 {
	doc := i.Get(path)
	i.require.NotNil(doc)
	result, err := strconv.ParseFloat(doc.InnerText(), 64)
	i.assert.NoError(err)
	return result
}

// Bool returns boolean found at given path.
func (i *JSONPayloadInspector) Bool(path string) bool {
	doc := i.Get(path)
	i.require.NotNil(doc)
	result, err := strconv.ParseBool(doc.InnerText())
	i.assert.NoError(err)
	return result
}
