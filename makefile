module ?= "go/hello-world"

sam-build:
	cd ${module} && sam build --use-container

debug: sam-build
	cd ${module} && sam local invoke -d 9999 --debug-args="-delveAPI=2" --debugger-path ${GOPATH}/bin/

invoke: sam-build
	cd ${module} && sam local invoke
