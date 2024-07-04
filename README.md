# README

This project compares the performance of different libraries to parse XML in Go.
  - [native package](https://pkg.go.dev/encoding/xml)
  - https://github.com/tamerh/xml-stream-parser
  - https://github.com/beevik/etree

Results can be found in the [`./results` directory](/results/README.md).

The program can also be run as standalone.

## Test data

The program and performance were always run with the same set of data (found inside ./pkg/data/input.xml).
It's a file with 1000 objects, each with 5 fields, replicating the largest response from a Ceph cluster
during the listobjects operation.

## In-house implementation

Whether to rely on an external library or to implement our own solution is a trade-off. For now, this option
was not considered but I am leaving some resources here for future reference.

- https://github.com/dps/go-xml-parse
