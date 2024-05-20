FROM golang:1.20-alpine
RUN mkdir /app
WORKDIR /app

LABEL "co.elastic.logs/enabled"=true
LABEL "co.elastic.logs/json.keys_under_root"=true
LABEL "co.elastic.logs/json.overwrite_keys"=true
LABEL "co.elastic.logs/json.add_error_key"=true
LABEL "co.elastic.logs/json.expand_keys"=true

COPY ./api/go.mod ./
COPY ./api/go.sum ./
RUN go mod download
RUN go mod verify

COPY . ./

RUN go build -o myapp

ENTRYPOINT ["./myapp"]