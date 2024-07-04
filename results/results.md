# Results

Overall the `xml_stream_parser` lib was the one that showed quite significant improvements in terms of performance. The `etree` and `native` packages were quite similar in terms of performance.

## Benchmarks
```
$ benchstat 10_runs_individual_bench.txt 30_runs_individual_bench.txt
goos: linux
goarch: amd64
pkg: xml-perf-test/pkg/xml
cpu: Intel(R) Xeon(R) Gold 6348 CPU @ 2.60GHz
                     │ 10_runs_individual_bench.txt │      30_runs_individual_bench.txt      │
                     │            sec/op            │   sec/op     vs base                   │
_etree-8                               15.66m ±  9%   14.98m ± 2%        ~ (p=0.123 n=10+30)
_native-8                              15.36m ± 14%   15.30m ± 2%        ~ (p=0.469 n=10+30)
_xml_stream_parser-8                   6.836m ±  5%   6.140m ± 2%  -10.17% (p=0.000 n=10+30)
geomean                                11.80m         11.21m        -5.04%

                     │ 10_runs_individual_bench.txt │      30_runs_individual_bench.txt      │
                     │             B/op             │     B/op      vs base                  │
_etree-8                               4.783Mi ± 0%   4.783Mi ± 0%       ~ (p=0.528 n=10+30)
_native-8                              3.558Mi ± 0%   3.557Mi ± 0%       ~ (p=0.281 n=10+30)
_xml_stream_parser-8                   4.178Mi ± 0%   4.177Mi ± 0%  -0.01% (p=0.000 n=10+30)
geomean                                4.142Mi        4.142Mi       -0.00%

                     │ 10_runs_individual_bench.txt │      30_runs_individual_bench.txt       │
                     │          allocs/op           │  allocs/op   vs base                    │
_etree-8                                128.1k ± 0%   128.1k ± 0%       ~ (p=0.603 n=10+30)
_native-8                               108.1k ± 0%   108.1k ± 0%       ~ (p=1.000 n=10+30) ¹
_xml_stream_parser-8                    52.06k ± 0%   52.06k ± 0%       ~ (p=1.000 n=10+30) ¹
geomean                                 89.67k        89.67k       +0.00%
```

The `xml_stream_parser` package is the fastest in terms of duration of operations and also the one with less memory allocations. It's also the only one with some degree of variation between the runs with 10 and 30 executions but that varied while generating this data.

Both `etree` and the `native` packages sit closely in terms of execution time and memory allocations.

## Practical results

From running this program against the dataset of 1000 records we get:

```
Starting...
time to load file: 220.072µs
time to parse XML (etree): 19.926844ms
time to parse XML (native): 19.815031ms
time to parse XML (xml-stream-parser): 7.777412ms
Stoping...
```

Summarizing, the `xml_stream_parser` package is the fastest in terms of execution time.

## cpuprofile

```
(pprof) list Benchmark_etree
Total: 165.67s
ROUTINE ======================== xml-perf-test/pkg/xml.Benchmark_etree in /home/rosorio/code/xml/pkg/xml/etree_test.go
         0     36.13s (flat, cum) 21.81% of Total
         .          .      9:func Benchmark_etree(b *testing.B) {
         .          .     10:   rawXML, err := data.LoadEntireFile("../data/input.xml")
         .          .     11:   if err != nil {
         .          .     12:           b.FailNow()
         .          .     13:   }
         .          .     14:
         .          .     15:   // run the function b.N times
         .          .     16:   for n := 0; n < b.N; n++ {
         .     36.13s     17:           Parse_etree(rawXML)
         .          .     18:   }
         .          .     19:}
(pprof) list Benchmark_native
Total: 165.67s
ROUTINE ======================== xml-perf-test/pkg/xml.Benchmark_native in /home/rosorio/code/xml/pkg/xml/native_test.go
         0     37.67s (flat, cum) 22.74% of Total
         .          .      8:func Benchmark_native(b *testing.B) {
         .       20ms      9:   rawXML, err := data.LoadEntireFile("../data/input.xml")
         .          .     10:   if err != nil {
         .          .     11:           b.FailNow()
         .          .     12:   }
         .          .     13:
         .          .     14:   // run the function b.N times
         .          .     15:   for n := 0; n < b.N; n++ {
         .     37.65s     16:           Parse_native(rawXML)
         .          .     17:   }
         .          .     18:}
(pprof) list Benchmark_xml_stream_parser
Total: 165.67s
ROUTINE ======================== xml-perf-test/pkg/xml.Benchmark_xml_stream_parser in /home/rosorio/code/xml/pkg/xml/xml-stream-parser_test.go
         0      3.98s (flat, cum)  2.40% of Total
         .          .      8:func Benchmark_xml_stream_parser(b *testing.B) {
         .       20ms      9:   rawXML, err := data.LoadEntireFile("../data/input.xml")
         .          .     10:   if err != nil {
         .          .     11:           b.FailNow()
         .          .     12:   }
         .          .     13:
         .          .     14:   // run the function b.N times
         .          .     15:   for n := 0; n < b.N; n++ {
         .      3.96s     16:           Parse_xml_stream_parser(rawXML)
         .          .     17:   }
         .          .     18:}
```

