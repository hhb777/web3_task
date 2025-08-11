// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ArrayTest{
    address[] adds;
    function Sumarray(uint[] calldata nums) public pure returns(uint sum){
        for (uint i=0;i<nums.length;i++){
            sum += nums[i];
        }
    }
    function setArray(address[] memory maddrs) public {
        adds = maddrs;
    }
    function delTest(address target) public returns(address[] memory){
        for (uint i=0;i<adds.length;i++){
            if (adds[i] == target){
                (adds[i], adds[adds.length-1]) = (adds[adds.length-1], adds[i]);
                break;
            }
            
        }
        adds.pop();
        return adds;
    
    }

    function opstr(bytes memory bs, string memory str) public pure returns(string memory,string memory){
        bs[0] = bytes1("H");
        // str[0] = 'M'; // ERROR
        return (string(bs), str);

    }
}