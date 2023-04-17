package database

import (
	"strings"

	"gorm.io/gorm/schema"
)

type namingStrategy struct {
	schema.NamingStrategy
}

func (ns namingStrategy) TableName(str string) string {
	str = strings.Replace(str, "Model", "", 1)
	return ns.NamingStrategy.TableName(str)
}
