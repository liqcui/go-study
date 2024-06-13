package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Go语言字符串用双引号包裹
// Go语言用单引号包裹的是字符
func main() {
	s := "hello redhat"
	c1 := 'h'
	c2 := 'e'

	fmt.Println(s)
	fmt.Printf("%c and %c", c1, c2)
	//多行字符串
	s2 := `
	东临碣石
	以观沧海
	水何澹澹
	山岛耸峙`
	fmt.Println(s2)
	//字符串拼接
	s3 := "你好"
	s4 := "Golang"
	fmt.Println(s3 + " " + s4)
	ss1 := fmt.Sprintf("%s %s", s3, s4)
	fmt.Println(ss1)
	//分割
	s5 := "Golang, World"
	ret := strings.Split(s5, ",")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s5, "Go"))
	//前缀
	fmt.Println(strings.HasPrefix(s5, "Go"))
	//后缀
	fmt.Println(strings.HasSuffix(s5, "Go"))
	//索引
	s6 := "abcdef"
	fmt.Println(strings.Index(s6, "c"))
	//join
	fmt.Println(strings.Join(ret, "+"))
	for _, c := range s {
		fmt.Printf("%c\n", c)
	}
	//rune类型处理中文和日文
	ss2 := "big"
	//强制类型转换
	byteS1 := []byte(ss2)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))
	ss3 := "白萝卜"
	runeS2 := []rune(ss3)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
	ss4 := `Good Morning
	simple-kmod
	simple-procs
	have a good day
	`
	fmt.Println(ss4)
	resourceList := `
	good
	bad
	no`
	list := strings.Split(resourceList, "\n")
	fmt.Printf("Type:%T,%v", list, list)
	for i := 0; i < len(list); i++ {

		fmt.Println(list[i])
	}
	var emptystr string
	fmt.Println(len(emptystr))
	abc := "liqcui-oc4103-2pzp4-worker-a-lltt8.c.openshift-qe.internal aaaaaa ready,reschedule"
	def := strings.Split(abc, " ")
	fmt.Println(def[0])
	// for i, _ := range abc {
	// 	fmt.Printf("%c\n", abc[i])
	// }
	str1 := "xadeevdde33ddd"
	str2 := "xadeevdde33ddd"
	str3 := "addeeceeeee"
	if str1 == str2 {
		fmt.Println("str1 is equal str2")
	}

	if str1 == str3 {
		fmt.Println("str1 is equal str3")
	} else {
		fmt.Println("str1 isn't equal str3")
	}

	removeQuota := `'v1.10'`
	aaa := strings.Trim(removeQuota, "'")
	fmt.Println(aaa)

	splitString := "gpu-operator-certified.v1.10.1 gpu-operator-certified.v1.10.1"
	bbb := strings.Split(splitString, " ")
	fmt.Println(bbb[0])

	defaultSMPAffinity := "f,ffffffff"
	defaultSMPAffinity = strings.ReplaceAll(defaultSMPAffinity, ",", "")
	fmt.Println(defaultSMPAffinity)

	base, _ := strconv.ParseInt(defaultSMPAffinity, 16, 64)
	fmt.Println(strconv.FormatInt(base, 10))
	str4 := `["registry.redhat.io/openshift4/ose-kube-rbac-proxy","registry.redhat.io/openshift4/ose-cluster-nfd-operator@sha256:d8f110c97b4c8ffbc6811c738c4b39a3cb78afd7db8727b7c277f2f6b767dc48","registry.redhat.io/openshift4/ose-node-feature-discovery@sha256:ca5c7471a0c596879ced290aba99f7a36d01482c7e03abffed4ff21f3a37dc0a"] ["registry.redhat.io/openshift4/ose-kube-rbac-proxy@sha256:9bda1737e8c4f5eb1760e314641f4a75feab981c3b36c51e03fa4eb354ab49f9","registry.redhat.io/openshift4/ose-cluster-nfd-operator@sha256:3a0dcf177f57a3e55042fde30d408f946b81a73cdbbd5b868ad424038887cc20","registry.redhat.io/openshift4/ose-node-feature-discovery@sha256:70759e1ffb9aca0237027811b9f0c27005358915f3c31371f15027715fda3a36"]`
	fmt.Println(str4)
	var finalresult string
	str5 := strings.ReplaceAll(str4, "[", ",")
	str6 := strings.ReplaceAll(str5, "]", ",")
	str7 := strings.ReplaceAll(str6, `"`, "")
	str8 := strings.Split(str7, ",")
	fmt.Println(" ------------------------------ ")
	fmt.Println(str8)
	for i := 0; i < len(str8); i++ {
		if strings.Contains(str8[i], "node-feature-discovery") {
			finalresult = str8[i]
		}
	}
	fmt.Println(" ------------------------------ ")
	fmt.Println(finalresult)

	str9 := "2.4 2.4-el7 2.4-el8 2.4-ubi8 2.4-ubi9"
	arr1 := strings.Split(str9, " ")
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
		jsonpathstr := fmt.Sprintf(`-ojsonpath='{.status.tags[%v].conditions[?(@.status=="False")]}{.status.tags[%v].tag}'`, i, i)
		fmt.Println(jsonpathstr)

	}
	//ProcessNum()

	endPointStr := `Endpoints:         [fd01:0:0:1::11]:60000`
	endPointStr1 := strings.ReplaceAll(endPointStr, "         ", ",")
	endPointStrArr := strings.Split(endPointStr1, ",")
	fmt.Println(endPointStrArr[1])

	fmt.Println(strconv.Itoa(3))
	str10 := "00020"
	if strings.HasPrefix(str10, "0") {
		fmt.Println(strings.TrimLeft(str10, "0"))
	}

	stalldStdOut := `Aug 04 13:45:05 ip-10-0-16-195 systemd[1]: Starting Stall Monitor...
	Aug 04 13:45:05 ip-10-0-16-195 systemd[1]: Started Stall Monitor.
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: lockdown mode is off
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: /sys/kernel/debug/sched/features exists
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: /sys/kernel/debug/sched/debug exists
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: boosted pid 0 (undef) using SCHED_DEADLINE
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: using SCHED_DEADLINE for boosting
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: initial config_buffer_size set to 286720
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: detected new task format
	Aug 04 13:45:05 ip-10-0-16-195 stalld[239426]: single threaded mode`

	exp := regexp.MustCompile(`\[.*?\]`)

	// Matches would contain a slice of all successive matching text
	// based on the regular expression.
	matches := exp.FindAllString(stalldStdOut, -1)
	fmt.Println("All matches : " + strings.Join(matches, " "))
	nodeName := "worker01"
	debugNodeNamespace := "nto"
	cmd := []string{"ps", "-ef"}
	cargs := []string{"node/" + nodeName + "--", "chroot", "/host"}
	cargs = []string{"node/" + nodeName, "--to-namespace=" + debugNodeNamespace, "--", "chroot", "/host"}
	cargs = append(cargs, cmd...)
	fmt.Println(cargs[:])

	time1 := "05:37"
	minSec := strings.Split(time1, ":")
	timeMin := minSec[0]
	if strings.HasPrefix(timeMin, "0") {
		timeMin = strings.TrimLeft(timeMin, "0")
	}
	fmt.Println(timeMin)

	findStr := fmt.Sprintf(`find /sys/class/net -type l -not -lname *virtual* -a -not -name enP*" -printf %%f"\n"`)
	fmt.Println(findStr)

	xxx := 4
	a := strconv.Itoa(xxx)
	fmt.Println(a)
	masterNodeNamesStr := `node/maxu-ibm-xjnnk-master-0
	node/maxu-ibm-xjnnk-master-1
	node/maxu-ibm-xjnnk-master-2`
	// masterNodeNames := strings.Trim(masterNodeNamesStr, "'")
	masterNodeNamesArray := strings.Split(masterNodeNamesStr, "\n")
	firstMasterNodeName := strings.Split(masterNodeNamesArray[0], "/")
	fmt.Println(masterNodeNamesArray[0])
	fmt.Println(firstMasterNodeName[1])
}
