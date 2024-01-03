.PHONY: all run clean help

APP = lancer

## linux: compile and package linux files
.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-linux64 ./main.go

## win: compile and package win files
.PHONY: win
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-win64.exe ./main.go

## mac: compile and package mac files
.PHONY: mac
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-darwin64 ./main.go

build:
	@go build -o ${APP}

## compile win，linux，mac platform
.PHONY: all
all:win linux mac

run:
	@go run ./

.PHONY: tidy
tidy:
	@go mod tidy

## test: Run unit test.
.PHONY: test
test:
	@$(MAKE) go.test

## Clean up binary files
clean:
	@if [ -f ./bin/${APP}-linux64 ] ; then rm ./bin/${APP}-linux64; fi
	@if [ -f ./bin/${APP}-win64.exe ] ; then rm ./bin/${APP}-win64.exe; fi
	@if [ -f ./bin/${APP}-darwin64 ] ; then rm ./bin/${APP}-darwin64; fi

help:
	@echo "make - format Go code and compile to generate binary files"
	@echo "make mac - compile Go code to generate binary files for mac"
	@echo "make linux - compile Go code to generate linux binary files"
	@echo "make win - compile Go code to generate windows binary files"
	@echo "make tidy - execute go mod tidy"
	@echo "make run - run Go code directly"
	@echo "make clean - remove compiled binary files"
	@echo "make all - compile binary files for multiple platforms"