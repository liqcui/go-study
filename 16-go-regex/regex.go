package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	// 	str := `
	// Module                  Size  Used by
	// simple_procfs_kmod     16384  0
	// simple_kmod            16384  0
	// xt_REDIRECT            16384  1
	// veth                   28672  0
	// nf_conntrack_netlink    49152  0
	// ipt_REJECT             16384  164
	// nf_reject_ipv4         16384  1 ipt_REJECT
	// xt_nat                 16384  3
	// geneve                 36864  0
	// ip6_udp_tunnel         16384  1 geneve
	// udp_tunnel             20480  1 geneve
	// xt_CT                  16384  4
	// ip6t_MASQUERADE        16384  1`

	// 	str4 := `
	// 2021/10/28 12:29:25 194 Sending: Pong
	// 2021/10/28 12:29:27 195 Received: Ping
	// 2021/10/28 12:29:27 195 Sending: Pong
	// 2021/10/28 12:29:29 196 Received: Ping
	// 2021/10/28 12:29:29 196 Sending: Pong
	// 2021/10/28 12:29:31 197 Received: Ping
	// 2021/10/28 12:29:31 197 Sending: Pong
	// 2021/10/28 12:29:33 198 Received: Ping
	// 2021/10/28 12:29:33 198 Sending: Pong`

	// 	match, err := regexp.MatchString(`^simple`, str)
	// 	fmt.Println("Match: ", match, " Error: ", err)
	// 	regexpstr, _ := regexp.Compile(`simple.*`)
	// 	fmt.Println(regexpstr.FindAllString(str, 2))
	// 	// keywords := "ip.*"
	// 	keywords := ".*Pong.*"
	// 	regstr, _ := regexp.Compile(keywords)
	// 	fmt.Printf("%T", regstr.FindAllString(str4, -1))
	// 	fmt.Println(regstr.FindAllString(str4, -1))

	// 	str2 := "Golang regular expressions example"
	// 	match2, err := regexp.MatchString(`^Golang`, str2)
	// 	fmt.Println("Match: ", match2, " Error: ", err)

	// 	str3 := "Golang expressions example"
	// 	regexp, _ := regexp.Compile("Gola([a-z]+)g")
	// 	fmt.Println(regexp.FindString(str3))

	//这个测试一个字符串是否符合一个表达式。
	match3, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match3)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。

	r, _ := regexp.Compile("p([a-z]+)ch")
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	fmt.Println(r.MatchString("peach"))
	//这是查找匹配字符串的。
	fmt.Println(r.FindString("peach punch"))
	//这个也是查找第一次匹配的字符串的，但是返回的匹配开始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindStringIndex("peach punch"))
	//Submatch 返回完全匹配和局部匹配的字符串。例如，这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch"))
	//类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	//带 All 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//All 同样可以对应到上面的所有函数。
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	//这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//上面的例子中，我们使用了字符串作为参数，并使用了如 MatchString 这样的方法。我们也可以提供 []byte参数并将 String 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach")))
	//创建正则表示式常量时，可以使用 Compile 的变体MustCompile 。因为 Compile 返回两个值，不能用语常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)
	//regexp 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	//Func 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	//https://wizardforcel.gitbooks.io/golang-stdlib-ref/content/107.html
	//https://yourbasic.org/golang/regexp-cheat-sheet/

	string1 := "nodepool.hypershift.openshift.io/psap-qe-gcluster01-us-east-2a nodepool.hypershift.openshift.io/psap-qe-gcluster02-us-east-2a"
	reg1 := regexp.MustCompile(".*psap-qe-gcluster02.*")
	nodepoolName := reg1.FindAllString(string1, -1)
	nodepoolName2 := reg1.FindString(string1)
	fmt.Println(nodepoolName[0])
	fmt.Println(nodepoolName2)
}
