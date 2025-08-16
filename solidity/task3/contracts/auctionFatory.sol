// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./auction.sol";


contract AuctionFactory {
    address[] public auctions;
    mapping(uint256 tokenId => Auction) public auctionMap;

    event AuctionCreated(address indexed auctionAddress, uint256 tokenId);

    function createAuction(address nftAddress, uint256 tokenId, uint256 startPrice,  uint256 duration) external returns (address){
        Auction newAuction = new Auction();
        newAuction.initialize();
        newAuction.createAuction(nftAddress, tokenId, startPrice, duration);
        auctions.push(address(newAuction));
        auctionMap[tokenId] = newAuction;
        emit AuctionCreated(address(newAuction), tokenId);
        return address(newAuction);
    }

    function getAuctions() external view returns(address[] memory){
        return auctions;
    }

    function getAuctions(uint256 tokenId) external view returns(address){
        require(tokenId < auctions.length, "tokenId out of range");
        return auctions[tokenId];
    }
}
