FROM golang:1.13 as build
WORKDIR /go/src/app
COPY . .
#RUN go mod tidy && go mod verify && go mod vendor
RUN go build -o server

FROM ubuntu
RUN apt-get update &&\
	apt-get install -y \
	curl \
	git-core \
	make \
	wget \
	vim
COPY --from=build /go/src/app/server .
CMD ["./server"]
