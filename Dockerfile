# Dockerfile
FROM golang:1.20

# Install air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /usr/local/bin

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go install -v golang.org/x/tools/gopls@latest

CMD ["air"]
