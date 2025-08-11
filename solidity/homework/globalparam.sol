// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ABItest{
    function Getblockinfo() public view returns(uint256 blocknum, uint time, address blockaddr, uint gasprice){
        return (block.number, block.timestamp, block.coinbase, tx.gasprice);
    }
}