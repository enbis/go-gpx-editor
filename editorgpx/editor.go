package editorgpx

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/spf13/viper"
	gpx "github.com/sudhanshuraheja/go-garmin-gpx"
)

var layout = "2006-01-02T15:04:05.000Z"

const prefix string = "<trkpt "
const closure string = "</trkseg>\n</trk>\n</gpx>"

type conf struct {
	closure string `yml: "closure"`
	prefix  string `yml: "prefix"`
}

func vipersetting() {

}

func KeepFrom(filepath string, starttime string) {

	var wg sync.WaitGroup

	g, err := gpx.ParseFile(filepath)

	checkerror(err)

	t0trkp := g.Tracks[0].TrackSegments[0].TrackPoint[0].Timestamp
	t0, err := time.Parse(layout, t0trkp)
	checkerror(err)

	stopIndex := iterateTrackPoint(g, t0, starttime)
	fmt.Printf("index %v ", stopIndex)

	wg.Add(1)

	go keepFromProcess(filepath, stopIndex, wg)

	wg.Wait()
	fmt.Println("Fine")

}

func KeepUntil(filepath string, stoptime string) {

	var wg sync.WaitGroup

	g, err := gpx.ParseFile(filepath)

	checkerror(err)

	t0trkp := g.Tracks[0].TrackSegments[0].TrackPoint[0].Timestamp
	t0, err := time.Parse(layout, t0trkp)
	checkerror(err)

	stopIndex := iterateTrackPoint(g, t0, stoptime)
	fmt.Printf("index %v ", stopIndex)

	wg.Add(1)

	go keepUntilProcess(filepath, stopIndex, wg)
	wg.Wait()
	fmt.Println("Fine")
}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func keepUntilProcess(filepath string, stop int, wg sync.WaitGroup) {
	filein, newFile := openF(filepath)
	fileout := writeF(newFile)
	algoKeepUntil(filein, fileout, stop, wg)
}

func keepFromProcess(filepath string, start int, wg sync.WaitGroup) {
	filein, newFile := openF(filepath)
	fileout := writeF(newFile)
	algoKeepFrom(filein, fileout, start, wg)
}

func openF(filepath string) (*os.File, string) {
	file, err := os.Open(filepath)
	checkerror(err)

	dir, filename := path.Split(filepath)

	filenamenoext := strings.Split(filename, ".")
	newfile := fmt.Sprintf("%s%s%s%s", dir, filenamenoext[0], "_edited", ".gpx")
	return file, newfile
}

func writeF(newfile string) *os.File {
	fwrite, err := os.Create(newfile)
	checkerror(err)
	return fwrite
}

func algoKeepFrom(filein *os.File, fileout *os.File, start int, wg sync.WaitGroup) {

	prefix := viper.GetString("Prefix")

	i := 0

	var notw bool
	scanner := bufio.NewScanner(filein)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(text, prefix) {
			if i < start {
				notw = true
			} else {
				_, err := fileout.WriteString(fmt.Sprintf("%s\n", text))
				checkerror(err)
			}
			i++
		} else {
			if notw {
				notw = !strings.HasPrefix(text, "</trkpt>")
			} else {
				_, err := fileout.WriteString(fmt.Sprintf("%s\n", text))
				checkerror(err)
			}
		}
	}

	filein.Close()
	fileout.Sync()
	wg.Done()
}

func algoKeepUntil(filein *os.File, fileout *os.File, stop int, wg sync.WaitGroup) {

	prefix := viper.GetString("Prefix")
	closure := viper.GetString("Closure")

	i := 0

	scanner := bufio.NewScanner(filein)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, prefix) {
			if i == stop {
				fmt.Printf(text)
				break
			}
			i++
		}
		_, err := fileout.WriteString(fmt.Sprintf("%s\n", text))
		checkerror(err)
	}
	_, err := fileout.WriteString(fmt.Sprintf("%s\n", closure))
	checkerror(err)

	filein.Close()
	fileout.Sync()
	wg.Done()
}

func iterateTrackPoint(g *gpx.GPX, t0 time.Time, stoptime string) int {

	retindex := 0
	stopTimeDuration, err := time.ParseDuration(stoptime)
	if err != nil {
		fmt.Println("ParseDuration error ", err)
	}
	stopTimeS := int64(stopTimeDuration) / 1000000000

	for i, val := range g.Tracks[0].TrackSegments[0].TrackPoint {
		t, err := time.Parse(layout, val.Timestamp)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(t.Sub(t0))
		timeS := int64(t.Sub(t0)) / 1000000000

		if timeS > stopTimeS {
			retindex = i
			break
		}

	}

	return retindex
}
