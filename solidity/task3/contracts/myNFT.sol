// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStora?ge.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721Enumerable, Ownable{
    string private _tokenURI;

    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}
    //铸币
    function mintNFT(address to, uint256 tokenId) public onlyOwner {
        // tokenId++;
        _mint(to, tokenId);
        // _setTokenURI(tokenId, tokenURL);
    }
    //设置tokenURI
    // function setTokenURI(uint256 tokenId, string memory _tokenURI) public onlyOwner{
    //     _setTokenURI(tokenId, _tokenURI);
    // }
    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        return _tokenURI;
    }

    function setTokenURI(string memory newTokenURI) external onlyOwner {
        _tokenURI = newTokenURI;
    }
    //转移
    // function transferNFT(address from, address to, uint256 id) public onlyOwner {
    //     require(ownerOf(id) == from, "From address does not own this NFT");
    //     require(to != address(0), "Cannot transfer to the zero address");
    //     _transfer(from, to, id);
    // }

    
}