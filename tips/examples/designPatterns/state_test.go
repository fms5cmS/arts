package designPatterns

import (
	"fmt"
	"testing"
	"time"
)

// 定义红绿灯的状态接口
type LightState interface {
	// 亮灯
	Light()
	// 用于转到新状态
	EnterState()
	// 可以转的下一个状态
	NextLight(*TrafficLight)
	// 检查车速
	CarPassingSpeed(*TrafficLight, int, string)
}

// 状态接口中通用方法的实现，用于嵌入到具体状态中，减少重复代码
type DefaultLightState struct {
	StateName string
}

func (state *DefaultLightState) CarPassingSpeed(road *TrafficLight, speed int, licensePlate string) {
	if speed > road.SpeedLimit {
		fmt.Printf("Car with license %s was speeding\n", licensePlate)
	}
}

func (state *DefaultLightState) EnterState() {
	fmt.Println("changed state to:", state.StateName)
}

func (tl *TrafficLight) TransitionState(newState LightState) {
	tl.State = newState
	tl.State.EnterState()
}

// 红灯状态
type redState struct {
	DefaultLightState
}

func NewRedState() *redState {
	state := &redState{}
	state.StateName = "RED"
	return state
}

func (state *redState) Light() {
	fmt.Println("红灯亮起，不可行驶")
}

func (state *redState) CarPassingSpeed(light *TrafficLight, speed int, licensePlate string) {
	// 红灯时不能行驶， 所以这里要重写覆盖 DefaultLightState 里定义的这个方法
	if speed > 0 {
		fmt.Printf("Car with license \"%s\" ran a red light!\n", licensePlate)
	}
}

func (state *redState) NextLight(light *TrafficLight) {
	light.TransitionState(NewGreenState())
}

// 绿灯状态
type greenState struct {
	DefaultLightState
}

func NewGreenState() *greenState {
	state := &greenState{}
	state.StateName = "GREEN"
	return state
}

func (state *greenState) Light() {
	fmt.Println("绿灯亮起，请行驶")
}

func (state *greenState) NextLight(light *TrafficLight) {
	light.TransitionState(NewAmberState())
}

// 黄灯状态
type amberState struct {
	DefaultLightState
}

func NewAmberState() *amberState {
	state := &amberState{}
	state.StateName = "AMBER"
	return state
}

func (state *amberState) Light() {
	fmt.Println("黄灯亮起，请注意")
}

func (state *amberState) NextLight(light *TrafficLight) {
	light.TransitionState(NewRedState())
}

// Context，提供调用状态行为的接口，内部维护了一个状态实例，并负责状态的切换
type TrafficLight struct {
	State      LightState
	SpeedLimit int
}

func NewSimpleTrafficLight(speedLimit int) *TrafficLight {
	return &TrafficLight{
		State:      NewRedState(),
		SpeedLimit: speedLimit,
	}
}

func TestState(t *testing.T) {
	trafficLight := NewSimpleTrafficLight(500)

	interval := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-interval.C:
			trafficLight.State.Light()
			trafficLight.State.CarPassingSpeed(trafficLight, 25, "CN1024")
			trafficLight.State.NextLight(trafficLight)
		default:
		}
	}
}
