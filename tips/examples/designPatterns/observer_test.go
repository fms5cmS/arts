package designPatterns

import (
	"fmt"
	"testing"
)

// 仅列出了核心的方法声明，实际还会有其他的方法
// 如 UnSubscribe 取消订阅
type Subject interface {
	Subscribe(observer Observer) // 添加订阅着
	Notify(msg string)           // 发布通知
}

type Observer interface {
	Read(msg string)
}

type Publisher struct {
	observers []Observer
}

func (p *Publisher) Subscribe(observer Observer) {
	p.observers = append(p.observers, observer)
}

func (p *Publisher) Notify(msg string) {
	for _, observer := range p.observers {
		observer.Read(msg)
	}
}

type Subscriber1 struct{}

func (s *Subscriber1) Read(msg string) {
	fmt.Println("Observer1 read: ", msg)
}

type Subscriber2 struct{}

func (s *Subscriber2) Read(msg string) {
	fmt.Println("Observer2 read: ", msg)
}

func TestObserverPattern(t *testing.T) {
	trafficSubjectPublisher := &Publisher{}
	trafficSubjectPublisher.Subscribe(&Subscriber1{})
	trafficSubjectPublisher.Subscribe(&Subscriber2{})
	trafficSubjectPublisher.Notify("publish traffic message, xxx.")
}
