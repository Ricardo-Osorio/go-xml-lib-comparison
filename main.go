package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"xml-perf-test/pkg/data"
	"xml-perf-test/pkg/xml"
)

func main() {
	f, err := os.Create("cpuprofile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	fmt.Printf("Starting...\n")

	// === Read XML ===

	timeNow := time.Now()

	rawXML, err := data.LoadEntireFile("")
	if err != nil {
		return
	}

	fmt.Printf("time to load file: %s\n", time.Since(timeNow))

	// === Native XML parsing ===

	timeNow = time.Now()

	_, err = xml.Parse_native(rawXML)
	if err != nil {
		fmt.Printf("failed to parse xml (native): %s\n", err.Error())
		return
	}

	fmt.Printf("time to parse XML (native): %s\n", time.Since(timeNow))

	// === github.com/tamerh/xml-stream-parser ===

	timeNow = time.Now()

	_ = xml.Parse_xml_stream_parser(rawXML)

	fmt.Printf("time to parse XML (xml-stream-parser): %s\n", time.Since(timeNow))

	// === https://github.com/beevik/etree ===

	timeNow = time.Now()

	_, err = xml.Parse_etree(rawXML)
	if err != nil {
		fmt.Printf("failed to parse xml (etree): %s\n", err.Error())
		return
	}

	fmt.Printf("time to parse XML (etree): %s\n", time.Since(timeNow))

	fmt.Printf("Stoping...\n")
}
