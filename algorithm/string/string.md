
KMP(Knuth Morris Pratt) 算法主要应用在字符串匹配上。其经典思想：**当出现字符串不匹配时，可以记录一部分之前已经匹配的文本内容，利用这些信息避免从头再去做匹配**。

所以，如何记录已经匹配的文本内容，就是 KMP 的重点！也是前缀表（next 数组）的任务。

前缀表是用来回退的，它记录了模式串与主串（文本串）不匹配的时候，模式串应该从哪里个位置开始重新匹配。

见 28, 459


