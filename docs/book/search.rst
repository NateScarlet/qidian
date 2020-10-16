搜索
===========


分类搜索
------------------------------

抓取 ``https://www.qidian.com/all``


请求参数:

page

  类型: int

  第几页

style

  类型: int
  
  可能的值:

    1：封面模式, 一页 20 条。

    2：列表模式, 一页 50 条。

pageSize

  类型: int

  好像没用，更改不影响结果。

  对应 style 条数数值。

orderId

  类型： int | undefined

  排序方式，可能的值：

    undefined: 人气

    1：周点击

    2：总推荐

    3：总字数

    4：更新时间，数据不可见。

    5：更新时间

    6：完本时间

    7：月点击，但是顺序不对。

    8：总点击，数据不可见。

    9：周推荐

    10：月推荐

    11：总收藏

chanId

  类型：int | undefined

  分类，可能的值为 :doc:`categories` 中的分类ID。

tag

  类型: string | undefined
  
  标签


关键字搜索
------------------------------

使用 https://www.qidian.com/search
