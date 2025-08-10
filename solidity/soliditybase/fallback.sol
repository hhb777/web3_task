// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

error sendFailed();
error CallFailed();
// contract fallbacktest{
//     event fallbackCalled(address sender, uint value, bytes data);

//     fallback() external payable{
//         emit fallbackCalled(msg.sender, msg.value, msg.data);
//     }

// }

contract  SendETH {
    constructor() payable{}
    receive() external payable { }

    function callETH(address payable _to, uint256 amount) external payable{
        (bool success,) = _to.call{value:amount}("");
        if (!success){
            revert CallFailed();
        }
    }
}

contract receiveETH{
    event LOG(uint amount, uint gas);
    receive() external payable {
        emit LOG(msg.value,gasleft());
    }
}