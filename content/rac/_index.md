---
alwaysopen: false
date: "2023-04-24T16:42:11.812Z"
description: RAC
head: <hr/>
title: RAC
weight: 3
---

Rubix Assert Contract (RAC) are rules / guidelines to enable users to develop tokens on top of Rubix Blockchain network resulting in Secondary Layer solutions, Decentralized Applications, NFTs and much more. It is a standard by which users can mint/transfer/verify secondary tokens or NFTs transfer/transact through the Rubix validator network.

Tokens minted in the Rubix network need to adhere to rules stated in the Rubix Asset Contract. These rules/guidelines may include details on what the token structure should look like, how they can be transferred in the network, validated by the quorum and so on.


|  Key  |    Type     | Description             |
| :---: |    :----:   |  :---                   |
| 1     | int         | Type of the RAC         |
| 2     | int         | RAC version             |
| 3     | string      | DID of the creator      |
| 4     | int         | Token sequence in the current supply      |
| 5     | int         | Total token supply      |
| 6     | string      | Token created timestamp      |
| 7     | string      | Creator ID, could be linked to user ID of the application   |
| 8     | string      | Creator input, could be json string will have all detials of the token   |
| 9     | string      | Content ID, json object mapping mutiple content IDs   |
| 10    | string      | Content Hash, json object mapping mutiple content hash  |
| 11    | string      | Content Url, json object mapping mutiple content url  |
| 12    | string      | Transaction information, json object mapping mutiple content transaction information  |

