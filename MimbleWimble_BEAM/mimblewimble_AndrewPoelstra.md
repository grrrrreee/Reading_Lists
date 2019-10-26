# Mimblewimble 

#### Andrew Poelstra
#### 2016-10-06(commit e9f45ec)

### Abstract

tom elvis 이름을 사용하는 익명의 사람이 비트코인 조사 채널 IRC에 서명을 하여 Tor hidden service[Jed16] 호스트된 문서를 삭제하였다 Mimblewimble이라는 제목의 이 문서는 비트코인과 근본적으로 다른 거래 구조를 가진 블록체인을 묘사했다. 단절적 합병 와 거래중단, 비밀거래 그리고 새유저들이 코인의이력 증명을 요구하지 않는 현재 chainstate의 완벽한 증명 을 지원한다 불행하게도 그 논문은 주요 아이디어들을 간결하게 적혀 있지만 그것은 보안의 대한 논쟁도 없고 심지어 실수도 하나 있다. 이 논문의 목적은 원래의 아이디어를 구체화 시키는 것이고, 저자가 개발한 확장 개선사항을 추가하는 것이다 특히 Mimblewimbled은 비트코인 기록이 있는 체인의 모든거래 을 기록하는데 15기가 데이터밖에 안 든다 ( 범위증명이 포함되고 100기가 가 필요한 UTXO set은 포함시키지 않는다) Jedusor는 이것을 어떻게 줄일 것인가에 대한 문제를 열어뒀다 우리는 이것을 해결하고, 15Gb를 1메가바이트 이하로 줄이기 위해 작업 증명 블록체인 압축에 대한 연구와 결합한다

License. This work is released into the public domain.

### contents

#### 1 Introduction 
##### 1.1 Overview and Goals 
##### 1.2 Trust Model

#### 2 Preliminaries
##### 2.1 Cryptographic Primitives 
###### 2.1.1 Standard Primitives 
###### 2.1.2 Sinking Signatures
##### 2.2 Mimblewimble Primitives

#### 3. The Mimblewimble Payment System
##### 3.1 Fundamental Theorem
##### 3.2 The Blockchain
##### 3.3 Consensus
###### 3.3.1 Block Headers
###### 3.3.2 Proven and Expected Work
###### 3.3.3 Sinking Signatures and Compact Chains 

#### 4. Extensions and Future Research

#### 5. Acknowledgements


### 1. Introduction 

