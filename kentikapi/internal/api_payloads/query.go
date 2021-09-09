package api_payloads

import (
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// QuerySQLRequest represents QueryAPI SQL JSON request.
type QuerySQLRequest struct {
	Query string `json:"query"`
}

// QuerySQLResponse represents QueryAPI SQL JSON response.
type QuerySQLResponse struct {
	Rows []interface{} `json:"rows"` // contents depend on used sql query
}

func (r QuerySQLResponse) ToQuerySQLResult() models.QuerySQLResult {
	return models.QuerySQLResult{Rows: r.Rows}
}

// QueryObjectRequest represents QueryAPI Data/Chart/URL JSON request.
type QueryObjectRequest queryObjectPayload

// QueryDataResponse represents QueryAPI Data JSON response.
type QueryDataResponse struct {
	Results []interface{} `json:"results"` // contents depend on used query object
}

func (r QueryDataResponse) ToQueryDataResult() models.QueryDataResult {
	return models.QueryDataResult{Results: r.Results}
}

// QueryURLResponse represents QueryAPI URL JSON response.
type QueryURLResponse string

func (r QueryURLResponse) ToQueryURLResult() models.QueryURLResult {
	unquotedURL := strings.Trim(string(r), `"`) // received url is in quotation marks
	return models.QueryURLResult{URL: unquotedURL}
}

// QueryChartResponse represents QueryAPI Chart JSON response.
type QueryChartResponse struct {
	DataURI string `json:"dataUri"` // like: "data:image/png;base64,iVBORw0KGgoAAAA..."
}

func (r QueryChartResponse) ToQueryChartResult() (models.QueryChartResult, error) {
	imageType, err := r.getImageType()
	if err != nil {
		return models.QueryChartResult{}, err
	}
	imageDataBase64, err := r.getImageDataBase64()
	if err != nil {
		return models.QueryChartResult{}, err
	}
	imageData, err := base64.StdEncoding.DecodeString(imageDataBase64)
	if err != nil {
		return models.QueryChartResult{}, err
	}
	return models.QueryChartResult{ImageType: imageType, ImageData: imageData}, nil
}

func (r QueryChartResponse) getImageType() (models.ImageType, error) {
	mime, _, _, err := parseURI(r.DataURI)
	if err != nil {
		return models.ImageTypeUnknown, err
	}

	switch mime {
	case "image/png":
		return models.ImageTypePNG, nil
	case "image/jpeg":
		return models.ImageTypeJPG, nil
	case "image/svg+xml":
		return models.ImageTypeSVG, nil
	case "application/pdf":
		return models.ImageTypePDF, nil
	default:
		return models.ImageTypeUnknown, errors.New("Expected mime type png/jpeg/svg+xml/pdf, got: " + mime)
	}
}

func (r QueryChartResponse) getImageDataBase64() (string, error) {
	_, encoding, payload, err := parseURI(r.DataURI)
	if err != nil {
		return "", err
	}
	if encoding != "base64" {
		return "", errors.New("Expected base64 encoding, got: " + encoding)
	}
	return payload, nil
}

func parseURI(uriString string) (mime string, encoding string, payload string, err error) {
	uri := uriString // eg. "data:image/png;base64,iVBORw0KGgoAAAA..."
	dataType, uri := cutHead(uri, ":")
	mime, uri = cutHead(uri, ";")
	encoding, uri = cutHead(uri, ",")
	payload = uri
	if dataType != "data" || mime == "" || encoding == "" || payload == "" {
		return mime, encoding, payload, errors.New("Invalid URI: " + uriString[:50])
	}
	return mime, encoding, payload, nil
}

func cutHead(s string, until string) (head string, remainder string) {
	pos := strings.Index(s, until)
	if pos == -1 {
		return "", s
	}
	return s[:pos], s[pos+1:]
}

type queryObjectPayload struct {
	Queries   []queryArrayItemPayload `json:"queries"`
	ImageType *string                 `json:"imageType,omitempty"`
}

//nolint:revive // queryObjectPayLoad doesn't need to be exported
func QueryObjectToPayload(q models.QueryObject) (queryObjectPayload, error) {
	var queries []queryArrayItemPayload
	if err := utils.ConvertList(q.Queries, queryArrayItemToPayload, &queries); err != nil {
		return queryObjectPayload{}, err
	}

	var imageType *string
	if q.ImageType != nil {
		imageType = new(string)
		*imageType = q.ImageType.String()
	}

	return queryObjectPayload{
		Queries:   queries,
		ImageType: imageType,
	}, nil
}

type queryArrayItemPayload struct {
	Query       queryPayload `json:"query"`
	Bucket      string       `json:"bucket"`
	BucketIndex *int         `json:"bucketIndex,omitempty"`
	IsOverlay   *bool        `json:"isOverlay,omitempty"`
}

//nolint:nilerr
func queryArrayItemToPayload(i models.QueryArrayItem) (queryArrayItemPayload, error) {
	query, err := queryToPayload(i.Query)
	if err != nil {
		return queryArrayItemPayload{}, nil
	}

	return queryArrayItemPayload{
		Query:       query,
		Bucket:      i.Bucket,
		BucketIndex: i.BucketIndex,
		IsOverlay:   i.IsOverlay,
	}, nil
}

type queryPayload struct {
	Metric       string               `json:"metric"`
	Dimension    []string             `json:"dimension,omitempty"`
	FiltersObj   *filtersPayload      `json:"filters_obj,omitempty"`
	SavedFilters []savedFilterPayload `json:"saved_filters,omitempty"`
	MatrixBy     []string             `json:"matrixBy" request:"post"` // matrixBy is required in request even if empty.
	// Otherwise Chart query hangs
	CIDR            *int               `json:"cidr,omitempty"`
	CIDR6           *int               `json:"cidr6,omitempty"`
	PPSThreshold    *int               `json:"pps_threshold,omitempty"`
	TopX            int                `json:"topx"`
	Depth           int                `json:"depth"`
	FastData        string             `json:"fastData"`
	TimeFormat      string             `json:"time_format"`
	HostnameLookup  bool               `json:"hostname_lookup"`
	LookbackSeconds int                `json:"lookback_seconds"`
	StartingTime    *string            `json:"starting_time,omitempty"` // format YYYY-MM-DD HH:mm:00
	EndingTime      *string            `json:"ending_time,omitempty"`   // format YYYY-MM-DD HH:mm:00
	AllSelected     *bool              `json:"all_selected,omitempty"`
	DeviceName      []string           `json:"device_name" request:"post"` // device_name is required in request even if empty
	Descriptor      string             `json:"descriptor"`
	Aggregates      []aggregatePayload `json:"aggregates,omitempty"`
	Outsort         *string            `json:"outsort,omitempty"`
	QueryTitle      string             `json:"query_title"`
	VizType         *string            `json:"viz_type,omitempty"`
	ShowOverlay     *bool              `json:"show_overlay,omitempty"`
	OverlayDay      *int               `json:"overlay_day,omitempty"`
	SyncAxes        *bool              `json:"sync_axes,omitempty"`
}

func queryToPayload(q models.Query) (queryPayload, error) {
	var err error

	dimensions := make([]string, len(q.Dimension))
	for i, d := range q.Dimension {
		dimensions[i] = d.String()
	}

	var aggregates []aggregatePayload
	err = utils.ConvertList(q.Aggregates, aggregateToPayload, &aggregates)
	if err != nil {
		return queryPayload{}, err
	}

	var viztype *string
	if q.VizType != nil {
		viztype = new(string)
		*viztype = q.VizType.String()
	}

	var savedFiltersPayloads []savedFilterPayload
	for _, i := range q.SavedFilters {
		savedFiltersPayloads = append(savedFiltersPayloads, SavedFilterToPayload(i))
	}

	return queryPayload{
		Metric:          q.Metric.String(),
		Dimension:       dimensions,
		FiltersObj:      filtersToPayload(q.FiltersObj),
		SavedFilters:    savedFiltersPayloads,
		MatrixBy:        q.MatrixBy,
		CIDR:            q.CIDR,
		CIDR6:           q.CIDR6,
		PPSThreshold:    q.PPSThreshold,
		TopX:            q.TopX,
		Depth:           q.Depth,
		FastData:        q.FastData.String(),
		TimeFormat:      q.TimeFormat.String(),
		HostnameLookup:  q.HostnameLookup,
		LookbackSeconds: q.LookbackSeconds,
		StartingTime:    FormatQueryTime(q.StartingTime),
		EndingTime:      FormatQueryTime(q.EndingTime),
		AllSelected:     q.AllSelected,
		DeviceName:      q.DeviceName,
		Descriptor:      q.Descriptor,
		Aggregates:      aggregates,
		Outsort:         q.Outsort,
		QueryTitle:      q.QueryTitle,
		VizType:         viztype,
		ShowOverlay:     q.ShowOverlay,
		OverlayDay:      q.OverlayDay,
		SyncAxes:        q.SyncAxes,
	}, nil
}

func filtersToPayload(f *models.Filters) *filtersPayload {
	if f == nil {
		return nil
	}

	var filterGroupsPayloads []filterGroupsPayload

	for _, i := range f.FilterGroups {
		filterGroupsPayloads = append(filterGroupsPayloads, filterGroupsToPayload(i))
	}

	return &filtersPayload{
		Connector:    f.Connector,
		FilterGroups: filterGroupsPayloads,
		Custom:       f.Custom,
		FilterString: f.FilterString,
	}
}

type aggregatePayload struct {
	Name       string `json:"name"`
	Column     string `json:"column"`
	Fn         string `json:"fn"`
	SampleRate int    `json:"sample_rate"`
	Rank       *int   `json:"rank,omitempty"` // valid: number 5..99; only used when Fn == Percentile
	Raw        *bool  `json:"raw,omitempty"`  // required for topxchart queries
}

func aggregateToPayload(a models.Aggregate) (aggregatePayload, error) {
	return aggregatePayload{
		Name:       a.Name,
		Column:     a.Column,
		Fn:         a.Fn.String(),
		SampleRate: a.SampleRate,
		Rank:       a.Rank,
		Raw:        a.Raw,
	}, nil
}

func FormatQueryTime(t *time.Time) *string {
	if t == nil {
		return nil
	}
	layout := "2006-01-02 15:04"
	result := t.Format(layout) + ":00" // "YYYY-MM-DD HH:mm:00"
	return &result
}
