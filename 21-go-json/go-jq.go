package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func remove() string {
	jsonstring := `[
    {
      "apiVersion": "nvidia.com/v1",
      "kind": "ClusterPolicy",
      "metadata": {
        "name": "gpu-cluster-policy"
      },
      "spec": {
        "dcgmExporter": {
          "config": {
            "name": ""
          }
        },
        "dcgm": {
          "enabled": true
        },
        "daemonsets": {
        },
        "devicePlugin": {
        },
        "driver": {
          "enabled": true,
          "use_ocp_driver_toolkit": true,
          "repoConfig": {
            "configMapName": ""
          },
          "certConfig": {
            "name": ""
          },
          "licensingConfig": {
            "nlsEnabled": false,
            "configMapName": ""
          },
          "virtualTopology": {
            "config": ""
          },
          "kernelModuleConfig": {
            "name": ""
          }
        },
        "gfd": {
        },
        "migManager": {
          "enabled": true
        },
        "nodeStatusExporter": {
          "enabled": true
        },
        "operator": {
          "defaultRuntime": "crio",
          "deployGFD": true,
          "initContainer": {
          }
        },
        "mig": {
          "strategy": "single"
        },
        "toolkit": {
          "enabled": true
        },
        "validator": {
          "plugin": {
            "env": [
              {
                "name": "WITH_WORKLOAD",
                "value": "true"
              }
            ]
          }
        }
      }
    }
  ]`
	// fmt.Printf("%T", jsonstring)
	// dst := &bytes.Buffer{}
	// if err := json.Indent(dst, []byte(jsonstring), " ", " "); err != nil {
	// 	panic(err)
	// }
	// fmt.Println(dst.String())
	// 序列化
	// b, err := json.Marshal(jsonstring)
	// if err != nil {
	// 	fmt.Printf("marshal failed,err:%v", err)
	// }

	//fmt.Println(string(b))
	//反序列化
	tmp1 := strings.TrimLeft(jsonstring, "[")
	output := strings.TrimRight(tmp1, "]")
	stringify_jq_output := strings.TrimSpace(string(output))
	return stringify_jq_output
}

//Define Struct
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
}

func main() {
	//	str := remove()
	//fmt.Println(str)
	//Input data.
	text := "{\"Bird\":10,\"Cat\":\"Fuzzy\"}"
	bytes := []byte(text)

	//Get struct from string.
	var animal AnimalData
	json.Unmarshal(bytes, &animal)

	//Print result
	fmt.Printf("BIRD = %v ,CAT=%v", animal.Bird, animal.Cat)

	//Create an instance of the Box struct
	box := Box{
		Width:  10,
		Height: 20,
		Color:  "Blue",
		Open:   false,
	}
	//Create JSON from the  instance data
	b, _ := json.Marshal(box)
	//Convert bytes to string
	fmt.Println(string(b))

	// String contains two JSON rows.
	text = "[{\"Id\": 100, \"Name\": \"Go\"}, {\"Id\": 200, \"Name\": \"Java\"}]"

	//Get byte slice from string.
	bytes = []byte(text)

	//Umarshal string into structs.
	var languages []Language
	json.Unmarshal(bytes, &languages)

	//Loop over structs and display them
	for l := range languages {
		fmt.Printf("Id = %v, Name = %v", languages[l].Id, languages[l].Name)
		fmt.Println()

		birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
		var birds []Bird
		json.Unmarshal([]byte(birdJson), &birds)
		fmt.Printf("Birds : %+v", birds)
		//Birds : [{Species:pigeon Description:} {Species:eagle Description:bird of prey}]

		type csvbody struct {
			body string
		}
		jsonstring := `[
    {
      "apiVersion": "nvidia.com/v1",
      "kind": "ClusterPolicy",
      "metadata": {
        "name": "gpu-cluster-policy"
      },
      "spec": {
        "dcgmExporter": {
          "config": {
            "name": ""
          }
        },
        "dcgm": {
          "enabled": true
        },
        "daemonsets": {
        },
        "devicePlugin": {
        },
        "driver": {
          "enabled": true,
          "use_ocp_driver_toolkit": true,
          "repoConfig": {
            "configMapName": ""
          },
          "certConfig": {
            "name": ""
          },
          "licensingConfig": {
            "nlsEnabled": false,
            "configMapName": ""
          },
          "virtualTopology": {
            "config": ""
          },
          "kernelModuleConfig": {
            "name": ""
          }
        },
        "gfd": {
        },
        "migManager": {
          "enabled": true
        },
        "nodeStatusExporter": {
          "enabled": true
        },
        "operator": {
          "defaultRuntime": "crio",
          "deployGFD": true,
          "initContainer": {
          }
        },
        "mig": {
          "strategy": "single"
        },
        "toolkit": {
          "enabled": true
        },
        "validator": {
          "plugin": {
            "env": [
              {
                "name": "WITH_WORKLOAD",
                "value": "true"
              }
            ]
          }
        }
      }
    }
  ]`
		// var result map[string]interface{}
		// bytes = []byte(jsonstring)
		// json.Unmarshal(bytes, &result)
		// fmt.Println(result)
		gpu_csv_tmp := strings.TrimLeft(jsonstring, "[")
		gpu_csv_string := strings.TrimRight(gpu_csv_tmp, "]")
		stringify_csv_output := strings.TrimSpace(string(gpu_csv_string))
		jsIndent, _ := json.MarshalIndent(&stringify_csv_output, "", "")
		fmt.Printf("%s", string(jsIndent))

	}

}
