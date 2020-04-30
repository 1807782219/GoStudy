package main

import (
	"fmt"
	"encoding/json"
)

/**
Object 转换成 JSON：
json.Marshal(data)
1、布尔值转化后还是布尔类型
2、浮点数和整型都会转换为JSON里面的常规数字
3、字符串将以UTF-8 编码转化为Unicode字符集的字符串，特殊字符也会转
4、数组和切片会转化为Json里面的数组，但[]byte 类型的值将会被转化为 Base64编码后的字符串，slice 类型的零值会转化为null
5、结构体会转化为JSON对象，并且只有结构体里面以大写字母开头的可被导出的字段才会被转换输出，而这些可导出的字段会作为JSON对象的字符串索引
6、转换一个map类型的数据结构时，该数据的类型必须是 map[string] T

JSON 转 Object ：
json.Unmarshal([]byte(data), obj)： data 》 json 字符串 obj 》 要接收的对象
 */
type Student struct {
	//tag JSON中头字母变小写
	Name string `json:"name"`
	Age int `json:"age"`
	// tag 此字段bool类型转换为字符串输出
	Sex bool `json:",string"`
	Address string `json:"address"`
	// tag 此字段不会显示输出
	// Likes []string `json:"-"`
	Likes []string `json:"likes"`
}

func main() {

	testNotKnowTypeJsonToObject()
}

// 一个对象转成数组
func testObjectToJson()  {
	student := Student{"lisi",12,true,"China", []string{"football","TV"}}
	str := ObjectToJson(student)
	if str == "" {
		fmt.Println("json 转换失败 ")
	} else {
		fmt.Println(str)
		// 结果： {"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]}
	}
}

//多个对象转为数组
func testObjectsToJson()  {
	students := make([]*Student, 2)
	students[0] = &Student{"lisi",12,true,"China", []string{"football","TV"}}
	students[1] = &Student{"zhangsan",22,true,"England", []string{"football","Dress"}}

	str := ObjectToJson(students)
	if str == "" {
		fmt.Println("json 转换失败 ")
	} else {
		fmt.Println(str)
		// 结果： [{"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]},{"name":"zhangsan","age":22,"Sex":"true","address":"England","likes":["football","Dress"]}]
	}
}

// 一个map 转json
func testMapToJson() {
	data := make(map[string]interface{}, 10)
	data["jsonData"] =  Student{"lisi",12,true,"China", []string{"football","TV"}}
	students := make([]*Student, 2)
	students[0] = &Student{"lisi",12,true,"China", []string{"football","TV"}}
	students[1] = &Student{"zhangsan",22,true,"England", []string{"football","Dress"}}
	data["jsonDatas"] = students

	str := ObjectToJson(data)
	if str == "" {
		fmt.Println("json 转换失败 ")
	} else {
		fmt.Println(str)
		// 结果：{"jsonData":{"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]},"jsonDatas":[{"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]},{"name":"zhangsan","age":22,"Sex":"true","address":"England","likes":["football","Dress"]}]}
	}

}

func ObjectToJson(data interface{}) string {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json 转换错误 》 ", err)
		return ""
	} else {
		return string(bytes)
	}
}
// json 转单个对象
func testJsonToObject()  {
	str := `{"name":"lisi","age":12,"Sex":"true","address":"China", "likes":["football","TV"]}`
	student := Student{}
	err := JsonToObject(str, &student)
	if err != nil {
		fmt.Println("转换失败", err)
	}

	fmt.Println(student)
	// 结果 ：{lisi 12 true China [football TV]}
}

// json 转多个对象
func testJsonToObjects()  {
	str := `[{"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]},{"name":"zhangsan","age":22,"Sex":"true","address":"England","likes":["football","Dress"]}]`
	students := []Student{}
	err := JsonToObject(str, &students)
	if err != nil {
		fmt.Println("转换失败", err)
	}

	fmt.Println(students)
	// 结果 ：[{lisi 12 true China [football TV]} {zhangsan 22 true England [football Dress]}]
}

// json 转map
func testJsonToMap()  {
	str := `{"jsonDatas":[{"name":"lisi","age":12,"Sex":"true","address":"China","likes":["football","TV"]},{"name":"zhangsan","age":22,"Sex":"true","address":"England","likes":["football","Dress"]}]}`
	maps := map[string][]Student{}

	err := JsonToObject(str, &maps)
	if err != nil {
		fmt.Println("转换失败", err)
	}

	// fmt.Println(maps) 结果：map[jsonDatas:[{lisi 12 true China [football TV]} {zhangsan 22 true England [football Dress]}]]

	for key, values := range maps {
		fmt.Printf("key = %s 对应值为： \n",key)
		for _, value := range values {
			fmt.Println(value)
		}
	}
	/**
	结果：
	key = jsonDatas 对应值为：
	{lisi 12 true China [football TV]}
	{zhangsan 22 true England [football Dress]}
	 */
}

func JsonToObject(data string, obj interface{})  error {
	return json.Unmarshal([]byte(data), obj)
}


/**
解码未知结构的JSON 数据
Go语言中的空接口可以定义任意类型，是通用类型，如果要解码一段未知结构的JSOn，只需要将这段JSON数据解码输出到一个空接口即可。
在解码JSON数据的过程中，JSON数据里面的元素类型将做如下转换：
1、 JSON 中的布尔值经会转换为bool类型
2、 数值将会转换为Go中的float64 类型
3、 字符串转换后还是string 类型
4、 JSON数组会转换为[]interface{}
5、 JSON 对象会转换为map[string]interface{} 类型
6、 null值会转换为nil

 */

func testNotKnowTypeJsonToObject()  {
	data := []byte(`{
		 "Title": "Go语言编程",
		 "Authors": ["XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan",
 		 "XuDaoli"],
 		 "Publisher": "ituring.com.cn",
 		 "IsPublished": true,
		 "Price": 9.99,
		 "Sales": 1000000
		}`)
	var obj interface{}
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("转换错误",err)
		return
	}
	gobook, ok := obj.(map[string]interface{})
	if ok {
		for k,v := range gobook {
			switch rv := v.(type) {
			case string:
					fmt.Println(k, "is string", rv)
			case int:
					fmt.Println(k, "is int", rv)
			case bool:
					fmt.Println(k, "is bool", rv)
			case []interface{}:
					fmt.Println(k, "is an array", rv)
					for i, iv:=range rv {
						fmt.Println(i,iv)
					}
			default:
					fmt.Println(k, "is another type not handle yet")

			}
		}
	}
}