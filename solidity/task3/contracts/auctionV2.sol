// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "./auction.sol";

contract AuctionV2 is Auction{
    function V2Print() public pure returns (string memory) {
        return "This is Auction V2";
    }
}