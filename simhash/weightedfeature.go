package simhash

import (
	"hash/fnv"
	"regexp"
)

// weightedfeature implements feature type
type WeightedFeature struct {
	sum    [16]byte
	weight int
}

// Sum returns the 128-bit hash of this WeightedFeature
func (f WeightedFeature) Sum() [16]byte {
	return f.sum
}

// Weight returns the weight of this WeightedFeature
func (f WeightedFeature) Weight() int {
	return f.weight
}

// Returns a new WeightedFeature representing the given byte slice, using a weight of 1
func newWeightedFeature(f []byte) WeightedFeature {
	weight := weightBytes(f)
	if unicodeBoundaries.Match(f) {
		f = md5sumUnicodeByte(f)
	}
	h := fnv.New128()
	h.Write(f)
	b := []byte{}
	b = h.Sum(b)
	c := [16]byte{}
	for index, bi := range b {
		c[index] = bi
	}
	return WeightedFeature{c, weight}
}

// Splits the given []byte using the given regexp, then returns a slice
// containing a Feature constructed from each piece matched by the regexp
func getWeightedFeatures(b []byte, r *regexp.Regexp) []Feature {
	//r = regexp.MustCompile("[\u4e00-\u9fa5]|[a-zA-Z0-9]+")
	words := r.FindAll(b, -1)
	//TODO should make it more solid
	//features := make([]Feature, len(words))
	features := []Feature{}
	for _, w := range words {
		if len(w) < 1 {
			continue
		}
		features = append(features, newWeightedFeature(w))
	}
	return features
}

// weight it by simple strategy
func weightBytes(b []byte) int {
	// simple level
	if len(b) <= 2 {
		return 1
	}
	str := string(b)
	// level 0, the de character exists too many times
	if str == "的" {
		return 0
	}
	levelonecharacters := []string{"中", "一", "是", "不", "有", "上", "在", "人", "国"}
	for _, v := range levelonecharacters {
		if str == v {
			return 1
		}
	}
	leveltwocharacters := []string{"也",
		"化",
		"现",
		"记",
		"得",
		"同",
		"法",
		"用",
		"前",
		"方",
		"第",
		"对",
		"公",
		"分",
		"以",
		"于",
		"能",
		"主",
		"个",
		"到",
		"要",
		"党",
		"北",
		"出",
		"发",
		"年",
		"股",
		"后",
		"我",
		"时",
		"作",
		"汉",
		"这",
		"全",
		"来",
		"文",
		"为",
		"了",
		"会",
		"语",
		"大",
		"和",
		"字",
		"球"}
	for _, v := range leveltwocharacters {
		if str == v {
			return 2
		}
	}

	return 5
}

/*
中文统计规律如下:
35000字中
去除词
1078	的
一类word:
223	国
226	人
238	在
252	上
275	有
284	不
346	是
353	一
353	中
二类word:
101	也
102	化
102	现
102	记
103	得
104	同
106	法
108	用
114	前
114	方
120	第
122	对
123	公
124	分
126	以
127	于
127	能
129	主
131	个
131	到
131	要
135	党
136	北
137	出
138	发
139	年
139	股
143	后
145	我
146	时
149	作
158	汉
158	这
160	全
166	来
176	文
178	为
180	了
182	会
188	语
190	大
195	和
215	字
218	球
统计文本增加于附件中
因为引入了一片百度百科文字wiki和一片足球新闻，所以字和球反常得多
*/
