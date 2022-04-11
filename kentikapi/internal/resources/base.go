package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/validation"
)

// BaseAPI provides marshall/unmarshall + validation functionality for all resource APIs.
type BaseAPI struct {
	Transport   api_connection.Transport
	LogPayloads bool
}

// GetAndValidate retrieves JSON at "path", unmarshalls and validates it against required fields
// defined in struct tags of "output".
// Output must be pointer to object or nil.
func (b BaseAPI) GetAndValidate(ctx context.Context, path string, output interface{}) error {
	b.logRequest(http.MethodGet, path, "")
	responseBody, err := b.Transport.Get(ctx, path)
	b.logResponse(http.MethodGet, path, string(responseBody), err)
	if err != nil {
		return err
	}

	if output == nil {
		return nil
	}

	if err = json.Unmarshal(responseBody, &output); err != nil {
		return fmt.Errorf("decode response body: %v", err)
	}

	if err = validation.CheckResponseRequiredFields("get", output); err != nil {
		return err
	}

	return nil
}

// PostAndValidate validates input against required fields defined in struct tags of "input",
// retrieves JSON at "path", unmarshalls and validates against required fields in "output"
// Output must be pointer to object or nil.
//nolint: dupl
func (b BaseAPI) PostAndValidate(ctx context.Context, path string, input interface{}, output interface{}) error {
	if err := validation.CheckRequestRequiredFields("post", input); err != nil {
		return err
	}
	payload, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("encode request body: %v", err)
	}

	b.logRequest(http.MethodPost, path, string(payload))
	responseBody, err := b.Transport.Post(ctx, path, payload)
	b.logResponse(http.MethodPost, path, string(responseBody), err)
	if err != nil {
		return err
	}

	if output == nil {
		return nil
	}

	if err = json.Unmarshal(responseBody, &output); err != nil {
		return fmt.Errorf("decode response body: %v", err)
	}

	if err = validation.CheckResponseRequiredFields("post", output); err != nil {
		return err
	}

	return nil
}

// UpdateAndValidate validates input against required fields defined in struct tags of "input",
// retrieves JSON at "path", unmarshalls and validates against required fields in "output"
// Output must be pointer to object or nil.
//nolint: dupl
func (b BaseAPI) UpdateAndValidate(ctx context.Context, path string, input interface{}, output interface{}) error {
	if err := validation.CheckRequestRequiredFields("put", input); err != nil {
		return err
	}
	payload, err := json.Marshal(input)
	if err != nil {
		return fmt.Errorf("encode request body: %v", err)
	}

	b.logRequest(http.MethodPut, path, string(payload))
	responseBody, err := b.Transport.Put(ctx, path, payload)
	b.logResponse(http.MethodPut, path, string(responseBody), err)
	if err != nil {
		return err
	}

	if output == nil {
		return nil
	}

	if err = json.Unmarshal(responseBody, &output); err != nil {
		return fmt.Errorf("decode response body: %v", err)
	}

	if err = validation.CheckResponseRequiredFields("put", output); err != nil {
		return err
	}

	return nil
}

// DeleteAndValidate retrieves JSON at "path", unmarshalls and validates
// against required fields defined in struct tags of "output"
// Output must be pointer to object or nil.
func (b BaseAPI) DeleteAndValidate(ctx context.Context, path string, output interface{}) error {
	b.logRequest(http.MethodDelete, path, "")
	responseBody, err := b.Transport.Delete(ctx, path)
	b.logResponse(http.MethodDelete, path, string(responseBody), err)
	if err != nil {
		return err
	}

	if output == nil {
		return nil
	}

	if err = json.Unmarshal(responseBody, &output); err != nil {
		return fmt.Errorf("decode response body: %v", err)
	}

	if err = validation.CheckResponseRequiredFields("get", output); err != nil {
		return err
	}

	return nil
}

func (b BaseAPI) logRequest(method, path, payloadJSON string) {
	if b.LogPayloads {
		log.Printf("Kentik API request: method=%v path=%v payload=%v", method, path, sanitizePayload(payloadJSON))
	}
}

func (b BaseAPI) logResponse(method, path, payloadJSON string, err error) {
	if b.LogPayloads {
		log.Printf(
			"Kentik API response: method=%v path=%v payload=%v error=%v",
			method, path, sanitizePayload(payloadJSON), err,
		)
	}
}

func sanitizePayload(payloadJSON string) string {
	if payloadJSON == "" {
		return "<empty>"
	}

	p := compactJSON([]byte(payloadJSON))

	const maxLoggedPayloadSize = 10000
	if len(p) > maxLoggedPayloadSize {
		return fmt.Sprintf("<size: %v bytes>", len(p))
	}
	return p
}

// compactJSON makes unindented JSON from given JSON, if possible.
func compactJSON(jsonS []byte) string {
	var buffer bytes.Buffer
	err := json.Compact(&buffer, jsonS)
	if err != nil {
		return string(jsonS)
	}

	return buffer.String()
}
