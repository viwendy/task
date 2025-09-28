// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

contract BeggingContract {
    mapping (address => uint) public donations;
    address public owner;
    constructor() {
        owner = msg.sender;
    }
    function donate() public payable {
        require(msg.value > 0, "Donation must be greater than 0");
        donations[msg.sender] += msg.value;
    }
    function withdraw() public {
        require(msg.sender == owner, "Only owner can withdraw");
        payable(msg.sender).transfer(address(this).balance);
    }
    function getDonation(address _address) public view returns (uint) {
        return donations[_address];
    }
    receive() external payable {
        donate();
    }
    fallback() external payable {
        donate();
    }

}
