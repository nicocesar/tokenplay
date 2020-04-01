FROM alpine
ADD tokenplay-service /tokenplay-service
ENTRYPOINT [ "/tokenplay-service" ]
