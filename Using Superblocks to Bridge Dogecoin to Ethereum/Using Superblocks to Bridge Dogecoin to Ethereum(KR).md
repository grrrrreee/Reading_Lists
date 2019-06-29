# Using Superblocks to Bridge Dogecoin to Ethereum

July 31th, 2018 
Revision : 3 

Ismael Bejarano ismael.bejarano@coinfabrik.com 

Oscar Guindzberg oscar.guindzberg@gmail.com

## 개요 

The Doge->Eth bridge requires over 10.000 USD in gas (at 450 USD per ether) per day to keep up with the Doge blockchain. We propose a system where no Doge block headers are submitted to ethereum. Instead, all the headers from Doge blocks mined within the last hour are used to build a Merkle tree whose root is submitted to the Superblocks contract. Verifiers can start a challenge/response process to prove this root wrong.

## 동기

Our first approach to build a Dogecoin-Ethereum bridge was to port BTCRelay to Solidity. BTCRelay requires every new Bitcoin block header to be submitted to the contract in order to validate transactions belonging to a particular block.

For BTCRelay, storing an 80-byte block header costs around 200K gas. It needs to store 144 block headers per day. At a gas price of 10 gwei, this requires spending 0.288 ether per day in storage costs.

The success of Ethereum caused the costs to increase notably: at a valuation of 900 USD per ether, running BTCRelay for one day costs 259.2 USD. We could use a lower gas price, but the fees are higher with increased network usage and our transactions may take a long time to be mined.

Dogecoin generates one block per minute; even assuming the Doge->Eth bridge uses a challenge/response system for scrypt hash verification, this makes the cost of the Doge->Eth bridge ten times that of BTCRelay at the same block header size. In addition to this, Dogecoin is merge mined: block headers are larger, with an average size of 700 bytes, and this data requires extra validations, thus increasing the costs and making them prohibitive.

## Superblocks (슈퍼블록)

Efficiently Bridging EVM Blockchains 기사에서 영감을 받은 부분이 있다. 비용 절감 하는 방법 중 하나인데, 모든 블록 헤더를 저장하는 것이 아니라, 여러 블록들의 머클트리의 루트 값만을 저장하는 것이다. 

그래서 우리는 일정 범위의 블록들을 대표하는 "superblock(슈퍼블록)" 을 전송하고 저장하는 방법을 제안한다. 이느 아래의 정보들을 포함할 것이다: 

- 블록 해시값들로 형성된 머클트리의 루트 값 
- 블록들의 누적 난이도 값 
- 마지막 superblock 의 해시 값
- 마지막 superblock 의 타임 스탬프 값 
- 부모 superblock 의 데이터의 해시 값 

여러 개의 블록들을 그룹화하여 하나의 superblock 을 만들며 비용을 최소화 될 것이다. 매 시간 60개의 700 바이트 블록들을 전송하고, 검증하고 저장하느 대신 오직, 200 바이트가 채 안되는 1개의 superblock 만을 저장하면 된다. 

tradeoff 가 존재하기는 한다 -- 더 많은 블록들이 하나의 superblock 에 담기면, 더 많은 비용이 절감되지만 거래 하나를 relaying 하는데 더 많은 시간이 소비된다. The one hour per superblock is an arbitrary decision that will be reviewed when more data is available.

Almost no validations will be done on-chain on superblocks. Superblock들은 challenge/response 시스템에 의하여 검증될 것이다.

이전 superblock 의 정보는 superblock 들끼리 연결하기 위해서 포함되어 있다; superblock 들의 체인이라고 생각하면 될 것이다. 이전 superblock 들을 confirm 하는 것을 촉진시키고 공격 받을 것이라고 고려되는 작은 fork 들의 superblock 들은 무시하게 한다.

### Relaying transactions 

Doge 토큰을 받기 위해서, 사용자는 lock 주소로 보낸 거래의 SPV 증명 superblock 컨트랙트 뿐 아니라, 해당 거래가 속해 있는 블록의 SPV 증명이 특정 superblock 의 일부라는 것을 증명해야 한다. 

