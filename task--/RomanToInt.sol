// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.30;

//用solidity实现整数转罗马数字

contract RomanToInt{

    mapping(bytes1 => int) private maps;

    constructor(){
        maps['I']=1;
        maps["V"]=5;
        maps["X"]=10;
        maps["L"]=50;
        maps ["C"]=100;
        maps["D"]=500;
        maps["M"]=1000;
    }

    function romanToInt(string memory s) public view returns(int,uint){
        bytes memory strs = bytes(s);
        int res = 0;
        for(uint i=0;i<strs.length;i++){
            int value = maps[strs[i]];
            if(i<strs.length-1&&value < maps[strs[i+1]]){
                res -= value;
            }else{
                res += value;
            }
        }
        return (res, strs.length);
    }
}