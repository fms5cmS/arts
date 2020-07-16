双指针的用法：

- l、r 两个指针从左往右相向移动，当两者相遇时结束。
    - 通常时间复杂度为 O(n)

示例：arrayRela/11_mostWater

- boundary、explore 两个指针分工合作
    - boundary 作为分界线
    - explore 则在分隔后的某一侧遍历，从而将满足某一条件的数据移动到另一侧

示例：arrayRela/283_moveZeros