Go Exam
====

This project just exam to enhancement my about Go Language as RestAPI.  

## How To
This project is already using Go Module, So i recommended to put this source code outside GOPATH.  

### Testing
To testing you can run with this command :
```bash
exam $ go test -v -cover -covermode=atomic ./...
```

or you prever using make:
```bash
exam $ make test
```

### Running application

Here is the steps to run it with `docker-compose`
```bash
# Build docker image
exam $ make docker

# Run application
exam $ make run

# Stop application
exam $ make stop
```

