package dao

import (
	"context"
	"kratos_frame/internal/model"
)

//go:generate kratos tool genbts
// Bts  interface  用于生成缓存回溯代码
type _bts interface {
	// bts: -nullcache=&model.Article{ID:-1} -check_null_code=$!=nil&&$.ID==-1
	Article(c context.Context, id int64) (*model.Article, error)
}
