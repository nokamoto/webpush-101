
protoc:
	protoc -I . protobuf/*.proto --go_out=plugins=grpc:webpush-go

format:
	prototool format -w protobuf
	gofmt -w ./webpush-go/push-subscription
	gofmt -w ./webpush-go/webpush
	gofmt -w ./webpush-go/webpush-lib
	cd webpush-scala/front && sbt scalafmt test:scalafmt sbt:scalafmt

vapid-new:
	openssl ecparam -name prime256v1 -genkey -noout -out vapid_private.pem

vapid-priv:
	openssl ec -in vapid_private.pem -text -noout -conv_form uncompressed | grep '^priv:' -A 3 | tail -n 3 | tr -d ' \n:' | xxd -r -p | base64

vapid-pub:
	openssl ec -in vapid_private.pem -text -noout -conv_form uncompressed | grep '^pub:' -A 5 | tail -n 5 | tr -d ' \n:' | xxd -r -p | base64
