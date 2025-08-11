// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract dselector{
    bytes4 storedSelector;
    //0x7b2929090000000000000000000000000000000000000000000000000000000000000004
    function square(uint x) public pure returns(uint , bytes4){
        bytes4 selector = bytes4(keccak256("square(uint256)"));
        return (x**2, selector);
    }
    //0xeee972060000000000000000000000000000000000000000000000000000000000000017
    function double(uint x) public pure returns(uint, bytes4){
        bytes4 selector = bytes4(keccak256("double(uint256)"));
        return (x*2, selector);
    }

    function executeFunction(bytes4 selector, uint x) public returns(uint z){
        (bool suc, bytes memory data) = address(this).call(abi.encodeWithSelector(selector,x));
        require(suc, "call failed");
        z = abi.decode(data, (uint));

    }
    function storeSelector(bytes4 selector) public{
        storedSelector = selector;
    }

    function executeStoreSelector(uint x) public returns(uint z){
        (bool suc, bytes memory data) = address(this).call(abi.encodeWithSelector(storedSelector, x));
        require(suc, "execute failed");
        z = abi.decode(data, (uint));
    }
}