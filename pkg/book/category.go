// Code generated from category.go.gotmpl, DO NOT EDIT.

package book

type Category string

type SubCategory string

const (
    C全部 Category = "-1"
    C奇幻 Category = "1"
    C悬疑 Category = "10"
    C轻小说 Category = "12"
    C现实 Category = "15"
    C武侠 Category = "2"
    C短篇 Category = "20076"
    C玄幻 Category = "21"
    C仙侠 Category = "22"
    C都市 Category = "4"
    C历史 Category = "5"
    C军事 Category = "6"
    C游戏 Category = "7"
    C体育 Category = "8"
    C科幻 Category = "9"
    SC都市生活 SubCategory = "12"
    SC青春校园 SubCategory = "130"
    SC娱乐明星 SubCategory = "151"
    SC商战职场 SubCategory = "153"
    SC都市异能 SubCategory = "16"
    SC修真文明 SubCategory = "18"
    SC影视剧本 SubCategory = "20075"
    SC评论文集 SubCategory = "20077"
    SC生活随笔 SubCategory = "20078"
    SC美文游记 SubCategory = "20079"
    SC历史神话 SubCategory = "20092"
    SC另类幻想 SubCategory = "20093"
    SC民间传说 SubCategory = "20094"
    SC古今传奇 SubCategory = "20095"
    SC短篇小说 SubCategory = "20096"
    SC诗歌散文 SubCategory = "20097"
    SC人物传记 SubCategory = "20098"
    SC古武未来 SubCategory = "20099"
    SC史诗奇幻 SubCategory = "201"
    SC武侠同人 SubCategory = "20100"
    SC古典仙侠 SubCategory = "20101"
    SC游戏系统 SubCategory = "20102"
    SC游戏主播 SubCategory = "20103"
    SC社会乡土 SubCategory = "20104"
    SC生活时尚 SubCategory = "20105"
    SC文学艺术 SubCategory = "20106"
    SC成功励志 SubCategory = "20107"
    SC青春文学 SubCategory = "20108"
    SC黑暗幻想 SubCategory = "202"
    SC国术无双 SubCategory = "206"
    SC神话修真 SubCategory = "207"
    SC现实百态 SubCategory = "209"
    SC古武机甲 SubCategory = "21"
    SC架空历史 SubCategory = "22"
    SC上古先秦 SubCategory = "220"
    SC两晋隋唐 SubCategory = "222"
    SC五代十国 SubCategory = "223"
    SC两宋元明 SubCategory = "224"
    SC清史民国 SubCategory = "225"
    SC外国历史 SubCategory = "226"
    SC抗战烽火 SubCategory = "230"
    SC谍战特工 SubCategory = "231"
    SC游戏异界 SubCategory = "240"
    SC未来世界 SubCategory = "25"
    SC超级科技 SubCategory = "250"
    SC时空穿梭 SubCategory = "251"
    SC进化变异 SubCategory = "252"
    SC末世危机 SubCategory = "253"
    SC诡秘悬疑 SubCategory = "26"
    SC探险生存 SubCategory = "260"
    SC篮球运动 SubCategory = "28"
    SC衍生同人 SubCategory = "281"
    SC搞笑吐槽 SubCategory = "282"
    SC武侠幻想 SubCategory = "30"
    SC历史传记 SubCategory = "32"
    SC奇妙世界 SubCategory = "35"
    SC现代魔法 SubCategory = "38"
    SC幻想修仙 SubCategory = "44"
    SC秦汉三国 SubCategory = "48"
    SC传统武侠 SubCategory = "5"
    SC军旅生涯 SubCategory = "54"
    SC体育赛事 SubCategory = "55"
    SC侦探推理 SubCategory = "57"
    SC王朝争霸 SubCategory = "58"
    SC爱情婚姻 SubCategory = "6"
    SC原生幻想 SubCategory = "60"
    SC剑与魔法 SubCategory = "62"
    SC现代修真 SubCategory = "64"
    SC军事战争 SubCategory = "65"
    SC青春日常 SubCategory = "66"
    SC星际文明 SubCategory = "68"
    SC电子竞技 SubCategory = "7"
    SC虚拟网游 SubCategory = "70"
    SC异世大陆 SubCategory = "73"
    SC异术超能 SubCategory = "74"
    SC高武世界 SubCategory = "78"
    SC东方玄幻 SubCategory = "8"
    SC战争幻想 SubCategory = "80"
    SC足球运动 SubCategory = "82"
)

