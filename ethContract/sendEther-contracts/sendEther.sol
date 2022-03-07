pragma solidity ^0.8.0;

contract sendether{
    address private owner;
    mapping(address => uint) private _balances;

    constructor() {
        owner = msg.sender;
    }

    function getBalances() view public returns (address, uint) {
        return (msg.sender, _balances[msg.sender]);
    }

    function receiveEther() payable public{
        _balances[msg.sender] += msg.value;
    }

    function sendEther(address payable _address, uint value) payable public{
        require(msg.sender == owner);
        _balances[_address] -= value;
        _address.transfer(value);
    }

function receive() external payable {

}

}