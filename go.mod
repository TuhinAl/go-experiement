// module go-experiment //module unique path
module github.com/tuhinal/go-experiment

go 1.23.0 // minimum compatible version of code

require (
	github.com/pborman/uuid v1.2.1
	github.com/tuhinal/go-experiment/go-basic v0.0.0-00010101000000-000000000000
)

require github.com/google/uuid v1.0.0 // indirect

replace github.com/tuhinal/go-experiment/go-basic => ./go-basic
