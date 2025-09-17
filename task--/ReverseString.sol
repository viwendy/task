//SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

//反转字符串 (Reverse String)
//题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"

contract ReverseString {

    function reverse(string memory str) public pure returns (string memory) {
        bytes memory strBytes = bytes(str);
        uint len = strBytes.length;
        for (uint i = 0; i < len / 2; i++) {
            (strBytes[i], strBytes[len - 1 - i]) = (strBytes[len - 1 - i], strBytes[i]);
        }
        return string(strBytes);
    }
}