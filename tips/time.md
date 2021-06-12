Go 语言的标准库的 time 包在使用时可能会遇到时区问题。

Go 语言使用 Location 来表示地区相关的时区，一个 Location 可能表示多个时区。

而在标准库 time 中，提供了 Location 的两个时区：Local、UTC，分别代表了本地时区和零时区，且默认使用 UTC。

那么该如何设置时区呢？

可以通过 `time.LoadLocation(name string) (*Location, error)` 来获取特定时区的 Location 实例。