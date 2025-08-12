// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

contract ERC20Test {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint256 public totalSupply;
    address public owner = msg.sender;
    constructor(string memory _name, string memory _symbol, uint8 _decimals){
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }
    mapping(address => uint256) private balances;
    mapping(address => mapping(address => uint256)) private allowance;
    event STransfer(address indexed _from , address indexed _to, uint256 _value);
    event SApproval(address indexed _owner, address indexed _spender, uint256 _value2);
    // 查询账户余额
    function getbalance(address addr) public view returns(uint256){
        return balances[addr];
    }
    // 转账
    function trans(address to, uint256 amount) public returns(bool){
        require(to != address(0), "to is zero address");
        require(balances[msg.sender] >= amount, "left less than current balance");
        emit STransfer(msg.sender, to, amount);
        balances[msg.sender] -= amount;
        balances[to] += amount;
        return true;
    }
    //授权
    function appro(address spender, uint256 amount) public returns(bool){
        require(spender != address(0), "spender is zero address");
        emit SApproval(msg.sender, spender, amount);
        allowance[msg.sender][spender] = amount;
        return true;
    }
    //代扣转账
    function transfrom(address from, address to, uint256 amount) public returns(bool){
        require(from != address(0), "from is zero address");
        require(to != address(0), "to is zero address");
        require(balances[from] >= amount, "left less than current balance");
        require(allowance[msg.sender][from] >= amount, "auth left less than current balance");
        emit STransfer(from, to, amount);
        balances[from] -= amount;
        balances[to] += amount;
        allowance[msg.sender][from] -= amount;

        return true;
    }
    modifier Onlyowner{
        require(msg.sender == owner, "NOT is contract owner");
        _;
    }
    //合约所有者增发代币
    function mint(address to, uint amount) public Onlyowner{
        owner = to;
        totalSupply += amount;
        balances[msg.sender] += amount;
        emit STransfer(address(0), msg.sender, amount);
    }
}