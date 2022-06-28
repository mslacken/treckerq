

trecker:    
	cd cmd/trecker; GOOS=linux go build -mod vendor -o ../../trecker

clean:
	rm -f trecker

vendor:
	go mod tidy -v    
	go mod vendor

all: trecker

.PHONY: all
