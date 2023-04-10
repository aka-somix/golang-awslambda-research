module github.com/aka-somix/crud-api/create-todo

go 1.20

require (
	github.com/aka-somix/crud-api/_shared v1.0.0
	github.com/aws/aws-lambda-go v1.39.1
	github.com/valyala/fastjson v1.6.4
)

require github.com/google/uuid v1.3.0 // indirect

replace github.com/aka-somix/crud-api/_shared => ../_shared
