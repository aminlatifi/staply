FROM golang:1.11.9
#RUN pip install Flask==0.11.1
#RUN pip install redis==2.10.5
RUN useradd -ms /bin/bash amin
USER amin
ADD staply $GOPATH/src/staply
WORKDIR $GOPATH/src/staply
RUN go get ./...
CMD ["go", "run", "main.go"] 
