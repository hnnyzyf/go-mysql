# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean

build:
		$(GOBUILD) -v -o ./bin/binary-advisor ./cmd/advisor/binary.go
		$(GOBUILD) -v -o ./bin/http-advisor ./cmd/advisor/http.go

clean:
		$(GOCLEAN)
		rm -f ./bin/binary-advisor
		rm -f ./bin/http-advisor

