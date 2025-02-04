package fx

import (
	"errors"
	"fofax/internal/printer"
	"fofax/internal/table"
	"io/ioutil"
	"strings"

	"fofax/internal/utils"
	"gopkg.in/yaml.v2"
)

const (
	TypeInline FxType = iota
	TypeYaml
)

type FxType int

func (f FxType) String() string {
	switch f {
	case TypeYaml:
		return "文件"
	case TypeInline:
		return "内置"
	default:
		return "未知"
	}
}

type Plugin struct {
	Id          string   `table:"Id" yaml:"id"`
	Query       string   `table:"Query" yaml:"query"`              // 查询语法
	RuleName    string   `table:"RuleName" yaml:"rule_name"`       // 标题名
	RuleEnglish string   `table:"RuleEnglish" yaml:"rule_english"` // 规则英文名
	Description string   `table:"Description" yaml:"description"`  // 描述
	Author      string   `table:"Author" yaml:"author"`            // 作者
	FofaQuery   string   `table:"-" yaml:"fofa_query"`             // fofa查询
	Tag         []string `table:"Tag" yaml:"tag"`                  // 标签
	Type        FxType   `table:"Type" yaml:"-"`                   // 类别
	Source      string   `table:"Source" yaml:"source"`            // 来源
	FileDir     string   `table:"Source" yaml:"-"`                 // 来源
}

func (q *Plugin) ShowInfoTable() {
	results := []Tinfo{
		{"ID", q.Id},
		{"Query", q.Query},
		{"RuleName", q.RuleName},
		{"RuleEnglish", q.RuleEnglish},
		{"Author", q.Author},
		{"FofaQuery", q.FofaQuery},
		{"Tag", strings.Join(q.Tag, ",")},
		{"Type", q.Type.String()},
		{"Description", q.Description},
		{"FileDir", q.FileDir},
	}
	table.Output(results)
}

func (f Plugin) QueryString() string {
	return f.FofaQuery
}

func (base *Plugin) GenPlugin(pluginFile string) error {
	if utils.FileExist(pluginFile) {
		return errors.New("File exist please check: " + pluginFile)
	}
	data, _ := yaml.Marshal(base)
	printer.Infof("Will Write Plugin file: %s\n", pluginFile)
	err := ioutil.WriteFile(pluginFile, data, 0644) // 写入
	if err != nil {
		printer.Fatalf("%s can't  write Plugin file\n", pluginFile)
		return err
	}
	return nil

}
