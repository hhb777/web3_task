// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract Voting{
    // 合约基础任务
    mapping(string name => uint votecounts) public votemap;
    string[] names;
    function vote(string memory name) public {
        votemap[name] += 1;
        names.push(name);
    }

    function getVotes(string calldata name) public view returns(uint){
        return votemap[name];
    }
    function resetVotes() public {
        for (uint i=0; i<names.length; i++){
            delete votemap[names[i]];
        }
    }
    
}


contract SolidityBase{
    // 反转字符串
    function reverseString(string memory str) public pure returns(string memory) {
        bytes memory bstr = bytes(str);
        for (uint i = 0; i < bstr.length /2; i++) {
            (bstr[i], bstr[bstr.length -i - 1]) = (bstr[bstr.length -i -1], bstr[i]);
        }
        return string(bstr);
    }

    //整数转罗马数字
    function IntToRoma(uint num) public pure returns(string memory){
        if (num < 1 || num > 3999) {
            revert("Invalid input");
        }
        uint th = num /1000;
        uint hu = (num - th*1000)/100;
        uint sh = (num - th*1000 - hu*100) / 10;
        uint ge = num - th*1000 - hu*100 - sh*10;

        bytes memory result = new bytes(16);
        uint index = 0;
        if (th > 0) {
            for (uint i=0; i<th;i++){
                result[index++] = bytes1("M");
            }
        }
        if (hu > 0) {
            if (hu == 9) {
                result[index++] = bytes1("C");
                result[index++] = bytes1("M");
            } else if (hu >= 5) {
                result[index++] = bytes1("D");
                for (uint i = 0; i < hu - 5; i++) {
                    result[index++] = bytes1("C");
                }
            } else if (hu == 4) {
                result[index++] = bytes1("C");
                result[index++] = bytes1("D");
            } else {
                for (uint i = 0; i < hu; i++) {
                    result[index++] = bytes1("C");
                }
            }
        }
        if (sh > 0) {
            if (sh == 9) {
                result[index++] = bytes1("X");
                result[index++] = bytes1("C");
            } else if (sh >= 5) {
                result[index++] = bytes1("L");
                for (uint i = 0; i < sh - 5; i++) {
                    result[index++] = bytes1("X");
                }
            } else if (sh == 4) {
                result[index++] = bytes1("X");
                result[index++] = bytes1("L");
            } else {
                for (uint i = 0; i < sh; i++) {
                    result[index++] = bytes1("X");
                }
            }
        }
        if (ge > 0) {
            if (ge == 9) {
                result[index++] = bytes1("I");
                result[index++] = bytes1("X");
            } else if (ge >= 5) {
                result[index++] = bytes1("V");
                for (uint i = 0; i < ge - 5; i++) {
                    result[index++] = bytes1("I");
                }
            } else if (ge == 4) {
                result[index++] = bytes1("I");
                result[index++] = bytes1("V");
            } else {
                for (uint i = 0; i < ge; i++) {
                    result[index++] = bytes1("I");
                }
            }
        }
        bytes memory realres = new bytes(index);
        for (uint i = 0; i < index; i++) {
            realres[i] = result[i];
        }
        return string(realres);
    }


    //罗马数字转整数
    function RomaToInt(string memory str) public pure returns(uint) {
        bytes memory bstr = bytes(str);
        uint result = 0;
        for (uint i=0; i<bstr.length; i++) {
            if (bstr[i] == bytes1("M")){
                result += 1000;
            } else if (bstr[i] == bytes1("D")) {
                result += 500;
            } else if (bstr[i] == bytes1("C")) {
                if (i +1 < bstr.length && bstr[i+1] == bytes1("M")) {
                    result += 900;
                    i++;
                } else if (i +1 < bstr.length && bstr[i+1] == bytes1("D")) {
                    result += 400;
                    i++;
                } else {
                    result += 100;
                }
            } else if (bstr[i] == bytes1("L")) {
                result += 50;
            } else if (bstr[i] == bytes1("X")) {
                if (i +1 < bstr.length && bstr[i+1] == bytes1("C")) {
                    result += 90;
                    i++;
                } else if (i +1 < bstr.length && bstr[i+1] == bytes1("L")) {
                    result += 40;
                    i++;
                } else {
                    result += 10;
                }
            }else if (bstr[i] == bytes1("V")){
                result += 5;
            }else if (bstr[i] == bytes1("I")) {
                if (i +1 < bstr.length && bstr[i+1] == bytes1("X")) {
                    result += 9;
                    i++;
                } else if (i +1 < bstr.length && bstr[i+1] == bytes1("V")) {
                    result += 4;
                    i++;
                } else {
                    result += 1;
                }
            }
        }
        return result;
    }

    //合并两个有序数组
    function MergeArr(uint[] memory nums1, uint[] memory nums2) public pure returns(uint[] memory){
        uint[] memory result = new uint[](nums1.length + nums2.length);
        uint i = 0;
        uint j = 0;
        uint k = 0;
        while (i < nums1.length && j < nums2.length) {
            if (nums1[i] < nums2[j]) {
                result[k] = nums1[i];
                i++;
            } else {
                result[k] = nums2[j];
                j++;
            }
            k++;
        }
        while (i < nums1.length) {
            result[k] = nums1[i];
            i++;
            k++;
        }
        while (j < nums2.length) {
            result[k] = nums2[j];
            j++;
            k++;
        }
        return result;
    }

    //二分法查找
    function BinarySearch(uint[] memory nums, uint target) public pure returns(uint){
        uint r = 0;
        uint l = nums.length - 1;
        for (uint mid=(r+l)/2;mid>=r;mid=(r+l)/2){
            if (nums[mid] == target){
                return mid;
            } else if (nums[mid] < target) {
                r = mid + 1;
            } else {
                l = mid - 1;
            }
        }
        return 0;
    } 

}