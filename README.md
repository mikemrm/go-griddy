# go-griddy
A simple Go package for pulling live pricing from GoGriddy.com

## Requirements
You will need to get your Meter ID, Member ID and Settlement Point from https://app.gogriddy.com

Look for a network request ending in /getnow. The Request Payload will contain all this information.

## Usage
```shell
$ go run cmd/query.go -help
Usage of query:
  -config string
    	Path to toml configuration file.
  -memberid string
    	Your Member ID
  -meterid string
    	Your meter ID
  -settlement string
    	Settlement Point (ex: LZ_HOUSTON)
  -url string
    	URL to be queried
```

```shell
go run cmd/query.go -config config.toml
go run cmd/query.go -memberid abc123 -meterid your-uuid -settlement LZ_YOUR_CITY
```

## Example

```shell
$ go run cmd/query.go -config config.toml
Server Time   : 2018-07-23T23:10:00Z
Settlement    : LZ_CITY
Price Type    : lmp
Price         : 3.30 ¢
Value Score   : 10
Mean Price    : 17.98 ¢
Diff Mean     : -14.67 ¢
High          : 206.04 ¢
Low           : 1.02 ¢
Local Time    : 2018-07-23T18:10:00-07:00
---------------
Data TTL      : 28 Sec.
```