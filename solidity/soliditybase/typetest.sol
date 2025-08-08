// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract addressTest{

    function test() public payable {
        address addr = 0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2;
        payable(addr).transfer(9999999999999999999);
    }

    mapping(string => uint) private t;
    function test2() public view returns(bytes1){
        bytes1 s;
        return s;
    }
}