
工厂模式包含有三种：简单工厂、工厂方法、抽象工厂。

[代码中公共类型定义](../examples/designPatterns/factoryBaseModel.go)

- 简单工厂

简单工厂[代码示例](../examples/designPatterns/simpleFactory_test.go)

业务类比： 椅子加工厂提供了沙发、椅子、板凳、电竞椅等各种产品的生产能力，用户联系加工厂提供要求后会生产对应的产品。

- 工厂方法

工厂方法中[代码示例](../examples/designPatterns/factoryMethod_test.go)

业务类比：由于电竞行业近年来兴起，电竞椅等人体工学椅的需求加大，椅子加工厂成立了分别生产生活用椅、工作用椅的两个子公司，用户依旧仅联系椅子加工厂，加工厂则根据要求将需求分给子公司，子公司生产产品。

- 抽象工厂

抽象工厂[代码示例](../examples/designPatterns/abstract_factory_test.go)

业务类比：家具公司不仅想要生产椅子，还想要生产桌子、电视柜等物品。

代码中不常使用，需要引入众多接口、实现类，较为复杂
