gocrawl
======

GoCrawl is command line tool that does network discovery by crawling all network nodes an storing them in a database file.

Install
-------

go get github.com/ThomasZumsteg/gocrawl


Usage
-----

gocrawl <ip_address> 

Begin the crawl

options include
-background to run in the background
-database <file> use the specified database file
-debug <level> set the debug level
-ignore pattern skips hosts that match some pattern
-verbose run in verbose mode

