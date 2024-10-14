# Use the official Golang image
FROM golang:1.22.3-alpine

WORKDIR /app

ENV API_PORT=8081 

ENV TOKEN_SECRET=asdknasjdbakjbdiuawbeiybajkdnkasndmhkbfihagwiura 
ENV TOKEN_TIME_ACCESS=150 
ENV TOKEN_TIME_REFRESH=168 
ENV TOKEN_REMEMBER_REFRESH=8760 

ENV DB_NAME=didlydoodash 
ENV DB_HOST=localhost 
ENV DB_PORT=5432 
ENV DB_USER=didlydoodash 
ENV DB_PASSWORD=didlydoodash 
ENV DB_SSL=disable
ENV DB_TIMEZONE=Europe/Helsinki

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o didlydoodash-api ./src/cmd/api
RUN go build -o didlydoodash-drop ./src/cmd/drop
RUN go build -o didlydoodash-migrate ./src/cmd/migrate

# Add the startup script
COPY start.sh ./
RUN chmod +x start.sh

EXPOSE 8081

# Run the startup script
CMD ["./start.sh"]