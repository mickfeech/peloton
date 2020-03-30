# peloton
Peloton API Go Client

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
