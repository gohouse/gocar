package structAutoGenerate

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

// StructAutoGenerate
type StructAutoGenerate struct {
	*Option
}

// Option 各种配置
type Option struct {
	Obj         interface{}
	SavePath    string
	PackageName string
	ShortPre    string
}

// New 初始化
func New(arg *Option) *StructAutoGenerate {
	if arg.Obj == nil {
		panic("请传入要解析的结构体!")
	}
	var s = &StructAutoGenerate{Option: arg}

	// 是否指定了method的前缀短标记, 目的是防止与包名或其他变量冲突
	if s.ShortPre == "" {
		s.ShortPre = "o"
	}
	// 如果没有指定保存路径, 则保存当前运行目录
	if s.SavePath == "" {
		s.SavePath = "struct-auto-gerate.go"
	}

	// 如果没有指定包名, 则用main
	if s.PackageName == "" {
		s.PackageName = "main"
	}

	return s
}

// Generate 执行迁移
func (s *StructAutoGenerate) Generate() error {
	// 包名
	var packageName = fmt.Sprintf("package %s\n\n", s.PackageName)

	// struct 内容, set-get 方法
	IContent, newFunc, structContent, setGetContent := s.createContent()

	// 写入文件struct
	filePath := fmt.Sprintf("%s", s.SavePath)
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

	//var IContent string = fmt.Sprintf("type I%s interface {\n", structNameUpper)
	var IContent string = s.getFormatStrWithNotes("type I%s interface {\n",
		[]interface{}{structNameUpper}, []interface{}{"I" + structNameUpper, " 定义接口"})

	//var newFunc string = fmt.Sprintf("func New%s() *%s {\nreturn new(%s)}\n\n",
	//	structNameUpper, structName, structName)
	var newFunc string = s.getFormatStrWithNotes("func New%s() *%s {\nreturn new(%s)}\n\n",
		[]interface{}{structNameUpper, structName, structName}, []interface{}{"New" + structNameUpper, " 初始化"})

	//var structContent string = fmt.Sprintf("type %s struct {\n", structName)
	var structContent string = s.getFormatStrWithNotes("type %s struct {\n",
		[]interface{}{structName}, []interface{}{structName, " 定义结构体"})
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
			s.getFormatStrWithNotes("func (%s *%s) Set%s(arg %s) {\no.%s = arg\n}\n\n",
				[]interface{}{s.ShortPre, structName, funcName, fieldType, fieldName},
				[]interface{}{"Set" + funcName, " 设置该字段值"}))

		// get
		setGetContent = append(setGetContent,
			s.getFormatStrWithNotes("func (%s *%s) Get%s() %s {\nreturn o.%s\n}\n\n",
				[]interface{}{s.ShortPre, structName, funcName, fieldType, fieldName},
				[]interface{}{"Get" + funcName, " 获取该字段值"}))

		// interface
		IContent += fmt.Sprintf("Set%s(arg %s)\n", funcName, fieldType)
		IContent += fmt.Sprintf("Get%s() %s\n", funcName, fieldType)
	}
	IContent += "}\n"
	structContent += "}\n\n"

	return IContent, newFunc, structContent, setGetContent
}

func (s *StructAutoGenerate) getFormatStrWithNotes(fmtStr string, bindName []interface{}, notes []interface{}) string {
	return fmt.Sprint(
		"// ", fmt.Sprint(notes...), "\n",
		fmt.Sprintf(fmtStr, bindName...),
	)
}
