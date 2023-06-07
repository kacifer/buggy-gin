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

### FakeStatusCode

Pass "X-Fake-Status-Code: 200" in request header to fake status code.

Response body can also be faked by passing "X-Fake-Response: test" in request header.

### FakeResponseTime

Pass "X-Fake-Response-Seconds: 10" in request header to fake response time.