## Security Assumptions 

우리는 해당 솔루션을 구성하는 요소들(Doge 블록체인, Ethereum 블록체인, Challenge/Response 메카니즘, BtcRelay)은 안전하다고 가정한다. 

- Dogecoin 블록체인의 대다수의 마이너들은 정직하다 : 개인 혹은 특정 그룹의 채굴자들이 가장 긴 메인 Dogecoin 블록체인을 만들 정도로 많은 해시파워를 가지고 있지 않는다. This discourages long range attacks, as any possible fork will have less accumulated difficulty than the main chain.

- Ethereum 블록체인을 채굴하는 대부분의 주체는 정직하다. 

- 대규모 체인 reorg 는 일어나지 않는다. (100개 이상의)

- 그 어느 누구도 금전적인 피해를 입는 것을 원치 않는다 : 공격을 하는 이유는 이득을 얻기 위함이거나 최소한의 비용으로 피해를 입히기를 원한다. No one is going to cause a small denial of service or another type of network disruption just to lose a lot of money.

- 우리는 블록체인이 가지고 있는 고유의 문제는 해결하지 않으려고 한다. : eclipse attck

- challenge/response 메카니즘은 항상 작동하는 것으로 가정한다. 예를 들면, 적어도 1개의 정직한 검증자는 항상 온라인 인것으로 간주된다. 

- Dogecoin 은 PoW 알고리즘을 고치거나 PoS 로 변경하지 않는다. 

- Dogecoin 은 header 나 거래의 format 을 바꾸지 않는다. 

## Validation 

우리는 1시간에 1개의 Supberblock 을 제출할 것이다. (약 60개의 Dogecoin 블록을 대표한다.) 이는 그냥 추가적인 실험을 위한 임의의 결정이다. 

체인 reorg 를 피하기 위해서, superblock은 마지막 블록이 채굴되고 3시간 안에 새로운 superblock 이 채굴되지 않을 것이다. 예를 들어, 오후 2시에서 3시에 보낸 블록들이 오후 6시에 보내진 superblock 에 담기게 된다.

superblock 승인 최소 3시간 간격은 제출자가 실수로 small fork 를 보내는 것을 확률적으로 최소화한다 ; it is an arbitrary value subject to more evaluations. 

We will treat superblocks not containing exactly the main chain blocks for that period as an attack, 왜냐하면 누군가가 실수로 since the probability of someone submitting small forks by mistake after 3 hours is negligible.

superblock 들의 검증은 challenge-response 프로토콜을 사용함으로써 해결된다. submitter 와 challenger 모두 예치금을 내야 한다. 그래야 가짜 submission 을 막을 수 있기 때문이다. challenge 의 승자는 패자의 예치금을 가져가게 된다. 

제출자가 challenge 에서 승리해도 바로 예치금이 복구되지 않는다. 해당 superblock 이후에 몇 개의 추가적인 superblock 들이 검증된 후 복구된다. (See “Superblock with blocks not in the main chain” attack below).

## Superblock states

- New : superblock 이 막 도착했을 때 , 검증자는 challenge 를 걸 수 있다. 
- InBattle : 검증자가 challenege 를 걸어 놓았지만, 아직 battle 이 끝나지 않은 상태이다. 
- Semi-Approved : 제출자가 battle 에서 승리하였지만, 또 다른 superblock 이 검증되기를 기다리고 있는 상태이다. 
- Approved (Final State) : superblock 이 유효하다. 
- Invalid (Final State) : 제출자가 battle 에서 졌다. 

image 

### Superblock verification battle 

제출자(submitter) 와 도전자(challenger) will take turns to send messages to the contracts to advance the battle. A failure to reply in a timely manner will be considered as the forfeiture of the battle and the counterpart will be declared as the winner.

