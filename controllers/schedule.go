package controllers

import (
	"api-tally/models"
	"time"

	"github.com/astaxie/beego"
)

// ScheduleController Schedule
type ScheduleController struct {
	beego.Controller
}

// Add AddSchedule
// @Title schedule/add
// @Description 添加代办事项
// @Param	body		body 	models.Schedule	true		"事项内容"
// @Success 200 {bool} success 是否添加成功
// @router / [post]
func (s *ScheduleController) Add() {
	body := models.Schedule{
		Content: "test1",
		IDs:     []uint{1, 2, 3, 4, 5, 6},
		Rules: []models.Rule{
			{
				Type:     1,
				RuleDate: time.Now(),
			},
		},
	}
	// json.Unmarshal(s.Ctx.Input.RequestBody, &body)
	body.Add()
	s.Data["json"] = map[string]bool{"success": true}
	s.ServeJSON()
}

// Get getSchedule
// @Title Get
// @Description 获取代办详细
// @Param	body		id 	uint	true		"id"
// @Success 200 {models.Schedule} 事项详细
// @router / [get]
func (s *ScheduleController) Get() {
	r := (&models.Schedule{}).Get()
	s.Data["json"] = map[string]models.Schedule{"body": r}
	s.ServeJSON()
}
