// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/**
代币税功能：实现交易税机制，对每笔代币交易征收一定比例的税费，并将税费分配给特定的地址或用于特定的用途。
流动性池集成：设计并实现与流动性池的交互功能，支持用户向流动性池添加和移除流动性。
交易限制功能：设置合理的交易限制，如单笔交易最大额度、每日交易次数限制等，防止恶意操纵市场
 */

contract MyMeme is Ownable, ERC20 {
    // 自定义转账事件
    event MyTransfer(address indexed from, address indexed to, uint256 value);

    // 税率阶梯结构体
    struct TaxTier{
        uint32 threshold; // 阈值
        uint16 taxRate;   // 税率，单位为百分比的万分之一（例如，1000表示10%）
    }

    address private Tax_pool; // 税收池地址

    TaxTier[] private taxTiers; // 税率阶梯

    string public _name;
    string public _symbol;

    uint256 private maxTxAmount; // 最大交易额
    uint32 private dailyTxLimit; // 每日交易次数限制
    mapping(address => uint256) private lastTx; //最后一次交易时间
    mapping(address => uint32) private dailyTxCount; // 记录每日交易次数

    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) Ownable(msg.sender) {
        _name = name_;
        _symbol = symbol_;
    }

    // 铸造新代币
    function mint(address to, uint256 amount) external onlyOwner {
        _mint(to, amount);
    }

    // 设置税收池地址
    function setTaxPool(address taxPool) external onlyOwner {
        Tax_pool = taxPool;
    }

    // 设置最大交易额
    function setMaxTxAmount(uint256 amount) external onlyOwner {
        maxTxAmount = amount;
    }
    // 设置每日交易次数限制
    function setDailyTxLimit(uint32 limit) external onlyOwner {
        dailyTxLimit = limit;
    }

    //添加税率阶梯
    function addTaxTier(uint32 threshold, uint16 taxRate) external onlyOwner{
        require(taxTiers.length == 0 || threshold > taxTiers[taxTiers.length - 1].threshold, "Threshold must be greater than previous tier");
        taxTiers.push(TaxTier(threshold, taxRate));
    }

    // //获取税率阶梯
    // function getTaxTiers() external view returns (TaxTier[] memory){
    //     return taxTiers;
    // }

    //计算税额
    function calculateTax(uint256 amount) internal view returns (uint256){
        uint16 taxRate = 0;
        for(uint i = 0; i < taxTiers.length; i++){
            //检查是否达到当前阶梯的阈值
            if(amount >= taxTiers[i].threshold){
                taxRate = taxTiers[i].taxRate;
            }else{
                break;
            }
        }
        return amount * taxRate / 10000; // 假设税率是以百分比的万分之一表示的
    }
    // 应用税收
    function applyTax(address from, uint256 amount) internal {
        uint256 tax = calculateTax(amount);
        if(tax > 0){
            _transfer(from, Tax_pool, tax); // 将税款转移到税款地址
        }
    }
    // 转账应用税收
    function metransfer(address from, address to, uint256 amount) external {
        // 检查交易限制
        require(amount <= maxTxAmount, "Transfer amount exceeds the maxTxAmount.");
        // 每天交易次数清零
        if (block.timestamp >= lastTx[from] + 1 days) {
            dailyTxCount[from] = 0;
        }
        // 检查每日交易次数
        require(dailyTxCount[from] < dailyTxLimit, "Exceeds daily transaction limit.");
        dailyTxCount[from]++;
        // 应用税收
        applyTax(from, amount);
        super._transfer(from, to, amount - calculateTax(amount)); // 转移扣除税款后的金额
        emit MyTransfer(from, to, amount - calculateTax(amount));
    }





}