1. 제출자 (submitter)  : 예치금을 건다. 
2. 제출자 (submitter)  : superblock 을 보낸다, 예치금은 lock 된다. superblock 상태는 New 가 된다. 
3. 도전자 (challenger) : 예치금을 건다. 
4. 도전자 (challenger) : 도전을 건다, 예치금이 lock 된다. superblock 상태는 InBattle 이 된다. 
5. 제출자 (submitter)  : 슈퍼블록에 모든 블록 해시값들을 array 를 보낸다. (that are used to form the Merkle tree that relates them to superblock root hash). 블록 해시값들과 머클 트리 그리고 루트 해시 조합은 on-chain 에서 검증된다. 
6. 도전자 (challenger) : 블록 헤더를 요청한다. 
7. 제출자 (submitter)  : 요청된 블록 헤더와 scrypt hash 값을 전송한다. 블록 헤더는 on-chain 에서 검증된다. 
8. TrueBit interactive scrypt hash verification off-chain. 
9. 6,7,8단계는 모든 블록들이 전송될 때까지 지속적으로 반복된다. 
10. 모든 블록이 전송되면 : 
    a. 블록들의 on-chain 누적 난이도가 superblock 의 누적 난이도와 일치하는지 확인한다.
    b. superblock 이 이전 superblock 이 제출된지 3시간 후에 제출되었는지 확인한다.  
11. 만약 모든 단계가 성공적으로 작동되었고 모든 도전자들이 battle 에서 패배하면 해당 superblock 은 Semi-Approved 상태가 된다.
12. 만약 제출자가 패배하면, 해당 superblock 은 Invalid 상태가 된다. 

만약 제출자가 battle 에서 패배하면, superblock 은 즉시 유효하지 않은 것으로 판단되고 도전자는 제출자의 예치금을 상금으로 얻는다. 

만약 도전자가 포기하면, 새로운 도전자가 스스로의 예치금을 건 후에, 새로운 도전을 걸 수 있다.

만약 모든 도전자가 battle 에서 패배하면 superblock 은 semi-approved 상태가 된다. 하지만, 예치금들은 메인 체인에 들어갈 이후의 superblock 들에 의해서 검증될 때까지 lock 되어 있는다. 

새로운 superblock 들이 해당 superblock 을 조상으로 잡은 상태에서 도착하게 되거나 제안된 다른 superblock 들 보다 더 많은 축적된 PoW 가 쌓여있으면 semi-approved superblock 으로 고려될 수 있다. 

만약 superblock 이 main chain 에 속해 있지 않게 되면 제출자의 예치금은 도전자에게 지급된다. 여러 명의 도전자들이 있는 경우에는 모든 도전자들이 각자 예치금의 비율에 맞게 지급된다. 

## Possible attacks 

### Superblock submitted before 3 hours have passed 

We validate using the timestamp from the last block of the superblock. One of the reasons is to minimize the possibility of a small fork causes a honest submitter to send an invalid superblock.

This requirements also prevent when an attacker creates an arbitrarily long fork mining blocks with a timestamp in the future.

### Attacker submits superblocks while all verifiers are offline 

If all verifiers are offline for a short period of time, an attacker might submit several fake superblocks during that period.

In order to reduce the odds of success of this attack, no more than one valid superblock will be accepted every 30 minutes. This means that at least 30 minutes must have passed between a parent and a child.

This minimum time between superblocks should also discourage long range attacks. Someone can mine an arbitrarily long chain from a block in the past, but submitting will take a long time and the main chain should also grow at the same time.

### Superblock with blocks not in the main chain 

One possible attack is to send a superblock that is built in such a way that all the blocks are valid Dogecoin blocks but some of them are not in the current main chain; 예를 들어, superblock 의 마지막 블록은 공격자에 의해서 채굴된 고아 블록일수도 있다. 

If the superblock contains blocks from a temporary fork, it does not seem possible to challenge it successfully, because all the data will be a valid part of the Doge blockchain (regardless of whether it is part of the Doge main chain).

The attacker might even keep sending “fake” superblocks on top of the attack superblock to keep her “fake” chain growing. Each fake superblock may contain just 1 orphaned Doge block mined by the attacker. The cost of mining 1 orphan Doge block per hour is relatively low.

