package state

import "github.com/jinzhu/gorm"

type State struct {
	gorm.Model
	key  string
	name string
}
