package api_endpoints

import "fmt"

const (
	TenantPath  = "/mykentik/tenant"
	TenantsPath = "/mykentik/tenants"
)

func GetTenantPath(tenantID ResourceID) string {
	return fmt.Sprintf("%v/%v", TenantPath, tenantID)
}

func CreateTenantUserPath(tenantID ResourceID) string {
	return fmt.Sprintf("%v/%v/user", TenantPath, tenantID)
}

func DeleteTenantUserPath(tenantID ResourceID, userID ResourceID) string {
	return fmt.Sprintf("%v/%v/user/%v", TenantPath, tenantID, userID)
}
