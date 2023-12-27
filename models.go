package main

import "github.com/jinzhu/gorm"

type Memo struct {
	gorm.Model
	Title   string
	Content string
}
