package book

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBook_Fetch(t *testing.T) {
	var ctx = context.Background()

	var b = Book{ID: "1"}
	var err = b.Fetch(ctx)
	require.NoError(t, err)
	assert.Equal(t, "魔法骑士英雄传说", b.Title)
	assert.Equal(t, "1", b.Author.ID)
	assert.Equal(t, "宝剑锋", b.Author.Name)
	assert.Equal(t, []string{"连载", "签约", "VIP", "学院流", "特种兵", "轻松"}, b.Tags)
	assert.Equal(t, C玄幻, b.Category)
	assert.Equal(t, SC异世大陆, b.SubCategory)
	assert.Equal(t, "https://bookcover.yuewen.com/qdbimg/349573/1/180", b.CoverURL)
	assert.Equal(t, time.Date(2003, 10, 23, 0, 0, 0, 0, TZ), b.LastUpdated)
	assert.Equal(t, uint64(987300), b.WordCount)
	assert.LessOrEqual(t, uint64(94200), b.TotalRecommendCount)
	// assert.Equal(t, uint64(68), b.WeekRecommendCount)
	assert.Equal(t, "心潮澎湃，无限幻想，迎风挥击千层浪，少年不败热血！", b.Summary)
	assert.Equal(t,
		`作为最强特种穿越到异世界的猎人雷尔斯，因为被青梅竹马的女友茱莉亚甩掉，进入秘洞探险，遇到远古巨龙，将其屠杀服用龙血后，离开小山村，进入了双子星学院。
义父母的迷一般死亡牵扯出雷尔斯父亲的未亡之谜，“我们是神所摒弃的人……”父亲的信中是否预示着他今后不平凡的人生。怒闯公爵府后才知道，凶手竟然另有他人。兽人族风云后冒名从，不想却是重遇故人。罗德兰王国的叛乱震惊大陆，身为一朝之相竟然是篡权通敌之人，无形的命运之手终令雷尔斯踏上了历史的舞台…
意气相投的师姐、狡猾的公主、吃人不吐骨头的罗兰、异乡的双面女神医、深宫里面掌权的女皇后，雷尔斯生命中最重要的五个女人，也是五个麻烦…`, b.Introduction)
}

func TestBook_Fetch_Free(t *testing.T) {
	var ctx = context.Background()

	var b = Book{ID: "8361"}
	var err = b.Fetch(ctx)
	require.NoError(t, err)
	assert.Equal(t, "光之子", b.Title)
	assert.Equal(t, "4921", b.Author.ID)
	assert.Equal(t, "唐家三少", b.Author.Name)
	assert.Equal(t, []string{"完本", "签约", "免费"}, b.Tags)
	assert.Equal(t, C奇幻, b.Category)
	assert.Equal(t, SC现代魔法, b.SubCategory)
	assert.Equal(t, "https://bookcover.yuewen.com/qdbimg/349573/8361/180", b.CoverURL)
	assert.Equal(t, time.Date(2008, 01, 14, 0, 0, 0, 0, TZ), b.LastUpdated)
	assert.Equal(t, uint64(731700), b.WordCount)
	assert.LessOrEqual(t, uint64(476300), b.TotalRecommendCount)
	// assert.Equal(t, uint64(68), b.WeekRecommendCount)
	assert.Equal(t, "魔兽践踏，巨龙咆哮，巫师诅咒，魔法璀璨之光照耀知识灯塔！", b.Summary)
	assert.Equal(t,
		`一个懒惰的少年，因性格原因选学了无人问津的光系魔法，却无意中踏近了命运的巨轮，一步一步的成为了传说中的大魔导师。正是在他的努力下结束了东西大陆的分界，让整个大陆不再有种族之分，成为了后世各族共尊的光之子。`, b.Introduction)
}
