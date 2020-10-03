# [Share-Price-Export-Using-Go](https://github.com/VagueCoder/Share-Price-Export-Using-Go)
Export Bulk Share Prices Easily &amp; Quickly Using Golang's Executable.

## Base System Configurations :wrench:
**Sno.** | **Name** | **Version/Config.**
-------: | :------: | :------------------
1 | Operating System | Windows 10 x64 bit
2 | Language | Go Version 1.14.7 Windows/amd64
3 | IDE | Visual Studio Code Version 1.49.3

## What is [Go (Golang)](https://golang.org/)?
Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as Golang because of its domain name, golang.org, but the proper name is Go.

Go is an open source programming language that makes it easy to build simple, reliable, and efficient software. You might want to check their Git Repo. [Click Here](https://github.com/golang/go).

## What is Share Market?
Share market is where buying and selling of share happens. Share represents a unit of ownership of the company from where you bought it. By buying share, you are investing money in the company. As the company grows, the price of your share too will increase. You can get profit by selling the shares in the market. There are various factors that affect the price of a share. Sometimes the price can rise and sometimes it can fall. Long term investment will nullify the fall in price.

1. NSE (National Stock Exchange)
1. BSE (Bombay Stock Exchange)

These are the two major stock exchanges in India and are regulated by SEBI (Securities and Exchange Board of India). Brokers act as an intermediary between the stock exchange and the investors.

## Why is this Program?
The program is for scraping the list of companies and the corresponding BSE prices alone, from the site [MoneyControl.com](https://www.moneycontrol.com/). All the data hence scrapped will be exported to a CSV (Comma-Separated Values) file in the same directory as the executable file. The aim to to automate the fetching and boosting the speed maximum possible. Hence, Go!

## Imports :package:
The program in developed using Go and has the following imports used.
Sn | **Import** | **Purpose**
-: | :--------: | :-------
1 | [os](https://golang.org/pkg/os/) | For a platform-independent interface to operating system functionality.
2 | [fmt](https://golang.org/pkg/fmt/) | For implementing formatted I/O with functions analogous to C's printf and scanf.
3 | [log](https://golang.org/pkg/log/) | Simple logging package.
4 | [net/http](https://golang.org/pkg/net/http/) | Provides HTTP client and server implementations.
5 | [time](https://golang.org/pkg/time/) | Provides functionality for measuring and displaying time.
6 | [path](https://golang.org/pkg/path/) | Implements utility routines for manipulating slash-separated paths.
7 | [path/filepath](https://golang.org/pkg/path/filepath/) | Implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.
8 | [regexp](https://golang.org/pkg/regexp/) | Implements regular expression search.
9 | [strconv](https://golang.org/pkg/strconv/) | Implements conversions to and from string representations of basic data types.
10 | [encoding/csv](https://golang.org/pkg/encoding/csv/) | Reads and writes comma-separated values (CSV) files.
11 | [encoding/json](https://golang.org/pkg/encoding/json/) | Implements encoding and decoding of JSON as defined in RFC 7159.
12 | [github.com/PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery) | Brings a syntax and a set of features similar to jQuery to the Go language.

The list of imported modules/packages are fetched using the command `go list -f "{{.ImportPath}} {{.Imports}}" ./...` in the same directory.

## [Code](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export.go)
The code is present at [github.com/VagueCoder/Share-Price-Export-Using-Go/Share-Price-Export.go](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export.go)
#### How to Run?
You can directly run the code as follows:
```
go run Share-Price-Export.go
```

or can build the executable using
```
go build Share-Price-Export.go
```
This creates an executable as [Share-Price-Export.exe](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export.exe), which can be used cross-platform in absence of Go.

## [Functional Model (or) Executable](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export.exe)
The executable of the Go is build as explained above and is ready for use, cross-platform, and doesn't require Go installed. Please check [Share-Price-Export.exe](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export.exe).

All you have to do is to run this from directory or through command prompt as follows. As simple as that.
```
Share-Price-Export.exe
```

You'll get command line output similar to this:
```
Exported to: "Share-Price-Export 03-Oct-2020 193614.csv"
Time Elapsed: 17m35.403775s (1055 Seconds)
```
`17 minutes for 8,500+ share's details-fetching is relatively very fast when compared with programs of other languages.`

## [Output](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export%2003-Oct-2020%20193614.csv)
The Output file will be having few columns of the share market companies and corresponding prices as you can see in [Share-Price-Export 03-Oct-2020 193614.csv](https://github.com/VagueCoder/Share-Price-Export-Using-Go/blob/master/Share-Price-Export%2003-Oct-2020%20193614.csv).
