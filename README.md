# web app sqlite3 and Golang

Required:
go version go1.11


## Usage

1. To check the current value of your $GOPATH

```
echo $GOPATH
```

2. If you do not see a directory name, add the GOPATH variable to your environment:

```
export $GOPATH="/location/of/your/gopath/directory"
```

3. Inside the Go directory, create the directory struccture

```
mkdir src
mkdir pkg
mkdir bin
```

4. In src

```
go get github.com/kaarla/proyecto-DEMO-Node
```
```
dep ensure
```
```
go build
```
```
./proyecto-DEMO
```
