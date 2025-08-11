// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

// interface IVault {
//     function addamount(uint) external returns (uint);
//     function subtractamount(uint) external returns (uint);
// }

// contract Bank is IVault{
//     mapping(address => uint) public balances;
//     function addamount(uint amount) external override returns(uint balance){

//         balances[msg.sender] += amount;
//         return balances[msg.sender];
//     }
//     function subtractamount(uint amount) external override returns(uint balance){
//         balances[msg.sender] -= amount;
//         return balances[msg.sender];
//     }
// }


interface MyTokenTest{
    function goldtransfer(address, uint) external;
}

contract MyToken is MyTokenTest{
    mapping(address => uint) balances;
    constructor(){
        balances[msg.sender] = 10000000;
    }
    function goldtransfer(address addr, uint amount) public override{
        balances[msg.sender] -= amount;
        balances[addr] += amount;
    }
}

contract Reward{
    MyTokenTest mytoken;
    constructor(MyTokenTest _mytoken){
        mytoken = _mytoken;
    }

    function reward(address addr, uint amount) public{
        mytoken.goldtransfer(addr, amount);
    }
}