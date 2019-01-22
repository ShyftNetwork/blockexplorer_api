FROM golang:1.11.4-alpine3.7 as builder

RUN apk add --update git make gcc musl-dev linux-headers ca-certificates

COPY . /go/src/github.com/ShyftNetwork/blockexplorer_api
WORKDIR /go/src/github.com/ShyftNetwork/blockexplorer_api

ENV GO111MODULE=on

RUN make install
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shyft_api .


FROM scratch

COPY --from=builder ./go/src/github.com/ShyftNetwork/blockexplorer_api/shyft_api .
ADD shyft_api /
CMD ["/shyft_api"]
EXPOSE 8080