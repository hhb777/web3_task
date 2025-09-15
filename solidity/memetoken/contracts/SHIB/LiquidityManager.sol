// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract LiquidityManager is Ownable {
    IERC20 public tokenA;
    IERC20 public tokenB;

    uint256 public totalLiquidity;
    uint256 public feePercentage = 3; // 3% 交易手续费
    uint256 public constant FEE_DENOMINATOR = 100; // 手续费分母，用于计算百分比

    mapping(address => uint256) public liquidityOf;

    // Tick结构体，用于记录价格等级
    struct Tick {
        uint256 price; // 当前价格
        uint256 liquidity; // 当前价格下的流动性
        bool isActive; // Tick是否有效
    }

    mapping(uint256 => Tick) public ticks; // 价格等级
    uint256 public tickCount; // 当前价格等级数量

    event LiquidityAdded(address indexed provider, uint256 amountA, uint256 amountB);
    event LiquidityRemoved(address indexed provider, uint256 amountA, uint256 amountB);
    event Swapped(address indexed user, uint256 amountIn, uint256 amountOut, bool isAToB);
    event TickAdded(uint256 indexed tickIndex, uint256 price, uint256 liquidity);
    event TickUpdated(uint256 indexed tickIndex, uint256 liquidity);

    constructor(address _tokenA, address _tokenB) Ownable(msg.sender){
        tokenA = IERC20(_tokenA);
        tokenB = IERC20(_tokenB);
    }

    // 增加流动性
    function addLiquidity(uint256 amountA, uint256 amountB) external {
        require(amountA > 0 && amountB > 0, "Invalid amounts");

        tokenA.transferFrom(msg.sender, address(this), amountA);
        tokenB.transferFrom(msg.sender, address(this), amountB);

        totalLiquidity += amountA + amountB;
        liquidityOf[msg.sender] += amountA + amountB;

        emit LiquidityAdded(msg.sender, amountA, amountB);
    }

    // 移除流动性
    function removeLiquidity(uint256 amountA, uint256 amountB) external {
        require(liquidityOf[msg.sender] >= amountA + amountB, "Insufficient liquidity");

        totalLiquidity -= amountA + amountB;
        liquidityOf[msg.sender] -= amountA + amountB;

        tokenA.transfer(msg.sender, amountA);
        tokenB.transfer(msg.sender, amountB);

        emit LiquidityRemoved(msg.sender, amountA, amountB);
    }

    // 代币交换
    function swap(uint256 amountIn, bool isAToB) external {
        require(amountIn > 0, "Invalid amount");

        uint256 fee = amountIn * feePercentage / FEE_DENOMINATOR;
        uint256 amountAfterFee = amountIn - fee;

        if (isAToB) {
            require(tokenA.balanceOf(address(this)) >= amountIn, "Insufficient tokenA in pool");
            uint256 amountOut = getSwapAmount(amountAfterFee, true); // 计算输出金额
            require(tokenB.balanceOf(address(this)) >= amountOut, "Insufficient tokenB in pool");

            tokenA.transferFrom(msg.sender, address(this), amountIn);
            tokenB.transfer(msg.sender, amountOut);

            // 记录手续费
            tokenA.transfer(owner(), fee);

            emit Swapped(msg.sender, amountIn, amountOut, true);
        } else {
            require(tokenB.balanceOf(address(this)) >= amountIn, "Insufficient tokenB in pool");
            uint256 amountOut = getSwapAmount(amountAfterFee, false); // 计算输出金额
            require(tokenA.balanceOf(address(this)) >= amountOut, "Insufficient tokenA in pool");

            tokenB.transferFrom(msg.sender, address(this), amountIn);
            tokenA.transfer(msg.sender, amountOut);

            // 记录手续费
            tokenB.transfer(owner(), fee);

            emit Swapped(msg.sender, amountIn, amountOut, false);
        }
    }

    // 添加价格等级
    function addTick(uint256 price, uint256 liquidity) external onlyOwner {
        ticks[tickCount] = Tick(price, liquidity, true);
        emit TickAdded(tickCount, price, liquidity);
        tickCount++;
    }

    // 更新价格等级流动性
    function updateTick(uint256 tickIndex, uint256 liquidity) external onlyOwner {
        require(tickIndex < tickCount, "Tick does not exist");
        require(ticks[tickIndex].isActive, "Tick is not active");

        ticks[tickIndex].liquidity = liquidity;
        emit TickUpdated(tickIndex, liquidity);
    }

    // 简单的交换金额计算，基于当前 Tick 的价格
    function getSwapAmount(uint256 amountIn, bool isAToB) internal view returns (uint256) {
        uint256 price = getCurrentTickPrice();
        if (isAToB) {
            return amountIn * price / 10**18; // 假设价格以 1e18 为单位表示
        } else {
            return amountIn * 10**18 / price;
        }
    }

    // 获取当前 Tick 的价格，可以根据你的逻辑实现
    function getCurrentTickPrice() internal view returns (uint256) {
        require(tickCount > 0, "No ticks available");
        // 返回第一个活动的 Tick 的价格
        for (uint256 i = 0; i < tickCount; i++) {
            if (ticks[i].isActive) {
                return ticks[i].price;
            }
        }
        revert("No active ticks found");
    }

    receive() external payable {}
}