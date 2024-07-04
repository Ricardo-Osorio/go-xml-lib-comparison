package xml

import (
	"strconv"
	"strings"
	"xml-perf-test/pkg/data"
	"xml-perf-test/pkg/utils.go"

	xmlparser "github.com/beevik/etree"
)

func Parse_etree(input []byte) (*data.ListObjectsV2Response, error) {
	out := &data.ListObjectsV2Response{}

	parser := xmlparser.NewDocument()
	err := parser.ReadFromBytes(input)
	if err != nil {
		return nil, err
	}

	root := parser.SelectElement("ListBucketResult")

	// Note: this lib is more prone to nil panics due to having to
	// directly reference each element and not offering an iterable
	// option.

	gotElement := root.SelectElement("NextContinuationToken")
	if gotElement != nil {
		out.NextContinuationToken = utils.StringPointer(strings.Trim(gotElement.Text(), " \n"))
	}

	gotElement = root.SelectElement("KeyCount")
	if gotElement != nil {
		KeyCountStr := strings.Trim(gotElement.Text(), " \n")
		// convert string to int32
		KeyCountInt, _ := strconv.Atoi(KeyCountStr)
		out.KeyCount = int32(KeyCountInt)
	}

	IsTruncatedStr := strings.Trim(root.SelectElement("IsTruncated").Text(), " \n")
	IsTruncated, _ := strconv.ParseBool(IsTruncatedStr)
	out.IsTruncated = IsTruncated

	for _, topLevelEntry := range root.SelectElements("Contents") {
		var (
			key          string
			LastModified string
			ETag         string
			Size         string // need convertion
			StorageClass string
		)
		toAdd := &data.ObjectMeta{}

		key = topLevelEntry.SelectElement("Key").Text()
		key = strings.Trim(key, " \n")
		toAdd.Key = key

		LastModified = topLevelEntry.SelectElement("LastModified").Text()
		LastModified = strings.Trim(LastModified, " \n")
		toAdd.LastModified = LastModified

		ETag = topLevelEntry.SelectElement("ETag").Text()
		ETag = strings.Trim(ETag, " \n")
		toAdd.Etag = ETag

		Size = topLevelEntry.SelectElement("Size").Text()
		Size = strings.Trim(Size, " \n")
		toAdd.Size, _ = strconv.ParseUint(Size, 10, 64)

		StorageClass = topLevelEntry.SelectElement("StorageClass").Text()
		StorageClass = strings.Trim(StorageClass, " \n")
		toAdd.StorageClass = data.StorageClassType(StorageClass)

		out.Contents = append(out.Contents, toAdd)
		// fmt.Printf("Object found:\n\tKey: %s\n\tLastModified: %s\n\tETag: %s\n\tSize: %s\n\tStorageClass: %s\n", key, LastModified, ETag, Size, StorageClass)
	}

	return out, nil
}
