// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract Pair{
    address public factory;
    address public token0;
    address public token1;

    constructor() payable{
        factory = msg.sender;
    }

    function initialize(address _token0, address _token1) external{
        require(msg.sender == factory, 'UniswapV2: FORBIDDEN'); // subcall to uniswap library);
        token0 = _token0;
        token1 = _token1;
    }
}

contract PairFactory2{
    mapping(address => mapping(address => address)) public getpair;
    address[] public allpair; 

    function createPair2(address tokenA, address tokenB) external returns (address pairAddr){
        require(tokenA != tokenB, 'IDENTICAL_ADDRESS');
        (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA);
        bytes32 salt = keccak256(abi.encodePacked(token0,token1));
        Pair pair = new Pair{salt:salt}();
        pair.initialize(tokenA, tokenB);
        pairAddr = address(pair);
        allpair.push(pairAddr);
        getpair[token0][token1] = pairAddr;
        getpair[token1][token0] = pairAddr;
    }

    // 提前计算pair合约地址
function calculateAddr(address tokenA, address tokenB) public view returns(address predictedAddress){
    require(tokenA != tokenB, 'IDENTICAL_ADDRESSES'); //避免tokenA和tokenB相同产生的冲突
    // 计算用tokenA和tokenB地址计算salt
    (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA); //将tokenA和tokenB按大小排序
    bytes32 salt = keccak256(abi.encodePacked(token0, token1));
    // 计算合约地址方法 hash()
    predictedAddress = address(uint160(uint(keccak256(abi.encodePacked(
        bytes1(0xff),
        address(this),
        salt,
        keccak256(type(Pair).creationCode)
        )))));
}
}