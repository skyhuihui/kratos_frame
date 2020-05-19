package dao

import (
	"context"
	"kratos_frame/internal/model"
)

func (d *Dao) RawArticle(ctx context.Context, id int64) (art *model.Article, err error) {
	t := DbTrace(ctx, find, 0, "根据 id 查询 活动 当前操作用户 ： "+"123")
	//d.db[GetDbId(len(d.db))].Find()
	t.Finish(nil)
	return
}
