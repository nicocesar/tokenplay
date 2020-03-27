FROM ubuntu
ADD tokenplay-srv /tokenplay-srv
ENTRYPOINT [ "/tokenplay-srv" ]
