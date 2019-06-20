package structAutoGenerate

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

type StructAutoGenerate struct {
	*Option
}
type Option struct {
	Obj         interface{}
	SavePath    string
	PackageName string
}

func New(arg *Option) *StructAutoGenerate {
	return &StructAutoGenerate{Option: arg}
}

func (s *StructAutoGenerate) Generate() error {
	// 包名
	var packageName string
	if s.PackageName == "" {
		packageName = "package main\n\n"
	} else {
		packageName = fmt.Sprintf("package %s\n\n", s.PackageName)
	}

	// struct 内容, set-get 方法
	IContent, newFunc, structContent, setGetContent := s.createContent()

	// 写入文件struct
	var savePath = s.SavePath
	// 是否指定保存路径
	if savePath == "" {
		savePath = "struct-auto-gerate.go"
	}
	filePath := fmt.Sprintf("%s", savePath)
	f, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Can not write file")
		return err
	}
	defer f.Close()

	_, err = f.WriteString(packageName + IContent + structContent + newFunc + strings.Join(setGetContent, ""))
	if err != nil {
		fmt.Println("Can not write file")
		return err
	}

	cmd := exec.Command("gofmt", "-w", filePath)
	err = cmd.Run()

	return err
}

func (s *StructAutoGenerate) createContent() (string, string, string, []string) {
	// set get
	var setGetContent []string
	// 组装struct
	var ref = reflect.TypeOf(s.Obj)

	var structName = ref.Name()
	var structNameUpper = structName
	if len(structNameUpper) <= 1 {
		structNameUpper = strings.ToUpper(structNameUpper)
	} else {
		structNameUpper = strings.ToUpper(structNameUpper[0:1]) + structNameUpper[1:]
	}

	var IContent string = fmt.Sprintf("type I%s interface {\n", structNameUpper)

	var newFunc string = fmt.Sprintf("func New%s() *%s {\nreturn new(%s)}\n\n",
		structNameUpper, structName, structName)

	var structContent string = fmt.Sprintf("type %s struct {\n", structName)
	for i := 0; i < ref.NumField(); i++ {
		var fieldName = ref.Field(i).Name
		var fieldType = ref.Field(i).Type.String()
		var fieldTage string
		if reflect.ValueOf(ref.Field(i).Tag).String() == "" {
			fieldTage = "%s"
		} else {
			fieldTage = " `%s`"
		}
		structContent += fmt.Sprintf("%s %s"+fieldTage+"\n", fieldName, fieldType, ref.Field(i).Tag)

		var funcName string
		if len(fieldName) <= 1 {
			funcName = strings.ToUpper(fieldName)
		} else {
			funcName = strings.ToUpper(fieldName[0:1]) + fieldName[1:]
		}
		// set
		setGetContent = append(setGetContent,
			fmt.Sprintf("func (o *%s) Set%s(arg %s) {\no.%s = arg\n}\n\n",
				structName, funcName, fieldType, fieldName))

		// get
		setGetContent = append(setGetContent,
			fmt.Sprintf("func (o *%s) Get%s() %s {\nreturn o.%s\n}\n\n",
				structName, funcName, fieldType, fieldName))
		//setGetContent = append(setGetContent,
		//	fmt.Sprintf("GetBindName(arg %s) %s\n",
		//		structName, funcName, fieldType, fieldType, fieldName))
		// interface
		IContent += fmt.Sprintf("Set%s(arg %s)\n", funcName, fieldType)
		IContent += fmt.Sprintf("Get%s() %s\n", funcName, fieldType)
	}
	IContent += "}\n"
	structContent += "}\n\n"

	return IContent, newFunc, structContent, setGetContent
}
