## The Bitocin Lightning Network:
### Scalable Off-Chain Instant Payments 

##### Joseph Poon,  joseph@lightning.network
##### Thaddeus Dryja, rx@awsomnet.org

#### January 14, 2016 DRAFT Version 0.5.9.2

*를 하는 것은 재해석이 필요한 것

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

Bitcoin[1]은 그 분산원장으로서 엄청난 잠재력과 가능성을 지니고 있다. 하지만, 결제 플랫폼으로서는 가까운 미래에 전세계적인 수준의 거래양을 소화해낼 수는 없다.* 블록체인은 가십 프로토콜을 사용한다. 가십 프로토콜이라는 것은 모든 참여자들에게 원장에 상태 변화가 일어났다면 이를 알린다는 것이다. 대표적으로 모든 이의 잔고와 같은 상태에 대한 합의는 "가십 프로토콜"을 통하여 이루어진다. 만약 비트코인 네트워크의 모든 노드가 전역적으로* 발생하는 모든 각 거래의 정보를 전달 받는다면, 모든 전세계적인 금융 거래를 감당해내는 네트워크의 능력에 지장을 초래한다. 하지만 동시에 네트워크 자체가 제공하는 탈중앙화와 보안을 희생하지 않은 채, 모든 거래를 다룰 수 있는 방법을 찾는다면 이것이 바람직한 것일 것이다. 

결제 네트워크 비자는 지난 2013년 연휴기간 동안 초당 47,000개의 거래를 감당해냈고[2], 현재는 평균 하루에 수억건의 거래를 다루고 있다. 현재, 비트코인은 초당 7개 이하의 거래를 지원하고 동시에 블록의 크기를 1MB로 제한한다. 만약에 거래 1개당 용량을 300 바이트라고 가정하고 블록의 크기에 제한이 없다고 한다면 비자와 같은 수준의 성능 즉, 초당 47000개의 거래를 비트코인이 처리해야한다면 10분당 8기가 바이트의 블록이 생성될 것이다. 이를 연단위로 계산해보면 400테라 바이트가 된다. 

비자 수준의 능력을 비트코인 네트워크가 해낸다는 것은 사실상 불가능하다는 것은 명확하다. 세상 그 어느 가정용 컴퓨터도 이 정도 수준의 대역폭과 용량을 감당해낼 수 없다. 만약 비트콩닝 미래에 비자 뿐 아니라 모든 전자 결제를 대체한다면 비트코인 네트워크는 완전하게 붕괴할 것이고 기껏해야 엄청난 비용을 감당해낼 수 있는 자만이 노드를 운용하고 채굴을 하는 극도의 중앙화될 것이다. 하지만 비트코인이 중앙화되면 탈중앙화인 상태에서 각 주체들이 체인을 검증하여 원장의 정확성과 보안을 보장받는 장점이 상쇄된다. 

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