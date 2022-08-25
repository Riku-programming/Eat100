FROM golang:latest

WORKDIR /go/src/app

COPY ./backend/go.mod .
COPY ./backend/go.sum .

RUN go install gorm.io/gorm \
    && go install gorm.io/driver/mysql \
    && go install github.com/cosmtrek/air@latest \
    && go install github.com/go-delve/delve/cmd/dlv@latest \
    && go get -u github.com/gin-gonic/gin

COPY ./backend/ .
ENV root="ROOT"

CMD ["air", "-c", ".air.toml"]