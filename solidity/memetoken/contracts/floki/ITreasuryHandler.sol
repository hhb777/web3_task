// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


//金库处理程序接口
// 任何实现此接口的类都可以用于与资金相关的协议特定操作。
interface ITreasuryHandler{
    // 在执行传输之前执行操作。
    // benefactor - 捐助者地址
    // beneficiary - 受益人地址
    // amount - 转账中的代币数量  
    function beforeTransferHandler(address benefactor, address bemeficiary, uint256 amount) external;
    // 在执行传输之后执行操作。
    // benefactor - 捐助者地址
    // beneficiary - 受益人地址
    // amount - 转账中的代币数量 
    function afterTransferHandler(address benefactor, address bemeficiary, uint256 amount) external;
}