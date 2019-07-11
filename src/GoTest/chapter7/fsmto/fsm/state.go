package fsm

import "reflect"

type State interface {
	// 获取状态名称
	Name() string
	// 该状态是否允许同状态转移
	EnableSanmeTransit()  bool
	// 响应状态开始时
	OnBegin()
	// 响应状态结束时
	OnEnd()
	// 判断能否转移到某个状态
	CanTransitTo(name string) bool
}

func StateName(s State) string  {
	if s == nil{
		return "None"
	}
	//return s.Name()
	// 使用反射获取状态名称
	return reflect.TypeOf(s).Elem().Name()
}
