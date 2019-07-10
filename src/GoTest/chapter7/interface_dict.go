package main

import "fmt"

type Dictionary struct {
	data map[interface{}]interface{} // 键值都为interface类型
}

func (d *Dictionary) Get(key interface{}) (value interface{}) {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// 遍历， 如果毁掉返回值为false 停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

//将map的键值都置为空接口，即清空数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// 初始化创建字典
func newDictionary() *Dictionary {
	d := &Dictionary{}

	d.Clear()

	return d
}

func main() {
	// 创建字典实例
	dict := newDictionary()

	//填充数据
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don`t Hungry", 24)

	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite : ", favorite)

	// 遍历
	dict.Visit(func(k, v interface{}) bool {
		if v.(int) > 40 {
			fmt.Println(k, "is expensive")
			return true
		}
		fmt.Println(k, "is cheap")
		return true
	})
}
