//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;

contract Nuki {

    string public name = "Nuki home";
    uint16 homeCounter = 0;
    
    mapping(address => address[]) allowedVisitors; // home => visitors

    event NewVisitor(address home, address visitor);

    function addVisitor(address visitor) public {
        allowedVisitors[msg.sender].push(visitor);
        emit NewVisitor(msg.sender, visitor);
    }

    function removeVisitor(address visitor) public {
        for (uint i = 0; i < allowedVisitors[msg.sender].length; i++) {
            if(allowedVisitors[msg.sender][i] == visitor) {
                allowedVisitors[msg.sender][i] = address(0x0);
                break;
            }
        }
    }

    function isSenderAllowedToEnter(address home) public view returns (bool) {
        for (uint i = 0; i < allowedVisitors[home].length; i++) {
            if(allowedVisitors[home][i] == msg.sender) {
                return true;
            }
        }
        return false;
    }

}