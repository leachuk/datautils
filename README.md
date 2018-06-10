# datautils
Using [mappify geocode api](http://mappify.io/docs/#api-Geocoding-PostApiRpcAddressGeocode) to parse address data and output lat long to csv

## compile
`go build`

`GOOS=windows GOARCH=386 go build` to cross compile. See https://golang.org/doc/install/source#environment for alternative options.

## run
`./datautils -csvfile=datasample.csv -debug`

#### inputs
* `-h` or `--help` (optional) output help
* `-csvfile` (required) specify the data input file for parsing by mappify. The colunm headings must match those in the datasample.csv file
* `-debug` (optional) print debug output

#### output
* `output.csv` The output of the geocoded lat/long to a csv file.
