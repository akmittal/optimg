FROM golang:buster

WORKDIR /optimg


COPY . ./

RUN apt-get update -y && apt-get install -y curl libvips libvips-dev
RUN curl -sL https://deb.nodesource.com/setup_14.x |  bash -
RUN  apt-get install -y nodejs

RUN cd client && npm install --silent
RUN cd client && npm run build
RUN export CGO_CFLAGS_ALLOW=-Xpreprocessor
RUN cd server && go get ./...
RUN cd server && go install -v ./...
RUN cd server && go build src/main.go
CMD ["./server/main"]
