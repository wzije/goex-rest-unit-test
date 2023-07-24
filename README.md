# Go RestAPI Unit Test Example

This project is an example of rest api unit test in golang. It's for learning purpose.

## How to run

1. Run project

   ```go run main.go```

2. Run test

   ```test go test ./... -v```

3. coverage test and convert file to html

   ```go test ./... -v --coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html```

### How to Build

```go build main.go```