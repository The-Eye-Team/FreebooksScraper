package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/cavaliercoder/grab"
	"github.com/labstack/gommon/color"
)

var arguments = struct {
	Input       string
	Output      string
	Concurrency int
	RandomUA    bool
	Verbose     bool
	GetAlbums   bool
	GetVideos   bool
	StartID     int
	StopID      int
}{}

var client = http.Client{}

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func init() {
	// Disable HTTP/2: Empty TLSNextProto map
	client.Transport = http.DefaultTransport
	client.Transport.(*http.Transport).TLSNextProto =
		make(map[string]func(authority string, c *tls.Conn) http.RoundTripper)
}

func downloadFile(url string, index int, worker *sync.WaitGroup) {
	defer worker.Done()

	resp, err := grab.Get(arguments.Output, url)
	if err != nil {
		fmt.Println(crossPre +
			color.Yellow("[") +
			color.Red(index) +
			color.Yellow("]") +
			color.Red(" Failed downloading: ") +
			color.Yellow(url+" ") +
			color.Red(err.Error()))
		return
	}

	fmt.Println(checkPre +
		color.Yellow("[") +
		color.Green(index) +
		color.Yellow("]") +
		color.Green(" Downloaded: ") +
		color.Yellow(resp.Filename))
}

func main() {
	var worker sync.WaitGroup
	var count int
	var url string

	// Parse arguments and fill the arguments structure
	parseArgs(os.Args)

	// Set maxIdleConnsPerHost
	client.Transport.(*http.Transport).MaxIdleConnsPerHost = arguments.Concurrency

	// Create output dir
	os.MkdirAll(arguments.Output+"/", os.ModePerm)

	for index := arguments.StartID; index <= arguments.StopID; index++ {
		worker.Add(1)
		count++
		url = "http://www.freebooks.com/fetch.php?file=" + strconv.Itoa(index)
		go downloadFile(url, index, &worker)
		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}

	worker.Wait()
}
