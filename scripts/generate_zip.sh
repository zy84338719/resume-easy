#!/bin/sh
# 生成 UUID 并去除短横线
UUID=$(uuidgen | tr -d '-')

cp -r ./docs ./${UUID}
# 压缩为 ZIP 文件
zip -r "${UUID}.zip" ./${UUID}
rm -rf ./${UUID}
