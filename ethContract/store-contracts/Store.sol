pragma solidity ^0.8.0;

contract SimpleStorage {
    address public minter;
    uint private nowHeight;
    constructor() {
        minter = msg.sender;
    }
    mapping (uint => string) public sidechainHash;
    function setSidechainHash(uint number, string memory hash) public {
        require(msg.sender == minter);
        sidechainHash[number] = hash;
        nowHeight = number;
    }

    function getFromSidechainNum(uint number) public view returns (string memory) {
        return sidechainHash[number];
    }

    function getNewBlockNumber() public view returns (uint) {
        return nowHeight;
    }

    function getNewBlockHash() public view returns (string memory) {
        return sidechainHash[nowHeight];
    }
}