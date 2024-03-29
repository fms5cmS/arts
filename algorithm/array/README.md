
Q1：什么是数组？

数组是一种**线性表**数据结构，它用一组**连续的内存空间**来存储一组具有**相同类型的数据**。

Q2：为什么数组的下标是从 0 开始的呢？

从数组的内存模型来看，下标更确切的叫法应该是"偏移(offset)"，如 `a[i]` 表示其相对于 `a[0]` 偏移了 i 个数据类型的位置。

Q3：数组的访问特性！

由于数组的底层是连续的内存空间，且保存的数据都是相同类型，所以数组就有了一个很重要的特性：**随机访问**！可以通过寻址公式找到数组中的某个元素。

![数组的寻址公式](../../images/array_address.png)

注意：数组适合查找，但查找的时间复杂度并不为 $O(1)$，即使是排序后的数组，使用二分查找的时间复杂度也有 $O(logn)$，所以应说：

**数组支持随机访问，根据下标随机访问的时间复杂度为 $O(1)$！**。

Q4：数组的相关操作有哪些？

假设数组长度为 n，

- 插入操作
  - 有序数组：为了保持数组的有序，向数组中的第 k 个位置插入元素时，需要将 k～n 之间的所有元素都后移一位。时间复杂度为 $O(n)$
  - 无序数组：向数组中的第 k 个位置插入元素时，可以直接将di k 个位置的元素放到数组末尾，然后将新元素插入即可。时间复杂度为 $O(1)$
- 删除操作
  - 删除第 k 个位置的元素，为了保证内存的连续性，需要将 k～n 之间的元素全部前移。时间复杂度为 $O(n)$
  - 删除数组末尾的元素，不需要移动任何元素。时间复杂度为 $O(1)$
  - 在某些特殊场景下，如果不要求数组中数据的连续性。每次删除操作仅记录那个位置的元素已经被删除，只有当数组没有更多空间存储数据时，再触发一次真正的删除操作，这样可以减少数据的搬移。这就是标记清除算法的思想。

Q5：使用数组时有哪些注意点？

1. 小心数组下标越界！
2. 创建数组时最好预先指定大小，减少后续的扩容和数据搬移的开销

# 题目

- [两数之和](1_twoSum_test.go) 返回数组的索引，因此无法使用双指针
- [盛水最多的容器](11_mostWater_test.go)，双指针
- [三数之和](15_3Sum_test.go)，双指针


