// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract mappingTest{
    mapping(address => uint) balances;
    mapping(address => string[]) userhisop;
    // mapping(address => uint) userscore;
    function des(uint amount) public payable{
        balances[msg.sender]+=amount;
    }
    function getbalance(address addr) public view returns(uint balance){
        return balances[addr];
    }
    function addhisop(address addr,string memory hisop) public{
        userhisop[addr].push(hisop);
    }
    function gethisop(address addr) public view returns(string[] memory){
        return userhisop[addr];
    }
}