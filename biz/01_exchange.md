
# K 线图

K 线图由许多烛台图案组成，创建烛台图案需要：开盘价、最高价、最底价、收盘价，该数据集通常被称为“OHLC 值”，开盘价和收盘价的间距称为“实体”，而实体与最高/低价的距离称为烛芯或影线，蜡烛图案高低点的间距称为烛台范围。

> 每个蜡烛图代表了该时间段内的市场走势，每个烛台显示开盘价和收盘价（烛台主体）以及最高价和最低价点（烛台上方和下方的长线，也成烛芯）。
> 
> 上升烛台（收盘价 > 开盘价）通常用绿色或黑色（实心）表示，下降烛台（收盘价 < 开盘价）通常为红色或空心（白色）。

# 订单薄

## Market Depth

订单薄列出各价格点的要价和出价订单数量，即市场深度(Market Depth)。 

## Maker&Taker

在交易是会有两种身份：Maker(挂单)，Taker(吃单)。其中 Maker 是流动性提供者，而 Taker 则是流动性消耗者。

当用户挂出一定数量和价格的委托单，但市场上没有与之匹配的订单，那么这个委托单就会一致挂在交易所的盘口上，提供市场深度，等待其他用户与之成交，相当于为整个市场提供了流动性，因此称为 Maker。

而当用户下单后，该委托单立即与排队的反向委托单成交，拿走市场深度，相当于消耗了市场的流动性，称为 Taker。

为了促进数字货币交易并保持市场价的稳定，交易平台更加鼓励 Maker 提供流动性，所以 Maker 的手续费会比 Taker 的手续费要低。

注意：用户最终是 Maker 还是 Taker 并不是根据下单时的状态判断的，主要是看成交时的状态。假设挂单时市场没有匹配的委托单，则为 Maker，如果在确认提交时瞬间出现了匹配的委托单且成交，则身份为 Taker。

所以，来自市价单的交易都是 Taker。

## MA

[移动平均线(Moving Average，MA)](https://academy.binance.com/zh/articles/moving-averages-explained) 代表过去一段时间里的平均成交价格。

> 股市的三类均线：短期（5日、10日），中期（20日，60日），长期（120日、250日）
> 
> 均线的交叉点
> 
> - 均线的金叉：短期均线由下往上穿过长期均线。代表买入信号
>
> - 均线的死叉：短期均线由上往下穿过长期均线，代表卖出信号
>
> 均线的两种极端（至少三根以上的均线组成的均线系统来构筑）
> 
> - 多头排列：短期均线由下往上依次穿过中期均线、长期均线。属于上涨
> 
> - 空头排列：短期均线由上往下依次穿过中期均线、长期均线。属于下跌

MA 通常被分为两大类：简单移动平均线(SMA)、指数移动平均线(EMA)。

- SMA

从设定的时间段获取数据，从而的出其资产的平均价格。SMA 与基本的均价差别在于：对于 SMA，一旦输入新数据集，以前的数据集就会忽略不计。

- EMA

类似 SMA，也是根据过去的价格波动提供技术分析，但计算更为复杂，因为 EMA 为最近的价格输入分配了更多的权重和价值。EMA 对突然的价格波动和逆转反应更敏感，所以通常特别受短期交易者的青睐。

当两个不同的 MA 在图表中交叉时，会创建交叉信号。短期 MA 超过长期 MA 时，发生看涨交叉（黄金交叉），表示上升趋势的开始；相反，当短期 MA 低于长期 MA 时出现看跌交叉（死亡交叉），这表示下跌趋势的开始。