// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract CounterTest{
    event CounterAddEvent(uint indexed count);
    uint public count;
    function CountAdd() external {
        count++;
        emit CounterAddEvent(count);
    }
}