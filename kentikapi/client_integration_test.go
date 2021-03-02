package kentikapi_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

const (
	dummyAuthEmail = "email@example.com"
	dummyAuthToken = "api-test-token"

	createdDateA = "2020-12-09T14:48:42.187Z"
	createdDateB = "2021-01-05T12:49:21.306Z"
	updatedDateA = "2020-12-09T14:48:43.243Z"
	updatedDateB = "2021-02-05T11:40:09.258Z"
	lastLoginA   = "2021-02-05T11:40:09.257Z"
)

func TestClient_GetAllUsers(t *testing.T) {
	tests := []struct {
		name          string
		responseCode  int
		responseBody  string
		expected      []models.User
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
			expectedError: true,
		}, {
			name:         "empty response",
			responseCode: http.StatusOK,
			responseBody: "{}",
		}, {
			name:         "no users",
			responseCode: http.StatusOK,
			responseBody: `{"users": []}`,
			expected:     []models.User{},
		}, {
			name:         "single user",
			responseCode: http.StatusOK,
			responseBody: `{
				"users": [
					{
						"id": "145999",
						"username": "test@user.example",
						"user_full_name": "Test User",
						"user_email": "test@user.example",
						"role": "Member",
						"email_service": true,
						"email_product": true,
						"last_login": null,
						"created_date": "` + createdDateA + `",
						"updated_date": "` + updatedDateA + `",
						"company_id": "74333"
					}
				]
        	}`,
			expected: []models.User{{
				ID:           145999,
				Username:     "test@user.example",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *utils.ParseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *utils.ParseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
			}},
		}, {
			name:         "multiple users",
			responseCode: http.StatusOK,
			responseBody: `{
				"users": [
					{
						"id": "145999",
						"username": "test@user.example",
						"user_full_name": "Test User",
						"user_email": "test@user.example",
						"role": "Member",
						"email_service": true,
						"email_product": true,
						"last_login": null,
						"created_date": "` + createdDateA + `",
						"updated_date": "` + updatedDateA + `",
						"company_id": "74333"
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
						"company_id": "74333"
        			}
				]
        	}`,
			expected: []models.User{{
				ID:           145999,
				Username:     "test@user.example",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *utils.ParseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *utils.ParseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
			}, {
				ID:           666666,
				Username:     "Alice",
				UserFullName: "Alice Awesome",
				UserEmail:    "alice.awesome@company.com",
				Role:         "Administrator",
				EmailService: false,
				EmailProduct: false,
				LastLogin:    utils.ParseISO8601Timestamp(t, lastLoginA),
				CreatedDate:  *utils.ParseISO8601Timestamp(t, createdDateB),
				UpdatedDate:  *utils.ParseISO8601Timestamp(t, updatedDateB),
				CompanyID:    74333,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    s.URL + "/users",
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
			assert.Equal(t, tt.expected, result)

			assert.Equal(t, 1, h.requestsCount)
			assert.NoError(t, h.lastWriteError)
		})
	}
}

func TestClientGetUser(t *testing.T) {
	tests := []struct {
		name          string
		userID        int
		responseCode  int
		responseBody  string
		expected      *models.User
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
			expectedError: true,
		}, {
			name:         "empty response",
			responseCode: http.StatusOK,
			responseBody: "{}",
			expected:     &models.User{},
		}, {
			name:         "user returned",
			responseCode: http.StatusOK,
			responseBody: `{
				"user": {
					"id": "145999",
					"username": "test@user.example",
					"user_full_name": "Test User",
					"user_email": "test@user.example",
					"role": "Member",
					"email_service": true,
					"email_product": true,
					"last_login": null,
					"created_date": "` + createdDateA + `",
					"updated_date": "` + updatedDateA + `",
					"company_id": "74333"
				}
			}`,
			expected: &models.User{
				ID:           145999,
				Username:     "test@user.example",
				UserFullName: "Test User",
				UserEmail:    "test@user.example",
				Role:         "Member",
				EmailService: true,
				EmailProduct: true,
				LastLogin:    nil,
				CreatedDate:  *utils.ParseISO8601Timestamp(t, createdDateA),
				UpdatedDate:  *utils.ParseISO8601Timestamp(t, updatedDateA),
				CompanyID:    74333,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// arrange
			h := newSpyHTTPHandler(tt.responseCode, []byte(tt.responseBody))
			s := httptest.NewServer(h)
			defer s.Close()

			c := kentikapi.NewClient(kentikapi.Config{
				APIURL:    fmt.Sprintf("%v/user/%v", s.URL, tt.userID) + "/user/",
				AuthEmail: dummyAuthEmail,
				AuthToken: dummyAuthToken,
			})

			// act
			result, err := c.Users.Get(context.Background(), tt.userID)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, result)

			assert.Equal(t, 1, h.requestsCount)
			assert.NoError(t, h.lastWriteError)
		})
	}
}

type spyHTTPHandler struct {
	responseCode   int
	responseBody   []byte
	requestsCount  int
	lastWriteError error
}

func newSpyHTTPHandler(responseCode int, responseBody []byte) *spyHTTPHandler {
	return &spyHTTPHandler{
		responseCode: responseCode,
		responseBody: responseBody,
	}
}

func (s *spyHTTPHandler) ServeHTTP(rw http.ResponseWriter, _ *http.Request) {
	s.requestsCount++

	// TODO(dfurman): verify that client sends auth headers

	rw.WriteHeader(s.responseCode)
	_, s.lastWriteError = rw.Write(s.responseBody)
}
