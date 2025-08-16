// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract Mytoken is ERC721URIStorage, Ownable {

    uint256 private tokenId;
    constructor(string memory name, string memory symbol) ERC721URIStorage() ERC721(name, symbol) Ownable(msg.sender){}
    function mintNFT(address to, string memory tokenURI) external onlyOwner{
        require(to != address(0), "to is zero address");
        tokenId++;
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}