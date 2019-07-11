package main

import (
	"GoTest/chapter7/fsmto/fsm"
	"fmt"
)

// 闲置状态
type IdleState struct {
	fsm.StateInfo // 使用StateInfo实现基础接口
}

// 移动状态
type MoveState struct {
	fsm.StateInfo
}

// 跳跃状态
type JumpState struct {
	fsm.StateInfo
}

// 重新实现状态开始
func (i *IdleState) OnBegin()  {
	fmt.Println("IdleState begin")
}


// 重新实现状态结束
func (i *IdleState) OnEnd()  {
	fmt.Println("IdleState end")
}


func (m *MoveState) OnBegin()  {
	fmt.Println("MoveState begin")
}

// 允许移动状态相互转移
func (m *MoveState) EnableSanmeTransit() bool  {
	return true
}

func (m *JumpState) OnBegin()  {
	fmt.Println("JumpState begin")
}

// 跳跃状态不能转移到移动状态
func (m *JumpState) CanTransitTo(name string) bool  {
	return name != "MoveState"
}

// 转移状态的和输出日志
func transitAndReport(sm *fsm.StateManager, state string)  {
	if err := sm.Transit(state); err != nil{
		fmt.Printf("FAILED! %s--->%s, %s\n\n", sm.CurrState().Name(), state, err.Error())
	}
}

func main()  {
	sm := fsm.NewStateManager()

	//响应状态转移的通知回调
	sm.OnChange = func(from, to fsm.State) {
		fmt.Printf("%s---> %s\n\n", fsm.StateName(from), fsm.StateName(to))
	}

	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	//在不同状态间转换
	transitAndReport(sm, "IdleState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "MoveState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "JumpState")
	transitAndReport(sm, "IdleState")

}