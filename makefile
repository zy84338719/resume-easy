build:
	go run main.go
	cd html && yarn && yarn build
dev:
	go run main.go
	cd html && yarn dev

deploy:
	tcb hosting delete -e ${envId} y
	cd docs && tcb hosting deploy -e ${envId}
	tcb hosting list -e ${envId}