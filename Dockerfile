FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

USER 1000

COPY --from=builder --chown=1000:1000 /app/main /app/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

ENV MODE=prod \
    MONGO_HOST=mongo \
    MONGO_PORT=27017 \
    MONGO_USER=rtyop \
    MONGO_PASS=123456 \
    MONGO_DATABASE=rcbs \
    RCBS_PORT=3000 \
    RCBS_MAIL_DOMAIN=localhost \
    RCBS_BASE_URI=http://localhost:3000

CMD ["/app/main"]