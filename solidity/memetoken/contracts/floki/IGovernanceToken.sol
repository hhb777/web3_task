// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


interface IGovernanceToken {
    //标记给定区块投票数量的检查点结构体。
    struct Checkpoint{
        //检查点创建时的块号
        uint32 blockNumber;
        //出于优化目的，此类型设置为“uint224”（即，专门适合32字节的块）。它假设实现治理令牌的投票数永远不会超过224位数字的最大值。
        uint224 votes;
    }
    //确定截至区块编号的帐户投票数。
    //块编号必须是最终确定的块，否则此功能将恢复以防止错误信息。
    // account - 要检查的帐户地址。
    // blockNumber - 获得投票平衡的区块号。
    // return - 截至给定区块，该账户的投票数。
    function getVotesAtBlock(address account, uint32 blockNumber) external view returns(uint224);

    //每当为某个帐户设置新代理时触发。
    event DelegateChanged(address delegator, address currentDelegate, address newDelegate);

    //当代表的投票计数发生变化时触发。
    event DelegateVotesChanged(address delegatee, uint224 oldVotes, uint224 newVotes);
}