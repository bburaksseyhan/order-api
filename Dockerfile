FROM golang:1.17-alpine as build-env
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN  go build -o /order-api github.com/bburaksseyhan/orderapi/src/cmd/api   

FROM alpine:3.14

RUN apk update \
    && apk upgrade\
    && apk add --no-cache tzdata curl

#RUN apk --no-cache add bash
ENV TZ Europe/Istanbul

WORKDIR /app
COPY --from=build-env /order-api .
COPY --from=build-env /app/src/cmd/api /app/

EXPOSE 80
CMD [ "./order-api" ]