// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract deferencetest{
    uint[] data;
    function updateData(uint[] memory newData) public {
        data = newData;
    }
    function getData() public view returns(uint[] memory){
        return data;
    }
    function modifyStorageData(uint index, uint value) public {
        data[index] = value;
    }
    function modifyMemoryData(uint[] memory memData) public view returns(uint[] memory, uint[] memory){
        memData[1] = 10;
        memData[2] = 100;
        return (memData, data);
    }
}