## The Bitocin Lightning Network:
### Scalable Off-Chain Instant Payments 

##### Joseph Poon,  joseph@lightning.network
##### Thaddeus Dryja, rx@awsomnet.org

#### January 14, 2016 DRAFT Version 0.5.9.2

### Contents 

1. The Bitcoin Blockchain Scalability Problem
2. A Network of Micropayment Channels Can Solve Scalability
    
    2.1 Micropayment Channels Do Not Require Trust 
    2.2 A Network of Channels 

3. Bidirectional Payment Channels 

    3.1 The Problem of Blame in Channel Creation
    
    3.1.1 Creating an Unsigned Funding Transaction

    3.1.2 Spending from an Unsigned Transation
    
    3.1.3 Commitment Transactions : Unenforcible Construction

    3.1.4 Commitment Transactions : Ascribing Blame 

    3.2 Creating a Channel with Contract Revocation 

    3.3 Sequence Number Maturity 

    3.3.1 Timestop

    3.3.2 Revocable Commitment Transactions

    3.3.3 Redeeming Funds from the Channel : cooperative Counterparties

    3.3.4 Creating a new commitment Transaction nad Revoking Prior Commitments 

    3.3.5 Process for Creating Revocable Commitment Transactions 

    3.4 Cooperatively Closing Out a Channel

    3.5 Bidirectional Channel Implications and Summary 

4. Hashed Timelock Contract (HTLC)

    4.1 Non-revocable HTML Construction

    4.2 Off-chain Revocable HTLC

    4.2.1 HTLC when the Sender Broadcasts the Commitment Transaction 

    4.2.2 HTLC when the Receiver Broadcasts the Commitment Transaction

    4.3 HTLC Off-chain Termination

    4.4 HTLC Formation and Closing Order 

5. Key Storage 

6. Blockchain Transaction Fees for Bidirectional Channels 

7. Payto Contract 

8. The Bitcoin Lightning Network 

    8.1 Decrementing Timelocks 

    8.2 Payment Amount

    8.3 Clearing Failure and Rerouting 

    8.4 Payment Routing 

    8.5 Fees 

9. Risks

    9.1 Improper Timelocks 

    9.2 Forced Expiration Spam

    9.3 Coin Theft via Cracking 

    9.4 Data Loss 

    9.5 Forgetting to Broadcast the Transaction in Time 

    9.6 Inability to Make Necessary Soft-Forks 

    9.7 Colluding Miner Attacks 

10. Block Size Increases and Consensus 

11. Use Cases 

12. Conclusion

13. Acknowledgements

    Appendix A - Resolving Malleability 

    References 



## Abstract 

The  bitcoin  protocol  can  encompass  the  global  financial  transaction volume in all electronic payment systems today, without a single custodial third party holding funds or requiring participants to haveanything  more  than  a  computer  using  a  broadband  connection.   A decentralized  system  is  proposed  whereby  transactions  are  sent  over a  network  of  micropayment  channels  (a.k.a.    payment  channels  or transaction  channels)  whose  transfer  of  value  occurs  off-blockchain.If  Bitcoin  transactions  can  be  signed  with  a  new  sighash  type  that addresses  malleability,  these  transfers  may  occur  between  untrustedparties along the transfer route by contracts which, in the event of un-cooperative or hostile participants, are enforceable via broadcast overthe bitcoin blockchain in the event of uncooperative or hostile partici-pants, through a series of decrementing timelocks

## 1 The Bitcoin Blockchain Scalability Problem 

The Bitcoin[1] blockchain holds great promise for distributed ledgers, but the blockchain as a payment platform, by itself, cannot cover the world’s commerce anytime in the near future. The blockchain is a gossip protocol whereby all state modifications to the ledger are broadcast to all participants. It is through this “gossip protocol” that consensus of the state, everyone’s balances, is agreed upon. If each node in the bitcoin network must know about every single transaction that occurs globally, that may create a significant drag on the ability of the network to encompass all global financial transactions. It would instead be desirable to encompass all transactions in a way that doesn’t sacrifice the decentralization and security that the network provides. 

The payment network Visa achieved 47,000 peak transactions per second (tps) on its network during the 2013 holidays[2], and currently averages hundreds of millions per day. Currently, Bitcoin supports less than 7 transactions per second with a 1 megabyte block limit. If we use an average of 300 bytes per bitcoin transaction and assumed unlimited block sizes, an equivalent capacity to peak Visa transaction volume of 47,000/tps would be nearly 8 gigabytes per Bitcoin block, every ten minutes on average. Continuously, that would be over 400 terabytes of data per year.

