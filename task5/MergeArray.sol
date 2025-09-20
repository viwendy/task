// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

//合并两个有序数组 (Merge Sorted Array)
//题目描述：将两个有序数组合并为一个有序数组。

contract MergeArray {

    function merge(uint[] memory a, uint[] memory b) public pure returns (uint[] memory) {
        uint[] memory c = new uint[](a.length + b.length);
        for (uint i = 0; i < a.length; i++) {
            c[i] = a[i];
        }
        for (uint j = 0; j < b.length; j++) {
            c[a.length + j] = b[j];
        }
        return c;
    }

}