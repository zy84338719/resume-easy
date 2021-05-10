package main

import (
	"encoding/json"
	"fmt"
	"github.com/zy84338719/resume-easy/tools"
	"html"
	"html/template"
	"os"
)

func main() {

	s := tools.NewHtml()
	data, err := readData()
	if err!=nil{
		return
	}
	err=json.Unmarshal(data,&s)
	if err!=nil{
		fmt.Println("数据加载失败err=",err)
	}
	s.Now()
	loadTemplate(s)

}
func loadTemplate(data *tools.Html){
	file,err := os.Create("html/src/index.html")
	if err!=nil{
		fmt.Printf("创建文件异常 err=%+v",err)
	}
	if file!=nil{
		defer file.Close()
	}
	tpl:=template.Must(template.ParseFiles("tools/index.html"))
	if err:=tpl.Execute(file, data);err!=nil{
		fmt.Printf("写入出现异常，err=%+v \n",err)
		return
	}
	r, _ := os.ReadFile("html/src/index.html")
	s := html.UnescapeString(string(r))
	os.WriteFile("html/src/index.html",[]byte(s),0755)
	fmt.Println("数据写入成功")
}

func readData()([]byte,error){
	file, err := os.ReadFile("data.json")
	if err!=nil{
		fmt.Printf("读取数据失败,err=%+v \n" ,err)
		return nil,err
	}
	return file,err
}