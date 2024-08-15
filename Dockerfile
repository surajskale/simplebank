#Build stage
FROM golang:1.22.6-alpine3.19 AS builder 
# builder is name of the stage
WORKDIR /app 
COPY . .
RUN go build -o main main.go


#run stage
FROM alpine:3.19 
WORKDIR /app
# alpine3.19 as base image 

#copying main build to working dir to reduce image size
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
CMD ["/app/main"]
