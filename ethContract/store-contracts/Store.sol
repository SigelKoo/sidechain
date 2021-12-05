pragma solidity ^0.7.1;

contract SimpleStorage {
    address public minter;
    constructor() {
        minter = msg.sender;
    }
    mapping (uint => string) public sidechainHash;
    function set(uint number, string memory hash) public {
        sidechainHash[number] = hash;
    }

    function get(uint number) public view returns (string memory) {
        require(msg.sender == minter);
        return sidechainHash[number];
    }
}

