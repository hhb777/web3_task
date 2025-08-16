// SPDX-License-Identifier: MIT
pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import {AggregatorV3Interface} from "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";
import "hardhat/console.sol";

contract Auction is Initializable, UUPSUpgradeable, IERC721Receiver{
// contract Auction{
    struct Auctions{
        //NFT 合约地址
        address nftAddress;
        //NFT ID
        uint256 nftId;
        //卖家
        address owner;
        //起拍价
        uint256 startPrice;
        //一口价
        uint256 endPrice;
        //拍卖开始时间
        uint256 startTime;
        //拍卖持续时间
        uint256 duration;
        //最高出价
        uint256 highestBid;
        //最高出价者
        address highestBidder;
        //竞拍资产类型
        address tokenAddress;
        //拍卖是否结束
        bool auctionEnded;
        //手续费用
        uint256 fee;
    }
    address public admin;
    uint256 public nextAuctionId;
    mapping(uint256 => Auctions) public bids;

    mapping(address => AggregatorV3Interface) public priceFeeds;

    function initialize() public initializer {
        admin = msg.sender;
    }

    //设置价格预言机
    function setPriceFeed(address tokenAddress, address feedAddress) public{
        priceFeeds[tokenAddress] = AggregatorV3Interface(feedAddress);
    }

    //获取价格预言机的最新价格
    function getLastPrice(address tokenAddress) public view returns(int256) {
        AggregatorV3Interface priceFeed = priceFeeds[tokenAddress];
        (, int256 price, , , ) = priceFeed.latestRoundData();
        require(price > 0, "Invalid price");
        return price;
    }

    //创建拍卖
    function createAuction(address _nftaddress, uint256 _nftId, uint256 _startPrice, uint256 _duration) public {
        require(_nftaddress != address(0), "NFT address cannot be zero");
        require(_startPrice > 0, "Start price must be greater than 0");
        require(_duration > 0, "Duration must be greater than 0");
        require(msg.sender == admin, "Only admin can create auction");
        IERC721(_nftaddress).safeTransferFrom(msg.sender, address(this), _nftId);
        bids[nextAuctionId] = Auctions({
            nftId: _nftId,
            nftAddress: _nftaddress,
            owner: msg.sender,
            startPrice: _startPrice,
            endPrice: 0,
            startTime: block.timestamp,
            duration: _duration,
            highestBid: 0,
            highestBidder: address(0),
            auctionEnded: false,
            tokenAddress: address(0),
            fee: _startPrice / 100 // 1% fee
        });
        nextAuctionId++;
    }

    function startbid(uint256 auctionid, uint256 amount, address tokenaddr) external payable{

        //参与auctionid这个拍卖
        Auctions storage auctions = bids[auctionid];
        require(auctions.auctionEnded == false && (auctions.startTime + auctions.duration) > block.timestamp, "Auction has already ended");
        uint payValue;
        if (tokenaddr != address(0)){
            //金额转换
            payValue = amount * uint(getLastPrice(tokenaddr));
        } else {
            //ETH
            amount = msg.value;
            payValue = amount * uint(getLastPrice(address(0)));
        }

        //将拍卖中的初始价格和最高出价转换为ERC20代币价格进行比较
        uint mstartprice = auctions.startPrice * uint(getLastPrice(tokenaddr));
        uint mhighestbid = auctions.highestBid * uint(getLastPrice(tokenaddr));
        require(payValue > mstartprice && payValue > mhighestbid, "Bid amount is too low");

        //如果是ERC20代币，先转账
        if (tokenaddr != address(0)) {
            IERC20(tokenaddr).transferFrom(msg.sender, address(this), amount);
        }
        //如果有出价者，退还之前的最高出价
        if (auctions.highestBid > 0) {
            if (tokenaddr != address(0)){
                IERC20(tokenaddr).transfer(auctions.highestBidder, auctions.highestBid);
            } else{
                payable(auctions.highestBidder).transfer(auctions.highestBid);
            }
        }
        auctions.tokenAddress = tokenaddr;
        auctions.highestBid = amount;
        auctions.fee = amount / 100; // 1% fee
        auctions.highestBidder = msg.sender;

    } 

    //结束拍卖
    function endbid(uint256 auctionid) external{
        Auctions storage auctions = bids[auctionid];
        // console.log("endAuction", auctions.startTime, auctions.duration, block.timestamp);
        require(block.timestamp > (auctions.startTime + auctions.duration), "Auction has not ended yet");
        require(!auctions.auctionEnded, "Auction has already been ended");
        //转出ERC721 NFT
        IERC721(auctions.nftAddress).safeTransferFrom(address(this), auctions.highestBidder, auctions.nftId);
        //结束拍卖
        auctions.auctionEnded = true;
        //转出手续费
        if (auctions.fee > 0) {
            if (auctions.tokenAddress != address(0)) {
                IERC20(auctions.tokenAddress).transfer(admin, auctions.fee);
            } else {
                payable(admin).transfer(auctions.fee);
            }
        }
    }
    
    //升级合约
    function _authorizeUpgrade(address) internal view override{
        require(msg.sender == admin, "Only admin can upgrade the contract");
    }

    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes calldata data
    ) external pure override returns (bytes4) {
        return this.onERC721Received.selector;
    }
}