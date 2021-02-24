package api_endpoints

// QuerySQL runs SQL command against all configured devices. Returns top X datasets.
func QuerySQL() string {
	return "/query/sql"
}

// QueryURL returns URL which a logged in user can use to directly access this query in Data Explorer in the Kentik Detect portal.
func QueryURL() string {
	return "/query/url"
}

// QueryData returns results in a JSON results array
func QueryData() string {
	return "/query/topXdata"
}

// QueryChart returns result as a chart image
func QueryChart() string {
	return "/query/topXchart"
}
