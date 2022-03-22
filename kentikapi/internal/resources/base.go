package resources

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/validation"
)

// BaseAPI provides marshall/unmarshall + validation functionality for all resource APIs.
type BaseAPI struct {
	Transport   api_connection.Transport
	LogPayloads bool
}

// GetAndValidate retrieves json at "url", unmarshalls and validates against required fields defined in struct tags of "output"
// output must be pointer to object or nil.
func (b BaseAPI) GetAndValidate(ctx context.Context, url string, output interface{}) error {
	b.logPayload("Kentik API request - GET", url, "")
	responseBody, err := b.Transport.Get(ctx, url)
	b.logPayload("Kentik API response - GET", url, responseBody)
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
// retrieves json at "url", unmarshalls and validates against required fields in "output"
// output must be pointer to object or nil.
func (b BaseAPI) PostAndValidate(ctx context.Context, url string, input interface{}, output interface{}) error {
	if err := validation.CheckRequestRequiredFields("post", input); err != nil {
		return err
	}
	payload, err := json.Marshal(input)
	b.logPayload("Kentik API request - POST", url, payload)
	if err != nil {
		return fmt.Errorf("encode request body: %v", err)
	}

	responseBody, err := b.Transport.Post(ctx, url, payload)
	b.logPayload("Kentik API response - POST", url, responseBody)
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
// retrieves json at "url", unmarshalls and validates against required fields in "output"
// output must be pointer to object or nil.
func (b BaseAPI) UpdateAndValidate(ctx context.Context, url string, input interface{}, output interface{}) error {
	if err := validation.CheckRequestRequiredFields("put", input); err != nil {
		return err
	}
	payload, err := json.Marshal(input)
	b.logPayload("Kentik API request - PUT", url, payload)
	if err != nil {
		return fmt.Errorf("encode request body: %v", err)
	}

	responseBody, err := b.Transport.Put(ctx, url, payload)
	b.logPayload("Kentik API response - PUT", url, responseBody)
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

// DeleteAndValidate retrieves json at "url" unmarshalls and validates
// against required fields defined in struct tags of "output"
// output must be pointer to object or nil.
func (b BaseAPI) DeleteAndValidate(ctx context.Context, url string, output interface{}) error {
	b.logPayload("Kentik API request - DELETE", url, "")
	responseBody, err := b.Transport.Delete(ctx, url)
	b.logPayload("Kentik API response - DELETE", url, responseBody)
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

func (b BaseAPI) logPayload(msg, url string, payload interface{}) {
	if b.LogPayloads {
		if payload == "" {
			log.Printf("%s %s", msg, url)
		} else {
			log.Printf("%s %s - %s", msg, url, payload)
		}
	}
}
