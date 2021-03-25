package provider

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func diagError(summary string, err error, httpResp *http.Response) diag.Diagnostics {
	var details string
	if httpResp != nil {
		details = fmt.Sprintf("%v %v", err.Error(), httpResp.Body)
	} else {
		details = err.Error()
	}

	diags := diag.Diagnostic{
		Severity: diag.Error,
		Summary:  summary,
		Detail:   details,
	}
	return diag.Diagnostics{diags}
}
