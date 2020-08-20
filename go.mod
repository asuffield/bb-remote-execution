module github.com/buildbarn/bb-remote-execution

go 1.14

require (
	github.com/aws/aws-sdk-go v1.34.9
	github.com/bazelbuild/remote-apis v0.0.0-20200708200203-1252343900d9
	github.com/buildbarn/bb-storage v0.0.0-20200817065905-ee662747ade7
	github.com/golang/mock v1.4.4
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.7.1
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux v0.11.0
	go.opentelemetry.io/otel v0.11.0
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642
	google.golang.org/genproto v0.0.0-20200815001618-f69a88009b70
	google.golang.org/grpc v1.31.0
)
