# 常量定义

集中管理业务中不需要放入数据库中进行配置用的全局常量,如:

- 内容类别
- 文章类别
- 缓存过期时间
- redis key前缀

```
const (
	TypeText = iota // 文本
	TypeAnswer      // 回答
	TypeArticle     // 文章
	TypeVideo       // 视频
)

const (
	// UserKey 用户key
	UserKey = "user|%d"
	// ArticleKey 文章key
	ArticleKey = "article|%d"
)

const (
	// TenMinute 10分钟
	TenMinute  = time.Second * 60 * 10
	UserKeyTTL    = time.Minute * 30
	ArticleKeyTTL = time.Hour * 6
)


```
