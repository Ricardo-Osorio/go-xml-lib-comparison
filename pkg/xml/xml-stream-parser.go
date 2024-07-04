package xml

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
	"xml-perf-test/pkg/data"
	"xml-perf-test/pkg/utils.go"

	xmlparser "github.com/tamerh/xml-stream-parser"
)

func Parse_xml_stream_parser(input []byte) *data.ListObjectsV2Response {
	out := &data.ListObjectsV2Response{}

	// convert from []byte to io.Reader
	f := bytes.NewReader(input)

	br := bufio.NewReaderSize(f, 65536)
	parser := xmlparser.NewXMLParser(br, "NextContinuationToken", "KeyCount", "IsTruncated", "Contents")

	for xml := range parser.Stream() {
		switch xml.Name {
		case "NextContinuationToken":
			out.NextContinuationToken = utils.StringPointer(strings.Trim(xml.InnerText, " \n"))
		case "KeyCount":
			KeyCountStr := strings.Trim(xml.InnerText, " \n")
			// convert string to int32
			KeyCountInt, _ := strconv.Atoi(KeyCountStr)
			out.KeyCount = int32(KeyCountInt)
		case "IsTruncated":
			IsTruncatedStr := strings.Trim(xml.InnerText, " \n")
			IsTruncated, _ := strconv.ParseBool(IsTruncatedStr)
			out.IsTruncated = IsTruncated
		case "Contents":
			var (
				key          string
				LastModified string
				ETag         string
				Size         string // need convertion
				StorageClass string
			)
			toAdd := &data.ObjectMeta{}

			key = xml.Childs["Key"][0].InnerText
			key = strings.Trim(key, " \n")
			toAdd.Key = key

			LastModified = xml.Childs["LastModified"][0].InnerText
			LastModified = strings.Trim(LastModified, " \n")
			toAdd.LastModified = LastModified

			ETag = xml.Childs["ETag"][0].InnerText
			ETag = strings.Trim(ETag, " \n")
			toAdd.Etag = ETag

			Size = xml.Childs["Size"][0].InnerText
			Size = strings.Trim(Size, " \n")
			toAdd.Size, _ = strconv.ParseUint(Size, 10, 64)

			StorageClass = xml.Childs["StorageClass"][0].InnerText
			StorageClass = strings.Trim(StorageClass, " \n")
			toAdd.StorageClass = data.StorageClassType(StorageClass)

			out.Contents = append(out.Contents, toAdd)
			// fmt.Printf("Object found:\n\tKey: %s\n\tLastModified: %s\n\tETag: %s\n\tSize: %s\n\tStorageClass: %s\n", key, LastModified, ETag, Size, StorageClass)
		}
	}
	return out
}
