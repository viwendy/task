// SPDX-License-Identifier: MIT
pragma solidity ^0.8.30;

//一个mapping来存储候选人的得票数
//一个vote函数，允许用户投票给某个候选人
//一个getVotes函数，返回某个候选人的得票数
//一个resetVotes函数，重置所有候选人的得票数

contract Voting {

    mapping(string=>uint) voteMap;
    string[] userArr;
    address owner;

    constructor(){
        owner = msg.sender;
    }

    function vote(string memory user) public {
        userArr.push(user);
        voteMap[user] = voteMap[user] + 1;
    }

    function getVotes(string memory user) public view returns (uint) {
        return voteMap[user];
    }

    function resetVotes() public {
        require(msg.sender == owner, "no permisson!");
        for(uint i = 0;i<userArr.length;i++){
            delete voteMap[userArr[i]];
        }
        delete(userArr);
    }
}