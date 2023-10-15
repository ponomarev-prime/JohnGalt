# GO GO GO

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=/home/alex/go
```

```
curl -w "\n" localhost:8080/echo?data=Hellow%20World
```

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem
```

```
go install JohnGalt
```
```
> ll $GOPATH/bin
total 6752
drwxr-xr-x 2 alex alex    4096 Sep  6 23:39 ./
drwxr-xr-x 4 alex alex    4096 Sep  6 23:39 ../
-rwxr-xr-x 1 alex alex 6904712 Sep  6 23:39 JohnGalt*
```


# Args

```
PS C:\Users\ponom\Documents\CODE\JohnGalt> go run .\args.go Yo yo yo 1 11 111111111111 -1 0 //
Yo yo yo 1 11 111111111111 -1 0 //
```