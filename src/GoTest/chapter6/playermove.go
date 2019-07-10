package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	x float32
	y float32
}

type Player struct {
	currPos   Vec2    //当前位置
	targetPos Vec2    //目标位置
	speed     float32 //速度
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.x - other.x,
		v.y - other.y,
	}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.x + other.x,
		v.y + other.y,
	}
}

func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{
		v.x * s,
		v.y * s,
	}
}

func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.x - other.x
	dy := v.y - other.y

	return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

// 返回当前矢量的标准矢量
func (v Vec2) Normalize() Vec2 {
	mag := v.x*v.x + v.y*v.y
	if mag > 0 {
		oneOverMag := 1 / float32(math.Sqrt(float64(mag)))
		return Vec2{v.x * oneOverMag, v.y * oneOverMag}
	}
	return Vec2{0, 0}
}

// 设置玩家移动的目标位置
func (p *Player) MoveTo(v Vec2) {
	p.targetPos = v
}

//获取当前位置
func (p *Player) Pos() Vec2 {
	return p.currPos
}

//判断是否到达目的地
func (p *Player) IsArrived() bool {
	// 通过计算当前玩家位置与目标位置的距离不超过移动的步长，判断已经到达目标点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

//更新玩家的位置
func (p *Player) Update() {
	if !p.IsArrived() {
		// 计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		// 添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		// 移动完成后更新当前位置
		p.currPos = newPos
	}
}

//创建玩家的方法，也可以直接初始化
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}

func main() {
	//初始化玩家，速度为0.5
	p := NewPlayer(0.5)
	// 让玩家移动到3，1点
	p.MoveTo(Vec2{3, 1})
	// 如果没有达到就一直循环
	for !p.IsArrived() {
		// 更新玩家位置
		p.Update()
		// 打印每次移动后玩家的位置
		fmt.Println(p.Pos())
	}
}
