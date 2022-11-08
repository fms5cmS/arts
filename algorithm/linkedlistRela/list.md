链表的内存空间不需要是连续的，通过指针将多个零散的内存块串联起来使用，但丧失了随机访问的能力，通常将内存块称为链表的节点(Node)。

Q1: 数组 VS 链表

内存中的形式：

 - 数组在内存中是连续的内存空间 -> 可以利用 CPU 缓存机制来预读数据，访问效率更高；支持随机访问，根据下标随机访问的时间复杂度为 O(1)
 - 链表在内存中不是连续存储的 -> 无法利用 CPU 缓存机制；不支持随机访问
 - 数组在内存中大小固定 -> 不支持动态扩容，声明过小可能会频繁触发扩容，适合读多写少的场景
 - 链表本身在内存中没有大小限制，通过指针连接 Node -> 支持动态扩容；每个 Node 需要额外空间存储指针，内存消耗更大；频繁进行插入、删除操作，容易造成内存碎片


Q2: 常见链表操作

删除节点：203 -> 19
链表反转：206 -> 24
两个链表的交点：160
链表是否有环：141
链表环的入口：142