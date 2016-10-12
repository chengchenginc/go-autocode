package parse

import (
	"regexp"
	"strings"
)

type Table struct {
	Name       string
	Fields     []Field
	PrimaryKey string
}

type Field struct {
	Name         string
	DataType     string
	DefaultValue string
	Nullable     bool
}

type Template struct{
		Domain string
		DomainFields string
}

func ParseTemplate(sql string)(template Template){
		table := ParseSql(sql)
		template.Domain = table.Name
		for _,f :=range table.Fields{
				template.DomainFields += "\n    "+ f.Name+" "+f.DataType+" "+ "`orm:\"column("+f.Name+")\"` "
		}
		return
}

func ParseSql(sql string) (table Table) {
	table.Name = parseTableName(sql)
	table.Fields = parseFields(sql)
	return
}

func parseTableName(sql string) (name string) {
	r, err := regexp.Compile(`CREATE TABLE[^(]*\(`)
	if err != nil {
		return
	}
	name = r.FindString(sql)
	name = strings.TrimPrefix(name, "CREATE TABLE")
	name = strings.TrimRight(name, "(")
	name = strings.TrimSpace(name)
	name = strings.Trim(name, "`")
	return
}

func parseFields(sql string) (fields []Field) { //slice
	start := strings.Index(sql, "(")
	end := strings.LastIndex(sql, ")")
	fieldStr := sql[start+1 : end]
	strs := strings.Split(fieldStr, ",")
	for _, str := range strs {
		rows := strings.Fields(str)
		flag := isContainsDefaultValue(rows)
		if flag == true {
			if strings.ToUpper(rows[2]) == "DEFAULT" {
				fields = append(fields, Field{
					Name:         strings.Trim(rows[0], "`"),
					DataType:     "string",
					DefaultValue: "",
					Nullable:     true,
				})
			} else if strings.ToUpper(rows[4]) == "DEFAULT" {
				fields = append(fields, Field{
					Name:         strings.Trim(rows[0], "`"),
					DataType:     "string",
					DefaultValue: rows[5],
					Nullable:     false,
				})
			}
		} else {
			if strings.ToUpper(rows[0]) == "PRIMARY" {
				//
			} else {
				fields = append(fields, Field{
					Name:         strings.Trim(rows[0], "`"),
					DataType:     "string",
					DefaultValue: "",
					Nullable:     false,
				})
			}
		}
	}
	return
}

func isContainsDefaultValue(rows []string) bool {
	flag := false
	for _, row := range rows {
		if strings.ToUpper(row) == "DEFAULT" {
			flag = true
			break
		}
	}
	return flag
}
