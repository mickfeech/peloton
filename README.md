# peloton
Peloton API Go Client

[![Go](https://github.com/mickfeech/peloton/workflows/Go/badge.svg)](https://github.com/mickfeech/peloton/actions?query=workflow%3AGo) 
[![codecov](https://codecov.io/gh/mickfeech/peloton/branch/master/graph/badge.svg)](https://codecov.io/gh/mickfeech/peloton)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickfeech/peloton)](https://goreportcard.com/report/github.com/mickfeech/peloton)
## Features

* Get data about user as struct
* Get data about instructors as struct

## Usage

```go
	username := os.Getenv("username")
	password := os.Getenv("password")
	var client *peloton.PelotonClient
	client = peloton.NewPelotonClient(username, password)
	me, _ := client.Me()
	fmt.printf("%s" me)
```
