# task
Web3 learn task

# task1
    136. 【appearOnlyOnceNum.go】只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
         找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
         例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
         回文数【palindromicNum.go】
    
    
    20. 有效的括号【isValid.go】-------这个题目没搞明白 （堆栈关系）
        给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
        有效字符串需满足：
        左括号必须用相同类型的右括号闭合。
        左括号必须以正确的顺序闭合。
        每个右括号都有一个对应的相同类型的左括号。

    14. 最长公共前缀【longestCommonPrefix.go】
        编写一个函数来查找字符串数组中的最长公共前缀。
        如果不存在公共前缀，返回空字符串 ""。

    66. 加一【plusOne.go】
        给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。

    26. 删除有序数组中的重复项[removeDuplicates.go]
        给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
        考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
        更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
        返回 k 。

    56. 合并区间【mergeArr.go】-------这个题目需要再次打磨
        以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
        请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

# task2
    【Voting.sol】
    ✅ 创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数

    【ReverseString.sol】
    ✅ 反转字符串 (Reverse String)
    题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

    【RomanToInt.sol】
    ✅ 用solidity实现整数转罗马数字

    【MergeArray.sol】
    ✅ 合并两个有序数组 (Merge Sorted Array)
    题目描述：将两个有序数组合并为一个有序数组。

    【BinarySearch.sol】
    二分查找 (Binary Search)
    题目描述：在一个有序数组中查找目标值。