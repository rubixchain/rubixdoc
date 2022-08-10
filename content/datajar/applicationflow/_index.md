---
alwaysopen: false
date: "2022-08-07T10:42:11.812Z"
description: Overview of how Rubix Data Jar is connected with Rubix Connect and how enterpise can build application over it
head: <hr/>
title: Application Flow
weight: 4
---
Rubix Datum has a set of request calls for committing the transactions to Rubix blockchain. Rubix Datum are served by Rubix nodes which in turn processes and submit the block of transactions to the Rubix blockchain. All transactions to Rubix nodes are committed by Application as batch (multiple files, event, log etc.). The Batch Size will be decided by the Application using the overall batch size or predefined duration.

<img src="https://raw.githubusercontent.com/rubixchain/learn/main/static/images/datumChain.png">

To initiate a data commit, Batch Start action is started in which the request will check the status of the node, then initiate a batch process and return a Batch ID with success message if it is not processing any other batch at the point of time. If it is already processing another batch of transactions it will return busy, and failure messages if any error occurs.

Once a batch is initiated then the user can create a transaction under this created batch and then Create Transaction request that will take Batch ID, transaction info as either a file or a Json string as input. The Committer can submit a transaction or list of transactions and the request will return a success or failure message based on the process. All these transactions are listed under the batch that committers have created in the previous Batch Start request. Once a transcation is created we can go ahead with Commit Transcation request. Commit transcation request will take Batch ID as the input and commits the all the transactions listed under the batch from the node to Rubix Chain as transaction metadata file, the request will return success message on successful submission of Commit Transcation request along with metadata file or it will return failure message.

Datum commit on Rubix blockchain is backed by 1 RBT, i.e. each batch needs a RBT to commit the data/transaction into the chain. For a single token once it has reached the limit of 256 commits that RBT token has to do at least one P2P transaction in-order to move ahead with more commits.

{{% children description="true"   %}}