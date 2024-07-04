package xml

import (
	"testing"

	"xml-perf-test/pkg/data"
)

func Benchmark_etree(b *testing.B) {
	rawXML, err := data.LoadEntireFile("../data/input.xml")
	if err != nil {
		b.FailNow()
	}

	// run the function b.N times
	for n := 0; n < b.N; n++ {
		Parse_etree(rawXML)
	}
}