func (c Category) String() string {
    switch c {
    case C全部:
        return "全部"
    case C奇幻:
        return "奇幻"
    case C悬疑:
        return "悬疑"
    case C轻小说:
        return "轻小说"
    case C现实:
        return "现实"
    case C武侠:
        return "武侠"
    case C短篇:
        return "短篇"
    case C玄幻:
        return "玄幻"
    case C仙侠:
        return "仙侠"
    case C都市:
        return "都市"
    case C历史:
        return "历史"
    case C军事:
        return "军事"
    case C游戏:
        return "游戏"
    case C体育:
        return "体育"
    case C科幻:
        return "科幻"
    }
    return ""
}

func CategoryByName(v string) Category {
    switch v {
    case "全部":
        return C全部
    case "奇幻":
        return C奇幻
    case "悬疑":
        return C悬疑
    case "轻小说":
        return C轻小说
    case "现实":
        return C现实
    case "武侠":
        return C武侠
    case "短篇":
        return C短篇
    case "玄幻":
        return C玄幻
    case "仙侠":
        return C仙侠
    case "都市":
        return C都市
    case "历史":
        return C历史
    case "军事":
        return C军事
    case "游戏":
        return C游戏
    case "体育":
        return C体育
    case "科幻":
        return C科幻
    }
    return ""
}

func (sc SubCategory) Parent() Category {
    switch sc{
    case SC都市生活:
        return C都市
    case SC青春校园:
        return C都市
    case SC娱乐明星:
        return C都市
    case SC商战职场:
        return C都市
    case SC都市异能:
        return C都市
    case SC修真文明:
        return C仙侠
    case SC影视剧本:
        return C短篇
    case SC评论文集:
        return C短篇
    case SC生活随笔:
        return C短篇
    case SC美文游记:
        return C短篇
    case SC历史神话:
        return C奇幻
    case SC另类幻想:
        return C奇幻
    case SC民间传说:
        return C历史
    case SC古今传奇:
        return C悬疑
    case SC短篇小说:
        return C短篇
    case SC诗歌散文:
        return C短篇
    case SC人物传记:
        return C短篇
    case SC古武未来:
        return C武侠
    case SC史诗奇幻:
        return C奇幻
    case SC武侠同人:
        return C武侠
    case SC古典仙侠:
        return C仙侠
    case SC游戏系统:
        return C游戏
    case SC游戏主播:
        return C游戏
    case SC社会乡土:
        return C现实
    case SC生活时尚:
        return C现实
    case SC文学艺术:
        return C现实
    case SC成功励志:
        return C现实
    case SC青春文学:
        return C现实
    case SC黑暗幻想:
        return C奇幻
    case SC国术无双:
        return C武侠
    case SC神话修真:
        return C仙侠
    case SC现实百态:
        return C现实
    case SC古武机甲:
        return C科幻
    case SC架空历史:
        return C历史
    case SC上古先秦:
        return C历史
    case SC两晋隋唐:
        return C历史
    case SC五代十国:
        return C历史
    case SC两宋元明:
        return C历史
    case SC清史民国:
        return C历史
    case SC外国历史:
        return C历史
    case SC抗战烽火:
        return C军事
    case SC谍战特工:
        return C军事
    case SC游戏异界:
        return C游戏
    case SC未来世界:
        return C科幻
    case SC超级科技:
        return C科幻
    case SC时空穿梭:
        return C科幻
    case SC进化变异:
        return C科幻
    case SC末世危机:
        return C科幻
    case SC诡秘悬疑:
        return C悬疑
    case SC探险生存:
        return C悬疑
    case SC篮球运动:
        return C体育
    case SC衍生同人:
        return C轻小说
    case SC搞笑吐槽:
        return C轻小说
    case SC武侠幻想:
        return C武侠
    case SC历史传记:
        return C历史
    case SC奇妙世界:
        return C悬疑
    case SC现代魔法:
        return C奇幻
    case SC幻想修仙:
        return C仙侠
    case SC秦汉三国:
        return C历史
    case SC传统武侠:
        return C武侠
    case SC军旅生涯:
        return C军事
    case SC体育赛事:
        return C体育
    case SC侦探推理:
        return C悬疑
    case SC王朝争霸:
        return C玄幻
    case SC爱情婚姻:
        return C现实
    case SC原生幻想:
        return C轻小说
    case SC剑与魔法:
        return C奇幻
    case SC现代修真:
        return C仙侠
    case SC军事战争:
        return C军事
    case SC青春日常:
        return C轻小说
    case SC星际文明:
        return C科幻
    case SC电子竞技:
        return C游戏
    case SC虚拟网游:
        return C游戏
    case SC异世大陆:
        return C玄幻
    case SC异术超能:
        return C都市
    case SC高武世界:
        return C玄幻
    case SC东方玄幻:
        return C玄幻
    case SC战争幻想:
        return C军事
    case SC足球运动:
        return C体育
    }
    return ""
}

