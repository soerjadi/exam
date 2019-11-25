BINARY=engine
test: 
		go test -v -cover -covermode=atomic ./...

engine:
		go build -o ${BINARY} main.go router.go

clean:
		if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

docker:
		docker build -t go-exam .

run:
		docker-compose up -d

stop:
		docker-compose down

.PHONY: test engine clean docker run stop