package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//func String2Json() (string, string) {

// 	cpu_manager_state_str := `{
//     "policyName": "static",
//     "defaultCpuSet": "0,2-7,9-15",
//     "entries": {
//       "3a9fea56-8de4-4e57-b59e-70dbbba8a538": {
//         "busybox-2": "1"
//       },
//       "a08e53a0-8ae1-456b-aca9-18623fe3f7a0": {
//         "busybox-2": "8"
//       }
//     },
//     "checksum": 2112539122
//   }
//   `

//Go出现警告struct doesn‘t have any exported fields, nor custom marshaling
//Go语言要求所有结构体成员变量的首字母需要大写,如果首字母小写的话，
//则该字段无法被外部包访问和解析，比如，json解析。
//所以修改成员变量首字母为大写即可
//ID := "7423c8cd-9d8f-44db-a7f0-4d5f0827744b"
// type PODPinCPU struct {
// 	CPUUsedbyPod string `json:"busybox-2"`
// }

// // type PODUUID struct {
// // 	UUID PODPinCPU `json:"."`
// // }
// type PODUUID struct {
// 	UUID PODPinCPU `json:"7423c8cd-9d8f-44db-a7f0-4d5f0827744b"`
// }

// type CPUManagerState struct {
// 	PolicyName    string
// 	DefaultCpuSet string
// 	Entries       PODUUID
// 	Checksum      int
// }

// var cpuManagerState CPUManagerState
// json.Unmarshal([]byte(cpu_manager_state_str), &cpuManagerState)
// fmt.Println(cpuManagerState.DefaultCpuSet, cpuManagerState.PolicyName, cpuManagerState.Checksum, cpuManagerState.Entries.UUID.CPUUsedbyPod)

// var cpuManagerStateInfo map[string]interface{}
// json.Unmarshal([]byte(cpu_manager_state_str), &cpuManagerStateInfo)

// fmt.Println(cpuManagerStateInfo["defaultCpuSet"])
// Entries := fmt.Sprint(cpuManagerStateInfo["entries"])

// PODMapCPUs := strings.Split(Entries, " ")
// Len := len(PODMapCPUs)
// var PODCUPs string
// var B string
// for i := 0; i < Len; i++ {
// 	if strings.Contains(PODMapCPUs[i], "busybox-2") {
// 		IDs := PODMapCPUs[i]
// 		fmt.Println(PODMapCPUs[i])
// 		A := strings.Split(IDs, ":")
// 		B = strings.Trim(A[2], "]")
// 		//fmt.Printf("%v,%T\n", IDs, IDs)
// 		//fmt.Println(A[2])
// 		fmt.Println(B)
// 	}
// 	PODCUPs += B + " "
// }
// //println(PODCUPs)
// return "Abc", PODCUPs

// var obj map[string]interface{}
// err := json.Unmarshal([]byte(cpu_manager_state_str), &obj)
// if err != nil {
// 	fmt.Println("error:", err)
// 	return
// }
// for k, v := range obj {
// 	println()
// 	println(k, &v)
// }

// jsonString, err := json.Marshal(obj)
// if err != nil {
// 	fmt.Println("error:", err)
// 	return
// }
// fmt.Println(string(jsonString))

// var v interface{}
// json.Unmarshal(jsonString, &v)
// data := v.(map[string]interface{})

// for k, v := range data {
// 	switch v := v.(type) {
// 	case string:
// 		fmt.Println(k, v, "(string)")
// 	case float64:
// 		fmt.Println(k, v, "(float64)")
// 	case []interface{}:
// 		fmt.Println(k, "(array):")
// 		for i, u := range v {
// 			fmt.Println("    ", i, u)
// 		}
// 	default:
// 		fmt.Println(k, v, "(unknown)")
// 	}
// }

//}

// Define Struct
type AnimalData struct {
	Bird int
	Cat  string
}

type Box struct {
	Width  int
	Height int
	Color  string
	Open   bool
}

type Language struct {
	Id   int
	Name string
}

type Bird struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

type Dimensions struct {
	Height int
	Width  int
}

