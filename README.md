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
-b, --background to run in the background
-d, --database <file> use the specified database file
-v, --verbose run in verbose mode
--debug <level> set the debug level

