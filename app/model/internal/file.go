// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// File is the golang structure for table gf_file.
type File struct {
	Id        uint        `orm:"id,primary" json:"id"`         // 自增ID
	Name      string      `orm:"name"       json:"name"`       // 文件名称
	Src       string      `orm:"src"        json:"src"`        // 本地文件存储路径
	Url       string      `orm:"url"        json:"url"`        // URL地址，可能为空
	UserId    uint        `orm:"user_id"    json:"user_id"`    // 操作用户
	CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` //
}
