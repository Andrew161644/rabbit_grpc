FROM andrewfromgilead/dependency_go_image AS build_base

COPY go.mod go.sum /go/github.com/Andrew161644/currency_exchange_grpc/
WORKDIR /go/github.com/Andrew161644/currency_exchange_grpc
RUN go mod download
COPY . /go/github.com/Andrew161644/currency_exchange_grpc
WORKDIR /go/github.com/Andrew161644/currency_exchange_grpc/api

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/avicks_laba github.com/Andrew161644/currency_exchange_grpc/api

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build_base /go/github.com/Andrew161644/currency_exchange_grpc/api/build/avicks_laba /app/go-sample-app


EXPOSE 8080
CMD ["/app/go-sample-app"]