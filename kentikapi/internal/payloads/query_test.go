package payloads_test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueryChartResponsePNGToQueryChartResult(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ImagePNGEncodedBase64str"
	decodedData := base64Decode(t, data)
	response := payloads.QueryChartResponse{DataURI: "data:image/png;base64," + data}

	// act
	result, err := response.ToQueryChartResult()

	// assert
	require.NoError(t, err)
	assert.Equal(t, models.ImageTypePNG, result.ImageType)
	assert.Equal(t, decodedData, result.ImageData)
}

func TestQueryChartResponseJPEGToQueryChartResult(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ImageJPGEncodedBase64str"
	decodedData := base64Decode(t, data)
	response := payloads.QueryChartResponse{DataURI: "data:image/jpeg;base64," + data}

	// act
	result, err := response.ToQueryChartResult()

	// assert
	require.NoError(t, err)
	assert.Equal(t, models.ImageTypeJPG, result.ImageType)
	assert.Equal(t, decodedData, result.ImageData)
}

func TestQueryChartResponseSVGToQueryChartResult(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ImageSVGEncodedBase64str"
	decodedData := base64Decode(t, data)
	response := payloads.QueryChartResponse{DataURI: "data:image/svg+xml;base64," + data}

	// act
	result, err := response.ToQueryChartResult()

	// assert
	require.NoError(t, err)
	assert.Equal(t, models.ImageTypeSVG, result.ImageType)
	assert.Equal(t, decodedData, result.ImageData)
}

func TestQueryChartResponsePDFToQueryChartResult(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ApplicationPDFEncodedBase64str=="
	decodedData := base64Decode(t, data)
	response := payloads.QueryChartResponse{DataURI: "data:application/pdf;base64," + data}

	// act
	result, err := response.ToQueryChartResult()

	// assert
	require.NoError(t, err)
	assert.Equal(t, models.ImageTypePDF, result.ImageType)
	assert.Equal(t, decodedData, result.ImageData)
}

func TestUnknownFormatResultsInError(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ImageBMPEncodedBase64str=="
	response := payloads.QueryChartResponse{DataURI: "data:image/bmp;base64," + data}

	// act
	_, err := response.ToQueryChartResult()

	// assert
	require.Error(t, err)
}

func TestUnknownEncodingResultsInError(t *testing.T) {
	t.Parallel()

	// arrange
	data := "ImagePNGEncodedBase32str=="
	response := payloads.QueryChartResponse{DataURI: "data:image/png;base32," + data}

	// act
	_, err := response.ToQueryChartResult()

	// assert
	require.Error(t, err)
}

func TestFormatQueryTimeNonNil(t *testing.T) {
	t.Parallel()

	// arrange
	datetime := time.Date(2001, 3, 9, 6, 45, 53, 111, time.UTC)

	// act
	queryTime := payloads.FormatQueryTime(&datetime)

	// assert
	require.NotNil(t, queryTime)
	assert.Equal(t, "2001-03-09 06:45:00", *queryTime)
}

func TestFormatQueryTimeNil(t *testing.T) {
	t.Parallel()

	// arrange
	var datetime *time.Time

	// act
	queryTime := payloads.FormatQueryTime(datetime)

	// assert
	assert.Nil(t, queryTime)
}

func base64Decode(t *testing.T, s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	require.NoError(t, err)
	return data
}
