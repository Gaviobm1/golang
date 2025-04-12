module example.com/hello

go 1.24.2

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.7.0
)
