FROM golang:alpine
ENV POSTGRES_DB=graphql_api \
    POSTGRES_USER=graphql_api \
    POSTGRES_PASSWORD=graphql_api \
    POSTGRES_URL=database \
    API_KEY=18d71e3df5eeec9c5073a226c11d3136
WORKDIR /app
COPY src /app
RUN go mod download
WORKDIR /app/cmd/go-graphql
EXPOSE 8080
CMD ["go", "run", "server.go"]