package main

import "fmt"
func main() {
	func2()
}
/**
map 创建方式：
1. 直接创建
2. 赋值
 */
func func1()  {
	var map1 map[string]string
	map1 = make(map[string]string, 10)
	// 赋值
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	map1["key3"] = "value3"
	map1["key4"] = "value4"

	for key, value := range map1 {
		fmt.Printf("key = %s, value = %s \n", key, value)
	}

	map2 := map[int]string{1:"str1", 2:"str2", 3:"str3",4:"str4"} //初始化
	for key, value := range map2 {
		fmt.Printf("key = %v, value = %s \n", key, value)
	}
	fmt.Printf("map1的长度是 %d\n map2的长度是 %d\n", len(map1), len(map2) )

}

func func2() {
	map2 := map[int]string{1:"str1", 2:"str2", 3:"str3",4:"str4"} //初始化
	for key, value := range map2 {
		fmt.Printf("key = %v, value = %s \n", key, value)
	}

	if _, ok := map2[2]; ok { //  _, ok := map2[2] 判断 key = 2 在map2 中是否存在，如果存在ok为true，否则为false
		fmt.Println(map2[2])
	}

	if _,ok := map2[5];!ok {
		fmt.Println("map2 中没有key 为 2的键值对")
	}

	// delete() 删除一个key对应的map元素, 如果key没有在map 中，也不会报错
	fmt.Println("第一次删除map2.。。")
	delete(map2, 3) 
	for key, value := range map2 {
		fmt.Printf("key = %v, value = %s \n", key, value)
	}
	fmt.Println("第二次删除map2.。。")
	delete(map2, 3) 
	for key, value := range map2 {
		fmt.Printf("key = %v, value = %s \n", key, value)
	}

	/*
	结果：
	key = 3, value = str3
	key = 4, value = str4
	key = 1, value = str1
	key = 2, value = str2
	str2
	第一次删除map2.。。
	key = 1, value = str1
	key = 2, value = str2
	key = 4, value = str4
	第二次删除map2.。。
	key = 1, value = str1
	key = 2, value = str2
	key = 4, value = str4
	*/
}