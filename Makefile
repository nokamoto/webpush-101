
protoc:
	protoc -I . protobuf/*.proto --go_out=plugins=grpc:webpush-go

format:
	prototool format -w protobuf
	gofmt -w ./webpush-go/push-subscription
	cd webpush-scala/front && sbt scalafmt test:scalafmt sbt:scalafmt
