// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./ITaxHandler.sol";
import "./ITreasuryHandler.sol";
import "./IGovernanceToken.sol";

// Floki代币合约
//Floki代币具有税务和国库处理以及治理功能的模块化系统。
contract FLOKI is IERC20, Ownable, IGovernanceToken {
    // 注册用户token余额的映射
    mapping(address => uint256) private _balances;
    // 注册用户授权额度的映射
    mapping(address => mapping(address => uint256)) private _allowances;
    // 用于治理的用户委托注册表
    mapping(address => address) private _delegates;
    //非投票代表登记处
    mapping(address => uint256) public nonces;
    //帐户具有的余额检查点数量的注册表。
    mapping(address => uint32) public numCheckpoints;
    //每个帐户的余额检查点的注册表。
     mapping(address => mapping(uint32 => Checkpoint)) public checkpoints;
    //合约域的EIP-712类型哈希。
    bytes32 public constant DOMAIN_TYPEHASH = keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    //委托人结构的EIP-712类型哈希。
    bytes32 public constant DELEGATION_TYPEHASH = keccak256("Delegation(address delegatee,uint256 nonce,uint256 expiry)");
    //执行税务计算的合同
    ITaxHandler public taxHandler;
    //执行财务相关业务的合同
    ITreasuryHandler public treasuryHandler;
    //税务处理人合同变更时发出
    event TaxHandlerChanged(address indexed oldHandler, address indexed newHandler);
    //当资金处理人合同发生变更时触发
    event TreasuryHandlerChanged(address indexed oldHandler, address indexed newHandler);
    //代币名称
    string private _name;
    //代币符号
    string private _symbol;
    //name_ - 代币名称
    //symbol_ - 代币符号
    //taxHandlerAddress - 初始化税务处理程序合约的地址
    //treasuryHandlerAddress - 初始资金处理人合约的地址
    constructor(string memory name_, string memory symbol_, address taxHandlerAddress, address treasuryHandlerAddress) IERC20() Ownable(_msgSender()) {
        _name = name_;
        _symbol = symbol_;
        taxHandler = ITaxHandler(taxHandlerAddress);
        treasuryHandler = ITreasuryHandler(treasuryHandlerAddress);
        _balances[_msgSender()] = totalSupply();
        emit Transfer(address(0), _msgSender(), totalSupply());
    }

    // 获取代币名称
    function name() public view returns (string memory) {
        return _name;
    }
    // 获取代币符号
    function symbol() public view returns (string memory) {
        return _symbol;
    }
    // 获取代币小数位数
    function decimals() external pure returns (uint8) {
        return 9;
    }
    // 获取代币总供应量
    function totalSupply() public pure returns (uint256) {
        return 1e13 * 1e9;
    }
    // 获取账户余额
    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }
    //将token从呼叫者的地址转移到另一个地址
    //recipient 用于将呼叫者的发送token目标地址
    //amount 给接收者传输token数量
    //return 成功返回true,失败raise异常
    function transfer(address recipient, uint256 amount) public returns (bool) {
        _transfer(_msgSender(), recipient, amount);
        return true;
    }
    //获取owner授权spender
    //owner - “消费者”可以代表谁消费代币的地址。
    //spender - 被授权代表所有者消费代币的地址。
    //return - 允许spender代表所有者消费的代币数量。
    function allowance(address owner, address spender) external view override returns (uint256) {
        return _allowances[owner][spender];
    }

    // 授权spender代表呼叫者消费代币
    function approve(address spender, uint256 amount) external override returns (bool) {
        _approve(_msgSender(), spender, amount);
        return true;
    }

    function transferFrom(address sender, address recipient, uint256 amount) external override returns (bool) {
        _transfer(sender, recipient, amount);
        uint256 currentAllowance = _allowances[sender][_msgSender()];
        require(currentAllowance >= amount, "FLOKI: transferFrom:ALLOWANCE_EXCEEDED: Transfer amount exceeds allowance");
        //禁用块内溢出和下溢检查，允许操作直接执行
        unchecked {
            _approve(sender, _msgSender(), currentAllowance - amount);
        }
        return true;
    }
    //增加spender的授权额度
    function increaseAllowance(address spender, uint256 addedValue) external returns (bool) {
        _approve(_msgSender(), spender, _allowances[_msgSender()][spender] + addedValue);
        return true;
    }
    //减少spender的授权额度
    function decreaseAllowance(address spender, uint256 subtractedValue) external returns (bool) {
        uint256 currentAllowance = _allowances[_msgSender()][spender];
        require(currentAllowance >= subtractedValue, "FLOKI: decreaseAllowance:ALLOWANCE_UNDERFLOW: Decreased allowance below zero");
        //禁用块内溢出和下溢检查，允许操作直接执行
        unchecked {
            _approve(_msgSender(), spender, currentAllowance - subtractedValue);
        }
        return true;
    }

    // 给地址委托投票
    //应该指出的是，想要自己投票的用户也需要调用此方法，尽管他们有自己的地址。
    function delegate(address delegatee) external {
        _delegate(_msgSender(), delegatee);
    }
    // 通过签名给地址委托投票
    // delegatee - 被委托投票的地址
    // nonce - 需要与签名匹配的合约状态
    // expiry - 签名过期的时间戳
    // v - 签名的恢复字节
    // r - ECDSA签名对的一半
    // s - ECDSA签名对的一半
    function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) external {
        bytes32 domainSeparator = keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes(name())), block.chainid, address(this)));
        bytes32 structHash = keccak256(abi.encode(DELEGATION_TYPEHASH, delegatee, nonce, expiry));
        //"\x19\x01"是EIP-712规定的固定前缀,可以有效防止重放攻击
        //通过代币名称，合约地址，链ID和结构(委托地址，nonce,超时时间)哈希计算出最终的消息摘要,确保签名的唯一性和安全性
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));
        //ecrecover(digest, v, r, s) 的作用是基于提供的签名数据，从哈希过的消息（digest）中恢复出签名者的以太坊地址
        address signatory = ecrecover(digest, v, r, s);
        require(signatory != address(0), "FLOKI: delegateBySig:INVALID_SIGNATURE: Invalid signature");
        require(nonce == nonces[signatory]++, "FLOKI: delegateBySig:INVALID_NONCE: Invalid nonce");
        require(block.timestamp <= expiry, "FLOKI: delegateBySig:SIGNATURE_EXPIRED: Signature expired");
        _delegate(signatory, delegatee);
    }

    // 确定截至区块编号的帐户投票数
    function getVotesAtBlock(address account, uint32 blockNumber) public view override returns (uint224) {
        require(blockNumber < block.number, "FLOKI: getVotesAtBlock:NOT_FINALIZED: Block not yet finalized");
        uint32 nCheckpoints = numCheckpoints[account];
        if (nCheckpoints == 0) {
            return 0;
        }
        //检查点的最新块号小于或等于请求的块号
        if (checkpoints[account][nCheckpoints - 1].blockNumber <= blockNumber) {
            return checkpoints[account][nCheckpoints - 1].votes;
        }
        //检查点的第一个块号大于请求的块号
        if (checkpoints[account][0].blockNumber > blockNumber) {
            return 0;
        }
        //二进制搜索检查点以找到请求的块号
        uint32 lower = 0;
        uint32 upper = nCheckpoints - 1;
        while (upper > lower) {
            uint32 center = upper - (upper - lower) / 2; // ceil, avoiding overflow
            Checkpoint memory cp = checkpoints[account][center];
            if (cp.blockNumber == blockNumber) {
                return cp.votes;
            } else if (cp.blockNumber < blockNumber) {
                lower = center;
            } else {
                upper = center - 1;
            }
        }
        return checkpoints[account][lower].votes;
    }

    // 设置新的税务处理程序合约地址
    // newHandler - 新税务处理程序合约地址
    function setTaxHandler(address newHandler) external onlyOwner {
        require(newHandler != address(0), "FLOKI: setTaxHandler:ZERO_ADDRESS: New handler is the zero address");
        address oldHandler = address(taxHandler);
        taxHandler = ITaxHandler(newHandler);
        emit TaxHandlerChanged(oldHandler, newHandler);
    }
    // 设置新的资金处理人合约地址
    // newHandler - 新资金处理人合约地址
    function setTreasuryHandler(address newHandler) external onlyOwner {
        require(newHandler != address(0), "FLOKI: setTreasuryHandler:ZERO_ADDRESS: New handler is the zero address");
        address oldHandler = address(treasuryHandler);
        treasuryHandler = ITreasuryHandler(newHandler);
        emit TreasuryHandlerChanged(oldHandler, newHandler);
    }

    //从一个地址到另一个地址的委托投票
    function _delegate(address delegator, address delegatee) private {
        address currentDelegate = _delegates[delegator];
        uint256 delegatorBalance = _balances[delegator]; // 代币余额也就是投票权
        _delegates[delegator] = delegatee;
        emit DelegateChanged(delegator, currentDelegate, delegatee);
        _moveDelegates(currentDelegate, delegatee, uint224(delegatorBalance));
    }

    //在投票代表之间移动投票
    function _moveDelegates(address srcRep, address dstRep, uint224 amount) private {
        if (srcRep != dstRep && amount > 0) {
            if (srcRep != address(0)) {
                uint32 srcRepNum = numCheckpoints[srcRep];
                uint224 srcRepOld = srcRepNum > 0 ? checkpoints[srcRep][srcRepNum - 1].votes : 0;
                uint224 srcRepNew = srcRepOld - amount;
                _writeCheckpoint(srcRep, srcRepNum, srcRepOld, srcRepNew);
            }
            if (dstRep != address(0)) {
                uint32 dstRepNum = numCheckpoints[dstRep];
                uint224 dstRepOld = dstRepNum > 0 ? checkpoints[dstRep][dstRepNum - 1].votes : 0;
                uint224 dstRepNew = dstRepOld + amount;
                _writeCheckpoint(dstRep, dstRepNum, dstRepOld, dstRepNew);
            }
        }
    }

    //为代表写入检查点
    // delegatee - 用于写入检查点的地址
    // nCheckpoints - 代表的当前delegatee(被委派者)已经有的检查点数量
    // oldVotes - 此检查点之前的投票数
    // newVotes - “被委派者”现在拥有的投票数。
    function _writeCheckpoint(address delegatee, uint32 nCheckpoints, uint224 oldVotes, uint224 newVotes) private {
        uint32 blockNumber = uint32(block.number);
        //检查点数量不会溢出，因为每个块最多只能添加一个检查点
        if (nCheckpoints > 0 && checkpoints[delegatee][nCheckpoints - 1].blockNumber == blockNumber) {
            checkpoints[delegatee][nCheckpoints - 1].votes = newVotes;
        } else {
            checkpoints[delegatee][nCheckpoints] = Checkpoint(blockNumber, newVotes);
            numCheckpoints[delegatee] = nCheckpoints + 1;
        }
        emit DelegateVotesChanged(delegatee, oldVotes, newVotes);
    }

    // 代表owner批准spender
    function _approve(address owner, address spender, uint256 amount) private {
        require(owner != address(0), "FLOKI: approve:OWNER_ZERO_ADDRESS: approve from the zero address");
        require(spender != address(0), "FLOKI: approve:SPENDER_ZERO_ADDRESS: approve to the zero address");
        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    // 从一个地址到另一个地址的代币转移
    function _transfer(address sender, address recipient, uint256 amount) private {
        require(sender != address(0), "FLOKI: transfer:SENDER_ZERO_ADDRESS: transfer from the zero address");
        require(recipient != address(0), "FLOKI: transfer:RECIPIENT_ZERO_ADDRESS: transfer to the zero address");
        require(amount > 0, "FLOKI: transfer:ZERO_AMOUNT: Transfer amount must be greater than zero");
        require(_balances[sender] >= amount, "FLOKI: transfer:INSUFFICIENT_BALANCE: Transfer amount exceeds balance");
        //在转账之前调用金库处理程序
        treasuryHandler.beforeTransferHandler(sender, recipient, amount);
        uint256 taxAmount = taxHandler.getTax(sender, recipient, amount);
        uint256 netAmount = amount - taxAmount;
        _balances[sender] -= amount;
        _balances[recipient] += netAmount;
        _moveDelegates(_delegates[sender], _delegates[recipient], uint224(netAmount));
        
        if (taxAmount > 0) {
            _balances[address(treasuryHandler)] += taxAmount;
            _moveDelegates(_delegates[sender], _delegates[address(treasuryHandler)], uint224(taxAmount));
            emit Transfer(sender, address(treasuryHandler), taxAmount);
        }
        
        //在转账之后调用金库处理程序
        treasuryHandler.afterTransferHandler(sender, recipient, amount);
        emit Transfer(sender, recipient, netAmount);
    }
}