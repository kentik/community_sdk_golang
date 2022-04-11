package kentikapi_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

//nolint:gosec
const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
	dummyAuthEmail  = "email@example.com"
	dummyAuthToken  = "api-test-token"
	testUserID      = "145999"
)

type object = map[string]interface{}

func TestClient_GetAllUsers(t *testing.T) {
	tests := []struct {
		name           string
		responseCode   int
		responseBody   string
		expectedResult []models.User
		expectedError  bool
	}{
		{
			name:          "status bad request",
			responseCode:  http.StatusBadRequest,
			responseBody:  `{"error":"Bad Request"}`,
			expectedError: true,
		}, {
			name:          "invalid response format",
			responseCode:  http.StatusOK,
			responseBody:  "invalid JSON",
			expectedError: true,
		}, {
			name:         "empty response",
			responseCode: http.StatusOK,
			responseBody: "{}",
		}, {
			name:           "no users",
			responseCode:   http.StatusOK,
			responseBody:   `{"users": []}`,
			expectedResult: nil,
		}, {
			name:         "single user",
			responseCode: http.StatusOK,
			responseBody: `{
				"users": [
					{
						"id": "145999",
						"username": "testuser",
						"user_full_name": "Test User",
						"user_email": "test@user.example",
						"role": "Member",
						"email_service": true,
						"email_product": true,
						"last_login": null,
						"created_date": "2020-12-09T14:48:42.187Z",
						"updated_date": "2020-12-09T14:48:43.243Z",
						"company_id": "74333",
						"user_api_token": null,
						"filters": {},
						"saved_filters": []
					}
				]
        	}`,
			expectedResult: []models.User{{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			}},
		}, {
			name:         "multiple users",
			responseCode: http.StatusOK,
			responseBody: `{
				"users": [
					{
						"id": "145999",
						"username": "testuser",
						"user_full_name": "Test User",
						"user_email": "test@user.example",
						"role": "Member",
						"email_service": true,
						"email_product": true,
						"last_login": null,
						"created_date": "2020-12-09T14:48:42.187Z",
						"updated_date": "2020-12-09T14:48:43.243Z",
						"company_id": "74333",
						"user_api_token": null,
						"filters": {},
						"saved_filters": []
					},
					{
						"id": "666666",
						"username": "Alice",
						"user_full_name": "Alice Awesome",
						"user_email": "alice.awesome@company.com",
						"role": "Administrator",
						"email_service": false,
						"email_product": false,
						"last_login": "2021-02-05T11:40:09.257Z",
						"created_date": "2021-01-05T12:49:21.306Z",
						"updated_date": "2021-02-05T11:40:09.258Z",
						"company_id": "74333",
						"user_api_token": null,
						"filters": {},
						"saved_filters": []
        			}
				]
        	}`,
			expectedResult: []models.User{{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			}, {
				ID:           "666666",
				Username:     "Alice",
				UserFullName: "Alice Awesome",
				UserEmail:    "alice.awesome@company.com",
				Role:         "Administrator",
				EmailService: false,
				EmailProduct: false,
				LastLogin:    testutil.ParseISO8601Timestamp(t, "2021-02-05T11:40:09.257Z"),
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2021-01-05T12:49:21.306Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2021-02-05T11:40:09.258Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			}},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			assert.NoError(t, err)

			// act
			result, err := c.Users.GetAll(context.Background())

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodGet, h.LastMethod)
			assert.Equal(t, "/api/v5/users", h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		responseCode   int
		responseBody   string
		expectedResult *models.User
		expectedError  bool
	}{
		{
			name:          "status bad request",
			responseCode:  http.StatusBadRequest,
			responseBody:  `{"error":"Bad Request"}`,
			expectedError: true,
		}, {
			name:          "invalid response format",
			responseCode:  http.StatusOK,
			responseBody:  "invalid JSON",
			expectedError: true,
		}, {
			name:          "empty response",
			responseCode:  http.StatusOK,
			responseBody:  "{}",
			expectedError: true,
		}, {
			name:         "user returned",
			responseCode: http.StatusOK,
			responseBody: `{
				"user": {
					"id": "145999",
					"username": "testuser",
					"user_full_name": "Test User",
					"user_email": "test@user.example",
					"role": "Member",
					"email_service": true,
					"email_product": true,
					"last_login": null,
					"created_date": "2020-12-09T14:48:42.187Z",
					"updated_date": "2020-12-09T14:48:43.243Z",
					"company_id": "74333",
					"user_api_token": "****************************a997",
					"filters": {},
					"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: pointer.ToString("****************************a997"),
			},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			assert.NoError(t, err)

			// act
			result, err := c.Users.Get(context.Background(), testUserID)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodGet, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/user/%v", testUserID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		options []kentikapi.ClientOption
		user    models.User
		// expectedRequestBody is a map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responses           []testutil.HTTPResponse
		expectedResult      *models.User
		expectedError       bool
	}{
		{
			name: "user created",
			user: *models.NewUser(models.UserRequiredFields{
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
			}),
			expectedRequestBody: object{
				"user": object{
					"username":       "testuser",
					"user_full_name": "Test User",
					"user_email":     "test@user.example",
					"role":           "Member",
					"email_service":  true,
					"email_product":  false,
				},
			},
			responses: []testutil.HTTPResponse{
				{
					StatusCode: http.StatusCreated,
					Body: `{
						"user": {
							"id": "145999",
							"username": "testuser",
							"user_full_name": "Test User",
							"user_email": "test@user.example",
							"role": "Member",
							"email_service": "true",
							"email_product": "false",
							"last_login": null,
							"created_date": "2020-12-09T14:48:42.187Z",
							"updated_date": "2020-12-09T14:48:43.243Z",
							"company_id": "74333",
							"user_api_token": null,
							"filters": {},
							"saved_filters": []
						}
					}`,
				},
			},
			expectedResult: &models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			},
		}, {
			name:                "status bad request",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responses: []testutil.HTTPResponse{
				{StatusCode: http.StatusBadRequest, Body: `{"error":"Bad Request"}`},
			},
			expectedError: true,
		}, {
			name:                "invalid response format",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responses: []testutil.HTTPResponse{
				{StatusCode: http.StatusCreated, Body: "invalid JSON"},
			},
			expectedError: true,
		}, {
			name:                "empty response",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responses: []testutil.HTTPResponse{
				{StatusCode: http.StatusCreated, Body: "{}"},
			},
			expectedError: true,
		}, {
			name:                "retry 4 times and when status 429, 502, 503, 504 received and last status is 429",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusServiceUnavailable),
				testutil.NewErrorHTTPResponse(http.StatusGatewayTimeout),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
			},
			expectedError: true,
		}, {
			name:                "retry 5 times when status 429, 502, 503, 504 received and last status is 502",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			options:             []kentikapi.ClientOption{kentikapi.WithRetryMaxAttempts(5)},
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
				testutil.NewErrorHTTPResponse(http.StatusServiceUnavailable),
				testutil.NewErrorHTTPResponse(http.StatusGatewayTimeout),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusBadGateway),
			},
			expectedError: true,
		}, {
			name: "retry till success when status 429 too many requests received",
			user: *models.NewUser(models.UserRequiredFields{
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
			}),
			expectedRequestBody: object{
				"user": object{
					"username":       "testuser",
					"user_full_name": "Test User",
					"user_email":     "test@user.example",
					"role":           "Member",
					"email_service":  true,
					"email_product":  false,
				},
			},
			responses: []testutil.HTTPResponse{
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				testutil.NewErrorHTTPResponse(http.StatusTooManyRequests),
				{
					StatusCode: http.StatusCreated,
					Body: `{
						"user": {
							"id": "145999",
							"username": "testuser",
							"user_full_name": "Test User",
							"user_email": "test@user.example",
							"role": "Member",
							"email_service": "true",
							"email_product": "false",
							"last_login": null,
							"created_date": "2020-12-09T14:48:42.187Z",
							"updated_date": "2020-12-09T14:48:43.243Z",
							"company_id": "74333",
							"user_api_token": null,
							"filters": {},
							"saved_filters": []
						}
					}`,
				},
			},
			expectedResult: &models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewMultipleResponseSpyHTTPHandler(t, tt.responses, 0)
			s := httptest.NewServer(h)
			defer s.Close()

			options := []kentikapi.ClientOption{
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithRetryMinDelay(1 * time.Microsecond),
				kentikapi.WithRetryMaxDelay(10 * time.Microsecond),
				kentikapi.WithLogPayloads(),
			}
			c, err := kentikapi.NewClient(append(options, tt.options...)...)
			assert.NoError(t, err)

			// act
			result, err := c.Users.
				Create(context.Background(), tt.user)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, len(tt.responses), len(h.Requests), "invalid number of requests")
			for _, r := range h.Requests {
				assert.Equal(t, http.MethodPost, r.Method)
				assert.Equal(t, "/api/v5/user", r.URL.Path)
				assert.Equal(t, dummyAuthEmail, r.Header.Get(authEmailKey))
				assert.Equal(t, dummyAuthToken, r.Header.Get(authAPITokenKey))
				assert.Equal(t, tt.expectedRequestBody, testutil.UnmarshalJSONToIf(t, r.Body))
			}
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	tests := []struct {
		name         string
		user         models.User
		updateFields func(*models.User) *models.User
		// expectedRequestBody is a map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responseCode        int
		responseBody        string
		expectedResult      *models.User
		expectedError       bool
	}{
		{
			name:                "empty user given, status bad request received",
			user:                models.User{},
			updateFields:        func(u *models.User) *models.User { return u },
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusBadRequest,
			responseBody:        `{"error":"Bad Request"}`,
			expectedError:       true,
		}, {
			name:                "invalid response format",
			user:                models.User{},
			updateFields:        func(u *models.User) *models.User { return u },
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusOK,
			responseBody:        "invalid JSON",
			expectedError:       true,
		}, {
			name:                "empty response",
			user:                models.User{},
			updateFields:        func(u *models.User) *models.User { return u },
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusOK,
			responseBody:        "{}",
			expectedError:       true,
		}, {
			name: "subset of fields updated",
			user: models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
			},
			updateFields: func(u *models.User) *models.User {
				u.UserFullName = "Updated Username"
				u.EmailProduct = false
				return u
			},
			expectedRequestBody: object{
				"user": object{
					"username":       "testuser",
					"user_full_name": "Updated Username",
					"user_email":     "test@user.example",
					"role":           "Member",
					"email_service":  true,
					"email_product":  false,
				},
			},
			responseCode: http.StatusOK,
			responseBody: `{
				"user": {
					"id": "145999",
					"username": "testuser",
					"user_full_name": "Updated Username",
					"user_email": "test@user.example",
					"role": "Member",
					"email_service": "true",
					"email_product": "false",
					"last_login": null,
					"created_date": "2020-12-09T14:48:42.187Z",
					"updated_date": "2020-12-09T14:48:43.243Z",
					"company_id": "74333",
					"user_api_token": null,
					"filters": {},
					"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				ID:           "145999",
				Username:     "testuser",
				UserFullName: "Updated Username",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
				LastLogin:    nil,
				CreatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:42.187Z"),
				UpdatedDate:  *testutil.ParseISO8601Timestamp(t, "2020-12-09T14:48:43.243Z"),
				CompanyID:    "74333",
				UserAPIToken: nil,
			},
		},
	}
	//nolint:dupl
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			assert.NoError(t, err)

			// act
			user := tt.updateFields(&tt.user)
			result, err := c.Users.Update(context.Background(), *user)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodPut, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/user/%v", user.ID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))
			assert.Equal(t, tt.expectedRequestBody, testutil.UnmarshalJSONToIf(t, h.LastRequestBody))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_DeleteUser(t *testing.T) {
	tests := []struct {
		name          string
		responseCode  int
		responseBody  string
		expectedError bool
	}{
		{
			name:          "status bad request",
			responseCode:  http.StatusBadRequest,
			responseBody:  `{"error":"Bad Request"}`,
			expectedError: true,
		}, {
			name:          "invalid response format",
			responseCode:  http.StatusOK,
			responseBody:  "invalid JSON",
			expectedError: false, // response payload is discarded
		}, {
			name:         "user deleted",
			responseCode: http.StatusNoContent,
			responseBody: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := testutil.NewSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c, err := kentikapi.NewClient(
				kentikapi.WithAPIURL(s.URL),
				kentikapi.WithCredentials(dummyAuthEmail, dummyAuthToken),
				kentikapi.WithLogPayloads(),
			)
			assert.NoError(t, err)

			// act
			err = c.Users.Delete(context.Background(), testUserID)

			// assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodDelete, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/user/%v", testUserID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))
			assert.Equal(t, "", h.LastRequestBody)
		})
	}
}

func newEmptyUserRequestBody() map[string]interface{} {
	return object{
		"user": object{
			"username":       "",
			"user_full_name": "",
			"user_email":     "",
			"role":           "",
			"email_service":  false,
			"email_product":  false,
		},
	}
}
