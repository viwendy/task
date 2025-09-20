// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {
    // 二分查找函数
    // 返回目标值在数组中的索引，如果不存在返回 -1
    // [1, 3, 5, 7, 9] target = 5
    function binarySearch(uint[] memory arr, uint x) public pure returns (int) {
        int low = 0;
        int high = int(arr.length) - 1;

        while (low <= high) {
            int mid = (low + high) / 2;
            if (arr[uint(mid)] == x) {
                return int(mid); // 找到目标值，返回索引
            } else if (arr[uint(mid)] < x) {
                low = mid + 1; // 调整搜索范围到右半部分
            } else {
                high = mid - 1; // 调整搜索范围到左半部分
            }
        }
        return -1; // 未找到目标值，返回 -1
    }
}