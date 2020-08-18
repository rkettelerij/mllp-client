Golang MLLP client
==================

Simple CLI tool written in [Go](http://golang.org) to send messages over [MLLP](http://en.wikipedia.org/wiki/Health_Level_7#MLLP) (Minimal Lower Layer Protocol). MLLP is widely used in healthcare for amongst others HL7 messaging.

Usage
-----

```
Usage of mllp-client:
  -file="<filename>": path to file which contents will be send to the MLLP server
  -dir="<directory>": path to directory which will have all files within it sent to the MLLP server
  -host="localhost": hostname of MLLP server, default value is localhost
  -port=2575: portnumber of MLLP server, default value is 2575
```

Example
-------

Send file to localhost:2575

```
./mllp-client -file /mllp-client/src/adt_a01.txt
```

Send all files in a directory to localhost:2575

```
./mllp-client -dir /mllp-client/src/hl7_files/

Send file to messagebroker.example.com on port 7890

```
./mllp-client -file /mllp-client/src/adt_a01.txt -host messagebroker.example.com -port 7890
```

Building
--------

Use the excellent [goxc](https://github.com/laher/goxc) tool to create cross-platform builds. For example to
create binaries for Linux, Windows and Mac use:

```
git clone https://github.com/rkettelerij/mllp-client.git
cd mllp-client
goxc -bc "linux, windows, darwin"
```

License
-------

```
The MIT License (MIT)

Copyright (c) 2014 Richard Kettelerij

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
```
