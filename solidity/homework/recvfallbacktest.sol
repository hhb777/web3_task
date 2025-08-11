// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract recvfallback{
    event recvLOG(string s);
    event fallbackLOG(string s);
    uint balance;
    receive() external payable{
        emit recvLOG("recieve test");
    }
    fallback() external payable{
        emit fallbackLOG("fallback test");
    }
    // function test(uint amount) public payable{
    //     balance += amount;
    // }
    // function getbalance() public view returns(uint){
    //     return balance;
    // }
}