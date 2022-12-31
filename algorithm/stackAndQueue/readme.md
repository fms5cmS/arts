
Stack 是一种“操作受限”的线性表，只允许在一端(栈顶)插入、删除数据，且具有 LIFO(Last In First Out，后进先出) 特性。

Queue 也是一种“操作受限”的线性表，其特点：FIFO（First In First Out）即在队尾(tail)插入数据项，在队首(front)移除数据项。

Golang 中并未实现 Stack 和 Queue，可以通过 slice 来进行模拟，而在平时使用中通常更推荐使用双端队列，而非单纯的 Stack 或 Queue。

优先队列 Priority Queue，一种抽象数据结构，其插入操作为 O(1)，取出操作为 O(logn)，需要按元素的优先级取出，底层实现多样，可以用 Heap、BST 等实现

使用栈实现队列、使用队列实现栈。

单调栈的特殊使用：84

单调队列的使用：239