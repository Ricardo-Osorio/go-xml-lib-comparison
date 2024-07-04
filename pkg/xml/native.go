package xml

import (
	"encoding/xml"
	"xml-perf-test/pkg/data"
)

func Parse_native(input []byte) (*data.ListObjectsV2Response, error) {
	out := &data.ListObjectsV2Response{}
	err := xml.Unmarshal(input, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
