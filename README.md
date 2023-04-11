Shared libraries and tools for Go gRPC services

Usage: `go get github.com/riza-io/go-grpc`

#### Credentials

##### Bearer tokens

In your client:
```go
opts := []grpc.DialOption{
    grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
    grpc.WithPerRPCCredentials(bearer.NewPerRPCCredentials("your bearer token")),
}

conn, err := grpc.Dial(hostname+":443", opts...)
```

Get a token from within your service handler:
```go
token, err := bearer.TokenFromContext(ctx)
```