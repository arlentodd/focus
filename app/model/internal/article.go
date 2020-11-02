// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// Article is the golang structure for table gf_article.
type Article struct {
    Id         uint        `orm:"id,primary"  json:"id"`          // 自增ID                                                  
    Key        string      `orm:"key"         json:"key"`         // 文章唯一键名，用于程序硬编码使用，一般用不到            
    CategoryId uint        `orm:"category_id" json:"category_id"` // 栏目ID                                                  
    UserId     uint        `orm:"user_id"     json:"user_id"`     // 创建文章的用户ID                                        
    Title      string      `orm:"title"       json:"title"`       // 文章标题                                                
    Content    string      `orm:"content"     json:"content"`     // 文章内容                                                
    Sort       uint        `orm:"sort"        json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶  
    Brief      string      `orm:"brief"       json:"brief"`       // 摘要                                                    
    Thumb      string      `orm:"thumb"       json:"thumb"`       // 缩略图                                                  
    Tags       string      `orm:"tags"        json:"tags"`        // 标签名称列表，以JSON存储                                
    Referer    string      `orm:"referer"     json:"referer"`     // 内容来源                                                
    Status     uint        `orm:"status"      json:"status"`      // 状态 0: 正常, 1: 禁用                                   
    ZanCount   uint        `orm:"zan_count"   json:"zan_count"`   // 赞                                                      
    CaiCount   uint        `orm:"cai_count"   json:"cai_count"`   // 踩                                                      
    CreatedAt  *gtime.Time `orm:"created_at"  json:"created_at"`  // 创建时间                                                
    UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updated_at"`  // 修改时间                                                
}