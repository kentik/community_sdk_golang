package demos

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"

	"github.com/kentik/community_sdk_golang/examples"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

var ExitOnError = examples.PanicOnError
var PrettyPrint = examples.PrettyPrint
var NewClient = examples.NewClient

// untypedData allows trawersing untyped structures made of maps and slices
type untypedData struct {
	V interface{}
}

func makeUntypedData(v interface{}) untypedData {
	return untypedData{V: v}
}

func (d untypedData) object(key string) untypedData {
	m := d.V.(map[string]interface{})
	return untypedData{V: m[key]}
}

func (d untypedData) item(index int) untypedData {
	l := d.V.([]interface{})
	return untypedData{V: l[index]}
}

func Step(msg string) {
	blueBold := "\033[1m\033[34m"
	reset := "\033[0m"

	fmt.Println()
	fmt.Printf("%s%s%s\n", blueBold, msg, reset)
	fmt.Printf("Press enter to continue...")
	fmt.Scanln()
}

// DisplayQueryDataResult prints returned data series in form of a table
func DisplayQueryDataResult(r models.QueryDataResult) {
	const printoutTableFormat = "%v\t%v\t\n" // bits/sec, datetime

	timeSeries := makeUntypedData(r.Results[0]).object("data").item(0).object("timeSeries").object("both_bits_per_sec").object("flow").V.([]interface{})

	w := makeTabWriter(printoutTableFormat)

	// print table header
	fmt.Fprintf(w, printoutTableFormat, "avg_bits_per_sec", "time")
	fmt.Fprintf(w, printoutTableFormat, "----------------", "----")

	// print table rows
	for _, ts := range timeSeries {
		timeBitsPeriod := ts.([]interface{})
		unixTime := int64(timeBitsPeriod[0].(float64)) / 1000 // API returns time in milliseconds, we need seconds
		avgBitsPerSecond := int64(timeBitsPeriod[1].(float64))
		fmt.Fprintf(w, printoutTableFormat, avgBitsPerSecond, time.Unix(unixTime, 0))
	}
	w.Flush()
}

// makeTabWriter prepares tabwriter for writing file list in form: Name  Size  Type
func makeTabWriter(format string) *tabwriter.Writer {
	const minWidth = 0  // minimal cell width including any padding
	const tabWidth = 2  // width of tab characters (equivalent number of spaces)
	const padding = 4   // distance between cells
	const padchar = ' ' // SCII char used for padding
	const flags = 0     // formatting control
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, flags)
	return w
}

// DisplayQueryChartResult shows returned chart image
func DisplayQueryChartResult(r models.QueryChartResult) error {
	// generate temp filename
	file, err := ioutil.TempFile("", "chart_*."+r.ImageType.String())
	if err != nil {
		return err
	}

	r.SaveImageAs(file.Name())
	cmd := exec.Command("firefox", file.Name())
	return cmd.Run()
}
