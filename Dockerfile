
## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY . ./
RUN go mod download

RUN go build -o /main

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main
COPY --from=build /app/all-match.json .
COPY --from=build /app/all-season.json .
COPY --from=build /app/all-team.json .

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/main"]