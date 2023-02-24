

模版、策略、职责链这三个设计模式是解决业务流程复杂多变问题的利器。

策略模式和模版模式经常配合使用，策略模式是让完成某个任务的具体方式可以互相切换，而模版模式则是针对一个流程的共性梳理出固定的执行步骤，具体步骤的执行方式则由子类实现。

创建型模式：

- [单例模式](./singleton.md)
- [工厂模式](./factory.md)
- [原型模式](./prototype.md)
- [建造者/生成器模式](./builder.md)，对比 Go 中 Option 模式的[代码](../examples/designPatterns/option_test.go)

结构型模式

- [代理模式](./proxy.md)
- [装饰器模式](./decorator.md)

行为模式：

- [观察者模式](./observer.md)
- [模版/模版方法模式](./template.md)
- [策略模式](./strategy.md)
- [职责链模式](./chainOfResponsibility.md)

References：

[设计模式](https://refactoringguru.cn/design-patterns)

额外：https://github.com/mohuishou/go-design-pattern
