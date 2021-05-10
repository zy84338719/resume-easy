build:
	go run main.go
	cd html && yarn && yarn build
dev:
	go run main.go
	cd html && yarn dev