# buggy-gin

Gin middleware for debugging.

## Usage

```go
r := gin.Default()

buggy_gin.UseAll(r)

return r
```

See examples for more details.

## Available middlewares

### PrintHeaders

Pass "X-Print-Headers: true" in request header to print all headers.

If env "PRINT_HEADERS" is set, headers will always be printed.

### FakeStatusCode

Pass "X-Fake-Status-Code: 200" in request header to fake status code.

Response body can also be faked by passing "X-Fake-Response: test" in request header.

If env "FAKE_STATUS_CODE" is set, status code will always be faked.

### FakeResponseTime

Pass "X-Fake-Response-Milliseconds: 1000" in request header to fake response time.

If env "FAKE_RESPONSE_TIME" is set, response time will always be faked.
