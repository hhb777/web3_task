// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// 任何实现此接口的类都可以用于协议特定的税务计算。
interface ITaxHandler{
    // 获取要作为税款支付的代币数量。
    // benefactor - 捐助者地址
    // beneficiary - 受益人地址
    // amount - 转账中的代币数量
    // retrun - 需缴税的代币数量。
    function getTax(address benefactor, address beneficiary, uint256 amount) external view returns(uint256);

}