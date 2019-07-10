package main

// 使用匿名结构体 分离JSON数据

import (
	"encoding/json"
	"fmt"
)

// 定义手机屏幕
type Screen struct {
	Size float32
	ResX int
	ResY int
}

//定义电池
type Battery struct {
	Capacity int
}

//生产json
func Json() []byte {
	raw := &struct { // 匿名结构体
		Screen
		Battery
		HasTouchID bool // 序列化时添加字段：是否有指纹识别
	}{
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery: Battery{
			Capacity: 2910,
		},
		HasTouchID: true,
	}

	jsonDate, _ := json.Marshal(raw)
	return jsonDate
}

func main() {
	jsonDtta := Json()
	fmt.Println(string(jsonDtta))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}

	json.Unmarshal(jsonDtta, &screenAndTouch)

	fmt.Printf("%+v\n", screenAndTouch)

	battreyAndTouch := struct {
		Battery
		HasTouchID bool
	}{
		Battery: Battery{
			Capacity: 2000,
		},
	}

	json.Unmarshal(jsonDtta, battreyAndTouch)

	fmt.Printf("%+v\n", battreyAndTouch)

}
