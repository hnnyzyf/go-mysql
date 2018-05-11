# Go parameters
GOCMD = go
GOYACC = ./bin/goyacc
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean

build:
		$(GOBUILD) -v -o ./bin/goyacc ./vendor/github.com/golang/tools/cmd/goyacc/yacc.go
		$(GOYACC) -o ./advisor/parser/yacc.go -v ./advisor/parser/yacc.output ./advisor/parser/parser.y
		$(GOBUILD) -v -o ./bin/binary-advisor ./cmd/advisor/binary.go
		$(GOBUILD) -v -o ./bin/http-advisor ./cmd/advisor/http.go

clean:
		$(GOCLEAN)
		rm -f ./bin/binary-advisor
		rm -f ./bin/http-advisor

