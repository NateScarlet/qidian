# qidian

[![build status](https://github.com/NateScarlet/qidian/workflows/Go/badge.svg)](https://github.com/NateScarlet/qidian/actions)

起点中文网 go 客户端，基于网页版页面提取。

- [x] 小说分类搜索
- [ ] 小说关键词搜索
- [x] 小说详情页数据查询
- [x] 小说排行查询
- [x] 作者数据查询
- [ ] 用户数据查询
- [x] 反数据字体混淆

详细使用方法见代码注释

```go
package main

import (
    "context"
    "time"

    "github.com/NateScarlet/qidian/pkg/author"
    "github.com/NateScarlet/qidian/pkg/book"
    "github.com/NateScarlet/qidian/pkg/client"
)

// 默认客户端用环境变量 `PIXIV_PHPSESSID` 登录。
ctx := context.Background()

// 默认使用 http.DefaultClient
client.For(ctx) // http.DefaultClient

// 可对 context 设置自定义 http.Client 。
client.With(ctx, new(http.Client)) // context.Context

// 分类搜索
book.NewCategorySearch().
    SetSubCategory(SC未来世界).
    SetSize(book.SizeGt300kLt500k).
    SetSort(book.SortWeekRecommend).
    SetUpdate(book.UpdateIn3Day).
    SetTag("变身").
    SetPage(2).
    Execute(ctx) // []book.Book

// 书籍详情
b := &book.Book{ID: "1"}
b.Fetch(ctx) // error
b.Title // "魔法骑士英雄传说"
b.Author.ID // "1"
b.Author.Name // "宝剑锋"

// 书籍排行榜
book.Rank{
    Type:  RTMonthlyTicket,
    Year:  2020,
    Month: time.January,
}.Fetch(ctx) // []book.Book, error

// 作者详情
a := &author.Author{ID: "1"}
a.Fetch(ctx) // error
b.Name // "宝剑锋"
```
