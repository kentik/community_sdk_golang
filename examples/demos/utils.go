//nolint:forbidigo
package demos

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"

	"github.com/kentik/community_sdk_golang/examples"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

//nolint:gochecknoglobals
var (
	NewClient   = examples.NewClient
	PrettyPrint = examples.PrettyPrint
)

// untypedData allows traversing untyped structures made of maps and slices.
type untypedData struct {
	V interface{}
}

func makeUntypedData(v interface{}) untypedData {
	return untypedData{V: v}
}

func (d untypedData) object(key string) untypedData {
	if m, ok := d.V.(map[string]interface{}); ok {
		return untypedData{V: m[key]}
	}
	return untypedData{V: nil}
}

func (d untypedData) item(index int) untypedData {
	if l, ok := d.V.([]interface{}); ok {
		return untypedData{V: l[index]}
	}
	return untypedData{V: nil}
}

func Step(msg string) {
	blueBold := "\033[1m\033[34m"
	reset := "\033[0m"

	fmt.Println()
	fmt.Printf("%s%s%s\n", blueBold, msg, reset)
	fmt.Printf("Press enter to continue...")
	if _, err := fmt.Scanln(); err != nil {
		log.Print(err)
	}
}

// ExitOnError converts err into panic; use it to reduce the number of: "if err != nil { return err }" statements.
func ExitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

//nolint:gomnd
// DisplayQueryDataResult prints returned data series in form of a table.
func DisplayQueryDataResult(r models.QueryDataResult) error {
	const printoutTableFormat = "%v\t%v\t\n" // bits/sec, datetime

	if timeSeries, ok := makeUntypedData(r.Results[0]).
		object("data").
		item(0).
		object("timeSeries").
		object("both_bits_per_sec").
		object("flow").
		V.([]interface{}); ok {
		w := makeTabWriter()

		// print table header
		fmt.Fprintf(w, printoutTableFormat, "avg_bits_per_sec", "time")
		fmt.Fprintf(w, printoutTableFormat, "----------------", "----")

		// print table rows
		for _, ts := range timeSeries {
			if timeBitsPeriod, ok := ts.([]interface{}); ok {
				unixTimeMS, ok := timeBitsPeriod[0].(float64)
				timeBitsPeriodMS, ok2 := timeBitsPeriod[1].(float64)
				if !ok || !ok2 {
					return errors.New("cannot convert timeBitsPeriod properties to float64")
				}
				unixTime := int64(unixTimeMS) / 1000 // API returns time in milliseconds, we need seconds
				avgBitsPerSecond := int64(timeBitsPeriodMS)
				fmt.Fprintf(w, printoutTableFormat, avgBitsPerSecond, time.Unix(unixTime, 0))
			}
		}
		err := w.Flush()
		if err != nil {
			return err
		}
	}
	return nil
}

// makeTabWriter prepares tab writer for writing file list in form: Name  Size  Type.
func makeTabWriter() *tabwriter.Writer {
	const minWidth = 0  // minimal cell width including any padding
	const tabWidth = 2  // width of tab characters (equivalent number of spaces)
	const padding = 4   // distance between cells
	const padchar = ' ' // SCII char used for padding
	const flags = 0     // formatting control
	w := tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, flags)
	return w
}

//nolint:gosec // G204: Subprocess launched with a potential tainted input or cmd arguments - not an issue
// DisplayQueryChartResult shows returned chart image.
func DisplayQueryChartResult(r models.QueryChartResult) error {
	// generate temp filename
	file, err := ioutil.TempFile("", "chart_*."+r.ImageType.String())
	if err != nil {
		return err
	}

	err = r.SaveImageAs(file.Name())
	if err != nil {
		return err
	}
	cmd := exec.Command("firefox", file.Name())
	return cmd.Run()
}
