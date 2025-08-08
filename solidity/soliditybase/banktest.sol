// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Bank{

    event Calllog(bytes input, bytes output);
    receive() external payable { }
    function withdrawWithTransfer() external{
        payable(msg.sender).transfer(1 ether);
    }
    function withdrawWithSend() external{
        bool suc = payable(msg.sender).send(1 ether);
        require(suc, "send failed");
    }
    function withdrawWithCall(bytes memory input) external{
        (bool suc,bytes memory data) = payable(msg.sender).call{value: 1 ether}(input);
        require(suc, "call failed");
        emit Calllog(input, data);
    }

}

contract BankUser{
    Bank bank;
    constructor(address payable _bank){
        bank = Bank(_bank);
    }
    receive() external payable { }
    function withdrawWithTransfer() external{
        bank.withdrawWithTransfer();
    }
    function withdrawWithSend() external{
        bank.withdrawWithSend();
    }
    function withdrawWithCall(bytes memory input) external{
        bank.withdrawWithCall(abi.encodePacked(input));
    }

    function testPay() external payable returns(address){
        return 0x5B38Da6a701c568545dCfcB03FcB875f56beddC4;
    }

}