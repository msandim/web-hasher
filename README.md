# Web Hasher
[![GoDoc](https://godoc.org/github.com/msandim//web-hasher?status.svg)](https://godoc.org/github.com/msandim//web-hasher)
[![GoReportCard](https://goreportcard.com/badge/github.com/msandim/web-hasher)](https://goreportcard.com/report/github.com/msandim/web-hasher)

This simple program prints the MD5 hashes of the contents of the URLs provided as arguments.

It also uses a pool of goroutines to perform the hashing process of the contents of each URL independently.

## Run program

Example usage:
```
go get -u github.com/msandim/web-hasher
$GOPATH/bin/web-hasher sapo.pt
```

You can use an optional flag `parallel` to specify how many goroutines should be created and maintained in the pool of goroutines.
If not specified, the default stays at 10.

```
$GOPATH/bin/web-hasher -parallel=3 sapo.pt google.com facebook.com publico.pt
```

## Run tests
```
go test -cover ./...  
```