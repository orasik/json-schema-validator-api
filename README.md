# Data Validation API
Data validation API example to validate POST json request against Json Schema with go lang.


### Packages
- [Gin Framework](https://gin-gonic.github.io/gin/)
- [xeipuuv/gojsonschema](https://github.com/xeipuuv/gojsonschema)
- [logrus](https://github.com/sirupsen/logrus)

### How to run

#### 1) If you have golang installed:

- Clone the repo
- go get -v && go build -o api
- ./api --port 8080


Make a valid request:

```bash
curl -X POST \
  http://localhost:8080/api/person \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"firstName": "Go",
	"lastName": "lang"
}'
```

response would be:

```json
{
    "success": true,
    "errors": {}
}
```

Invalid request:

```bash
curl -X POST \
  http://localhost:8080/api/person \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
	"noName": "Go",
	"lastName": "lang"
}'
```

response:

```json
{
    "success": false,
    "errors": {
        "firstName": "firstName is required"
    }
}
```

### 2) With Docker (coming soon ...)


### Running test

```bash
go test -v
```