func (sc SubCategory) String() string {
    switch sc {
    case SC都市生活:
        return "都市生活"
    case SC青春校园:
        return "青春校园"
    case SC娱乐明星:
        return "娱乐明星"
    case SC商战职场:
        return "商战职场"
    case SC都市异能:
        return "都市异能"
    case SC修真文明:
        return "修真文明"
    case SC影视剧本:
        return "影视剧本"
    case SC评论文集:
        return "评论文集"
    case SC生活随笔:
        return "生活随笔"
    case SC美文游记:
        return "美文游记"
    case SC历史神话:
        return "历史神话"
    case SC另类幻想:
        return "另类幻想"
    case SC民间传说:
        return "民间传说"
    case SC古今传奇:
        return "古今传奇"
    case SC短篇小说:
        return "短篇小说"
    case SC诗歌散文:
        return "诗歌散文"
    case SC人物传记:
        return "人物传记"
    case SC古武未来:
        return "古武未来"
    case SC史诗奇幻:
        return "史诗奇幻"
    case SC武侠同人:
        return "武侠同人"
    case SC古典仙侠:
        return "古典仙侠"
    case SC游戏系统:
        return "游戏系统"
    case SC游戏主播:
        return "游戏主播"
    case SC社会乡土:
        return "社会乡土"
    case SC生活时尚:
        return "生活时尚"
    case SC文学艺术:
        return "文学艺术"
    case SC成功励志:
        return "成功励志"
    case SC青春文学:
        return "青春文学"
    case SC黑暗幻想:
        return "黑暗幻想"
    case SC国术无双:
        return "国术无双"
    case SC神话修真:
        return "神话修真"
    case SC现实百态:
        return "现实百态"
    case SC古武机甲:
        return "古武机甲"
    case SC架空历史:
        return "架空历史"
    case SC上古先秦:
        return "上古先秦"
    case SC两晋隋唐:
        return "两晋隋唐"
    case SC五代十国:
        return "五代十国"
    case SC两宋元明:
        return "两宋元明"
    case SC清史民国:
        return "清史民国"
    case SC外国历史:
        return "外国历史"
    case SC抗战烽火:
        return "抗战烽火"
    case SC谍战特工:
        return "谍战特工"
    case SC游戏异界:
        return "游戏异界"
    case SC未来世界:
        return "未来世界"
    case SC超级科技:
        return "超级科技"
    case SC时空穿梭:
        return "时空穿梭"
    case SC进化变异:
        return "进化变异"
    case SC末世危机:
        return "末世危机"
    case SC诡秘悬疑:
        return "诡秘悬疑"
    case SC探险生存:
        return "探险生存"
    case SC篮球运动:
        return "篮球运动"
    case SC衍生同人:
        return "衍生同人"
    case SC搞笑吐槽:
        return "搞笑吐槽"
    case SC武侠幻想:
        return "武侠幻想"
    case SC历史传记:
        return "历史传记"
    case SC奇妙世界:
        return "奇妙世界"
    case SC现代魔法:
        return "现代魔法"
    case SC幻想修仙:
        return "幻想修仙"
    case SC秦汉三国:
        return "秦汉三国"
    case SC传统武侠:
        return "传统武侠"
    case SC军旅生涯:
        return "军旅生涯"
    case SC体育赛事:
        return "体育赛事"
    case SC侦探推理:
        return "侦探推理"
    case SC王朝争霸:
        return "王朝争霸"
    case SC爱情婚姻:
        return "爱情婚姻"
    case SC原生幻想:
        return "原生幻想"
    case SC剑与魔法:
        return "剑与魔法"
    case SC现代修真:
        return "现代修真"
    case SC军事战争:
        return "军事战争"
    case SC青春日常:
        return "青春日常"
    case SC星际文明:
        return "星际文明"
    case SC电子竞技:
        return "电子竞技"
    case SC虚拟网游:
        return "虚拟网游"
    case SC异世大陆:
        return "异世大陆"
    case SC异术超能:
        return "异术超能"
    case SC高武世界:
        return "高武世界"
    case SC东方玄幻:
        return "东方玄幻"
    case SC战争幻想:
        return "战争幻想"
    case SC足球运动:
        return "足球运动"
    }
    return ""
}