Clearly, achieving Visa-like capacity on the Bitcoin network isn’t feasible today. No home computer in the world can operate with that kind of bandwidth and storage. If Bitcoin is to replace all electronic payments in the future, and not just Visa, it would result in outright collapse of the Bitcoin network, or at best, extreme centralization of Bitcoin nodes and miners to the only ones who could afford it. This centralization would then defeat aspects of network decentralization that make Bitcoin secure, as the ability for entities to validate the chain is what allows Bitcoin to ensure ledger accuracy and security. 

 Having fewer validators due to larger blocks not only implies fewer individuals ensuring ledger accuracy, but also results in fewer entities that would be able to validate the blockchain as part of the mining process, which results in encouraging miner centralization. Extremely large blocks, for example in the above case of 8 gigabytes every 10 minutes on average, would imply that only a few parties would be able to do block validation. This creates a great possibility that entities will end up trusting centralized parties. Having privileged, trusted parties creates a social trap whereby the central party will not act in the interest of an individual (principalagent problem), e.g. rentierism by charging higher fees to mitigate the incentive to act dishonestly. In extreme cases, this manifests as individuals sending funds to centralized trusted custodians who have full custody of customers’ funds. Such arrangements, as are common today, create severe counterparty risk. A prerequisite to prevent that kind of centralization from occurring would require the ability for bitcoin to be validated by a single consumer-level computer on a home broadband connection. By ensuring that full validation can occur cheaply, Bitcoin nodes and miners will be able to prevent extreme centralization and trust, which ensures extremely low transaction fees

While it is possible that Moore’s Law will continue indefinitely, and the computational capacity for nodes to cost-effectively compute multigigabyte blocks may exist in the future, it is not a certainty.

To achieve much higher than 47,000 transactions per second using Bitcoin requires conducting transactions off the Bitcoin blockchain itself. It would be even better if the bitcoin network supported a near-unlimited number of transactions per second with extremely low fees for micropayments. Many micropayments can be sent sequentially between two parties to enable any size of payments. Micropayments would enable unbunding, less trust and commodification of services, such as payments for per-megabyte internet service. To be able to achieve these micropayment use cases, however, would require severely reducing the amount of transactions that end
up being broadcast on the global Bitcoin blockchain.

While it is possible to scale at a small level, it is absolutely not possible to handle a large amount of micropayments on the network or to encompass all global transactions. For bitcoin to succeed, it requires confidence that if it were to become extremely popular, its current advantages stemming from decentralization will continue to exist. In order for people today to believe that Bitcoin will work tomorrow, Bitcoin needs to resolve the issue of block size centralization effects; large blocks implicitly create trusted custodians and significantly higher fees.

### 2 A Network of Micropayment Channels Can Solve Scalability 

“If a tree falls in the forest and no one is around to hear it, does it make a sound?”

The above quote questions the relevance of unobserved events —if nobody hears the tree fall, whether it made a sound or not is of no consequence. Similarly, in the blockchain, if only two participants care about an everyday recurring transaction, it’s not necessary for all other nodes in the bitcoin network to know about that transaction. It is instead preferable to only have the bare minimum of information on the blockchain. By deferring telling the entire world about every transaction, doing net settlement of their relationship at a later date enables Bitcoin users to conduct many transactions without bloating up the blockchain or creating trust in a centralized counterparty. An effectively trustless structure can be achieved by using time locks as a component to global consensus.

Currently the solution to micropayments and scalability is to offload the transactions to a custodian, whereby one is trusting third party custodians to hold one’s coins and to update balances with other parties. Trusting third parties to hold all of one’s funds creates counterparty risk and transaction costs.

Instead, using a network of these micropayment channels, Bitcoin can scale to billions of transactions per day with the computational power available on a modern desktop computer today. Sending many payments inside a given micropayment channel enables one to send large amounts of funds to another party in a decentralized manner. These channels are not a separate trusted network on top of bitcoin. They are real bitcoin transactions.

Micropayment channels[3][4] create a relationship between two parties to perpetually update balances, deferring what is broadcast to the blockchain in a single transaction netting out the total balance between those two parties. This permits the financial relationships between two parties to be trustlessly deferred to a later date, without risk of counterparty default. Micropayment channels use real bitcoin transactions, only electing to defer the broadcast to the blockchain in such a way that both parties can guarantee their current balance on the blockchain; this is not a trusted overlay network —payments in micropayment channels are real bitcoin communicated and exchanged off-chain.

#### 2.1 Micropayment Channels Do Not Require Trust 

### 12 Conclusion 



### 13 Acknowledgements 

### Appendix A Resolving Malleability 

### References 