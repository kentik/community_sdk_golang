package api_endpoints_test

import (
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDeviceRouter(t *testing.T) {
	// arrange
	getResponsePayload := `
    {
        "device": {
            "id": "42",
            "company_id": "74333",
            "device_name": "testapi_router_full_1",
            "device_type": "router",
            "device_status": "V",
            "device_description": "testapi router with full config",
            "site": {
                "id": 8483,
                "site_name": "marina gdańsk",
                "lat": 54.348972,
                "lon": 18.659791,
                "company_id": 74333
            },
            "plan": {
                "active": true,
                "bgp_enabled": true,
                "cdate": "2020-09-03T08:41:57.489Z",
                "company_id": 74333,
                "description": "Your Free Trial includes 6 devices (...)",
                "deviceTypes": [],
                "devices": [],
                "edate": "2020-09-03T08:41:57.489Z",
                "fast_retention": 30,
                "full_retention": 30,
                "id": 11466,
                "max_bigdata_fps": 30,
                "max_devices": 6,
                "max_fps": 1000,
                "name": "Free Trial Plan",
                "metadata": {}
            },
            "labels": [
                        {
                            "id": 2590,
                            "name": "AWS: terraform-demo-aws",
                            "description": null,
                            "edate": "2020-10-05T15:28:00.276Z",
                            "cdate": "2020-10-05T15:28:00.276Z",
                            "user_id": "133210",
                            "company_id": "74333",
                            "color": "#5340A5",
                            "order": null,
                            "_pivot_device_id": "77715",
                            "_pivot_label_id": "2590"
                        },
                        {
                            "id": 2751,
                            "name": "GCP: traffic-generator-gcp",
                            "description": null,
                            "edate": "2020-11-20T12:54:49.575Z",
                            "cdate": "2020-11-20T12:54:49.575Z",
                            "user_id": "136885",
                            "company_id": "74333",
                            "color": "#5289D9",
                            "order": null,
                            "_pivot_device_id": "77373",
                            "_pivot_label_id": "2751"
                        }
                    ],
            "all_interfaces": [
                {
                    "interface_description": "testapi-interface-1",
                    "initial_snmp_speed": null,
                    "device_id": "42",
                    "snmp_speed": "75"
                },
                {
                    "interface_description": "testapi-interface-2",
                    "initial_snmp_speed": "7",
                    "device_id": "42",
                    "snmp_speed": "7"
                }
            ],
            "device_flow_type": "auto",
            "device_sample_rate": "1001",
            "sending_ips": [
                "128.0.0.11",
                "128.0.0.12"
            ],
            "device_snmp_ip": "129.0.0.1",
            "device_snmp_community": "",
            "minimize_snmp": false,
            "device_bgp_type": "device",
            "device_bgp_neighbor_ip": "127.0.0.1",
            "device_bgp_neighbor_ip6": null,
            "device_bgp_neighbor_asn": "11",
            "device_bgp_flowspec": true,
            "device_bgp_password": "*********ass",
            "use_bgp_device_id": null,
            "custom_columns": "",
            "custom_column_data": [],
            "device_chf_client_port": null,
            "device_chf_client_protocol": null,
            "device_chf_interface": null,
            "device_agent_type": null,
            "max_flow_rate": 1000,
            "max_big_flow_rate": 30,
            "device_proxy_bgp": "",
            "device_proxy_bgp6": "",
            "created_date": "2020-12-17T08:24:45.074Z",
            "updated_date": "2020-12-17T08:24:45.074Z",
            "device_snmp_v3_conf": {
                "UserName": "John",
                "AuthenticationProtocol": "MD5",
                "AuthenticationPassphrase": "john_md5_pass",
                "PrivacyProtocol": "DES",
                "PrivacyPassphrase": "**********ass"
            },
            "bgpPeerIP4": "208.76.14.223",
            "bgpPeerIP6": "2620:129:1:2::1",
            "snmp_last_updated": null,
            "device_subtype": "router"
        }
    }`
	transport := api_connection.StubTransport{ResponseBody: getResponsePayload}
	devicesAPI := api_endpoints.NewDevicesAPI(transport)
	deviceID := models.ID(42)

	// act
	device, err := devicesAPI.Get(nil, deviceID)

	// assert
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	assert.Equal(models.ID(42), device.ID)
	assert.Equal(models.ID(74333), device.CompanyID)
	assert.Equal("testapi_router_full_1", device.DeviceName)
	assert.Equal(models.DeviceTypeRouter, device.DeviceType)
	assert.Equal("testapi router with full config", *device.DeviceDescription)
	require.NotNil(device.Site)
	assert.Equal(models.ID(8483), *device.Site.ID)
	assert.Equal("marina gdańsk", *device.Site.SiteName)
	assert.Equal(54.348972, *device.Site.Latitude)
	assert.Equal(18.659791, *device.Site.Longitude)
	assert.Equal(models.ID(74333), *device.Site.CompanyID)
	require.NotNil(device.Plan)
	assert.True(device.Plan.Active)
	assert.True(device.Plan.BGPEnabled)
	assert.Equal(time.Date(2020, 9, 3, 8, 41, 57, 489*1000000, time.UTC), device.Plan.CreatedDate)
	assert.Equal(models.ID(74333), device.Plan.CompanyID)
	assert.Equal("Your Free Trial includes 6 devices (...)", device.Plan.Description)
	assert.Equal(0, len(device.Plan.DeviceTypes))
	assert.Equal(0, len(device.Plan.Devices))
	assert.Equal(time.Date(2020, 9, 3, 8, 41, 57, 489*1000000, time.UTC), *device.Plan.UpdatedDate)
	assert.Equal(30, device.Plan.FastRetention)
	assert.Equal(30, device.Plan.FullRetention)
	assert.Equal(models.ID(11466), device.Plan.ID)
	assert.Equal(30, device.Plan.MaxBigdataFPS)
	assert.Equal(6, device.Plan.MaxDevices)
	assert.Equal(1000, device.Plan.MaxFPS)
	assert.Equal("Free Trial Plan", device.Plan.Name)
	assert.Equal(2, len(device.Labels))
	assert.Equal(models.ID(2590), device.Labels[0].ID)
	assert.Equal("AWS: terraform-demo-aws", device.Labels[0].Name)
	assert.Equal(time.Date(2020, 10, 5, 15, 28, 00, 276*1000000, time.UTC), device.Labels[0].UpdatedDate)
	assert.Equal(time.Date(2020, 10, 5, 15, 28, 00, 276*1000000, time.UTC), device.Labels[0].CreatedDate)
	assert.Equal(models.ID(133210), *device.Labels[0].UserID)
	assert.Equal(models.ID(74333), device.Labels[0].CompanyID)
	assert.Equal("#5340A5", device.Labels[0].Color)
	assert.Equal(models.ID(2751), device.Labels[1].ID)
	assert.Equal("GCP: traffic-generator-gcp", device.Labels[1].Name)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), device.Labels[1].UpdatedDate)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), device.Labels[1].CreatedDate)
	assert.Equal(models.ID(136885), *device.Labels[1].UserID)
	assert.Equal(models.ID(74333), device.Labels[1].CompanyID)
	assert.Equal("#5289D9", device.Labels[1].Color)
	assert.Equal(2, len(device.AllInterfaces))
	assert.Equal("testapi-interface-1", device.AllInterfaces[0].InterfaceDescription)
	assert.Nil(device.AllInterfaces[0].InitialSNMPSpeed)
	assert.Equal(models.ID(42), device.AllInterfaces[0].DeviceID)
	assert.Equal(75.0, device.AllInterfaces[0].SNMPSpeed)
	assert.Equal("testapi-interface-2", device.AllInterfaces[1].InterfaceDescription)
	assert.Equal(7.0, *device.AllInterfaces[1].InitialSNMPSpeed)
	assert.Equal(models.ID(42), device.AllInterfaces[1].DeviceID)
	assert.Equal(7.0, device.AllInterfaces[1].SNMPSpeed)
	assert.Equal("auto", *device.DeviceFlowType)
	assert.Equal(1001, *&device.DeviceSampleRate)
	assert.Equal(2, len(device.SendingIPS))
	assert.Equal("128.0.0.11", device.SendingIPS[0])
	assert.Equal("128.0.0.12", device.SendingIPS[1])
	assert.Equal("129.0.0.1", *device.DeviceSNMNPIP)
	assert.Equal("", *device.DeviceSNMPCommunity)
	assert.False(*device.MinimizeSNMP)
	assert.Equal(models.DeviceBGPTypeDevice, *device.DeviceBGPType)
	assert.Equal("127.0.0.1", *device.DeviceBGPNeighborIP)
	assert.Nil(device.DeviceBGPNeighborIPv6)
	assert.Equal("11", *device.DeviceBGPNeighborASN)
	assert.True(*device.DeviceBGPFlowSpec)
	assert.Equal("*********ass", *device.DeviceBGPPassword)
	assert.Nil(device.UseBGPDeviceID)
	assert.Equal(time.Date(2020, 12, 17, 8, 24, 45, 74*1000000, time.UTC), device.CreatedDate)
	assert.Equal(time.Date(2020, 12, 17, 8, 24, 45, 74*1000000, time.UTC), device.UpdatedDate)
	require.NotNil(device.DeviceSNMPv3Conf)
	assert.Equal("John", device.DeviceSNMPv3Conf.UserName)
	assert.Equal(models.AuthenticationProtocolMD5, *device.DeviceSNMPv3Conf.AuthenticationProtocol)
	assert.Equal("john_md5_pass", *device.DeviceSNMPv3Conf.AuthenticationPassphrase)
	assert.Equal(models.PrivacyProtocolDES, *device.DeviceSNMPv3Conf.PrivacyProtocol)
	assert.Equal("**********ass", *device.DeviceSNMPv3Conf.PrivacyPassphrase)
	assert.Equal("208.76.14.223", *device.BGPPeerIP4)
	assert.Equal("2620:129:1:2::1", *device.BGPPeerIP6)
	assert.Nil(device.SNMPLastUpdated)
	assert.Equal(models.DeviceSubtypeRouter, device.DeviceSubType)
}

