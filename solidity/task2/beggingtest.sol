// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract BeggingContract{
    event Donation(address addr, uint256 amount);
    address public owner = msg.sender;
    mapping(address account => uint256 amount) private forbeggingAccount;
    address[] addrs;
    uint permittime = block.timestamp - 30 seconds;

    function donate() external payable{
        require(block.timestamp >= permittime + 30 seconds, "Times not reach");
        permittime = block.timestamp;
        if (forbeggingAccount[msg.sender] == 0){
            addrs.push(msg.sender);
        }
        forbeggingAccount[msg.sender] += msg.value;

        emit Donation(msg.sender, msg.value);
    }

    modifier OnlyOwner{
        require(owner == msg.sender, "Only owner can withdraw");
        _;
    }

    function withdrow(address payable addr, uint256 amount) external payable OnlyOwner{
        require(address(this).balance != 0, "contract left is empty");
        addr.transfer(amount);
    }

    function getDonation(address addr) external view returns(uint256){
        require(addr != address(0), "Invalid address");
        require(forbeggingAccount[addr] != 0, "you didn't donate");

        return forbeggingAccount[addr];
    }

    function getsortthreeDonations() external returns(address[] memory){
        for(uint i=0;i<addrs.length;i++){
            for(uint j=i+1;j<addrs.length;j++){
                if(forbeggingAccount[addrs[i]]<forbeggingAccount[addrs[j]]){
                    (addrs[i],addrs[j]) = (addrs[j],addrs[i]);
                }
            }
        }
        address[] memory tmpaddr = new address[](3);
        uint length = addrs.length > 3? 3:addrs.length;
        for(uint i=0;i<length;i++){
            tmpaddr[i] = addrs[i];
        }
        return tmpaddr;
    }
    receive() external payable { }
    fallback() external payable { }
}