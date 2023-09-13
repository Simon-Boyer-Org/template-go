FROM golang:1.19

COPY app.go ./

RUN CGO_ENABLED=0 GOOS=linux go build app.go
RUN ls

# Run
CMD ["./app"]