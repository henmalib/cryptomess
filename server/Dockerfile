FROM cosmtrek/air

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

ENTRYPOINT [ "air"]
