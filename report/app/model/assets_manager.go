// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"assets/app/model/internal"
	"github.com/gogf/gf/os/gtime"
)

// AssetsManager is the golang structure for table assets_manager.
type AssetsManager internal.AssetsManager

// Fill with you ideas below.

// RequestAssetsManagerAdd 添加安全管理员所需信息
type RequestAssetsManagerAdd struct {
	ManagerName string `v:"required#安全管理员不能为空"`
}

// RsponseManager 安全管理员 模糊分页查询返回数据所需信息
type RsponseManager struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []RsponseManagerInfo `json:"data"`
}

// RsponseManagerInfo 安全管理 模糊分页查询返回数据所需详细信息
type RsponseManagerInfo struct{
	Id int `json:"id"`
	ManagerName string `json:"manager_name"`
	WebCount int `json:"web_count"`
	LevelNoCount int `json:"level_no_count"`
	LevelYesCount int `json:"level_yes_count"`
	LevelCount int `json:"level_count"`
	CusTime *gtime.Time `json:"cus_time"`
}

// ResponseAssetsManagerNameGroup 安全管理员Group分组
type ResponseAssetsManagerNameGroup struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []ResponseAssetsManagerNameGroupInfo `json:"data"`
}

// ResponseAssetsManagerNameGroupInfo 安全管理员Group分组详细信息
type ResponseAssetsManagerNameGroupInfo struct{
	ID int `json:"id"`
	ManagerName string `json:"manager_name"`
}