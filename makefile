module ?= "."

build:
	cd ${module} && sam build

debug: build-sam
	cd ${module} && sam local invoke -d 9999
