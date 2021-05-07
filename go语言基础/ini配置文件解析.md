```go
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Host string `ini:"host"`
	Port int `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	CharSet string `ini:"charset"`

}

type Config struct {
	MysqlConfig `ini:"mysql"`
}

// loadIni 读取ini配置
func loadIni(b []byte, s interface{}) (err error){
	// 参数校验，
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Ptr {
		err = errors.New("the second param must be a pointer")
		return
	}

	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("the second param must be a struct pointer")
		return
	}

	// 将第一个参数转为字符串，并按照换行符分隔
	slice := strings.Split(string(b), "\r\n")
	fmt.Printf("%v\n", slice)
	var structName string
	for idx, line := range slice{
		line = strings.TrimSpace(line)  // 去掉字符串的首尾空格
		if len(line) == 0 {
			// 忽略空行
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			// 忽略注释
			continue
		}
		if strings.HasPrefix(line, "["){
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d has syntax error\n", idx+1)
				return
			}

			sectionName := strings.Trim(line, "[ ]")
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d has syntax error\n", idx+1)
				return
			}
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 找到对应的节
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			// 以=分隔行
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "="){
				err = fmt.Errorf("line:%d has syntax error\n", idx+1)
				return
			}
			lineSlice := strings.Split(line, "=")
			key := strings.TrimSpace(lineSlice[0])
			value := strings.TrimSpace(lineSlice[len(lineSlice)-1])
			// 获取sectionName对应的structName结构体
			v := reflect.ValueOf(s)
			sValue := v.Elem().FieldByName(structName)  // 嵌套结构体的值信息
			sType := sValue.Type()  // 嵌套结构体的类型信息
			if sValue.Kind() != reflect.Struct {
				// 判断找到的是否是结构体
				err = fmt.Errorf("第二个参数对应中%s对应的字段应该是个结构体\n", structName)
				return
			}
			var fieldName string
			for i := 0; i < sType.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == key {
					// 找到对应字段
					fieldName = field.Name
				}
			}
			if len(fieldName) == 0{
				// 在结构体中找不到对应的字段
				continue
			}

			// 根据fieldName取出字段，并赋值
			fieldObj :=sValue.FieldByName(fieldName)
			switch fieldObj.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return
				}
				fieldObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					return
				}
				fieldObj.SetBool(valueBool)

			}
		}
	}
	return nil
}

func main() {
	var c Config
	//var a int
	file, _ := os.Open("./config.ini")
	ret, _ := ioutil.ReadAll(file)
	err := loadIni(ret, &c)
	if err != nil {
		fmt.Printf("load ini file faild; err:%v\n", err)
	}
	fmt.Printf(`
	Host:%s
	Port:%d
	Username:%s
	Password:%s
	CharSet:%s
`, c.MysqlConfig.Host, c.MysqlConfig.Port, c.MysqlConfig.Username, c.MysqlConfig.Password, c.MysqlConfig.CharSet)

}
```