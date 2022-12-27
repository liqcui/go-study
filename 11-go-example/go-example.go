package main

import (
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func CountChineseNums() {
	//1.判断字符串中汉子的数量
	s1 := "Hello北京你好"
	//1.依次拿到字符串中的字符
	var count int
	for _, c := range s1 {
		//2.判断当前这个字符是不是汉字,貌似有问题，需要进一步验证
		if unicode.Is(unicode.Javanese, c) {
			count++
		}
	}
	fmt.Println(count)
}

//回文判断
//字符哦那个左往右和从右往左读是一样的，那么就是回文
//上海自来水来自海上
//黄山落叶松叶落黄山
func loopreading() {
	ss := "山西运煤车煤运西山"

	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r, c)
	}
	fmt.Println(r)
	fmt.Println(len(ss) / 2)
	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
func ExtractServiceIP() {
	var serviceOutput string
	serviceOutput = `
	Name:              node-tuning-operator
	Namespace:         openshift-cluster-node-tuning-operator
	Labels:            name=node-tuning-operator
	Annotations:       include.release.openshift.io/ibm-cloud-managed: true
					   include.release.openshift.io/self-managed-high-availability: true
					   include.release.openshift.io/single-node-developer: true
					   service.alpha.openshift.io/serving-cert-signed-by: openshift-service-serving-signer@1647833506
					   service.beta.openshift.io/serving-cert-secret-name: node-tuning-operator-tls
					   service.beta.openshift.io/serving-cert-signed-by: openshift-service-serving-signer@1647833506
	Selector:          name=cluster-node-tuning-operator
	Type:              ClusterIP
	IP Family Policy:  SingleStack
	IP Families:       IPv4
	IP:                None
	IPs:               None
	Port:              <unset>  60000/TCP
	TargetPort:        60000/TCP
	Endpoints:         10.128.0.8:60000
	Session Affinity:  None
	Events:            <none>
	`
	endPointReg, _ := regexp.Compile(".*Endpoints.*")
	endPointIP := endPointReg.FindString(serviceOutput)
	a := strings.ReplaceAll(endPointIP, " ", "")
	b := []rune(endPointIP)
	endPointIPArr := strings.Split(a, ":")
	fmt.Println(a)
	fmt.Println(b)
	c := endPointIPArr[1] + ":" + endPointIPArr[2]
	fmt.Println(c)

	uEnc := base64.URLEncoding.EncodeToString([]byte(serviceOutput))
	uEnc2 := base64.StdEncoding.EncodeToString([]byte(serviceOutput))
	fmt.Println(uEnc)
	fmt.Println(uEnc2)
}

