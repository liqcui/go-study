package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func LoadIni(fileName string, data interface{}) (err error) {
	//1.参数的校验
	//1.1 传递进来的data参数必须是指针类型（因为需要在函数中对其赋值）
	t := reflect.TypeOf(data)
	t.Kind()
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("data should be a pointer.")
	}
	//1.2 传递进来的data参数必须是结构体类型指针（因为配置文件中各种赋值对需要赋值给结构体到字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("Data param should be a struct pointer")
		return err
	}
	//2. 对文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	//将字节类型的文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	fmt.Printf("%#v\n", lineSlice)
	//3. 一行一行的读数据
	var structName string
	for idx, line := range lineSlice {
		//去掉字符串首位的空格
		line = strings.TrimSpace(line)
		//3.1如果是注释就跳过
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		//3.2 如果是【开头的就表示字节
		if strings.HasPrefix(line, "[") {
			if line[0] == '[' && line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			if len(strings.TrimSpace(line[1:len(line)-1])) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			//把这一行首尾的[]去掉，取到中间的内容把首尾的空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}

			//根据反射
			reflect.ValueOf(data)
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
				}
			}
		} else {
			//2.3 如果不是[开头就是=分割的键值对
			//1.以等号分割这一行，等号左边死key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")
			key := strings.TrimSpace(line[:index])
			value := strings.TrimSpace(line[index+1:])
			//2.根据structName去data里面把对应的嵌套结构体给取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName)
			sType := sValue.Type()
			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data中的%s字段应该是一个结构体", structName)
				return
			}
			var fieldName string
			//3.遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i := 0; i < sValue.NumField(); i++ {
				field := sType.Field(i)
				if field.Tag.Get("ini") == key {
					fieldName = field.Name
				}
			}
			//4.如果key=tag，给这个字段赋值
			fileObj := sValue.FieldByName(fileName)
			fmt.Println(fileName, fileObj.Type().Kind())

		}
		return
	}
	//3.1 如果是注释就跳过
	//3.2 如果是[开头就表示是节（Section）
	//3.3 如果不是[开头就是=分割的键值对
	return
}

func main() {
	var mc MysqlConfig
	err := LoadIni("./config.ini", &mc)
	if err!=nil{

	}
}
