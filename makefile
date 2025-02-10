build:
	go run main.go
	cd html && yarn && yarn build
	# 调用脚本生成 ZIP 文件
	sh ./scripts/generate_zip.sh

dev:
	go run main.go
	cd html && yarn dev

clean:
	rm -rf *.zip
	rm -rf [0-9A-Fa-f]*

# deploy:
# 	tcb hosting delete -e ${envId} y
# 	cd docs && tcb hosting deploy -e ${envId}
# 	tcb hosting list -e ${envId}