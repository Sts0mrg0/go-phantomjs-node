docker run -ti --rm -p 6677:6677 -e "token=arkady" -e "server=192.168.4.193" node

#CGO_ENABLED=0 GOOS=linux go build -race -a -installsuffix cgo -o node .
#только с образом scratch
