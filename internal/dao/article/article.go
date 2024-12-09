package article

import (
	"context"
)

type Article struct {
	ID       int64  `json:"id"`        // 数据库自增ID
	AuthorID int64  `json:"author_id"` // 作者id
	Status   int    `json:"status"`    // 状态 0: banned，1: using, 2: preparing
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type ArticleDao interface {
	AddArticle(ctx context.Context, article Article) (lastInsertID int64, err error)
	// 更新创作分&等级
	UpdateArticle(ctx context.Context, status int) (err error)
}
