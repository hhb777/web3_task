// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Pair.sol";
import "./MyMeme.sol";

contract LiquidityManager is Ownable {
    IUniswapV2Router02 public uniswapRouter;
    MyMeme public memeToken;
    address public uniswapPair;

    event LiquidityAdded(uint256 tokenAmount, uint256 ethAmount);
    event LiquidityRemoved(uint256 liquidity);

    constructor(address _memeToken, address _uniswapRouter) Ownable(msg.sender) {
        memeToken = MyMeme(_memeToken);
        uniswapRouter = IUniswapV2Router02(_uniswapRouter);

        // 创建流动性池
        uniswapPair = IUniswapV2Factory(uniswapRouter.factory()).createPair(address(memeToken), uniswapRouter.WETH());
    }

    // 向流动性池添加流动性
    function addLiquidity(uint256 tokenAmount) external payable onlyOwner {
        memeToken.approve(address(uniswapRouter), tokenAmount);

        // 添加流动性
        (uint256 amountToken, uint256 amountETH, ) = uniswapRouter.addLiquidityETH{value: msg.value}(
            address(memeToken),
            tokenAmount,
            0, // 不接受滑点
            0, // 不接受滑点
            owner(),
            block.timestamp
        );

        emit LiquidityAdded(amountToken, amountETH);
    }

    // 交换代币为 ETH
    function swapExactTokensForETH(uint256 tokenAmount, uint256 minEthAmount, address to, uint256 deadline) external {
        require(memeToken.transferFrom(msg.sender, address(this), tokenAmount), "Transfer failed");

        // 批准路由器支取代币
        memeToken.approve(address(uniswapRouter), tokenAmount);

        // 进行交换
        address[] memory path = new address[](2);
        path[0] = address(memeToken);
        path[1] = uniswapRouter.WETH();

        uniswapRouter.swapExactTokensForETH(
            tokenAmount,
            minEthAmount,
            path,
            to,
            deadline
        );
    }
    

    // 从流动性池中移除流动性
    function removeLiquidity(uint256 liquidity) external onlyOwner {
        IUniswapV2Pair pair = IUniswapV2Pair(uniswapPair);
        pair.approve(address(uniswapRouter), liquidity);

        // 移除流动性
        uniswapRouter.removeLiquidityETH(
            address(memeToken),
            liquidity,
            0, // 不接受滑点
            0, // 不接受滑点
            owner(),
            block.timestamp
        );

        emit LiquidityRemoved(liquidity);
    }

    // 接收ETH
    receive() external payable {}
}