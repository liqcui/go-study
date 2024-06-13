// package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"strings"
// // )

// // func remove() string {
// // 	jsonstring := `[
// //     {
// //       "apiVersion": "nvidia.com/v1",
// //       "kind": "ClusterPolicy",
// //       "metadata": {
// //         "name": "gpu-cluster-policy"
// //       },
// //       "spec": {
// //         "dcgmExporter": {
// //           "config": {
// //             "name": ""
// //           }
// //         },
// //         "dcgm": {
// //           "enabled": true
// //         },
// //         "daemonsets": {
// //         },
// //         "devicePlugin": {
// //         },
// //         "driver": {
// //           "enabled": true,
// //           "use_ocp_driver_toolkit": true,
// //           "repoConfig": {
// //             "configMapName": ""
// //           },
// //           "certConfig": {
// //             "name": ""
// //           },
// //           "licensingConfig": {
// //             "nlsEnabled": false,
// //             "configMapName": ""
// //           },
// //           "virtualTopology": {
// //             "config": ""
// //           },
// //           "kernelModuleConfig": {
// //             "name": ""
// //           }
// //         },
// //         "gfd": {
// //         },
// //         "migManager": {
// //           "enabled": true
// //         },
// //         "nodeStatusExporter": {
// //           "enabled": true
// //         },
// //         "operator": {
// //           "defaultRuntime": "crio",
// //           "deployGFD": true,
// //           "initContainer": {
// //           }
// //         },
// //         "mig": {
// //           "strategy": "single"
// //         },
// //         "toolkit": {
// //           "enabled": true
// //         },
// //         "validator": {
// //           "plugin": {
// //             "env": [
// //               {
// //                 "name": "WITH_WORKLOAD",
// //                 "value": "true"
// //               }
// //             ]
// //           }
// //         }
// //       }
// //     }
// //   ]`
// // 	// fmt.Printf("%T", jsonstring)
// // 	// dst := &bytes.Buffer{}
// // 	// if err := json.Indent(dst, []byte(jsonstring), " ", " "); err != nil {
// // 	// 	panic(err)
// // 	// }
// // 	// fmt.Println(dst.String())
// // 	// 序列化
// // 	// b, err := json.Marshal(jsonstring)
// // 	// if err != nil {
// // 	// 	fmt.Printf("marshal failed,err:%v", err)
// // 	// }

// // 	//fmt.Println(string(b))
// // 	//反序列化
// // 	tmp1 := strings.TrimLeft(jsonstring, "[")
// // 	output := strings.TrimRight(tmp1, "]")
// // 	stringify_jq_output := strings.TrimSpace(string(output))
// // 	return stringify_jq_output
// // }

// // func func1() {
// // 	jsonstring := `[
// //     {
// //       "apiVersion": "nvidia.com/v1",
// //       "kind": "ClusterPolicy",
// //       "metadata": {
// //         "name": "gpu-cluster-policy"
// //       },
// //       "spec": {
// //         "dcgmExporter": {
// //           "config": {
// //             "name": ""
// //           }
// //         },
// //         "dcgm": {
// //           "enabled": true
// //         },
// //         "daemonsets": {
// //         },
// //         "devicePlugin": {
// //         },
// //         "driver": {
// //           "enabled": true,
// //           "use_ocp_driver_toolkit": true,
// //           "repoConfig": {
// //             "configMapName": ""
// //           },
// //           "certConfig": {
// //             "name": ""
// //           },
// //           "licensingConfig": {
// //             "nlsEnabled": false,
// //             "configMapName": ""
// //           },
// //           "virtualTopology": {
// //             "config": ""
// //           },
// //           "kernelModuleConfig": {
// //             "name": ""
// //           }
// //         },
// //         "gfd": {
// //         },
// //         "migManager": {
// //           "enabled": true
// //         },
// //         "nodeStatusExporter": {
// //           "enabled": true
// //         },
// //         "operator": {
// //           "defaultRuntime": "crio",
// //           "deployGFD": true,
// //           "initContainer": {
// //           }
// //         },
// //         "mig": {
// //           "strategy": "single"
// //         },
// //         "toolkit": {
// //           "enabled": true
// //         },
// //         "validator": {
// //           "plugin": {
// //             "env": [
// //               {
// //                 "name": "WITH_WORKLOAD",
// //                 "value": "true"
// //               }
// //             ]
// //           }
// //         }
// //       }
// //     }
// //   ]`
// // 	// var result map[string]interface{}
// // 	// bytes = []byte(jsonstring)
// // 	// json.Unmarshal(bytes, &result)
// // 	// fmt.Println(result)
// // 	gpu_csv_tmp := strings.TrimLeft(jsonstring, "[")
// // 	gpu_csv_string := strings.TrimRight(gpu_csv_tmp, "]")
// // 	stringify_csv_output := strings.TrimSpace(string(gpu_csv_string))
// // 	jsIndent, _ := json.MarshalIndent(&stringify_csv_output, "", "")
// // 	fmt.Printf("%s", string(jsIndent))
// // }

// /////////////////////////////////

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Name        string `json:"full_name"`
// 	Age         int    `json:"age,omitempty"`
// 	Active      bool   `json:"-"`
// 	lastLoginAt string
// }

// func main() {
// 	u, err := json.Marshal(User{Name: "Bob", Age: 0, Active: true, lastLoginAt: "today"})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(u)) // {"full_name":"Bob"}

// 	numberJson := "3"
// 	floatJson := "3.1412"
// 	stringJson := `"bird"`

// 	var n int
// 	var pi float64
// 	var str string

// 	json.Unmarshal([]byte(numberJson), &n)
// 	fmt.Println(n)
// 	// 3

// 	json.Unmarshal([]byte(floatJson), &pi)
// 	fmt.Println(pi)
// 	// 3.1412

// 	json.Unmarshal([]byte(stringJson), &str)
// 	fmt.Println(str)
// 	// bird

// }
