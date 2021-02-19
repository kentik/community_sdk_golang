package utils

import (
	"strconv"
	"strings"
	"testing"

	"github.com/antchfx/jsonquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type jsonPayloadInspector struct {
	assert  *assert.Assertions
	require *require.Assertions
	doc     *jsonquery.Node
}

// NewJSONPayloadInspector creates json inspector object for evaluating correctness of json document
// It uses XPath for addressing JSON fields
func NewJSONPayloadInspector(t *testing.T, jsonString string) *jsonPayloadInspector {
	doc, err := jsonquery.Parse(strings.NewReader(jsonString))
	if err != nil {
		t.Fatalf("%v\n%s\n", err, jsonString)
	}
	return &jsonPayloadInspector{assert: assert.New(t), require: require.New(t), doc: doc}
}

// Exists checks if field at given path exists
func (i *jsonPayloadInspector) Exists(path string) bool {
	return jsonquery.Find(i.doc, path) != nil
}

// Get returns whatever can be found at given path or nil if nothing is there
func (i *jsonPayloadInspector) Get(path string) *jsonquery.Node {
	return jsonquery.FindOne(i.doc, path)
}

// GetAll returns list of whatever can be found at given path or nil if nothing is there
func (i *jsonPayloadInspector) GetAll(path string) []*jsonquery.Node {
	return jsonquery.Find(i.doc, path)
}

// Count returns number of array elements at given path
func (i *jsonPayloadInspector) Count(path string) int {
	return len(jsonquery.Find(i.doc, path))
}

// String returns text found at given path
func (i *jsonPayloadInspector) String(path string) string {
	doc := i.Get(path)
	i.require.NotNil(doc)
	return doc.InnerText()
}

// Int returns integer found at given path
func (i *jsonPayloadInspector) Int(path string) int {
	doc := i.Get(path)
	i.require.NotNil(doc)
	result, err := strconv.Atoi(doc.InnerText())
	i.assert.NoError(err)
	return result
}

// Bool returns boolean found at given path
func (i *jsonPayloadInspector) Bool(path string) bool {
	doc := i.Get(path)
	i.require.NotNil(doc)
	result, err := strconv.ParseBool(doc.InnerText())
	i.assert.NoError(err)
	return result
}
