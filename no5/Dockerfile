#Build Stage
FROM golang:1.17.8-alpine as build
WORKDIR /app/
COPY . .
RUN go build -o main .

#Production Stage
FROM alpine
RUN apk --update add curl
COPY --from=build /app/main /main
EXPOSE 8080
CMD ["./main"]