2009 년 사토시 나카모토는 디지털 토큰의 p2p을 가능케하는 온라인 화폐시스템, 비트코인이라는 암호화폐를[Nak09](https://bitcoin.org/bitcoin.pdf) 세상에 소개했다. 이 토큰들의 소유권은 공개키를 통해서 ascertain 되고 전송은  UTXO라고 불리우는 전역적으로 복제된 거래의 리스트, 공개 원장의 mean 으로 성취되며 그리고 and create new ones of equal value but different keys. 

모든 비트코인은 코인베이스 거래에서 생성됩니다. 코인베이스 거래는 기존의 코인을 기반으로 새로운 출력값을 뽑아내지 않고 and which are limited in value to maintain a fixed inflation schedule. 한번 생성되면, they change hands by means of ordinary transactions. 이 말은 시스템이 문제없이 작동하고 있다는 것(absent theft or illegal inflation)을 검증하기 위해서 새로운 유저들은 반드시 2009년 제네시스 블록부터 지금까지 있는 모든 내역을 다운로드 받아서 검증해야 한다는 이야기이다. 다른말로, 그들은 모두 다운받아서 거래를 "재개"해야한다는 것이다. 

Today the Bitcoin blockchain sits just shy of 100Gb on the author’s disk. Had Bitcoin used Con-fidential Transactions[Max15] (CT), each of the approximately 430 million outputs would consume2.5Kb each, totalling over a terabyte of historical data.

이 백서에서는, 비트코인보다 더 높은 확장성과 개인 정보보호를 향상시킨 Bitcoin-like 한 암호화폐에 대해서 연구한 프로젝트에 대해서 설명합니다. 특히, 대부분의 블록체인 내역과 사용된 거래 출력값, 범위증명 등을 제거했음에도 여전히 사용자들은 체인을 검증할 수 있게 된다는 점이 큰 특징이다. 실제로, 블록체인 내역 데이터 대부분을 갖고있지 않아도, 시스템은 작동할 수 있다.

Further, this currency allows transactions to be noninteractively combined (a la Coinjoin[Max13a])and cut-through[Max13b], eliminating much of the topological structure of the transaction graph. Ifthis is done without publishing the original component transactions, the result is an enormous boonto user privacy, amplified by CT.

Unfortunately, it accomplishes these goals at cost of sacrificing functionality. 

#### 1.1 Overview and Goals

Mimblewimble 은 내역 자체를 compact 하게 줄여서 수년간의 체인 운영 후에도 trivial한 컴퓨팅 하드웨어로도 빠르게 검증할 수 있게 설계된 암호화폐이다. 두번째로 goal 로, 이는 confidential 거래와 난독하된 거래 그래프를 통해서 강력한 사용자 privacy 를 지원해야 한다. 그러나 이들을 지원하기 위해서는, Mimblewimble 은 비트코인처럼 일반적-목적의 스크립팅 시스템을 사용할 수 없다. This precludes such functionality as zero-knowledge contingent payments[Max16], cross-chain atomic swaps[Nol13] and micropayment channels[PD16]. 이러한 기능들을 emulate 하기 위해서는 추가적인 연구가 필요하며 이는 섹션 4에서 논의될 예정이다.

더 정확하게는,

• 시스템 상에서 가치의 전송이 한 주체(들)에서 다른 주체(들)간에 직접적으로 이루어져야 한다. 

• 모든 거래들은 confidential 거래를 사용하여 출력값 양을 가려야만 한다. 

• Transactions should be noninteractively aggregable[Mou13], i.e. support a noninteractive variant of CoinJoin[Max13a, Max13b], such that parties not privy to the original transactions are unable to separate them. This can improve censorship resistance and privacy, though it is unclear how to design a safe peer-to-peer network capable of exploiting this ability.

• For a new participant, the amount of bandwidth and processing power needed to catch up with the system should be proportional to the current state of the system, i.e. the size of the UTXO set plus rangeproofs, rather than the total size of all historical transactions.

다음 섹션 그리고 더 자세히는 섹션 3.3에서 잘 설명되어있지만 Mimblewimble 은 블록 체인 내역을 to a size polylog in the original size 으로 줄이는 스킴이 있다. which in practice for a system with Bitcoin’s scale should allow a decade or more of transaction history to be verified in several seconds using a couple megabytes of data.

On the other hand, Bitcoin’s current state of 40 million unspent outputs would grow to 100Gb and a few days of verification on a modern CPU, thanks to Mimblewimble’s use of confiden- tial transitions. It is an open problem how to improve this. We observe that even without cryptographic improvements, it would go a long way to simply cap the UTXO set size as a function of blockheight and let the transaction-fee market take care of it.

#### 1.2 Trust Model

비트코인 처럼, Mimblewimble 은 유저가 언제든지 시스템에 언제든지 참가하거나 떠날 수 있는 탈중앙화 방식으로 인스턴스화된 블록체인 기반의 암호화폐이다. 게다가, 네트워크에 (재)가입할 때, 사용자들은 최근 블록체인 상태를 파악할 수 있어야 하고, 제 3 기관을 신뢰하지 않고도 해당 상태가 일련의 유효한 상태변환의 결과물이라는 것을 검증할 수 있어야 한다. 

그러나, 비트코인의 신뢰 모델에는 2가지 departure 가 있다 : 

1.  Unlike Bitcoin, whose blockchain describes every transaction in its entirety, and thereforeallows all users to agree on and verify the precise series of transactions that led to the currentchainstate, Mimblewimble only allows users to verify the essential features:

- A transaction, once committed to the block, cannot be reversed without rewriting theblock (or by the owner(s) of its outputs explicitly undoing it).

- The current state of all coins was obtained by zero net theft or inflation: there are exactlyas many coins in circulation as there should be, and there for each unspent output there exists a path of transactions leading to it, all of which are committed in the chain andauthorized. Note that there may be other paths which have also been committed in the chain, during which  some  transactions  may  have  been  invalid  or  unauthorized  or  inflationary,  butsince a legitimate path exists, all these things must have netted out to zero.

2.  Like Bitcoin, verifiers may see multiple differing blockchains, and select the valid one as theone with the greatesttotal work. This is detailed in Section 3.3.2.

However,  in  Bitcoin  the  total  work  represents  both  theexpected  workto  produce  such  ablockchain as well as theproven workof the chain, in the sense that any party who expendssignificantly less than this much work will be unable to produce such a chain except withnegligible probability.

Mimblewimble, on the other hand, separates these.  The total work still represents the ex-pected  work  to  produce  the  blockchain,  and  therefore  incentivizes  rational  actors  to  contribute to the most-work chain rather than rewriting it. The proven work, on the other hand, iscapped at some fixed value independent of the length of the chain, and serves to make forgeryby irrational lucky actors prohibitively expensive.


### 2. Preliminaries 

#### 2.1 Cryptographic Primitives 

**Groups.** Throughout, G1,G2 will denote elliptic curve groups adorned with an efficiently computable bilinear pairing ˆe:G1×G2→GT, withGTequal to the multiplicative group ofFqkfor someprimeq, small positive integerk. We further require bothG1andG2to have prime orderr. LetG,Hbe fixed generators ofG1whose discrete logarithms relative to each other are unknown (i.e.they arenothing-up-my-sleeve (NUMS) points;G2does not need any canonical generators.  We will make90computational hardness assumptions about these groups as needed.  We writeZrfor the integersmodulor, and write+for the group operation in all groups.

All cryptographic schemes will have an implicit GenParams(1λ)phase which generatesr,G1,G2,GT,GandHuniformly randomly given a security parameter λ.

##### 2.1.1 Standard Primitives 

**Definition 1.** commitment scheme 은 한 쌍의 알고리즘 Commit(v,r)-> g , Open(v,r,C) -> {true.false} 이다. 여기서 Open(v,r Commit(v,r)) 은 Commit 의 도메인 안에 있는 모든 (v,r) 에서 만족해야 한다. 그리고 아래의 security properties 들을 만족해야 한다. 

* Binding. The scheme is binding if for all(v,r) in the domain of Commit, there is no (v′,r′)!=(v,r) such that Open(v′,r′,Commit(v,r)) accepts. It is computationally binding if no PPT algorithm can produce such a (v′,r′) with nonneglible probability.

* The scheme is (perfectly, computationally) hiding if the distribution of Commit(v,r) for uniformly random r is (equal, computationally indistinguishable) for different values of v.

**Definition 2.** We define a homomorphic commitment scheme as one for which there is a group operations on commitments and Commitis homomorphic in its value parameter. That is, one where commitments  to  v,  v′can  be  added  to  obtain  a  commitment to v+v′having the same security properties.

**Example 1.** Define aPedersen commitmentas the following scheme: Commit :Z^2r → G maps(v,r)to vH+rG, and Open : Z^2r×G → {true,false} checks that Commit(v,r)equals vH+rG. 
If the discrete logarithm problem inGis hard, then this is a computationally binding, perfectly110hiding homomorphic commitment scheme[Ped01].

**Definition 3.** Given a homomorphic encryption C, we define arangeproofon C as a cryptographicproof that the committed value of C lies in some given range[a,b].  We further require that range-proofs are zero-knowledge proofs of knowledge (zkPoK) of the opening information of the commitments. 

Unless otherwise stated, rangeproofs will commit to the range[0,2n]wherenis small enoughthat no practical amount of commitments can be summed to produce an overflow.We may use, for example, the rangeproofs described in [Max15] for Pedersen commitments,which satisfy these criteria
 
##### 2.1.2 Sinking Signatures

This brings us to the first new primitive needed for Mimblewimble.

**정의 4.** sinking signature 은 ***아래의 알고리즘 집단으로 정의될 수*** 있을 것이다: 

- Setup (1 <sup> λ </sup>) 은 키 쌍 (sk, pk)의 결과값을 낸다 ; 
- Sign(sk, h) 는 ***비밀키 sk 와 음수가 아닌 정수 "높이" h ~~~*** 
- Verify(pk,h,s) 는 ***공개키*** pk와 ***서명 s*** 를 취하고 {참, 거짓} 형태의 결과값을 낸다.

***아래의 correctness 와 security 요소를 만족해야 한다:***
- ***Correctness 모든 다항식 h, (sk,pk)←Setup(1 <sup> λ </sup>), s←Sign(sk,h)은 Verify(pk,h,s)을 만족한다*** 
- ***Security (·,pk)←Setup(1<sup>λ</sup>)이라고 해보자. 주어진 pk와 오라클 H~~~***

"sinking signature"라는 이름은 블록 높이 ***h***에 있는 서명 ***s***가 주어진 상태에서 forger가 서명 ***s'***를 ***h′≤h***의 조건을 만족하는 블록 높이 ***h'***에서 같은 공개키로 형성할 수 있어 서명의 "높이의 감소"가 가능다하는 그 사실에 기반하여 지어진 이름이다. 나중에 이 기능은 밈블윔블 기반의 체인이 확장성을 확보하는데에 도움을 준다. 

#### 2.2 Mimblewimble Primitives



### 3. The Mimblewimble Payment System

#### 3.1 Fundamental Theorem

**정리2.** (***밈블윔블의 기본적인 정리***) ***Suppose we have a binding,  hiding ho-momorphic commitment scheme. Then no algorithm A can win at the following game against a challengerCexcept with negligible probability:***

***1. (Setup.)C computes a finite list L of uniformly random homomorphic commitments. C sends L to A.***

***2. (Challenge.) At most polynomially many times, A selects some (integer) linear combination Ti of L and requests the opening of this combination from C. C obliges.***

***3. (Forgery.) A then chooses a new linear combination T which is not a linear combination of {Ti} and reveals the opening information of T.***

***Proof.*** Consider the lattice Λ of formal linear combinations generated by L, that is, the set {∑A∈LbA A: bA∈Z}. Consider the quotient lattice Λ/Γ where Γ is the sublattice of Λ generated by the queries {Ti}. We may consider every element of Λ/Γ to be a homomorphic commitment by using some canonical representative. In particular, every element of Λ/Γ is a homomorphic commitment whichsatisfies the same hiding/blinding properties that the original scheme did.

Then the projection of every{Ti}intoΛ/Γis zero, and the projection of T is nonzero. Therefore A has learned no information about the projection of T from its queries; however, if A knows the opening of T then it also knows the opening of the projection of T, contradicting the hiding property of the homomorphic commitment scheme.

#### 3.2 The Blockchain

Mimblewimble은 consists of a blockchain, which is a weighted-vertex directed rooted tree of blocks(defined below) 가장 높은 가중치를 보이는 경로가 합의 내역으로 간주되는 ~~~

**정의 14.** 밈블윔블에서의 블록은 ***아래와 같이 정의될 수 있다:***

- ***2가지 종류의 형태를 입력값으로 가진 canonical transaction:***

- ***이전 블록의 거래의 출력값을 참조하거나 ; 또는***
- ***An explicit (i.e. with zero blinding factor) input restricted by rules above those given in this paper <sup>5</sup>***

***만약 블록의 거래들이 유효하다면, 블록이 유효하다고 할 수 있을 것이다.***

섹션 3.3에서는, 블록들에 가중치가 어떻게 산정되고 어느 이전 블록들에 연결되어야 하는지에 대해서 설명한다. 지금으로선, 14번 정의가 유효함의 정의로 사용되나, commitments 추가적인 요구사항이 있을 것이라는 것을 명심하라.
~~

추가 내용

#### 3.3 Consensus

Mimblewimble 은 블록체인상의 모든 블록에 난이도라는 이름의 가중치가 라벨된 해시캐시 스타일의 블록체인을 사용한다. 블록체인은 모든 난이도 D 를 가진 모든 블록이 해시 전체 크기의 1/D 범위 크기안으로 헤더 해시값의 결과를 낸다면, 유효하다고 할 수 있다. 만약에 A 가 B 의 헤더에 커밋하고 각 블록이 그들의 고유 부모에게 커밋하는 경우, 우리는 이를 directed edge 라고 정의한다. (?)

그렇다면 이제 섹션 3.2 의 컨센서스 내역(consensus history)을 가장 높은 가중치를 가진 경로로 재정의 할 수 있을 것이다. 이는 비트코인에서의 설계와 동일합니다.

##### 3.3.1 Block Headers

However, in we also define a second graph structure on blocks as vertices, called the compact blockchain. We define the compact blockchain iteratively as follows. 

1. 제네시스 블록은 compact blockchain 에 속해있다.
2. 

##### 3.3.2 Proven and Expected Work



##### 3.3.3 Sinking Signatures and Compact Chains 




### 4. Extensions and Future Research

**Multisignature Outputs.** CT 범위 증명이 슈노 서명이 다중서명 출력값을 생성했듯이 interactive 하게 생성될 수 있다는 것을 관찰하였다. 비슷하게 sinking signature 역시 다자간 방식으로 간단하게 생성될 수 있다. 그래서 다자간 서명을 지원하기 위해서, 본 논문에서는 설명되지 않았지만, 단순 지갑의 문제이므로 시스템을 더 이상 변경시킬 필요는 없다. 

**Payment channels.** 비트코인의 스크립트 시스템은 off-chain payment 채널을 지원한다, 이는 라이트닝 네트워크에서 사용되는데 한정된 블록체인 공간에서 제한없는 거래 처리량을 지원한다. Mimblewimble 과 같은 확장성 중심의 블록체인 설계는 그러한 scheme 을 지원한다.(?) It is an open problem to produce payment channels that are as ergonomic and flexible as those in Bitcoin, but to show that this ought to be possible, here is an example of a primitive Mimblewimble payment channel. Suppose that a party A wants to send a series of payments to party B of maximum value V.

1.  First A creates  a  spend  of V coins  to  a  multisignature  output requiring both A’s and B’s signature to sign. A “timelocks” this by taking the highest 0 bitiin the current blockheighth, then asking B to sign a transaction with height h+2^i returning the coins to A. Only after receiving this signature, A publishes the spend.

The signing signature construction ensures that such a refund transaction cannot be spent inany block between hand h+2^i.  On the other hand, this means that ifAwants a locktime of 2^i blocks he must wait for a blockheight that is a multiple of 2^(i+1) to create it.

2.  Now that all V coins are in a multisignature output requiring both A’s and B’s signatures tospend (with the coins going back to A after 2^i blocks in case of stall),A can send an amount v to B by  signing a transaction from this output sending v to B and (V−n) to A.

3.  To increase the valuev,Asimply signs a new transaction with the new valuev.Bwill notsign and publish until the channel is near-closing, at which pointBcan publish one trans-action taking the whole valuev, even if it was actually produced over the course of manyinteractions.16


### 5. Acknowledgements

저자는 두 분께 특별히 감사의 인사를 보내고 싶습니다. 

• (저자의 기존 설계를 부수고) 해당 논문에 묘사된 sinking signature 구조를 설계한 Avi Kulkarni 

• 합의 구조의 명확한 그림이 나타날 때까지 Mimblewimble 의 보증에 대한 다양한 질문을 한 Gregory Sanders 

### References 

[Bac02]A. Back,Hashcash — a denial of service counter-measure, 2002,http://hashcash.org/papers/hashcash.pdf.

[BCD+14]  A.  Back,  M.  Corallo,  L.  Dashjr,  M.  Friedenbach,  G.  Maxwell,  A.  Miller,  A.  Poel-stra, J. Tim ́on, and P. Wuille,Enabling blockchain innovations with pegged sidechains,2014,https://www.blockstream.com/sidechains.pdf.

[CDG+15]  D.  Cabarcas,  D.  Demirel,  F.  G ̈opfert,  J.  Lancrenon,  and  T.  Wunderer,An  uncondi-tionally hiding and long-term binding post-quantum commitment scheme, CryptologyePrint Archive, Report 2015/628, 2015,http://eprint.iacr.org/2015/628.

[Jed16]T.E.Jedusor,Mimblewimble,2016,Defuncthiddenservice,http://5pdcbgndmprm4wud.onion/mimblewimble.txt.Redditdiscussionathttps://www.reddit.com/r/Bitcoin/comments/4vub3y/mimblewimble_noninteractive_coinjoin_and_better/.

[KLS16]A. Kiayias,  N. Lamprou,  and A.-P. Stouka,Proofs of proofs of work with sublinearcomplexity, Financial Cryptography and Data Security:  FC 2016 International Work-shops, BITCOIN, VOTING, and WAHC, Christ Church, Barbados, February 26, 2016,Revised Selected Papers (Berlin, Heidelberg) (J. Clark, S. Meiklejohn, Y.A.P. Ryan,D.  Wallach,  M.  Brenner,  and  K.  Rohloff,  eds.),  Springer  Berlin  Heidelberg,  2016,pp. 61–78.

[Max13a]    G.  Maxwell,CoinJoin:  Bitcoin  privacy  for  the  real  world,  2013,  BitcoinTalk  post,https://bitcointalk.org/index.php?topic=279249.0.

[Max13b],Transaction cut-through, 2013, BitcoinTalk post,https://bitcointalk.org/index.php?topic=281848.0.

[Max15],Confidential  transactions,  2015,  Plain  text,https://people.xiph.org/~greg/confidential_values.txt.

[Max16],Thefirstsuccessfulzero-knowledgecontingentpayment,2016,Blogpost,https://bitcoincore.org/en/2016/02/26/zero-knowledge-contingent-payments-announcement/.

[Mou13]Y. M. Mouton,Increasing anonymity in Bitcoin ... (possible alternative to Zerocoin?),2013, BitcoinTalk post,https://bitcointalk.org/index.php?topic=290971.

[Nak09]S. Nakamoto,Bitcoin:  A peer-to-peer electronic cash system, 2009,https://www.bitcoin.org/bitcoin.pdf.

[Nol13]T. Nolan,Re: Alt chains and atomic transfers,https://bitcointalk.org/index.php?topic=193281.msg2224949#msg2224949, 2013.

[PD16]J.  Poon  and  T.  Dryja,The  bitcoin  lightning  network,  2016,https://lightning.network/lightning-network-paper.pdf.

[Ped01]T. Pedersen,Non-interactive and information-theoretic secure verifiable secret sharing,Lecture Notes in Computer Science576(2001), 129–140.