package main

import (
	"fmt"

	"github.com/neoxue/simhash"
)

func main() {
	var docs = [16][]byte{
		[]byte("this is a test phrase"),
		[]byte("this is c test phrass"),
		[]byte("foo bar"),
		[]byte("看公司操盘。如果道德水平高，又有抱负，可能会继续平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太多.aaaa..我个人看来，2020年绝对会超预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大股东操守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊"),
		[]byte("看公司操盘。如果道德水平高，又有平滑业绩。如果只个人看来，2020年绝对会超预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大股东操守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊"),
		[]byte("为此我们需要一种应对于海量数据场景的去重方案，经过研究发现有种叫 local sensitive hash 局部敏感哈希 的东西，据说这玩意可以把文档降维到hash数字，数字两两计算运算量要小很多。"),
		[]byte("为此我们需要一种应对于海量数据场景的去重方案，经过研究发现有种叫 local sensitive hash 局部敏感哈希 的东西，据说这玩意可以把文档降维到hash数字，数字两两计算运算量要小很多。"),
		[]byte("s"),
		[]byte("看公司操盘。如果道德水平高，又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太多...我个人看来，2020年绝对会超预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大股东操守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊"),
		[]byte("大跌眼镜，居然计算耗费4秒。假设我们一天需要比较100w次，光是比较100w次的数据是否重复就需要4s，就算4s一个文档，单线程一分钟才处理15个文档，一个小时才900个，一天也才21600个文档，这个数字和一天100w相差甚远，需要多少机器和资源才能解决。"),
		[]byte("为此现有种叫 local sensitive hash 局部敏感哈希 的东西，据说量要小很多。"),
		[]byte("公"),
	}
	docs[11] = []byte("共")
	docs[12] = []byte("应该")
	docs[13] = []byte("a")
	docs[14] = []byte("s")
	//fmt.Println("hihihi")
	//bts := []byte("simple test上海")
	//fmt.Println(simhash.NewWordFeatureSet(bts).GetFeatures())
	//os.Exit(0)
	t := [16]byte{}
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("this is a")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("this Is ABBSS")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("this is abbss 2999 9")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("测")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("试")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	t = simhash.Simhash(simhash.NewWordFeatureSet([]byte("试人")))
	fmt.Printf("here is a simhash")
	fmt.Println(t)
	/*
		fmt.Printf("here is simhash")
		fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet([]byte("试人"))))
		fmt.Printf("here is simhash")
		fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet([]byte("试人"))))
		fmt.Printf("here is simhash")
		fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet([]byte("试人"))))
	*/
	//os.Exit(0)

	hashes := make([][16]byte, len(docs))
	for i, d := range docs {
		hashes[i] = simhash.Simhash(simhash.NewWordFeatureSet(d))
		//fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet(d)))
		fmt.Printf("Simhash of %x:%s\n", hashes[i], d)
		//os.Exit(0)
	}
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[11], docs[12], simhash.Compare(hashes[11], hashes[12]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[13], docs[14], simhash.Compare(hashes[13], hashes[14]))

	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[1], simhash.Compare(hashes[0], hashes[1]))
	fmt.Printf("Comparison of `%s` and `%s`: %d\n", docs[0], docs[2], simhash.Compare(hashes[0], hashes[2]))
	//fmt.Printf("Comparison is %d, `%s` and `%s`\n", simhash.Compare(hashes[3], hashes[3]), docs[3], docs[3])
	fmt.Printf("Comparison is %d, `%s` and `%s`\n", simhash.Compare(hashes[3], hashes[4]), docs[3], docs[4])
	fmt.Printf("Comparison is %d, `%s` and `%s`\n", simhash.Compare(hashes[5], hashes[6]), docs[5], docs[6])

	s := []byte("")
	s1 := []byte("")

	s = []byte("看公司操盘。如果道德水平高，又有平续平滑业绩，倒不是看好大股东操守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")
	s = []byte("看公司操盘。如果道德水平高,我们记者水平差，又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")
	s = []byte("看公司操盘。如果道德水平高,我们记者水平差,人类,特斯拉，又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")
	s1 = []byte("看公司操盘。如果道德水平高,我们记者水平差,人类,特死啦，又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")
	s1 = []byte("看公司操盘。如果水平差,人类,特斯拉，又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")
	s = []byte("看公司操盘。如果水平差,又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")

	s = []byte(`在做项目的过程中，使用正则表达式来匹配一段文本中的特定种类字符，是比较常用的一种方式，下面是对常用的正则匹配做了一个归纳整理。

匹配中文:[\u4e00-\u9fa5]

英文字母:[a-zA-Z]
数字:[0-9]

匹配中文，英文字母和数字及_:
^[\u4e00-\u9fa5_a-zA-Z0-9]+$

同时判断输入长度：
[\u4e00-\u9fa5_a-zA-Z0-9_]{4,10}

^[\w\u4E00-\u9FA5\uF900-\uFA2D]*$ 1、一个正则表达式，只含有汉字、数字、字母、下划线不能以下划线开头和结尾：
^(?!_)(?!.*?_$)[a-zA-Z0-9_\u4e00-\u9fa5]+$其中：
^与字符串开始的地方匹配
(?!_)　　不能以_开头
(?!.*?_$)　　不能以_结尾
[a-zA-Z0-9_\u4e00-\u9fa5]+　　至少一个汉字、数字、字母、下划线
$　　与字符串结束的地方匹配

放在程序里前面加@，否则需要\\进行转义 "^(?!_)(?!.*?_$)[a-zA-Z0-9_\u4e00-\u9fa5]+$"
（或者："^(?!_)\w*(? 
2、只含有汉字、数字、字母、下划线，下划线位置不限：
^[a-zA-Z0-9_\u4e00-\u9fa5]+$

3、由数字、26个英文字母或者下划线组成的字符串
^\w+$

4、2~4个汉字
"^[\u4E00-\u9FA5]{2,4}$";

5、
^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$

用：(Abc)+ 来分析：XYZAbcAbcAbcXYZAbcAb

XYZAbcAbcAbcXYZAbcAb6、
[^\u4E00-\u9FA50-9a-zA-Z_]
34555#5' -->34555#5'


publicbool RegexName(string str)
{
bool flag=Regex.IsMatch(str,@"^[a-zA-Z0-9_\u4e00-\u9fa5]+$");
returnflag;
}

Regex reg=new Regex("^[a-zA-Z_0-9]+$"); 
if(reg.IsMatch(s)) 
{ 
\\符合规则 
} 
else 
{ 
\\存在非法字符 
}

最长不得超过7个汉字，或14个字节(数字，字母和下划线)正则表达式
^[\u4e00-\u9fa5]{1,7}$|^[\dA-Za-z_]{1,14}$

常用正则表达式大全！

（例如：匹配中文、匹配html）

匹配中文字符的正则表达式： [u4e00-u9fa5] 
评注：匹配中文还真是个头疼的事，有了这个表达式就好办了

匹配双字节字符(包括汉字在内)：[^x00-xff]
评注：可以用来计算字符串的长度（一个双字节字符长度计2，ASCII字符计1）
匹配空白行的正则表达式：ns*r
评注：可以用来删除空白行
匹配HTML标记的正则表达式：<(S*?)[^>]*>.*?|<.*? />
评注：网上流传的版本太糟糕，上面这个也仅仅能匹配部分，对于复杂的嵌套标记依旧无能为力
匹配首尾空白字符的正则表达式：^s*|s*$
评注：可以用来删除行首行尾的空白字符(包括空格、制表符、换页符等等)，非常有用的表达式
匹配Email地址的正则表达式：^[a-zA-Z0-9][\w\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\w\.-]*[a-zA-Z0-9]\.[a-zA-Z][a-zA-Z\.]*[a-zA-Z]$

评注：表单验证时很实用

手机号：^((13[0-9])|(14[0-9])|(15[0-9])|(17[0-9])|(18[0-9]))\d{8}$

身份证：(^\d{15}$)|(^\d{17}([0-9]|X|x)$)

匹配网址URL的正则表达式：[a-zA-z]+://[^s]*
评注：网上流传的版本功能很有限，上面这个基本可以满足需求
匹配帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]{4,15}$
评注：表单验证时很实用
匹配国内电话号码：d{3}-d{8}|d{4}-d{7}
评注：匹配形式如 0511-4405222 或 021-87888822
匹配腾讯QQ号：[1-9][0-9]{4,}
评注：腾讯QQ号从10000开始
匹配中国邮政编码：[1-9]d{5}(?!d)
评注：中国邮政编码为6位数字
匹配身份证：d{15}|d{18}
评注：中国的身份证为15位或18位
匹配ip地址：d+.d+.d+.d+
评注：提取ip地址时有用
匹配特定数字：
^[1-9]d*$　 　 //匹配正整数
^-[1-9]d*$ 　 //匹配负整数
^-?[1-9]d*$　　 //匹配整数
^[1-9]d*|0$　 //匹配非负整数（正整数 + 0）
^-[1-9]d*|0$　　 //匹配非正整数（负整数 + 0）
^[1-9]d*.d*|0.d*[1-9]d*$　　 //匹配正浮点数
^-([1-9]d*.d*|0.d*[1-9]d*)$　 //匹配负浮点数
^-?([1-9]d*.d*|0.d*[1-9]d*|0?.0+|0)$　 //匹配浮点数
^[1-9]d*.d*|0.d*[1-9]d*|0?.0+|0$　　 //匹配非负浮点数（正浮点数 + 0）
^(-([1-9]d*.d*|0.d*[1-9]d*))|0?.0+|0$　　//匹配非正浮点数（负浮点数 + 0）
评注：处理大量数据时有用，具体应用时注意修正
匹配特定字符串：
^[A-Za-z]+$　　//匹配由26个英文字母组成的字符串
^[A-Z]+$　　//匹配由26个英文字母的大写组成的字符串
^[a-z]+$　　//匹配由26个英文字母的小写组成的字符串
^[A-Za-z0-9]+$　　//匹配由数字和26个英文字母组成的字符串
^w+$　　//匹配由数字、26个英文字母或者下划线组成的字符串
在使用RegularExpressionValidator验证控件时的验证功能及其验证表达式介绍如下:
只能输入数字：“^[0-9]*$”
只能输入n位的数字：“^d{n}$”
只能输入至少n位数字：“^d{n,}$”
只能输入m-n位的数字：“^d{m,n}$”
只能输入零和非零开头的数字：“^(0|[1-9][0-9]*)$”
只能输入有两位小数的正实数：“^[0-9]+(.[0-9]{2})?$”
只能输入有1-3位小数的正实数：“^[0-9]+(.[0-9]{1,3})?$”
只能输入非零的正整数：“^+?[1-9][0-9]*$”
只能输入非零的负整数：“^-[1-9][0-9]*$”
只能输入长度为3的字符：“^.{3}$”
只能输入由26个英文字母组成的字符串：“^[A-Za-z]+$”
只能输入由26个大写英文字母组成的字符串：“^[A-Z]+$”
只能输入由26个小写英文字母组成的字符串：“^[a-z]+$”
只能输入由数字和26个英文字母组成的字符串：“^[A-Za-z0-9]+$”
只能输入由数字、26个英文字母或者下划线组成的字符串：“^w+$”
验证用户密码:“^[a-zA-Z]w{5,17}$”正确格式为：以字母开头，长度在6-18之间，
只能包含字符、数字和下划线。
验证是否含有^%&',;=?$"等字符：“[^%&',;=?$x22]+”
只能输入汉字：“^[u4e00-u9fa5],{0,}$”
验证Email地址：“^w+[-+.]w+)*@w+([-.]w+)*.w+([-.]w+)*$”
验证InternetURL：“^http://([w-]+.)+[w-]+(/[w-./?%&=]*)?$”
验证身份证号（15位或18位数字）：“^d{15}|d{}18$”
验证一年的12个月：“^(0?[1-9]|1[0-2])$”正确格式为：“01”-“09”和“1”“12”
验证一个月的31天：“^((0?[1-9])|((1|2)[0-9])|30|31)$”
正确格式为：“01”“09”和“1”“31”。
匹配中文字符的正则表达式： [u4e00-u9fa5]
匹配双字节字符(包括汉字在内)：[^x00-xff]
匹配空行的正则表达式：n[s| ]*r
匹配HTML标记的正则表达式：/<(.*)>.*|<(.*) />/
匹配首尾空格的正则表达式：(^s*)|(s*$)
匹配Email地址的正则表达式：w+([-+.]w+)*@w+([-.]w+)*.w+([-.]w+)*
匹配网址URL的正则表达式：http://([w-]+.)+[w-]+(/[w- ./?%&=]*)?
(1)应用：计算字符串的长度（一个双字节字符长度计2，ASCII字符计1）
String.prototype.len=function(){return this.replace([^x00-xff]/g,"aa").length;}
(2)应用：javascript中没有像vbscript那样的trim函数，我们就可以利用这个表达式来实现
String.prototype.trim = function()
{
return this.replace(/(^s*)|(s*$)/g, "");
}
(3)应用：利用正则表达式分解和转换IP地址
function IP2V(ip) //IP地址转换成对应数值
{
re=/(d+).(d+).(d+).(d+)/g //匹配IP地址的正则表达式
if(re.test(ip))
{
return RegExp.$1*Math.pow(255,3))+RegExp.$2*Math.pow(255,2))+RegExp.$3*255+RegExp.$4*1
}
else
{
throw new Error("Not a valid IP address!")
}
}
(4)应用：从URL地址中提取文件名的javascript程序
s="http://www.juapk.com/forum.php";
s=s.replace(/(.*/){0,}([^.]+).*/ig,"$2") ;//Page1.htm
(5)应用：利用正则表达式限制网页表单里的文本框输入内容
用正则表达式限制只能输入中文：onkeyup="value=value.replace(/[^u4E00-u9FA5]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^u4E00-u9FA5]/g,'))"
用正则表达式限制只能输入全角字符：onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^uFF00-uFFFF]/g,'))"
用正则表达式限制只能输入数字：onkeyup="value=value.replace(/[^d]/g,') "onbeforepaste= "clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'))"
用正则表达式限制只能输入数字和英文：onkeyup="value=value.replace(/[W]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'`)

	s1 = []byte(`在做项目的过程中，使用正则表达式来匹配一段文本中的特定种类字符，是比较常用的一种方式，下面是对常用的正则匹配做了一个归纳整理。
看公司操盘。如果道德水平高，又有平续平滑业绩，倒不是看好大股东操守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊
匹配中文:[\u4e00-\u9fa5]

英文字母:[a-zA-Z]
数字:[0-9]

匹配中文，英文字母和数字及_:
^[\u4e00-\u9fa5_a-zA-Z0-9]+$

同时判断输入长度：
[\u4e00-\u9fa5_a-zA-Z0-9_]{4,10}

^[\w\u4E00-\u9FA5\uF900-\uFA2D]*$ 1、一个正则表达式，只含有汉字、数字、字母、下划线不能以下划线开头和结尾：
^(?!_)(?!.*?_$)[a-zA-Z0-9_\u4e00-\u9fa5]+$其中：
^与字符串开始的地方匹配
(?!_)　　不能以_开头
(?!.*?_$)　　不能以_结尾
[a-zA-Z0-9_\u4e00-\u9fa5]+　　至少一个汉字、数字、字母、下划线
$　　与字符串结束的地方匹配

放在程序里前面加@，否则需要\\进行转义 "^(?!_)(?!.*?_$)[a-zA-Z0-9_\u4e00-\u9fa5]+$"
（或者："^(?!_)\w*(? 
2、只含有汉字、数字、字母、下划线，下划线位置不限：
^[a-zA-Z0-9_\u4e00-\u9fa5]+$

3、由数字、26个英文字母或者下划线组成的字符串
^\w+$

4、2~4个汉字
"^[\u4E00-\u9FA5]{2,4}$";

5、
^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$

用：(Abc)+ 来分析：XYZAbcAbcAbcXYZAbcAb

XYZAbcAbcAbcXYZAbcAb6、
[^\u4E00-\u9FA50-9a-zA-Z_]
34555#5' -->34555#5'


publicbool RegexName(string str)
{
bool flag=Regex.IsMatch(str,@"^[a-zA-Z0-9_\u4e00-\u9fa5]+$");
returnflag;
}

Regex reg=new Regex("^[a-zA-Z_0-9]+$"); 
if(reg.IsMatch(s)) 
{ 
\\符合规则 
} 
else 
{ 
\\存在非法字符 
}

最长不得超过7个汉字，或14个字节(数字，字母和下划线)正则表达式
^[\u4e00-\u9fa5]{1,7}$|^[\dA-Za-z_]{1,14}$

常用正则表达式大全！

（例如：匹配中文、匹配html）

匹配中文字符的正则表达式： [u4e00-u9fa5] 
评注：匹配中文还真是个头疼的事，有了这个表达式就好办了

匹配双字节字符(包括汉字在内)：[^x00-xff]
评注：可以用来计算字符串的长度（一个双字节字符长度计2，ASCII字符计1）
匹配空白行的正则表达式：ns*r
评注：可以用来删除空白行
匹配HTML标记的正则表达式：<(S*?)[^>]*>.*?|<.*? />
评注：网上流传的版本太糟糕，上面这个也仅仅能匹配部分，对于复杂的嵌套标记依旧无能为力
匹配首尾空白字符的正则表达式：^s*|s*$
评注：可以用来删除行首行尾的空白字符(包括空格、制表符、换页符等等)，非常有用的表达式
匹配Email地址的正则表达式：^[a-zA-Z0-9][\w\.-]*[a-zA-Z0-9]@[a-zA-Z0-9][\w\.-]*[a-zA-Z0-9]\.[a-zA-Z][a-zA-Z\.]*[a-zA-Z]$

评注：表单验证时很实用

手机号：^((13[0-9])|(14[0-9])|(15[0-9])|(17[0-9])|(18[0-9]))\d{8}$

身份证：(^\d{15}$)|(^\d{17}([0-9]|X|x)$)

匹配网址URL的正则表达式：[a-zA-z]+://[^s]*
评注：网上流传的版本功能很有限，上面这个基本可以满足需求
匹配帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线)：^[a-zA-Z][a-zA-Z0-9_]{4,15}$
评注：表单验证时很实用
匹配国内电话号码：d{3}-d{8}|d{4}-d{7}
评注：匹配形式如 0511-4405222 或 021-87888822
匹配腾讯QQ号：[1-9][0-9]{4,}
评注：腾讯QQ号从10000开始
匹配中国邮政编码：[1-9]d{5}(?!d)
评注：中国邮政编码为6位数字
匹配身份证：d{15}|d{18}
评注：中国的身份证为15位或18位
匹配ip地址：d+.d+.d+.d+
评注：提取ip地址时有用
匹配特定数字：
^[1-9]d*$　 　 //匹配正整数
^-[1-9]d*$ 　 //匹配负整数
^-?[1-9]d*$　　 //匹配整数
^[A-Za-z0-9]+$　　//匹配由数字和26个英文字母组成的字符串
^w+$　　//匹配由数字、26个英文字母或者下划线组成的字符串
在使用RegularExpressionValidator验证控件时的验证功能及其验证表达式介绍如下:
只能输入数字：“^[0-9]*$”
只能输入n位的数字：“^d{n}$”
只能输入至少n位数字：“^d{n,}$”
只能输入m-n位的数字：“^d{m,n}$”
只能输入零和非零开头的数字：“^(0|[1-9][0-9]*)$”
只能输入由26个小写英文字母组成的字符串：“^[a-z]+$”
只能输入由数字和26个英文字母组成的字符串：“^[A-Za-z0-9]+$”
只能输入由数字、26个英文字母或者下划线组成的字符串：“^w+$”
验证用户密码:“^[a-zA-Z]w{5,17}$”正确格式为：以字母开头，长度在6-18之间，
用正则表达式限制只能输入数字：onkeyup="value=value.replace(/[^d]/g,') "onbeforepaste= "clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'))"
用正则表达式限制只能输入数字和英文：onkeyup="value=value.replace(/[W]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'`)

	//s = []byte("看公司操盘。如果水平差,又有平滑业绩。如果只想干一炮，业绩一定会让你目瞪口呆。A股这种货色不要太人啊多...我个人看来，2020年绝对会预期，但幅度不会太厉害。公司应该会继续平滑业绩，倒不是看好大守，主要是2020年只是员工持股解禁，如果是大股东解禁想跑，那就难说了,今天中午下雨啊")

	//s = []byte("test")
	//s = []byte(`test

	//`)
	fmt.Printf("Comparison is %d, `%s` and `%s`\n", simhash.Compare(simhash.Simhash(simhash.NewWordFeatureSet(s)), simhash.Simhash(simhash.NewWordFeatureSet(s1))), "s", "s1")

	s = []byte(`用正则表达式限制只能输入数字：onkeyup="value=value.replace(/[^d]/g,') "onbeforepaste= "clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'))"`)
	s1 = []byte(`用正则表达式限制只能输入数字和英文：onkeyup="value=value.replace(/[W]/g,') "onbeforepaste="clipboardData.setData('text',clipboardData.getData('text').replace(/[^d]/g,'`)
	s1 = []byte(`中文`)

	fmt.Println("hihihihi")
	fmt.Println("hihihihi")
	fmt.Println("hihihihi")
	fmt.Println("hihihihi")
	fmt.Println("hihihihi")
	fmt.Println("hihihihi")
	fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet(s)))
	fmt.Println(simhash.Simhash(simhash.NewWordFeatureSet(s1)))
	fmt.Printf("Comparison is %d, `%s` and `%s`\n", simhash.Compare(simhash.Simhash(simhash.NewWordFeatureSet(s)), simhash.Simhash(simhash.NewWordFeatureSet(s1))), "s", "s1")
}
