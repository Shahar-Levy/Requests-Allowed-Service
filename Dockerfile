FROM golang:latest
RUN mkdir /app
EXPOSE 80
ADD . /app
RUN cd /app; go get; go build; 
ENTRYPOINT [ "/app/RequestsAllowedService" ]