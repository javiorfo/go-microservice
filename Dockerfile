FROM alpine:latest  

WORKDIR /root/

COPY main .

EXPOSE 8080

CMD ["./main"]
