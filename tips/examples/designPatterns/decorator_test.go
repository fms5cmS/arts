package designPatterns

import (
	"fmt"
	"testing"
)

// 披萨接口
type IPizza interface {
	getPrice() int
}

// 蔬菜披萨
type VeggeMania struct{}

func (p *VeggeMania) getPrice() int {
	return 15
}

// 装饰1：披萨额外加上番茄
type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// 装饰 2：披萨额外加上奶酪
type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func TestDecorator(t *testing.T) {
	pizza := &VeggeMania{}
	pizzaWithChees := &CheeseTopping{pizza: pizza}
	pizzaWithCheeseAndTomato := &TomatoTopping{pizza: pizzaWithChees}
	fmt.Println("Price of veggeMania with tomato and cheese topping is ", pizzaWithCheeseAndTomato.getPrice())
}
