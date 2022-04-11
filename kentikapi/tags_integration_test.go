package kentikapi_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/testutil"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
)

const (
	testTagID = "42"
)

func TestClient_GetAllTags(t *testing.T) {
	tests := []struct {
		name           string
		responseCode   int
		responseBody   string
		expectedResult []models.Tag
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
			name:           "no tags",
			responseCode:   http.StatusOK,
			responseBody:   `{"tags": []}`,
			expectedResult: nil,
		}, {
			name:         "multiple tags",
			responseCode: http.StatusOK,
			responseBody: `{
				"tags": [
					` + testTagOneResponseJSON + `,
					{
						"id": 452718,
						"flow_tag": "DNS_TRAFFIC",
						"addr_count": 0,
						"port": "53",
						"user": "39242",
						"created_date": "2018-10-04T23:39:29.158284Z",
						"updated_date": "2018-10-04T23:39:29.158284Z",
						"company_id": "26393",
						"mac_count": 0,
						"edited_by": "el.celebes@acme.com"
					}
				]
			}`,
			expectedResult: []models.Tag{
				*newTestTagOne(t),
				{
					FlowTag:     "DNS_TRAFFIC",
					Port:        pointer.ToString("53"),
					ID:          "452718",
					UserID:      "39242",
					CompanyID:   "26393",
					AddrCount:   0,
					MACCount:    0,
					EditedBy:    "el.celebes@acme.com",
					CreatedDate: *testutil.ParseISO8601Timestamp(t, "2018-10-04T23:39:29.158284Z"),
					UpdatedDate: *testutil.ParseISO8601Timestamp(t, "2018-10-04T23:39:29.158284Z"),
				},
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
			result, err := c.Tags.GetAll(context.Background())

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodGet, h.LastMethod)
			assert.Equal(t, "/api/v5/tags", h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_GetTag(t *testing.T) {
	tests := []struct {
		name           string
		responseCode   int
		responseBody   string
		expectedResult *models.Tag
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
			name:         "tag returned",
			responseCode: http.StatusOK,
			responseBody: `{
				"tag": ` + testTagOneResponseJSON + `
			}`,
			expectedResult: newTestTagOne(t),
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
			result, err := c.Tags.Get(context.Background(), testTagID)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodGet, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/tag/%v", testTagID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_CreateTag(t *testing.T) {
	tests := []struct {
		name string
		tag  models.Tag
		// expectedRequestBody is a map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responseCode        int
		responseBody        string
		expectedResult      *models.Tag
		expectedError       bool
	}{
		{
			name:                "empty tag given, status bad request received",
			tag:                 models.Tag{},
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusBadRequest,
			responseBody:        `{"error":"Bad Request"}`,
			expectedError:       true,
		}, {
			name:                "invalid response format",
			tag:                 models.Tag{},
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusCreated,
			responseBody:        "invalid JSON",
			expectedError:       true,
		}, {
			name:                "empty response",
			tag:                 models.Tag{},
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusCreated,
			responseBody:        "{}",
			expectedError:       true,
		}, {
			name: "minimal tag created",
			tag: func() models.Tag {
				tag := models.NewTag("TEST-TAG")
				tag.DeviceName = pointer.ToString("test-device")
				return *tag
			}(),
			expectedRequestBody: object{
				"tag": object{
					"flow_tag":    "TEST-TAG",
					"device_name": "test-device",
				},
			},
			responseCode: http.StatusCreated,
			responseBody: `{
				"tag": {
					"id": 1550982519,
					"flow_tag": "TEST-TAG",
					"device_name": "test-device",
					"addr_count": 0,
					"user": "149492",
					"created_date": "2021-03-04T14:32:04.864641Z",
					"updated_date": "2021-03-04T14:32:04.864641Z",
					"company_id": "74333",
					"mac_count": 0,
					"edited_by": "foo.bar@baz.com"
				}
			}`,
			expectedResult: &models.Tag{
				FlowTag:       "TEST-TAG",
				DeviceName:    pointer.ToString("test-device"),
				DeviceType:    nil,
				Site:          nil,
				InterfaceName: nil,
				Addr:          nil,
				Port:          nil,
				TCPFlags:      nil,
				Protocol:      nil,
				ASN:           nil,
				LasthopAsName: nil,
				NexthopAsn:    nil,
				NexthopAsName: nil,
				Nexthop:       nil,
				BGPAspath:     nil,
				BGPCommunity:  nil,
				MAC:           nil,
				Country:       nil,
				VLANs:         nil,
				ID:            "1550982519",
				UserID:        "149492",
				CompanyID:     "74333",
				AddrCount:     0,
				MACCount:      0,
				EditedBy:      "foo.bar@baz.com",
				CreatedDate:   *testutil.ParseISO8601Timestamp(t, "2021-03-04T14:32:04.864641Z"),
				UpdatedDate:   *testutil.ParseISO8601Timestamp(t, "2021-03-04T14:32:04.864641Z"),
			},
		}, {
			name: "tag created",
			tag: models.Tag{
				FlowTag:       "APITEST-TAG-1",
				DeviceName:    pointer.ToString("192.168.5.100,device1"),
				DeviceType:    pointer.ToString("router,switch"),
				Site:          pointer.ToString("site1,site2"),
				InterfaceName: pointer.ToString("interface1,interface2"),
				Addr:          pointer.ToString("192.168.0.1/32,192.168.0.2/32"),
				Port:          pointer.ToString("9000,9001"),
				TCPFlags:      pointer.ToString("7"),
				Protocol:      pointer.ToString("6,17"),
				ASN:           pointer.ToString("101,102,103"),
				LasthopAsName: pointer.ToString("as1,as2,as3"),
				NexthopAsn:    pointer.ToString("51,52,53"),
				NexthopAsName: pointer.ToString("as51,as52,as53"),
				Nexthop:       pointer.ToString("192.168.7.1/32,192.168.7.2/32"),
				BGPAspath:     pointer.ToString("201,202,203"),
				BGPCommunity:  pointer.ToString("301,302,303"),
				MAC:           pointer.ToString("FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF"),
				Country:       pointer.ToString("ES,IT"),
				VLANs:         pointer.ToString("4001,4002,4003"),
			},
			expectedRequestBody: object{
				"tag": object{
					"flow_tag":        "APITEST-TAG-1",
					"device_name":     "192.168.5.100,device1",
					"device_type":     "router,switch",
					"site":            "site1,site2",
					"interface_name":  "interface1,interface2",
					"addr":            "192.168.0.1/32,192.168.0.2/32",
					"port":            "9000,9001",
					"tcp_flags":       "7",
					"protocol":        "6,17",
					"asn":             "101,102,103",
					"lasthop_as_name": "as1,as2,as3",
					"nexthop_asn":     "51,52,53",
					"nexthop_as_name": "as51,as52,as53",
					"nexthop":         "192.168.7.1/32,192.168.7.2/32",
					"bgp_aspath":      "201,202,203",
					"bgp_community":   "301,302,303",
					"mac":             "FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF",
					"country":         "ES,IT",
					"vlans":           "4001,4002,4003",
				},
			},
			responseCode: http.StatusCreated,
			responseBody: `{
				"tag": ` + testTagOneResponseJSON + `
			}`,
			expectedResult: newTestTagOne(t),
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
			result, err := c.Tags.Create(context.Background(), tt.tag)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodPost, h.LastMethod)
			assert.Equal(t, "/api/v5/tag", h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))
			assert.Equal(t, tt.expectedRequestBody, testutil.UnmarshalJSONToIf(t, h.LastRequestBody))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

//nolint:dupl
func TestClient_UpdateTag(t *testing.T) {
	tests := []struct {
		name         string
		tag          models.Tag
		updateFields func(*models.Tag) *models.Tag
		// expectedRequestBody is a map for the granularity of assertion diff
		expectedRequestBody map[string]interface{}
		responseCode        int
		responseBody        string
		expectedResult      *models.Tag
		expectedError       bool
	}{
		{
			name:                "empty tag given, status bad request received",
			tag:                 models.Tag{},
			updateFields:        func(t *models.Tag) *models.Tag { return t },
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusBadRequest,
			responseBody:        `{"error":"Bad Request"}`,
			expectedError:       true,
		}, {
			name:                "invalid response format",
			tag:                 models.Tag{},
			updateFields:        func(t *models.Tag) *models.Tag { return t },
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusOK,
			responseBody:        "invalid JSON",
			expectedError:       true,
		}, {
			name:                "empty response",
			tag:                 models.Tag{},
			updateFields:        func(t *models.Tag) *models.Tag { return t },
			expectedRequestBody: newEmptyTagRequestBody(),
			responseCode:        http.StatusOK,
			responseBody:        "{}",
			expectedError:       true,
		}, {
			name: "subset of fields updated",
			tag: models.Tag{
				FlowTag: "APITEST-TAG-1",
				ID:      "42",
			},
			updateFields: func(t *models.Tag) *models.Tag {
				t.FlowTag = "APITEST-TAG-2"
				t.DeviceName = pointer.ToString("device2,192.168.5.200")
				t.DeviceType = pointer.ToString("router2,switch2")
				t.Site = pointer.ToString("site3,site4")
				t.InterfaceName = pointer.ToString("interface3,interface4")
				t.TCPFlags = pointer.ToString("8")
				t.ASN = pointer.ToString("111,112,113")
				t.Country = pointer.ToString("ES,IT")
				t.VLANs = pointer.ToString("4011,4012,4013")
				return t
			},
			expectedRequestBody: object{
				"tag": object{
					"flow_tag":       "APITEST-TAG-2",
					"device_name":    "device2,192.168.5.200",
					"device_type":    "router2,switch2",
					"site":           "site3,site4",
					"interface_name": "interface3,interface4",
					"tcp_flags":      "8",
					"asn":            "111,112,113",
					"country":        "ES,IT",
					"vlans":          "4011,4012,4013",
				},
			},
			responseCode: http.StatusOK,
			responseBody: `{
				"tag": {
					"id": 42,
					"flow_tag": "APITEST-TAG-2",
					"device_name": "192.168.5.200,device2",
					"interface_name": "interface3,interface4",
					"addr": "192.168.0.1/32,192.168.0.2/32",
					"addr_count": 2,
					"port": "9000,9001",
					"tcp_flags": "8",
					"protocol": "6,17",
					"asn": "111,112,113",
					"nexthop": "192.168.7.1/32,192.168.7.2/32",
					"nexthop_asn": "51,52,53",
					"bgp_aspath": "201,202,203",
					"bgp_community": "301,302,303",
					"user": "144319",
					"created_date": "2020-12-10T11:53:48.752418Z",
					"updated_date": "2020-12-10T11:53:48.752418Z",
					"company_id": "74333",
					"device_type": "router2,switch2",
					"site": "site3,site4",
					"lasthop_as_name": "as1,as2,as3",
					"nexthop_as_name": "as51,as52,as53",
					"mac": "FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF",
					"mac_count": 2,
					"country": "ES,IT",
					"edited_by": "john.doe@acme.com",
					"vlans": "4011,4012,4013"
				}
			}`,
			expectedResult: &models.Tag{
				FlowTag:       "APITEST-TAG-2",
				DeviceName:    pointer.ToString("192.168.5.200,device2"),
				DeviceType:    pointer.ToString("router2,switch2"),
				Site:          pointer.ToString("site3,site4"),
				InterfaceName: pointer.ToString("interface3,interface4"),
				Addr:          pointer.ToString("192.168.0.1/32,192.168.0.2/32"),
				Port:          pointer.ToString("9000,9001"),
				TCPFlags:      pointer.ToString("8"),
				Protocol:      pointer.ToString("6,17"),
				ASN:           pointer.ToString("111,112,113"),
				LasthopAsName: pointer.ToString("as1,as2,as3"),
				NexthopAsn:    pointer.ToString("51,52,53"),
				NexthopAsName: pointer.ToString("as51,as52,as53"),
				Nexthop:       pointer.ToString("192.168.7.1/32,192.168.7.2/32"),
				BGPAspath:     pointer.ToString("201,202,203"),
				BGPCommunity:  pointer.ToString("301,302,303"),
				MAC:           pointer.ToString("FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF"),
				Country:       pointer.ToString("ES,IT"),
				VLANs:         pointer.ToString("4011,4012,4013"),
				ID:            "42",
				UserID:        "144319",
				CompanyID:     "74333",
				AddrCount:     2,
				MACCount:      2,
				EditedBy:      "john.doe@acme.com",
				CreatedDate:   *testutil.ParseISO8601Timestamp(t, "2020-12-10T11:53:48.752418Z"),
				UpdatedDate:   *testutil.ParseISO8601Timestamp(t, "2020-12-10T11:53:48.752418Z"),
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
			tag := tt.updateFields(&tt.tag)
			result, err := c.Tags.Update(context.Background(), *tag)

			// assert
			t.Logf("Got result: %v, err: %v", result, err)
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodPut, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/tag/%v", tag.ID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))
			assert.Equal(t, tt.expectedRequestBody, testutil.UnmarshalJSONToIf(t, h.LastRequestBody))

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}

func TestClient_DeleteTag(t *testing.T) {
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
			name:         "tag deleted",
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
			err = c.Tags.Delete(context.Background(), testTagID)

			// assert
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, 1, h.RequestsCount)
			assert.Equal(t, http.MethodDelete, h.LastMethod)
			assert.Equal(t, fmt.Sprintf("/api/v5/tag/%v", testTagID), h.LastURL.Path)
			assert.Equal(t, dummyAuthEmail, h.LastHeader.Get(authEmailKey))
			assert.Equal(t, dummyAuthToken, h.LastHeader.Get(authAPITokenKey))
			assert.Equal(t, "", h.LastRequestBody)
		})
	}
}

const testTagOneResponseJSON = `{
	"id": 42,
	"flow_tag": "APITEST-TAG-1",
	"device_name": "192.168.5.100,device1",
	"interface_name": "interface1,interface2",
	"addr": "192.168.0.1/32,192.168.0.2/32",
	"addr_count": 2,
	"port": "9000,9001",
	"tcp_flags": "7",
	"protocol": "6,17",
	"asn": "101,102,103",
	"nexthop": "192.168.7.1/32,192.168.7.2/32",
	"nexthop_asn": "51,52,53",
	"bgp_aspath": "201,202,203",
	"bgp_community": "301,302,303",
	"user": "144319",
	"created_date": "2020-12-10T11:53:48.752418Z",
	"updated_date": "2020-12-10T11:53:48.752418Z",
	"company_id": "74333",
	"device_type": "router,switch",
	"site": "site1,site2",
	"lasthop_as_name": "as1,as2,as3",
	"nexthop_as_name": "as51,as52,as53",
	"mac": "FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF",
	"mac_count": 2,
	"country": "ES,IT",
	"edited_by": "john.doe@acme.com",
	"vlans": "4001,4002,4003"
}`

//nolint:dupl
func newTestTagOne(t *testing.T) *models.Tag {
	return &models.Tag{
		FlowTag:       "APITEST-TAG-1",
		DeviceName:    pointer.ToString("192.168.5.100,device1"),
		DeviceType:    pointer.ToString("router,switch"),
		Site:          pointer.ToString("site1,site2"),
		InterfaceName: pointer.ToString("interface1,interface2"),
		Addr:          pointer.ToString("192.168.0.1/32,192.168.0.2/32"),
		Port:          pointer.ToString("9000,9001"),
		TCPFlags:      pointer.ToString("7"),
		Protocol:      pointer.ToString("6,17"),
		ASN:           pointer.ToString("101,102,103"),
		LasthopAsName: pointer.ToString("as1,as2,as3"),
		NexthopAsn:    pointer.ToString("51,52,53"),
		NexthopAsName: pointer.ToString("as51,as52,as53"),
		Nexthop:       pointer.ToString("192.168.7.1/32,192.168.7.2/32"),
		BGPAspath:     pointer.ToString("201,202,203"),
		BGPCommunity:  pointer.ToString("301,302,303"),
		MAC:           pointer.ToString("FF:FF:FF:FF:FF:FE,FF:FF:FF:FF:FF:FF"),
		Country:       pointer.ToString("ES,IT"),
		VLANs:         pointer.ToString("4001,4002,4003"),
		ID:            "42",
		UserID:        "144319",
		CompanyID:     "74333",
		AddrCount:     2,
		MACCount:      2,
		EditedBy:      "john.doe@acme.com",
		CreatedDate:   *testutil.ParseISO8601Timestamp(t, "2020-12-10T11:53:48.752418Z"),
		UpdatedDate:   *testutil.ParseISO8601Timestamp(t, "2020-12-10T11:53:48.752418Z"),
	}
}

func newEmptyTagRequestBody() map[string]interface{} {
	return object{
		"tag": object{
			"flow_tag": "",
		},
	}
}