type containerMemoryRSSType struct {
	Data struct {
		Result []struct {
			Metric struct {
				MetricName string `json:"__name__"`
				Container  string `json:"container"`
				Endpoint   string `json:"endpoint"`
				ID         string `json:"id"`
				Image      string `json:"image"`
				Instance   string `json:"instance"`
				Job        string `json:"job"`
				MetricPath string `json:"metrics_path"`
				Name       string `json:"name"`
				Namespace  string `json:"namespace"`
				Node       string `json:"node"`
				Pod        string `json:"pod"`
				Service    string `json:"service"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
		ResultType string `json:"resultType"`
	} `json:"data"`
	Status string `json:"status"`
}

func (memRSS *containerMemoryRSSType) String2Json2() {
	metricResults := `{
		"status": "success",
		"data": {
		  "resultType": "vector",
		  "result": [
			{
			  "metric": {
				"__name__": "container_memory_rss",
				"container": "kube-apiserver",
				"endpoint": "https-metrics",
				"id": "/kubepods.slice/kubepods-burstable.slice/kubepods-burstable-pode4e6004e827fdf45941cf777618f60ce.slice/crio-192f5a90b165665d969bb0137b8573f2214c8f75b728ecc58dbbc7fccd64af28.scope",
				"image": "quay.io/openshift-release-dev/ocp-v4.0-art-dev@sha256:4ae6c4725722b7f0040d58c455a0a320e86cbaec5a9899b392750dcdadd7d5e7",
				"instance": "10.0.29.97:10250",
				"job": "kubelet",
				"metrics_path": "/metrics/cadvisor",
				"name": "k8s_kube-apiserver_kube-apiserver-ip-10-0-29-97.us-east-2.compute.internal_openshift-kube-apiserver_e4e6004e827fdf45941cf777618f60ce_0",
				"namespace": "openshift-kube-apiserver",
				"node": "ip-10-0-29-97.us-east-2.compute.internal",
				"pod": "kube-apiserver-ip-10-0-29-97.us-east-2.compute.internal",
				"service": "kubelet"
			  },
			  "value": [
				1700491106.539,
				"1608310784"
			  ]
			}
		  ]
		}
	  }`

	json.Unmarshal([]byte(metricResults), &memRSS)
	fmt.Println("-------------------------------")
	fmt.Println(memRSS.Status, memRSS.Data.Result[0].Metric.MetricName, memRSS.Data.Result[0].Value[1])

	// var perfMemMetricsInfo map[string]interface{}
	// json.Unmarshal([]byte(metricResults), &perfMemMetricsInfo)

	// fmt.Println(perfMemMetricsInfo["status"])
	// fmt.Println(perfMemMetricsInfo["data"])

	// a := "1374389534720000"
	// b, err := strconv.Atoi(a)
	// fmt.Println("---------------------")
	// fmt.Println(b)
	// if err == nil {
	// 	fmt.Println(b)
	// }

}

func main() {
	// //	str := remove()
	// //fmt.Println(str)
	// //Input data.
	// text := "{\"Bird\":10,\"Cat\":\"Fuzzy\"}"
	// bytes := []byte(text)

	// //Get struct from string.
	// var animal AnimalData
	// json.Unmarshal(bytes, &animal)

	// //Print result
	// fmt.Printf("BIRD = %v ,CAT=%v", animal.Bird, animal.Cat)

	// //Create an instance of the Box struct
	// box := Box{
	// 	Width:  10,
	// 	Height: 20,
	// 	Color:  "Blue",
	// 	Open:   false,
	// }
	// //Create JSON from the  instance data
	// b, _ := json.Marshal(box)
	// //Convert bytes to string
	// fmt.Println(string(b))

	// // String contains two JSON rows.
	// text = "[{\"Id\": 100, \"Name\": \"Go\"}, {\"Id\": 200, \"Name\": \"Java\"}]"

	// //Get byte slice from string.
	// bytes = []byte(text)

	// //Umarshal string into structs.
	// var languages []Language
	// json.Unmarshal(bytes, &languages)

	// //Loop over structs and display them
	// for l := range languages {
	// 	fmt.Printf("Id = %v, Name = %v", languages[l].Id, languages[l].Name)
	// 	fmt.Println()

	// 	birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	// 	var birds []Bird
	// 	json.Unmarshal([]byte(birdJson), &birds)
	// 	fmt.Printf("Birds : %+v", birds)
	// 	//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]

	// 	type csvbody struct {
	// 		body string
	// 	}
	// }

	birdJson := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds Bird
	json.Unmarshal([]byte(birdJson), &birds)
	fmt.Printf("%s,%d", birds.Species, birds.Dimensions.Height)
	// A, B := String2Json2()
	// fmt.Println(A)
	// fmt.Println(B)
	containerIDStr := "cri-o://c861a00a8e5beb2b124faf22adb42eb952bf89f5c7932253ffede56cc65abff6"
	containerIDArr := strings.Split(containerIDStr, "/")
	fmt.Println(containerIDArr[2])

	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)
	// 3

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)
	// 3.1412

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)
	// bird
	var memRSS containerMemoryRSSType
	memRSS.String2Json2()
}