and the following (same that can be observed in the PNG files):

```
(pprof) top Benchmark_etree
Active filters:
   focus=Benchmark_etree
Showing nodes accounting for 21.61s, 13.04% of 165.67s total
Dropped 145 nodes (cum <= 0.83s)
Showing top 10 nodes out of 77
      flat  flat%   sum%        cum   cum%
     4.24s  2.56%  2.56%      7.27s  4.39%  encoding/xml.(*Decoder).getc
     3.08s  1.86%  4.42%     11.94s  7.21%  runtime.mallocgc
     2.90s  1.75%  6.17%      2.90s  1.75%  unicode/utf8.DecodeRune
     2.66s  1.61%  7.77%      2.76s  1.67%  bufio.(*Reader).ReadByte
     2.39s  1.44%  9.22%      9.84s  5.94%  encoding/xml.(*Decoder).text
     1.82s  1.10% 10.32%      1.82s  1.10%  runtime.nextFreeFast (inline)
     1.32s   0.8% 11.11%      1.39s  0.84%  runtime.writeHeapBits.flush
     1.30s  0.78% 11.90%      1.73s  1.04%  bytes.(*Buffer).WriteByte
     1.08s  0.65% 12.55%      1.08s  0.65%  runtime.memclrNoHeapPointers
     0.82s  0.49% 13.04%      0.83s   0.5%  unicode.is16
(pprof) top Benchmark_native
Active filters:
   focus=Benchmark_native
Showing nodes accounting for 20.57s, 12.42% of 165.67s total
Dropped 178 nodes (cum <= 0.83s)
Showing top 10 nodes out of 83
      flat  flat%   sum%        cum   cum%
     4.56s  2.75%  2.75%      6.52s  3.94%  encoding/xml.(*Decoder).getc
     3.05s  1.84%  4.59%      3.05s  1.84%  unicode/utf8.DecodeRune
     2.72s  1.64%  6.24%      8.37s  5.05%  runtime.mallocgc
     2.63s  1.59%  7.82%      9.82s  5.93%  encoding/xml.(*Decoder).text
     1.81s  1.09%  8.92%      1.81s  1.09%  bytes.(*Reader).ReadByte
     1.34s  0.81%  9.72%     37.65s 22.73%  encoding/xml.(*Decoder).unmarshal
     1.27s  0.77% 10.49%      1.27s  0.77%  runtime.nextFreeFast (inline)
     1.15s  0.69% 11.18%      1.53s  0.92%  bytes.(*Buffer).WriteByte
     1.07s  0.65% 11.83%      1.82s  1.10%  reflect.(*rtype).Implements
     0.97s  0.59% 12.42%         1s   0.6%  runtime.writeHeapBits.flush
(pprof) top Benchmark_xml_stream_parser
Active filters:
   focus=Benchmark_xml_stream_parser
Showing nodes accounting for 3.04s, 1.83% of 165.67s total
Dropped 51 nodes (cum <= 0.83s)
Showing top 10 nodes out of 75
      flat  flat%   sum%        cum   cum%
     1.19s  0.72%  0.72%      3.96s  2.39%  xml-perf-test/pkg/xml.Parse_xml_stream_parser
     0.59s  0.36%  1.07%      0.75s  0.45%  runtime.mapaccess1_faststr
     0.31s  0.19%  1.26%      0.64s  0.39%  strings.Trim
     0.23s  0.14%  1.40%      0.23s  0.14%  runtime.memclrNoHeapPointers
     0.22s  0.13%  1.53%      0.22s  0.13%  strings.(*asciiSet).contains (inline)
     0.15s 0.091%  1.62%      0.42s  0.25%  runtime.chanrecv
     0.11s 0.066%  1.69%      0.11s 0.066%  memeqbody
     0.11s 0.066%  1.76%      0.13s 0.078%  runtime.lock2
     0.07s 0.042%  1.80%      0.79s  0.48%  runtime.mallocgc
     0.06s 0.036%  1.83%      0.06s 0.036%  runtime.madvise
```
