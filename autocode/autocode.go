package autocode

import (
	"github.com/chengchenginc/go-autocode/parse"
	"io/ioutil"
	"strings"
)

func Gen(sql string, tpl string) (template string, err error) {
	bytes, err := ioutil.ReadFile(tpl)
	if err != nil {
		return
	}
	template = string(bytes)
	pt := parse.ParseTemplate(sql)
	template = strings.Replace(template, "@{domain}", pt.Domain, -1)
	template = strings.Replace(template, "@{domain_fields}", pt.DomainFields, -1)
	return
}
