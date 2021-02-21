package models

import "gorm.io/gorm/schema"

type ActiveRecord interface {
	schema.Tabler
	SetCreateBy(createBy int)
	SetUpdateBy(updateBy int)
	Generate() ActiveRecord
	GetId() interface{}
}
type ActiveRecord2 interface {
	schema.Tabler
	Generate() ActiveRecord2
	GetId() interface{}
}