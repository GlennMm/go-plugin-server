package models

import "time"

type Todo struct {
	Created     time.Time `xorm:"'created'"`
	Updated     time.Time `xorm:"'updated'"`
	Name        string    `xorm:"'name'"`
	Description string    `xorm:"'description'"`
	Id          int64     `xorm:"'id' pk autoincr"`
	Done        bool      `xorm:"'done'"`
}
