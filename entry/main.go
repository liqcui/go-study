package main

import (
	"fmt"

	"github.com/liqcui/openshift-psap-qe/config"
)

func main() {
	var mc config.MysqlConfig
	err := config.LoadIni("./config.ini", &mc)
	if err != nil {
		fmt.Printf("Load ini failed, err:%v\n", err)
		return
	}
	fmt.Println(mc.Address, mc.Port, mc.Username, mc.Password)
}
