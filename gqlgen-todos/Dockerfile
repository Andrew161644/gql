FROM andrewfromgilead/dependency_go_image AS build_base

COPY go.mod go.sum /go/app/
WORKDIR /go/app
RUN go mod download
COPY . /go/app
WORKDIR /go/app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/app andrew.com/bff

FROM alpine:3.9
RUN apk add ca-certificates
COPY --from=build_base /go/app/build/app /app/go-sample-app

EXPOSE 8080
CMD ["/app/go-sample-app"]