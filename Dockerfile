FROM golang:1.11.4-alpine3.8 as builder

RUN apk --no-cache add git linux-headers ca-certificates
COPY . /go/src/github.com/ShyftNetwork/blockexplorer_api
WORKDIR /go/src/github.com/ShyftNetwork/blockexplorer_api
ENV GO111MODULE=on
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blockx_api .


FROM scratch

WORKDIR /
COPY --from=builder ./go/src/github.com/ShyftNetwork/blockexplorer_api/blockx_api .
CMD ["./blockx_api"]
EXPOSE 8080