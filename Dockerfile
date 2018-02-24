FROM golang:1.9

ENV PATH ${GOPATH}/bin:$PATH
WORKDIR ${GOPATH}/src/LampIO
COPY . .

RUN go build -o LampIO CLI/Main.go
EXPOSE 7778

ARG TYPE="server"
RUN ./LampIO --$TYPE