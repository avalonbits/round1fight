FROM golang:1.21 as builder

WORKDIR /app
COPY . ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -tags 'osusergo,netgo,static' --ldflags '-linkmode external -extldflags "-static"' -o /app/round1fight ./cmd/round1fight

FROM alpine:latest
RUN apk update && apk add --no-cache libc6-compat

EXPOSE 1323

COPY --from=builder /app/round1fight .
COPY --from=builder /app/.env .
COPY --from=builder /app/run.sh .

CMD ["./run.sh"]
