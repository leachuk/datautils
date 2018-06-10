# datautils
Using [mappify geocode api](http://mappify.io/docs/#api-Geocoding-PostApiRpcAddressGeocode) to parse address data and output lat long to csv

## usage
`./main -csvfile=datasample.csv -debug`

### inputs
* `-csvfile` (required) specify the data input file for parsing by mappify. The colunm headings must match those in the datasample.csv file
* `-debug` (optional) print debug output

### output
* `output.csv` The output of the geocoded lat/long to a csv file.