func trunctFromString() {
	originString := `
	-----BEGIN CERTIFICATE-----
MIIE6TCCA9GgAwIBAgIIIA1B9kFfwm0wDQYJKoZIhvcNAQELBQAwNjE0MDIGA1UE
Awwrb3BlbnNoaWZ0LXNlcnZpY2Utc2VydmluZy1zaWduZXJAMTY0Nzg1ODUyMjAe
Fw0yMjAzMjExNTU3MzdaFw0yNDAzMjAxNTU3MzhaMEwxSjBIBgNVBAMMQSoubm9k
ZS10dW5pbmctb3BlcmF0b3Iub3BlbnNoaWZ0LWNsdXN0ZXItbm9kZS10dW5pbmct
b3BlcmF0b3Iuc3ZjMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAr8fS
ImvOrFbYsVbHBz6G3FH51IN79JbLHdmQKlScMefTSgUW/y3/tcao+JX/GG6f1dlt
k2cu570Eahq81E8KTwxQfrC8FrQkTiMjkyjJt2v5SdG6zSXj3Fr/6OEE2D4w2JxJ
5Px7I4SO6xWZM7OIXIs5abPIYeed74nDj6D7okI7G3CHdIAuMgpv9yJXe18EeJuv
lwY/bxA87DIhne0VJj6lMH5aRhgDqR0njBLxFg6JKWFoCT8TE51zy/sMPgXcWka9
uVCZvu8c4mxuRsKcPnyyQNlmQDv8lOXnBgmECusgAyMDiI2OYj8tP4rgNwm6UOgG
SqJYuq90xuWC7qnRMwIDAQABo4IB4zCCAd8wDgYDVR0PAQH/BAQDAgWgMBMGA1Ud
JQQMMAoGCCsGAQUFBwMBMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFE2U7PqmSvBA
iDAfEA9PJ7XyupuEMB8GA1UdIwQYMBaAFJBYvDBPFG/DzUlZhQ14EOKyGC2vMIIB
MQYDVR0RBIIBKDCCASSCQSoubm9kZS10dW5pbmctb3BlcmF0b3Iub3BlbnNoaWZ0
LWNsdXN0ZXItbm9kZS10dW5pbmctb3BlcmF0b3Iuc3Zjgk8qLm5vZGUtdHVuaW5n
LW9wZXJhdG9yLm9wZW5zaGlmdC1jbHVzdGVyLW5vZGUtdHVuaW5nLW9wZXJhdG9y
LnN2Yy5jbHVzdGVyLmxvY2Fsgj9ub2RlLXR1bmluZy1vcGVyYXRvci5vcGVuc2hp
ZnQtY2x1c3Rlci1ub2RlLXR1bmluZy1vcGVyYXRvci5zdmOCTW5vZGUtdHVuaW5n
LW9wZXJhdG9yLm9wZW5zaGlmdC1jbHVzdGVyLW5vZGUtdHVuaW5nLW9wZXJhdG9y
LnN2Yy5jbHVzdGVyLmxvY2FsMDUGCysGAQQBkggRZAIBBCYTJGQyMzQ1NTQ5LTE2
ZWItNDBmZi1iMDg2LTRjYzkwY2IzNGQ2YzANBgkqhkiG9w0BAQsFAAOCAQEA1oVy
IN8earL/qKZ1gezeknr21rEAiSnCa8Hb9OpQtUZNVoly/SxhGsS5KRp966phQgFl
LYkriOXHPHYwQCIu8Wx4g7k4wiPHsyfpyiv6ijZj/A+ytu1QP93ioEb8T6/LI/fw
3XjcsJaDmfbOGqaB/cvXBIw+O70zZxAFXaJeLQtZKquI0b+5jEdnrBpTvO+PstHs
FUjDNJ5XMJop40EYyIUWX8u36uwzwiuFYcLBRrv/pkGfdYcp/Aq87aIzRj5nsZRD
54e6yIMpjmuSGWTO+jkFxpI4g9fWt8+OtDXGHSUCx+iOTttGE/gSk4QheUjpzWVB
zovT0mQgMUyBviVs3w==
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIDUTCCAjmgAwIBAgIILFJYOqLi8OswDQYJKoZIhvcNAQELBQAwNjE0MDIGA1UE
Awwrb3BlbnNoaWZ0LXNlcnZpY2Utc2VydmluZy1zaWduZXJAMTY0Nzg1ODUyMjAe
Fw0yMjAzMjExMDI4NDFaFw0yNDA1MTkxMDI4NDJaMDYxNDAyBgNVBAMMK29wZW5z
aGlmdC1zZXJ2aWNlLXNlcnZpbmctc2lnbmVyQDE2NDc4NTg1MjIwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDjq3cdNOC8uhO33f7uDaVmvTDfjfYTSwx8
SFT7fgQEZb3zYzCK+jyzN8TauIYDaFreNO+oA1sB3Qm29SkJgynp58aePHYIAYtX
7MhxxJlwbLuk+SeHCzSMlt5kiXS5RE7XdN24qIXeoQvsZvvhn2RKDlmT7D19Rm7a
TPfJHD3kSpUaYFI+4r9u0Z7lMy9QrUgVgTti3zOT8041Ne9/2OW/gtbYuudJ916+
0j2neDPRqbAtodNTE3FO0qxFkORlJTAeT1Km5+a3fadXZHBm5y5EO7rStUse576k
ks2k0AR1unYU7w0EF44p/t2sPQEtdv1h+asuR69rUW+CgRlTDv+fAgMBAAGjYzBh
MA4GA1UdDwEB/wQEAwICpDAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBSQWLww
TxRvw81JWYUNeBDishgtrzAfBgNVHSMEGDAWgBSQWLwwTxRvw81JWYUNeBDishgt
rzANBgkqhkiG9w0BAQsFAAOCAQEA31zH713lbDUUkDajftmFMLnL7HZ4jTSdQJoi
0PB+OOLycQZ2HiXgS1dzGl/8tGVXomS7H09Lf4T+47KIfvBmMUyR0jZjFza1bxdL
CBNM4OSGyaHZmtS9ffUM5kF3i/g4s96i38asKEoA+ZsOfVqkA0ceJgv9m8cbKom3
wqf3kKBpXO0MdqGkcKw97ryRIHlV6ay9EhbYuZfLVj7it0jF+Z1n6wq75AHDRSFf
5dEKvGG2u6xQprT5Nj6/ygeGzW/qsAQZ/eeC9TDyKx7WOAWgjfJEOCVx09z0BI4f
yDcZeHquwPNs5Tm4Zg8D3u1SeCsGvJBGi/eh7yjeRHS9vh+7nw==
-----END CERTIFICATE-----`

	trimL := strings.Trim(originString, "-----END CERTIFICATE-----")
	fmt.Println(trimL)
}

// Base64 Decode
func BASE64DecodeStr(src string) string {
	plaintext, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return ""
	}
	return string(plaintext)
}

/*
有50枚金币，需要分配给以下几个人：Mattew,Sarah,Augustus, Heidi, Emiller, Giana
1. 名字中包含一个e或者E分一枚金币
2. 名字中包含一个i或者I分2枚金币
3. 名字中包含一个o或者O分3枚金币
4. 名字中包含一个u或者U分4枚金币
*/

var (
	coins = 5000
	users = []string{
		"Mattew", "Sarah", "Augustus", "Heidi", "Emiller", "Giana",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (left int) {

	// 1.依次拿到每个人的名字
	for _, name := range users {
		// 2. 拿到一个人的名字根据分金币的规则去分金币
		for _, c := range name {
			switch c {
			case 'e', 'E':
				// 2.1每个人分的金币应该保存到distribution中
				// 2.2还要记录下剩余的金币数
				distribution[name] += 1
				coins -= 1
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3
			case 'u', 'U':
				distribution[name] += 4
				coins -= 4
			}
		}
		// 3.整个第二步执行完就能得到最终每个人分的金币数和剩余金币数

	}
	left = coins
	return left
}

func main() {
	// CountChineseNums()
	// loopreading()
	left := dispatchCoin()
	fmt.Println("剩下金币：", left)
	for k, v := range distribution {
		fmt.Printf("%s:%d\n", k, v)
	}
	//ExtractServiceIP()
	//trunctFromString()

	// 	fmt.Printf("%T", st1)
	// 	pt := BASE64DecodeStr(str1)
	// 	fmt.Println(pt)
}
