FROM golang:latest as stage
#RUN go get -u github.com/pressly/goose/cmd/goose
ADD . /go/goose
WORKDIR /go/goose
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o goose ./cmd
RUN chmod 775 ./goose

#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /go/goose
#COPY  --from=stage /go/goose .