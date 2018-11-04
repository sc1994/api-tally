package models

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// ScheduleList 列表
	ScheduleList map[uint]*Schedule
)

// Schedule 待办事项
type Schedule struct {
	gorm.Model
	Content string
	Rules   []Rule
	IDs     []uint
}

// Rule 待办事项规则
type Rule struct {
	gorm.Model
	ScheduleID uint      // 关联到Schedule
	Type       int       // 代办事项的规则，1:每天。2:每周。3:每月。4:每年。5:固定日期。
	RuleDate   time.Time // 规则的日期，依据规则类型选择性的取日期的有效成分
}

// Add add
func (s *Schedule) Add() {
	db := GetDB()
	defer db.Close()
	db.Create(s)
}

// Get get
func (s *Schedule) Get() (result Schedule) {
	db := GetDB()
	defer db.Close()
	var rules []Rule
	db.Find(&result).Related(&rules)
	result.Rules = rules
	return result
}
