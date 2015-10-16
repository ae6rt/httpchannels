### Build and run

```
$ go build
$ ./httpchannels
```

### Do some reads

```
$ curl http://localhost:8080/read 
2015-10-10 17:20:41.63585846 -0700 PDT

$ curl http://localhost:8080/read 
2015-10-10 17:20:41.63585846 -0700 PDT
```

### Do an update

```
$ curl http://localhost:8080/update 
```

### Read update

```
$ curl http://localhost:8080/read 
2015-10-10 17:21:02.912332905 -0700 PDT

$ curl http://localhost:8080/read 
2015-10-10 17:21:02.912332905 -0700 PDT
```
