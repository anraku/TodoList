# TodoList
![sample](https://user-images.githubusercontent.com/1418311/41823600-4f37587c-783d-11e8-9db3-858f0f32cf40.png)

# Dependency
```
$ go version
go version go1.10 darwin/amd64

$ node -v
v9.8.0

$ npm -v
5.8.0

// you need to install https://github.com/golang/dep
$ dep version
dep:
 version     : v0.4.1
 build date  : 2018-01-27
 git hash    : 37d9ea0
 go version  : go1.9.3
 go compiler : gc
 platform    : darwin/amd64
 
```

# Setup
Please build .go file in backend/service
```
cd TodoList/backend/service
GOOS=linux go build -o api main.go handler.go
```

# Usage
```
cd TodoList/infra
docker-compose up --build -d
```

# Licence
MIT

# Authors
Dai Omori (a.k.a anraku)

# References
https://mattn.kaoriya.net/software/lang/go/20180330093346.htm

https://github.com/mattn/go-vue-example
