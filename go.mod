// module go-experiment //module unique path
module github.com/tuhinal/go-experiment

go 1.23.0 // minimum compatible version of code

require github.com/tuhinal/go-experiment/go-basic v0.0.0-00010101000000-000000000000

replace github.com/tuhinal/go-experiment/go-basic => ./go-basic

replace github.com/tuhinal/go-experiment/concurrency => ./concurrency
