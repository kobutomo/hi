FROM golang:1.20.3-alpine as build
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/app ./

FROM gcr.io/distroless/base-debian10 as deploy
COPY --from=build /go/bin/app /
EXPOSE 8080
CMD ["/app"]