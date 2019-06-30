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

	gpx "github.com/sudhanshuraheja/go-garmin-gpx"
)

var layout = "2006-01-02T15:04:05.000Z"

const prefix string = "<trkpt "
const closure string = "</trkseg>\n</trk>\n</gpx>"

func KeepFrom(filepath string, starttime string) {
	//TODO implementare algo
}

func KeepUntil(filepath string, stoptime string) {

	var wg sync.WaitGroup

	g, err := gpx.ParseFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	t0trkp := g.Tracks[0].TrackSegments[0].TrackPoint[0].Timestamp
	t0, err := time.Parse(layout, t0trkp)
	checkerror(err)

	stopIndex := iterateTrackPoint(g, t0, stoptime)
	fmt.Printf("index %v ", stopIndex)

	wg.Add(1)

	go openFile(filepath, stopIndex, wg)
	wg.Wait()
	fmt.Println("Fine")
}

func checkerror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func openFile(filepath string, stop int, wg sync.WaitGroup) {

	i := 0
	file, err := os.Open(filepath)
	checkerror(err)
	defer file.Close()

	dir, filename := path.Split(filepath)
	fmt.Println(dir)
	fmt.Println(filename)
	filenamenoext := strings.Split(filename, ".")
	newfile := fmt.Sprintf("%s%s%s%s", dir, filenamenoext[0], "_edited", ".gpx")

	fwrite, err := os.Create(newfile)
	checkerror(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(text, prefix) {
			if i == stop {
				fmt.Printf(text)
				break
			}
			i++
		}
		_, err = fwrite.WriteString(fmt.Sprintf("%s\n", text))
	}
	_, err = fwrite.WriteString(fmt.Sprintf("%s\n", closure))
	fwrite.Sync()
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
		fmt.Println(t.Sub(t0))
		timeS := int64(t.Sub(t0)) / 1000000000

		if timeS > stopTimeS {
			retindex = i
			break
		}

	}

	return retindex
}
