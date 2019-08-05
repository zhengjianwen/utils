package main

import (
	"fmt"
	"strings"
)
func main() {

	fmt.Println(" strings.Contains:查找子串是否在指定的字符串中")
	fmt.Println(strings.Contains("HaiRui", "Hai")) //true
	fmt.Println(strings.Contains("HaiRui", "MeiZi")) //false
	fmt.Println(strings.Contains("HaiRui", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true 这里要特别注意
	fmt.Println(strings.Contains("我是中国人", "我"))     //true

	fmt.Println(" ContainsAny 函数的用法")
	fmt.Println(strings.ContainsAny("HaiRui", "b"))        // false
	fmt.Println(strings.ContainsAny("HaiRui", "u & i"))    // true
	fmt.Println(strings.ContainsAny("HaiRui", ""))         // false
	fmt.Println(strings.ContainsAny("", ""))               // false

	fmt.Println(" ContainsRune 函数的用法")
	fmt.Println(strings.ContainsRune("我是中国", '我')) // true 注意第二个参数，用的是字符

	fmt.Println(" Count:统计方法")
	fmt.Println(strings.Count("cheese", "e")) // 3
	fmt.Println(strings.Count("five", ""))    // before & after each rune result: 5 , 源码中有实现

	fmt.Println(" EqualFold:比较2个参数忽略大小写比较")
	fmt.Println(strings.EqualFold("Go", "go")) //true

	fmt.Println(" Fields：将字符串拆成列表，根据英文单词边界拆分")
	fmt.Println("Fields are:", strings.Fields("  foo bar  baz   ")) //["foo" "bar" "baz"] 返回一个列表

	fmt.Println(" HasPrefix：字符串前缀是以什么开头")
	fmt.Println(strings.HasPrefix("HaiRui", "Hai")) //前缀是以Hai开头的

	fmt.Println(" HasSuffix:字符串前缀是以什么结尾")
	fmt.Println(strings.HasSuffix("HaiRui", "i")) //后缀是以i开头的
	fmt.Println("")
	fmt.Println(" Index：查询字符串的索引位置并返回")
	fmt.Println(strings.Index("HaiRui", "a")) // 返回第一个匹配字符的位置，这里是1
	fmt.Println(strings.Index("HaiRui", "i")) // 第一个匹配的位置 2
	fmt.Println(strings.Index("HaiRui", "s")) // 不存在返回 -1
	fmt.Println(strings.Index("我是中国人", "中")) // 在存在返回 6

	fmt.Println(" IndexAny ：索引")
	fmt.Println(strings.IndexAny("我是中国人", "中")) // 在存在返回 6
	fmt.Println(strings.IndexAny("我是中国人", "和")) // 在存在返回 -1
	fmt.Println("")
	fmt.Println(" IndexRune：在rune类型下的索引")
	fmt.Println(strings.IndexRune("NLT_abc", 'b')) // 返回第一个匹配字符的位置，这里是4
	fmt.Println(strings.IndexRune("NLT_abc", 's')) // 在存在返回 -1
	fmt.Println(strings.IndexRune("我是中国人", '中'))   // 在存在返回 6
	fmt.Println(" Join：将列表拼接成字符串")
	s := []string{"I", "Learn", "Go"}
	fmt.Println(strings.Join(s, " ")) // 返回字符串：I Learn Go

	fmt.Println(" LastIndex:查询最后出现的索引")
	fmt.Println(strings.LastIndex("go gopher", "go")) // 3

	fmt.Println(" LastIndexAny：查询最后出现的索引 ")
	fmt.Println(strings.LastIndexAny("go gopher", "go")) // 4
	fmt.Println(strings.LastIndexAny("我是中国人", "中"))  // 6

	fmt.Println(" Repeat：返回一个重复字符串")
	fmt.Println("ba" + strings.Repeat("na", 2)) //banana
	fmt.Println("")
	fmt.Println(" Replace：替换字符串")
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	//oinky oinky oink
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
	//moo moo moo  最后最后一个参数小于0则全部替换
	fmt.Println(" Split：根据指定字符分割，")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	//["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	//["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	// 没有边界的分割 [" " "x" "y" "z" " "]
	fmt.Printf("%q\n", strings.Split("", "Hairui"))
	// [""] 没有也会有一个

	fmt.Println(" SplitAfter:分割包含分隔符")
	fmt.Printf("%q\n", strings.SplitAfter("/home/m_ta/src", "/")) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" SplitAfterN:分割包含分隔符，分割次数 -1为全部")
	fmt.Printf("%q\n", strings.SplitAfterN("/home/m_ta/src", "/", 2))  //["/" "home/m_ta/src"]
	fmt.Printf("%q\n", strings.SplitAfterN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]
	fmt.Println("")
	fmt.Println(" SplitN:根据特殊字符串分割，-1为全部分割")
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 1))
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", 2))  //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("/home/m_ta/src", "/", -1)) //["" "home" "m_ta" "src"]
	fmt.Printf("%q\n", strings.SplitN("home,m_ta,src", ",", 2))   //["/" "home/" "m_ta/" "src"]
	fmt.Printf("%q\n", strings.SplitN("#home#m_ta#src", "#", -1)) //["/" "home/" "m_ta/" "src"]

	fmt.Println(" Title：单词首字母变大写") //这个函数，还真不知道有什么用
	fmt.Println(strings.Title("i am hairui"))

	fmt.Println(" ToLower：字符串全部变小写")
	fmt.Println(strings.ToLower("Gopher")) //gopher

	fmt.Println(" ToLowerSpecial 函数的用法")


	fmt.Println(" ToTitle：全部变大写")
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("loud 中国"))

	fmt.Println(" Replace ：替换字符串，最后-1为全部替换，其他传几个就为几个")
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", 2)) // aBaACEDF
	//第四个参数小于0，表示所有的都替换， 可以看下golang的文档
	fmt.Println(strings.Replace("ABAACEDF", "A", "a", -1)) // aBaaCEDF

	fmt.Println(" ToUpper：字符串全部大写")
	fmt.Println(strings.ToUpper("Gopher")) //GOPHER

	fmt.Println(" Trim：去除2边指定字符")
	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! ")) // ["Achtung"]

	fmt.Println(" TrimLeft：去除左边指定字符")
	fmt.Printf("[%q]", strings.TrimLeft(" !!! Achtung !!! ", "! ")) // ["Achtung !!! "]

	fmt.Println(" TrimSpace：去除2边空格")
	fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n")) // a lone gopher
}