If a superblock is challenged and the submitter (i.e. the attacker) wins the battle, the superblock will be considered “semi-approved”. Deposits will be held and it will not yet be possible to use the superblock to relay transactions, but superblocks on top of it will be accepted.

Assuming the legit “superblock” chain keeps growing, after 24 superblocks (i.e. after 24 hours), the challenger can request to get her deposit back and the “semi-approved” superblock to be considered “invalid” because it is not part of the superblock mainchain. On the other hand, the submitter can request to get her deposit back and the “semi-approved” superblock to be considered “approved” if it IS part of the superblock mainchain.

## Questions and Answers 

- How are the deadlines set for each stage?

The deadlines for each stage are not defined yet. There's a trade-off -- with higher timeouts it’s easier to cause a denial of service by always replying in the last second, a lower timeout can cause an honest participant to miss replying in time.

- Do the deadlines increase on every battle turn?

deadline 은 매 battle 마다 reset 됩니다. 

- Are the deadlines fixed? 

There's no max time for the battle to resolve.

- What about new superblocks that arrive when a battle is being held?


If a superblock arrives and its parent is in battle, then the new superblock is rejected. Otherwise it is accepted and marked as "New".

For example, an attacker can challenge an honest superblock and send a competing fake superblock at the same time. It is supposed an honest challenger to challenge the fake superblock. The assumption is that the fake superblock will eventually lose the battle and the honest superblock will win.

## Fine tuning 

### Dealing with block timestamps going backwards 

Doge 블록의 timestamp 는 꼭 부모 블록의 timestamp 보다 높을 필요는 없다. (단지 최근 블록들의 중간값(median)보다만 높으면 된다.) 
그러므로, two valid competing superblocks could potentially coexist, leading to two valid competing superblock chains.
아래의 Doge 체인을 가정해보자 

B1 (16:59:58), B2 (17:00:01), B3 (16:59:59), B4 (17:00:02), B5 (18:00:01).

그리고 superblock 체인은 다음과 같다

Superblock chain A: S1 (B1), S2 (B2, B3, B4), S3(B5)
Superblock chain B: S1 (B1, B2, B3), S2 (B4), S3(B5)

2개의 superblock chain 모두 같은 누적 PoW 를 가지고 있다, 그래서 이들은 지속적으로 경쟁을 할 것이다.
이러한 문제를 피하기 위해서 ,as benevolent dictators, we establish the rule that superblock chain A is valid and chain B is invalid.
For example, if we are processing the superblock for blocks between 4pm and 5pm, no blocks are allowed with a timestamp after 5pm.
Note: The superblock might contain (most likely in the first couple of blocks) blocks whose timestamp is before 4pm.

### Dealing with long periods without blocks 

In the unlikely event that no blocks are produced during an entire hour, the superblock for that hour will be skipped.
Consider this Doge blockchain:
B1 (16:30:00), B2 (16:45:00), B3 (19:20:00).

The Superblock chain should be: S1 (B1, B2), S2 (B3).

## Conclusion and next steps 

제시된 해결책은 Doge to Ethereum 페그를 만들 때 가스 사용량을 최소화 할 수 있는 실행 가능한 방법 중 하나이다. 다음 단계는 prototype 을 만들고 추가 공격 방법에 대해서 조사하기이다.

## Acknowledgements 

Many thanks to Jason Teutsch, Sina Habibian and Sergio Lerner for providing feedback. And Catalina Juarros for proofreading and editing.

## References 

- “Efficiently Bridging EVM Blockchains”, https://blog.gridplus.io/efficiently-bridging-evm-blockchains-8421504e9ced

- “A bridge between the Bitcoin blockchain & Ethereum smart contracts”, http://btcrelay.org/

- “Reduce gas usage, in particular when adding doge block headers”, https://github.com/dogethereum/dogethereum-contracts/issues/16

- Scrypt interactive verification source code https://github.com/TrueBitFoundation/scrypt-interactive

- Doge<->Eth bridge documentation https://github.com/dogethereum/docs