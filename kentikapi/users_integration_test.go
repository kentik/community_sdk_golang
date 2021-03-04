package kentikapi_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

const (
	authEmailKey    = "X-CH-Auth-Email"
	authAPITokenKey = "X-CH-Auth-API-Token"
	dummyAuthEmail  = "email@example.com"
	dummyAuthToken  = "api-test-token"
	testUserID      = 145999

	createdDateA = "2020-12-09T14:48:42.187Z"
	createdDateB = "2021-01-05T12:49:21.306Z"
	updatedDateA = "2020-12-09T14:48:43.243Z"
	updatedDateB = "2021-02-05T11:40:09.258Z"
	lastLoginA   = "2021-02-05T11:40:09.257Z"
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
						"created_date": "` + createdDateA + `",
						"updated_date": "` + updatedDateA + `",
						"company_id": "74333",
						"user_api_token": null,
						"filters": {},
						"saved_filters": []
					}
				]
        	}`,
			expectedResult: []models.User{{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
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
						"created_date": "` + createdDateA + `",
						"updated_date": "` + updatedDateA + `",
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
						"last_login": "` + lastLoginA + `",
						"created_date": "` + createdDateB + `",
						"updated_date": "` + updatedDateB + `",
						"company_id": "74333",
						"user_api_token": null,
						"filters": {},
						"saved_filters": []
        			}
				]
        	}`,
			expectedResult: []models.User{{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: nil,
			}, {
				ID:           666666,
				Username:     "Alice",
				UserFullName: "Alice Awesome",
				UserEmail:    "alice.awesome@company.com",
				Role:         "Administrator",
				EmailService: false,
				EmailProduct: false,
				LastLogin:    parseISO8601Timestamp(t, lastLoginA),
				CreatedDate:  *parseISO8601Timestamp(t, createdDateB),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateB),
				CompanyID:    74333,
				UserAPIToken: nil,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL,
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

			// act
			result, err := c.Users.GetAll(context.Background())

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.requestsCount)
			assert.Equal(t, http.MethodGet, h.lastMethod)
			assert.Equal(t, "/users", h.lastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.lastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.lastHeader.Get(authAPITokenKey))

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
					"created_date": "` + createdDateA + `",
					"updated_date": "` + updatedDateA + `",
					"company_id": "74333",
					"user_api_token": "****************************a997",
        			"filters": {},
        			"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: stringPointer("****************************a997"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL,
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

			// act
			result, err := c.Users.Get(context.Background(), testUserID)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.requestsCount)
			assert.Equal(t, http.MethodGet, h.lastMethod)
			assert.Equal(t, fmt.Sprintf("/user/%v", testUserID), h.lastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.lastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.lastHeader.Get(authAPITokenKey))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_CreateUser(t *testing.T) {
	tests := []struct {
		name                string
		user                models.User
		expectedRequestBody interface{}
		responseCode        int
		responseBody        string
		expectedResult      *models.User
		expectedError       bool
	}{
		{
			name:                "empty user given, status bad request received",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusBadRequest,
			responseBody:        `{"error":"Bad Request"}`,
			expectedError:       true,
		}, {
			name:                "invalid response format",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusCreated,
			responseBody:        "invalid JSON",
			expectedError:       true,
		}, {
			name: "empty response",
			user: *models.NewUser(models.UserRequiredFields{
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
			}),
			expectedRequestBody: object{
				"user": object{
					"username":       "testuser",
					"user_full_name": "Test User",
					"user_email":     "test@user.example",
					"role":           "Member",
					"email_service":  true,
					"email_product":  true,
				},
			},
			responseCode:  http.StatusCreated,
			responseBody:  "{}",
			expectedError: true,
		}, {
			name:                "empty user created",
			user:                models.User{},
			expectedRequestBody: newEmptyUserRequestBody(),
			responseCode:        http.StatusCreated,
			responseBody: `{
				"user": {
					"id": "0",
					"username": "",
					"user_full_name": "",
					"user_email": "",
					"role": "",
					"email_service": "false",
					"email_product": "false",
					"last_login": null,
					"created_date": "` + createdDateA + `",
					"updated_date": "` + updatedDateA + `",
					"company_id": "0",
					"user_api_token": null,
					"filters": {},
					"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				CreatedDate: *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate: *parseISO8601Timestamp(t, updatedDateA),
			},
		}, {
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
			responseCode: http.StatusCreated,
			responseBody: `{
				"user": {
					"id": "145999",
					"username": "testuser",
					"user_full_name": "Test User",
					"user_email": "test@user.example",
					"role": "Member",
					"email_service": "true",
					"email_product": "false",
					"last_login": null,
					"created_date": "` + createdDateA + `",
					"updated_date": "` + updatedDateA + `",
					"company_id": "74333",
					"user_api_token": null,
					"filters": {},
					"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL,
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

			// act
			result, err := c.Users.Create(context.Background(), tt.user)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.requestsCount)
			assert.Equal(t, http.MethodPost, h.lastMethod)
			assert.Equal(t, "/user", h.lastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.lastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.lastHeader.Get(authAPITokenKey))
			assert.Equal(t, tt.expectedRequestBody, unmarshalJSONToIf(t, h.lastRequestBody))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_UpdateUser(t *testing.T) {
	tests := []struct {
		name                string
		user                models.User
		updateFields        func(*models.User) *models.User
		expectedRequestBody interface{}
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
			name: "empty response",
			user: models.User{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: nil,
			},
			updateFields: func(u *models.User) *models.User { return u },
			expectedRequestBody: object{
				"user": object{
					"username":       "testuser",
					"user_full_name": "Test User",
					"user_email":     "test@user.example",
					"role":           "Member",
					"email_service":  true,
					"email_product":  true,
				},
			},
			responseCode:  http.StatusOK,
			responseBody:  "{}",
			expectedError: true,
		}, {
			name: "subset of fields updated",
			user: models.User{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: nil,
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
					"created_date": "` + createdDateA + `",
					"updated_date": "` + updatedDateA + `",
					"company_id": "74333",
					"user_api_token": null,
					"filters": {},
					"saved_filters": []
				}
			}`,
			expectedResult: &models.User{
				ID:           145999,
				Username:     "testuser",
				UserFullName: "Updated Username",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: false,
				LastLogin:    nil,
				CreatedDate:  *parseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *parseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
				UserAPIToken: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL,
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

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

			assert.Equal(t, 1, h.requestsCount)
			assert.Equal(t, http.MethodPut, h.lastMethod)
			assert.Equal(t, fmt.Sprintf("/user/%v", user.ID), h.lastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.lastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.lastHeader.Get(authAPITokenKey))
			assert.Equal(t, tt.expectedRequestBody, unmarshalJSONToIf(t, h.lastRequestBody))

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
			h := newSpyHTTPHandler(t, tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL,
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

			// act
			err := c.Users.Delete(context.Background(), testUserID)

			// assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.requestsCount)
			assert.Equal(t, http.MethodDelete, h.lastMethod)
			assert.Equal(t, fmt.Sprintf("/user/%v", testUserID), h.lastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.lastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.lastHeader.Get(authAPITokenKey))
			assert.Equal(t, "", h.lastRequestBody)
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

type spyHTTPHandler struct {
	t            testing.TB
	responseCode int
	responseBody []byte

	requestsCount   int
	lastMethod      string
	lastURL         *url.URL
	lastHeader      http.Header
	lastRequestBody string
}

func newSpyHTTPHandler(t testing.TB, responseCode int, responseBody []byte) *spyHTTPHandler {
	return &spyHTTPHandler{
		t:            t,
		responseCode: responseCode,
		responseBody: responseBody,
	}
}

func (s *spyHTTPHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	s.requestsCount++
	s.lastMethod = r.Method
	s.lastURL = r.URL
	s.lastHeader = r.Header

	body, err := ioutil.ReadAll(r.Body)
	assert.NoError(s.t, err)
	s.lastRequestBody = string(body)

	err = r.Body.Close()
	assert.NoError(s.t, err)

	rw.WriteHeader(s.responseCode)
	_, err = rw.Write(s.responseBody)
	assert.NoError(s.t, err)
}

func parseISO8601Timestamp(t testing.TB, timestamp string) *time.Time {
	const iso8601Layout = "2006-01-02T15:04:05Z0700"
	ts, err := time.Parse(iso8601Layout, timestamp)
	assert.NoError(t, err)

	return &ts
}

func unmarshalJSONToIf(t testing.TB, jsonString string) interface{} {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	assert.NoError(t, err)
	return data
}

func stringPointer(s string) *string {
	return &s
}
