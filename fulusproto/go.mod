module github.com/khatibomar/fulus/fulusproto

go 1.27.2

replace github.com/khatibomar/fulus => ../

require (
	github.com/khatibomar/fulus v0.0.0
	google.golang.org/genproto v0.0.0-20261005182115-fad411399dd8
)

require google.golang.org/protobuf v1.36.12 // indirect
