package fsm

import "errors"

// 状态没有找到的错误
var ErrStateNotFound = errors.New("state not found")

// 禁止在同状态间转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")

// 不能转移到指定状态
var ErrCannotTransitToState = errors.New("cannot transit to state")


// 状态管理器
type StateManager struct {
	// 已经添加的状态
	stateByNaame map[string]State
	// 状态改变时的回调
	OnChange func(from, to State)
	// 当前状态
	curr State
}

// 添加一个状态到管理器
func (sm *StateManager) Add (s State)  {
	// 获取状态名称
	name := StateName(s)
	// 将s转换为能设置名字的接口，然后调用该接口
	s.(interface{
		SetName(name string)
	}).SetName(name)

	// 根据状态名称获取已经添加的状态，检查该状态是否存在
	if sm.Get(name) !=nil{
		panic("duplicate state :" + name)
	}

	// 根据名称保存到map
	sm.stateByNaame[name] =s

}

//根据名字获取指定状态
func (sm *StateManager) Get(name string) State {
	if v,ok := sm.stateByNaame[name]; ok{
		return v
	}
	return nil
}

//获取当前状态
func (sm *StateManager) CurrState() State {
	return sm.curr
}

//当前状态能否转移到目标状态
func (sm *StateManager) CanCurrTransitTo(name string) bool {
	if sm.curr == nil{
		return true
	}

	// 相同的状态不用转换
	if sm.curr.Name() == name && !sm.curr.EnableSanmeTransit(){
		return false
	}
	// 使用当前状态， 检查能否转移到指定名字的状态
	return sm.curr.CanTransitTo(name)
}

// 转移到指定状态
func (sm *StateManager) Transit(name string) error  {
	// 获取目标状态
	next := sm.Get(name)

	// 目标状态不存在
	if next == nil{
		return ErrStateNotFound
	}

	// 记录转移前的状态
	pre := sm.curr

	//当前有状态
	if sm.curr != nil{
		// 相同的状态不用转换
		if sm.curr.Name() == name && !sm.curr.EnableSanmeTransit(){
			return ErrForbidSameStateTransit
		}

		// 不能转移到目标状态
		if !sm.curr.CanTransitTo(name){
			return ErrCannotTransitToState
		}
		// 结束当前状态
		sm.curr.OnEnd()
	}

	// 将当前的状态切换为目标状态
	sm.curr = next

	// 调用新状态的开始
	sm.curr.OnBegin()

	// 通知回调
	if sm.OnChange != nil{
		sm.OnChange(pre, sm.curr)
	}

	return nil

}


//初始化状态管理器
func NewStateManager() *StateManager {
	return &StateManager{
		stateByNaame:make(map[string]State),
	}
}