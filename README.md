# Routing service

To read more about the project, please read the official [design plan](https://github.com/Greennav/greennav.github.io/blob/master/wiki/Roadmap.md#design-plan). This repository consists of a simple mockup server for testing purposes.

## Next steps

- Agree on a communication protocol (refer to the [mockup-response.json](mockup-response.json) as an example)
- Establish mockup connectivity to the front end and the database service mockup

## Running the server

1. Make sure you have Go installed `sudo apt-get install golang`
2. Either run with `go run mockup.go` or with 
```bash
    go build mockup.go
    ./mockup (or mockup.exe)
```
