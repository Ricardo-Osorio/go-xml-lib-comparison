package xml

import (
	"sync"
	"testing"

	"xml-perf-test/pkg/data"
)

// Having all functions called together helps with drawing a better cpuprofile image for comparison purposes
func Benchmark_all(b *testing.B) {
	rawXML, err := data.LoadEntireFile("../data/input.xml")
	if err != nil {
		b.FailNow()
	}

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		// run the function b.N times
		for n := 0; n < b.N; n++ {
			Parse_etree(rawXML)
		}
	}()

	go func() {
		defer wg.Done()
		// run the function b.N times
		for n := 0; n < b.N; n++ {
			Parse_native(rawXML)
		}
	}()

	go func() {
		defer wg.Done()
		// run the function b.N times
		for n := 0; n < b.N; n++ {
			Parse_xml_stream_parser(rawXML)
		}
	}()

	wg.Wait()
}
