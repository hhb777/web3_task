// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

library MathLib{
    // function add(uint x,uint y) internal pure returns(uint){
    //     return x+y;
    // }
    // function sub(uint x,uint y) internal pure returns(uint){
    //     (x, y) = x >y ? (x,y):(y,x);
    //     return x-y;
    // }
    // function mul(uint x,uint y) internal pure returns(uint){
    //     return x*y;
    // }
    function add(uint x,uint y) external  pure returns(uint){
        return x+y;
    }
    function sub(uint x,uint y) external pure returns(uint){
        (x, y) = x >y ? (x,y):(y,x);
        return x-y;
    }
    //0xc8a4ac9c00000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000003
    function mul(uint x,uint y) external pure returns(uint){
        return x*y;
    }
}

contract MathContract{
    function add(uint x,uint y) public pure returns(uint){
        return MathLib.add(x,y);
    }
    function sub(uint x,uint y) public pure returns(uint){
        return MathLib.sub(x,y);
    }
    function mul(uint x,uint y) public pure returns(uint){
        return MathLib.mul(x,y);
    }
}

contract CalcTest{
    function calop(address addr, bytes4 selector, uint x, uint y) public returns(uint256){
        bytes memory data = abi.encodeWithSelector(selector, x,y);
        (bool suc, bytes memory resdata) = address(addr).call(data);
        require(suc);
        return abi.decode(resdata, (uint256));
    }
}

library ArrayOP{
    function contains(uint[] memory nums, uint num) external pure returns(bool){
        for(uint i=0;i<nums.length;i++){
            if(nums[i]==num){
                return true;
            }
        }
        return false;
    }
}

contract ArrayTest{
    using ArrayOP for uint[];
    uint[] test;
    function setarray(uint num) public {
        test.push(num);
    }
    function tst(uint nums) public view returns(bool){
        return test.contains(nums);
    }
}