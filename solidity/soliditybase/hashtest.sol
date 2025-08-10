// SPDX-License-Identifier: MIT
pragma solidity ^0.8;


contract hashTest{
    bytes32 _msg = keccak256(abi.encodePacked("0xAA"));
    function weak(string memory string1) public view returns(bool){
        return keccak256(abi.encodePacked(string1)) == _msg;
    }

    // 强抗碰撞性
function strong(string memory string1, string memory string2) public pure returns (bool){
    return keccak256(abi.encodePacked(string1)) == keccak256(abi.encodePacked(string2));
}
}