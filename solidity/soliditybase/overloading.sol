// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract overloading{
    // 0xb0c74430
    function saysomething() public pure returns(string memory){
        return ("Nothing");
    }
    //0xe3b34a9d000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000067a68616e67330000000000000000000000000000000000000000000000000000
     function saysomething(string memory _name) public pure returns(string memory){
        return (_name);
    }

    function f(uint8 _in) public pure returns(uint8 out){
        out = _in;
    }
    function f(uint256 _in) public pure returns(uint256 out){
        out = _in;
    }
}