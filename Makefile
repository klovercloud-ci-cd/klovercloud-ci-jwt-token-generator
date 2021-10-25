build:
	go build -o kcpctl
install:
  export PATH=$PATH:/app/klovercloud-ci-jwt-token-generator
run:
	./kcpctl