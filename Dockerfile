FROM golang:1.18.2-alpine as build

WORKDIR /src

COPY build/app/go.mod go.mod
COPY build/app/go.sum go.sum

RUN go mod download

COPY build/app/cmd cmd/
COPY build/app/models models/
COPY build/app/restapi restapi/
COPY build/app/twitter twitter/

ENV CGO_LDFLAGS "-static -w -s"

RUN go build -tags osusergo,netgo -o /application cmd/twitter-server/main.go; 

FROM ubuntu:22.04

RUN apt-get update && apt-get install ca-certificates curl ruby -y
RUN gem install twurl

# DON'T CHANGE BELOW 
COPY --from=build /application /bin/application

EXPOSE 8080

CMD ["/bin/application", "--port=8080", "--host=0.0.0.0", "--write-timeout=0"]