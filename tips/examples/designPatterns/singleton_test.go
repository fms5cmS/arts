package designPatterns

import (
	"sync"
	"testing"
)

// 类型为非导出类型
type singleton struct{}

/** 饿汉式实现单例模式 */

var singletonInstance *singleton

// 保证程序初始化时就创建了实例
func init() {
	singletonInstance = &singleton{}
}

// 对外暴露的方法，用于获取实例，但每次仅返回全局变量所保存的实例
func GetSingletonInstance() *singleton {
	return singletonInstance
}

func TestGetInstance(t *testing.T) {
	if GetSingletonInstance() != GetSingletonInstance() {
		t.Error("not singleton!")
	}
}

func BenchmarkGetInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetSingletonInstance() != GetSingletonInstance() {
				b.Errorf("test fail")
			}
		}
	})
}

/** 懒汉式+双重检测 实现单例模式*/

var (
	lazySingleton *singleton
	once          sync.Once
)

// once.Do() 中本身做了双重检测
func GetLazySingletonInstance() *singleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &singleton{}
		})
	}
	return lazySingleton
}

func TestGetLazyInstance(t *testing.T) {
	if GetLazySingletonInstance() != GetLazySingletonInstance() {
		t.Error("not singleton!")
	}
}

func BenchmarkGetLazyInstanceParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingletonInstance() != GetLazySingletonInstance() {
				b.Errorf("test fail")
			}
		}
	})
}
