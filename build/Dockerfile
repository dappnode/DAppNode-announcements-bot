# BUILDER
FROM golang as builder

WORKDIR /usr/src/app
RUN apt update && apt install gcc
COPY . .
WORKDIR /usr/src/app/src
RUN go get -v -t ./...
RUN go build -o main main.go

# RUNNER
FROM golang as runner
WORKDIR /usr/src/app
ENV GO_ENV=development
COPY --from=builder /usr/src/app/src/main .
COPY --from=builder /usr/src/app/src/*.env .
ENTRYPOINT [ "/usr/src/app/main" ]