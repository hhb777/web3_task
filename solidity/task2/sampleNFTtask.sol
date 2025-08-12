// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SampleNFTTest is ERC721URIStorage, Ownable{

    uint256 public tokenId;
    constructor(string memory _name, string memory _symbol) ERC721URIStorage() ERC721(_name, _symbol) Ownable(msg.sender) {}

    function mintNFT(address to, string memory tokenURI) external onlyOwner{
        require(to != address(0), "to is zero address");
        tokenId++;
        _mint(to, tokenId);
        _setTokenURI(tokenId, tokenURI);
    }
}