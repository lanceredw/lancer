FROM golang:1.20-alpine

WORKDIR /app

COPY bin/lancer-linux64 /app/lancer-linux64

COPY settings.yml /app/settings.yml

EXPOSE 8080

CMD ["./lancer-linux64"]