func SubCategoryByName(v string) SubCategory {
    switch v {
    case "都市生活":
        return SC都市生活
    case "青春校园":
        return SC青春校园
    case "娱乐明星":
        return SC娱乐明星
    case "商战职场":
        return SC商战职场
    case "都市异能":
        return SC都市异能
    case "修真文明":
        return SC修真文明
    case "影视剧本":
        return SC影视剧本
    case "评论文集":
        return SC评论文集
    case "生活随笔":
        return SC生活随笔
    case "美文游记":
        return SC美文游记
    case "历史神话":
        return SC历史神话
    case "另类幻想":
        return SC另类幻想
    case "民间传说":
        return SC民间传说
    case "古今传奇":
        return SC古今传奇
    case "短篇小说":
        return SC短篇小说
    case "诗歌散文":
        return SC诗歌散文
    case "人物传记":
        return SC人物传记
    case "古武未来":
        return SC古武未来
    case "史诗奇幻":
        return SC史诗奇幻
    case "武侠同人":
        return SC武侠同人
    case "古典仙侠":
        return SC古典仙侠
    case "游戏系统":
        return SC游戏系统
    case "游戏主播":
        return SC游戏主播
    case "社会乡土":
        return SC社会乡土
    case "生活时尚":
        return SC生活时尚
    case "文学艺术":
        return SC文学艺术
    case "成功励志":
        return SC成功励志
    case "青春文学":
        return SC青春文学
    case "黑暗幻想":
        return SC黑暗幻想
    case "国术无双":
        return SC国术无双
    case "神话修真":
        return SC神话修真
    case "现实百态":
        return SC现实百态
    case "古武机甲":
        return SC古武机甲
    case "架空历史":
        return SC架空历史
    case "上古先秦":
        return SC上古先秦
    case "两晋隋唐":
        return SC两晋隋唐
    case "五代十国":
        return SC五代十国
    case "两宋元明":
        return SC两宋元明
    case "清史民国":
        return SC清史民国
    case "外国历史":
        return SC外国历史
    case "抗战烽火":
        return SC抗战烽火
    case "谍战特工":
        return SC谍战特工
    case "游戏异界":
        return SC游戏异界
    case "未来世界":
        return SC未来世界
    case "超级科技":
        return SC超级科技
    case "时空穿梭":
        return SC时空穿梭
    case "进化变异":
        return SC进化变异
    case "末世危机":
        return SC末世危机
    case "诡秘悬疑":
        return SC诡秘悬疑
    case "探险生存":
        return SC探险生存
    case "篮球运动":
        return SC篮球运动
    case "衍生同人":
        return SC衍生同人
    case "搞笑吐槽":
        return SC搞笑吐槽
    case "武侠幻想":
        return SC武侠幻想
    case "历史传记":
        return SC历史传记
    case "奇妙世界":
        return SC奇妙世界
    case "现代魔法":
        return SC现代魔法
    case "幻想修仙":
        return SC幻想修仙
    case "秦汉三国":
        return SC秦汉三国
    case "传统武侠":
        return SC传统武侠
    case "军旅生涯":
        return SC军旅生涯
    case "体育赛事":
        return SC体育赛事
    case "侦探推理":
        return SC侦探推理
    case "王朝争霸":
        return SC王朝争霸
    case "爱情婚姻":
        return SC爱情婚姻
    case "原生幻想":
        return SC原生幻想
    case "剑与魔法":
        return SC剑与魔法
    case "现代修真":
        return SC现代修真
    case "军事战争":
        return SC军事战争
    case "青春日常":
        return SC青春日常
    case "星际文明":
        return SC星际文明
    case "电子竞技":
        return SC电子竞技
    case "虚拟网游":
        return SC虚拟网游
    case "异世大陆":
        return SC异世大陆
    case "异术超能":
        return SC异术超能
    case "高武世界":
        return SC高武世界
    case "东方玄幻":
        return SC东方玄幻
    case "战争幻想":
        return SC战争幻想
    case "足球运动":
        return SC足球运动
    }
    return ""
}
