package dao

import (
	"context"
	"kratos_frame/internal/model"
)

//go:generate kratos tool genmc
// 用于生成 memcached 缓存代码
type _mc interface {
	// mc: -key=keyArt -type=get
	CacheArticle(c context.Context, id int64) (*model.Article, error)
	// mc: -key=keyArt -expire=d.demoExpire
	AddCacheArticle(c context.Context, id int64, art *model.Article) (err error)
	// mc: -key=keyArt
	DeleteArticleCache(c context.Context, id int64) (err error)
}