func TestGetAllDevices(t *testing.T) {
	// arrange
	getResponsePayload := `
    {
        "devices": [
            {
                "id": "42",
                "company_id": "74333",
                "device_name": "testapi_router_full_1",
                "device_type": "router",
                "device_status": "V",
                "device_description": "testapi router with full config",
                "site": {
                    "id": 8483,
                    "site_name": "marina gdańsk",
                    "lat": 54.348972,
                    "lon": 18.659791,
                    "company_id": 74333
                },
                "plan": {
                    "active": true,
                    "bgp_enabled": true,
                    "cdate": "2020-09-03T08:41:57.489Z",
                    "company_id": 74333,
                    "description": "Your Free Trial includes 6 devices (...)",
                    "deviceTypes": [],
                    "devices": [],
                    "edate": "2020-09-03T08:41:57.489Z",
                    "fast_retention": 30,
                    "full_retention": 30,
                    "id": 11466,
                    "max_bigdata_fps": 30,
                    "max_devices": 6,
                    "max_fps": 1000,
                    "name": "Free Trial Plan",
                    "metadata": {}
                },
                "labels": [
                            {
                                "id": 2590,
                                "name": "AWS: terraform-demo-aws",
                                "description": null,
                                "edate": "2020-10-05T15:28:00.276Z",
                                "cdate": "2020-10-05T15:28:00.276Z",
                                "user_id": "133210",
                                "company_id": "74333",
                                "color": "#5340A5",
                                "order": null,
                                "_pivot_device_id": "77715",
                                "_pivot_label_id": "2590"
                            },
                            {
                                "id": 2751,
                                "name": "GCP: traffic-generator-gcp",
                                "description": null,
                                "edate": "2020-11-20T12:54:49.575Z",
                                "cdate": "2020-11-20T12:54:49.575Z",
                                "user_id": null,
                                "company_id": "74333",
                                "color": "#5289D9",
                                "order": null,
                                "_pivot_device_id": "77373",
                                "_pivot_label_id": "2751"
                            }
                        ],
                "all_interfaces": [],
                "device_flow_type": "auto",
                "device_sample_rate": "1001",
                "sending_ips": [
                    "128.0.0.11",
                    "128.0.0.12"
                ],
                "device_snmp_ip": "129.0.0.1",
                "device_snmp_community": "",
                "minimize_snmp": false,
                "device_bgp_type": "device",
                "device_bgp_neighbor_ip": "127.0.0.1",
                "device_bgp_neighbor_ip6": null,
                "device_bgp_neighbor_asn": "11",
                "device_bgp_flowspec": true,
                "device_bgp_password": "*********ass",
                "use_bgp_device_id": null,
                "custom_columns": "",
                "custom_column_data": [],
                "device_chf_client_port": null,
                "device_chf_client_protocol": null,
                "device_chf_interface": null,
                "device_agent_type": null,
                "max_flow_rate": 1000,
                "max_big_flow_rate": 30,
                "device_proxy_bgp": "",
                "device_proxy_bgp6": "",
                "created_date": "2020-12-17T08:24:45.074Z",
                "updated_date": "2020-12-17T08:24:45.074Z",
                "device_snmp_v3_conf": {
                    "UserName": "John",
                    "AuthenticationProtocol": "MD5",
                    "AuthenticationPassphrase": "john_md5_pass",
                    "PrivacyProtocol": "DES",
                    "PrivacyPassphrase": "**********ass"
                },
                "bgpPeerIP4": "208.76.14.223",
                "bgpPeerIP6": "2620:129:1:2::1",
                "snmp_last_updated": null,
                "device_subtype": "router"
            },
            {
                "id": "43",
                "company_id": "74333",
                "device_name": "testapi_dns_minimal_1",
                "device_type": "host-nprobe-dns-www",
                "device_status": "V",
                "device_description": "testapi dns with minimal config",
                "site": {
                    "id": null,
                    "site_name": null,
                    "lat": null,
                    "lon": null,
                    "company_id": null
                },
                "plan": {
                    "active": true,
                    "bgp_enabled": true,
                    "cdate": "2020-09-03T08:41:57.489Z",
                    "company_id": 74333,
                    "description": "Your Free Trial includes 6 devices (...)",
                    "deviceTypes": [],
                    "devices": [],
                    "edate": "2020-09-03T08:41:57.489Z",
                    "fast_retention": 30,
                    "full_retention": 30,
                    "id": 11466,
                    "max_bigdata_fps": 30,
                    "max_devices": 6,
                    "max_fps": 1000,
                    "name": "Free Trial Plan",
                    "metadata": {}
                },
                "labels": [],
                "all_interfaces": [],
                "device_flow_type": "auto",
                "device_sample_rate": "1",
                "sending_ips": [],
                "device_snmp_ip": null,
                "device_snmp_community": "",
                "minimize_snmp": false,
                "device_bgp_type": "none",
                "device_bgp_neighbor_ip": null,
                "device_bgp_neighbor_ip6": null,
                "device_bgp_neighbor_asn": null,
                "device_bgp_flowspec": false,
                "device_bgp_password": null,
                "use_bgp_device_id": null,
                "custom_columns": "",
                "custom_column_data": [],
                "device_chf_client_port": null,
                "device_chf_client_protocol": null,
                "device_chf_interface": null,
                "device_agent_type": null,
                "max_flow_rate": 1000,
                "max_big_flow_rate": 30,
                "device_proxy_bgp": "",
                "device_proxy_bgp6": "",
                "created_date": "2020-12-17T12:53:01.025Z",
                "updated_date": "2020-12-17T12:53:01.025Z",
                "device_snmp_v3_conf": null,
                "cdn_attr": "Y",
                "snmp_last_updated": null,
                "device_subtype": "aws_subnet"
            }
        ]
    }`
	transport := api_connection.StubTransport{ResponseBody: getResponsePayload}
	devicesAPI := api_endpoints.NewDevicesAPI(transport)

	// act
	devices, err := devicesAPI.GetAll(nil)

	// assert
	assert := assert.New(t)
	require := require.New(t)

	require.NoError(err)
	require.Equal(2, len(devices))
	// device 0
	device := devices[0]
	assert.Equal(models.ID(42), device.ID)
	assert.Equal(models.ID(74333), device.CompanyID)
	assert.Equal("testapi_router_full_1", device.DeviceName)
	assert.Equal(models.DeviceTypeRouter, device.DeviceType)
	assert.Equal("testapi router with full config", *device.DeviceDescription)
	require.NotNil(device.Site)
	assert.Equal(models.ID(8483), *device.Site.ID)
	assert.Equal("marina gdańsk", *device.Site.SiteName)
	assert.Equal(54.348972, *device.Site.Latitude)
	assert.Equal(18.659791, *device.Site.Longitude)
	assert.Equal(models.ID(74333), *device.Site.CompanyID)
	require.NotNil(device.Plan)
	assert.True(device.Plan.Active)
	assert.True(device.Plan.BGPEnabled)
	assert.Equal(time.Date(2020, 9, 3, 8, 41, 57, 489*1000000, time.UTC), device.Plan.CreatedDate)
	assert.Equal(models.ID(74333), device.Plan.CompanyID)
	assert.Equal("Your Free Trial includes 6 devices (...)", device.Plan.Description)
	assert.Equal(0, len(device.Plan.DeviceTypes))
	assert.Equal(0, len(device.Plan.Devices))
	assert.Equal(time.Date(2020, 9, 3, 8, 41, 57, 489*1000000, time.UTC), *device.Plan.UpdatedDate)
	assert.Equal(30, device.Plan.FastRetention)
	assert.Equal(30, device.Plan.FullRetention)
	assert.Equal(models.ID(11466), device.Plan.ID)
	assert.Equal(30, device.Plan.MaxBigdataFPS)
	assert.Equal(6, device.Plan.MaxDevices)
	assert.Equal(1000, device.Plan.MaxFPS)
	assert.Equal("Free Trial Plan", device.Plan.Name)
	assert.Equal(2, len(device.Labels))
	assert.Equal(models.ID(2590), device.Labels[0].ID)
	assert.Equal("AWS: terraform-demo-aws", device.Labels[0].Name)
	assert.Equal(time.Date(2020, 10, 5, 15, 28, 00, 276*1000000, time.UTC), device.Labels[0].UpdatedDate)
	assert.Equal(time.Date(2020, 10, 5, 15, 28, 00, 276*1000000, time.UTC), device.Labels[0].CreatedDate)
	assert.Equal(models.ID(133210), *device.Labels[0].UserID)
	assert.Equal(models.ID(74333), device.Labels[0].CompanyID)
	assert.Equal("#5340A5", device.Labels[0].Color)
	assert.Equal(models.ID(2751), device.Labels[1].ID)
	assert.Equal("GCP: traffic-generator-gcp", device.Labels[1].Name)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), device.Labels[1].UpdatedDate)
	assert.Equal(time.Date(2020, 11, 20, 12, 54, 49, 575*1000000, time.UTC), device.Labels[1].CreatedDate)
	assert.Nil(device.Labels[1].UserID)
	assert.Equal(models.ID(74333), device.Labels[1].CompanyID)
	assert.Equal("#5289D9", device.Labels[1].Color)
	require.Equal(0, len(device.AllInterfaces))
	assert.Equal("auto", *device.DeviceFlowType)
	assert.Equal(1001, *&device.DeviceSampleRate)
	assert.Equal(2, len(device.SendingIPS))
	assert.Equal("128.0.0.11", device.SendingIPS[0])
	assert.Equal("128.0.0.12", device.SendingIPS[1])
	assert.Equal("129.0.0.1", *device.DeviceSNMNPIP)
	assert.Equal("", *device.DeviceSNMPCommunity)
	assert.False(*device.MinimizeSNMP)
	assert.Equal(models.DeviceBGPTypeDevice, *device.DeviceBGPType)
	assert.Equal("127.0.0.1", *device.DeviceBGPNeighborIP)
	assert.Nil(device.DeviceBGPNeighborIPv6)
	assert.Equal("11", *device.DeviceBGPNeighborASN)
	assert.True(*device.DeviceBGPFlowSpec)
	assert.Equal("*********ass", *device.DeviceBGPPassword)
	assert.Nil(device.UseBGPDeviceID)
	assert.Equal(time.Date(2020, 12, 17, 8, 24, 45, 74*1000000, time.UTC), device.CreatedDate)
	assert.Equal(time.Date(2020, 12, 17, 8, 24, 45, 74*1000000, time.UTC), device.UpdatedDate)
	require.NotNil(device.DeviceSNMPv3Conf)
	assert.Equal("John", device.DeviceSNMPv3Conf.UserName)
	assert.Equal(models.AuthenticationProtocolMD5, *device.DeviceSNMPv3Conf.AuthenticationProtocol)
	assert.Equal("john_md5_pass", *device.DeviceSNMPv3Conf.AuthenticationPassphrase)
	assert.Equal(models.PrivacyProtocolDES, *device.DeviceSNMPv3Conf.PrivacyProtocol)
	assert.Equal("**********ass", *device.DeviceSNMPv3Conf.PrivacyPassphrase)
	assert.Equal("208.76.14.223", *device.BGPPeerIP4)
	assert.Equal("2620:129:1:2::1", *device.BGPPeerIP6)
	assert.Nil(device.SNMPLastUpdated)
	assert.Equal(models.DeviceSubtypeRouter, device.DeviceSubType)

	// device 1
	device = devices[1]
	assert.Equal(models.ID(43), device.ID)
	assert.Equal(models.DeviceTypeHostNProbeDNSWWW, device.DeviceType)
	assert.Equal(models.DeviceSubtypeAwsSubnet, device.DeviceSubType)
	assert.Equal(models.DeviceBGPTypeNone, *device.DeviceBGPType)
}
