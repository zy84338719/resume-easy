package main

import (
	"encoding/json"
	"fmt"
	"github.com/zy84338719/resume-easy/tools" //nolint:gci
	"html"
	"html/template"
	"os" //nolint:gci
)

const (
	DATA_READ   = "data.json"
	CREATE_DATA = "html/src/index.html"
	TEMPLATE    = "tools/index.tpl"
)

func main() {
	s := tools.NewHtml()
	data, err := readData()
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &s) //nolint:wsl
	if err != nil {                //nolint:wsl
		fmt.Println("数据加载失败err=", err)
	}
	s.Now()
	loadTemplate(s)
}
func loadTemplate(data *tools.Html) {
	file, err := os.Create(CREATE_DATA)
	if err != nil {
		fmt.Printf("创建文件异常 err=%+v", err) //nolint:forbidigo
	}
	if file != nil {
		defer file.Close()
	}
	tpl := template.Must(template.ParseFiles(TEMPLATE))
	if err := tpl.Execute(file, data); err != nil {
		fmt.Printf("写入出现异常，err=%+v \n", err)
		return
	}
	r, err := os.ReadFile("html/src/index.html")
	if err != nil {
		fmt.Printf("读取数据失败 err=%+v", err)
	}
	s := html.UnescapeString(string(r))
	err = os.WriteFile(CREATE_DATA, []byte(s), 0755)
	if err != nil {
		fmt.Printf("写入数据失败 err=%+v", err) //nolint:forbidigo
	}
	fmt.Println("数据写入成功") //nolint:wsl,forbidigo
}

func readData() ([]byte, error) {
	file, err := os.ReadFile(DATA_READ)
	if err != nil {
		fmt.Printf("读取数据失败,err=%+v \n", err)
		return nil, err
	}
	return file, err
}
