package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
	"time"
	"{{projectName}}/pkg/util"
)

type Common struct {
	Id        string     `gorm:"size:32;primary_key" json:"id"`
	CreatedAt int64      `json:"created_at"`
	UpdatedAt int64      `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func buildWhere(rawQuery string, db *gorm.DB) (*gorm.DB, error) {
	if rawQuery != "" {
		querys := strings.Split(rawQuery, ",")
		for _, query := range querys {
			oneQuery := strings.Split(query, ":")
			if len(oneQuery) != 2 && len(oneQuery) != 3 {
				return db, errors.New("parseRawQuery error, rawQuery should like: 'title:=:golang,name:like:%jason%,id:<:100' , if the whereType is '=', you can omitted it: title:golang, notice: '%' after encode is %25")
			}
			if len(oneQuery) == 2 {
				field := oneQuery[0]
				whereType := "="
				value := oneQuery[1]
				db = db.Where(field+" "+whereType+" "+"?", value)
			}
			if len(oneQuery) == 3 {
				field := oneQuery[0]
				whereType := oneQuery[1]
				value := oneQuery[2]
				db = db.Where(field+" "+whereType+" "+"?", value)
			}
		}
	}
	return db, nil
}

func buildOrder(rawOrder string, db *gorm.DB) (*gorm.DB, error) {

	if rawOrder != "" {
		orders := strings.Split(rawOrder, ",")
		for _, order := range orders {
			oneOrder := strings.Split(order, ":")
			if len(oneOrder) != 1 && len(oneOrder) != 2 {
				return db, errors.New("parse rawOrder error, rawOrder should like:'created:desc,id:asc,name', orderType default is asc")
			}

			if len(oneOrder) == 1 {
				field := oneOrder[0]
				db = db.Order(field)
			}
			if len(oneOrder) == 2 {
				field := oneOrder[0]
				orderType := oneOrder[1]
				db = db.Order(field + " " + orderType)
			}
		}
	}
	return db, nil
}

func (b *Common) BeforeCreate(scope *gorm.Scope) error {
	//scope.SetColumn("Id", uuid.NewV4().String())
	scope.SetColumn("Id", util.NewUuid())
	return nil
}
