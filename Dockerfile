
## Build
FROM golang:1.17-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /main

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main
COPY --from=build /app/.env /
COPY --from=build /app/lck-data-project-firebase-adminsdk-x03sw-7009ebcef2.json /

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]