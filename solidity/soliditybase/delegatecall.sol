// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract C{
    uint public num;
    address public sender;

    function setVars(uint _num) public payable{
        num = _num;
        sender = msg.sender;
    }
}


contract B{
    uint public num;
    address public sender;
    event LOG(bytes data);
    function CallsetVars(address addr, uint _num) public {
        (bool suc, bytes memory data) = addr.call(abi.encodeWithSignature("setVars(uint256)", _num));
        if (!suc){
            revert();
        }
        emit LOG(data);
    }

    function DelegateCallsetVars(address addr,uint _num) public {
        (bool suc, bytes memory data) = addr.delegatecall(abi.encodeWithSignature("setVars(uint256)", _num));
        if (!suc){
            revert();
        }
        emit LOG(data);
    }
}