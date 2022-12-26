.PHONY: build maria

image:
	docker build --no-cache \
	-t sample-go-api:test -f Dockerfile .

container:
	docker run -p:8081:8081  --env-file ./dev.env \
	--name samplegoapi sample-go-api:test

run 

docs:	
	swag init -d ./cmd/http --parseDependency  -o docs 
	
# swag init --dir ./cmd/api/main.go --parseDependency --output docs

installvegeta:
	go install github.com/tsenart/vegeta@latest
vegeta:
	echo "GET http://:8081/limitz" | vegeta attack -rate=10/s -duration=1s | vegeta report
load:
	echo "GET http://:8081/limitz" | vegeta attack -rate=10/s -duration=1s > results.10qps.bin
plot:
	 cat results.10qps.bin | vegeta plot > plot.10qps.html
hist:
	cat results.10qps.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"

# APP_ENV=dev go run main.go
# make image, make container

# APP_ENV=dev go run ./cmd/http/main.go   
# go mod init macus  
# go get -u -v -f all
