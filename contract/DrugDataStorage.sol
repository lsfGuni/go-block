// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DrugDataStorage {
    struct DrugData {
        uint256 seq;
        bytes32 hashCode;
    }

    mapping(uint256 => DrugData) public storedData;

    function storeData(uint256 seq, bytes32 hashCode) public {
        storedData[seq] = DrugData(seq, hashCode);
    }

    function getData(uint256 seq) public view returns (uint256, bytes32) {
        DrugData memory data = storedData[seq];
        return (data.seq, data.hashCode);
    }
}
