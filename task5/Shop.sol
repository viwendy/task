// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

//部署合约，向本合约充值100wei

contract Shop {
    uint256 public price = 100;
    bool public isPaid;
    function buy() public payable {
        require(msg.value >= price, "Not enough Ether provided.");
        isPaid = true;
    }
    function withdraw() public {
        require(isPaid, "Item is not paid.");
        (bool sent, ) = msg.sender.call{value: address(this).balance}("");
        require(sent, "Failed to send Ether");
    }
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
    function getPrice() public view returns (uint256) {
        return price;
    }
}
