// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"go-blog/internal/dao/internal"
)

// settingDao is the data access object for table gf_setting.
// You can define custom methods on it to extend its functionality as you wish.
type settingDao struct {
	*internal.SettingDao
}

var (
	// Setting is globally public accessible object for table gf_setting operations.
	Setting = settingDao{
		internal.NewSettingDao(),
	}
)

// Fill with you ideas below.
