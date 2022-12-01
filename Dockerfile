#-------------- Build ----------
FROM golang:1.16-buster AS build

WORKDIR /app

## copy code from local dir to container
COPY . ./

RUN go build -o /simple-web-server

#-------------- Deploy ----------
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /simple-web-server /simple-web-server

COPY ./static ./static

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT ["/simple-web-server"]



