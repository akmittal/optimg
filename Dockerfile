FROM ubuntu

WORKDIR /pixer

RUN wget -q -O - https://git.io/vQhTU | bash

COPY . ./

RUN apt-get update -y
RUN apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_14.x |  bash -
RUN  apt-get install -y nodejs

RUN cd client && npm install --silent
RUN cd client && npm run build
RUN export CGO_CFLAGS_ALLOW=-Xpreprocessor
RUN cd server && go get ./...
RUN go install -v ./...
CMD ["server"]
