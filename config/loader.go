package config

import (
	"bytes"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

/*
	caeate by yanggang
*/

func appendByte(buff *bytes.Buffer, b []byte) {
	buff.Write(b)
	if len(b) > 0 && b[len(b)-1] != '\n' {
		buff.WriteByte('\n')
	}
}

type Section interface {
	SectionName() string
}

func (c *config) get(key string) interface{} {
	return (*c)[key]
}

func Load(section Section) {
	s := loader.get(section.SectionName())
	data, _ := yaml.Marshal(s)
	yaml.Unmarshal(data, section)
}

/* 声明一个叫 config 的 map 结构 */
type config map[string]interface{}

/* 声明一个叫 loader 的 config */
var loader = &config{}

/**/
func init() {
	// 构建一个空bytes,等待拼接全部配置文件内容
	buff := bytes.Buffer{}
	// 从系统配置项寻找config配置节点
	env := os.Getenv("config")
	if env == "" {
		// config 节点没找到， 从index.yml 查找想加载那些配置文件
		app, e := ioutil.ReadFile("./config/index.yml")
		if e != nil {
			fmt.Printf("File error: %v\n", e)
			os.Exit(1)
		} else {
			appendByte(&buff, app)
			yaml.Unmarshal(app, &loader)
			if c := loader.get("config"); c != nil {
				env = c.(string)
			}
		}
	}
	if env != "" {
		for _, file := range strings.Split(env, ",") {
			b, e := ioutil.ReadFile("./config/" + file + ".yml")
			if e != nil {
				fmt.Printf("File error: %v\n", e)
			} else {
				appendByte(&buff, b)
			}
		}
	}

	yaml.Unmarshal(buff.Bytes(), &loader)
}
