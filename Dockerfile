FROM golang:1.9.2-stretch

RUN apt-get update && \
    apt-get upgrade -y

# install revel and revel-cli
RUN go get github.com/revel/revel && \
    go get github.com/revel/cmd/revel

# Add our code
ADD . /go/src/near_me_server
WORKDIR /go/src/near_me_server

# Commentize it before deploy(only used for dev)
# EXPOSE 9000

CMD revel run near_me_server prod $PORT
