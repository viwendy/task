// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.10;

import { ERC721 } from "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract NftContract is ERC721 {

    using Strings for uint256;


    constructor() ERC721("NftContract", "NFT") {
        _mint(msg.sender, 1);
    }

    function mint(address to, uint256 tokenId) public {
        _mint(to, tokenId);
    }
}