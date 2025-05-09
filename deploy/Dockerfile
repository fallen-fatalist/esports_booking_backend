FROM golang:alpine as builder

WORKDIR /app

COPY . /app

RUN go build -o /app/bin/booking-service /app

FROM alpine:latest as runner

WORKDIR /app 

# Copy binary
COPY --from=builder /app/bin /app



EXPOSE 8080

CMD ["/app/booking-service", "--port", "8080"]