package models

import (
  "github.com/astaxie/beego/orm"
  _ "github.com/go-sql-driver/mysql"
)

type @{domain} struct{@{domain_fields}
}

func (d *@{domain}) tableName() string {
      return @{domain}
}

func init(){
  orm.RegisterModel(new(@{domain}))
}
