# Log format parser
*logparse* is a library written in [golang](http://golang.org) to parse the most common log formats.

# Usage
First install the library with:
```
go get xojoc.pw/logparse
```
*logparse* is simple to use. First parse a string with either [logparse.Common](http://godoc.org/xojoc.pw/logparse#Common) or [logparse.Combined](http://godoc.org/xojoc.pw/logparse#Combined) and then access the field of [logparse.Entry](http://godoc.org/xojoc.pw/logparse#Entry) for the required information. Examples:
 * [Common log format](http://godoc.org/xojoc.pw/logparse#example-Common)
 * [Combined log format](http://godoc.org/xojoc.pw/logparse#example-Combined)

see [godoc](http://godoc.org/xojoc.pw/logparse) for the complete documentation.

# Log formats
Right now *logparse* can parse the common and combined log formats. Support is under way for the extended log format.

# Who?
*logparse* was written by [Alexandru Cojocaru](http://xojoc.pw).

# License
*logparse* is released under the GPLv3 or later, see [COPYING](COPYING).
