# Truebit Light's Incentive System - Work in Progress - 

Christian Reitwiebner
chris@ethereum.org

## page 1

블록체인은 탈중앙화된 상태에서 실행되는 프로그램들이 정확하게 실행되는 것을 합리적으로 확신시켜준다. 거래가 승인되었다는 것을 보장하지 않는다, 또한 승인된 거래가 지속적으로 승인된 상태로 남는다는 것 역시 보장하지 않는다. 하지만 시간이 흐름에 따라서 상태가 바뀔 확률은 지속적으로 감소한다. 

블록체인의 가장 주된 문제점은 그들의 확장성이다: 시간당 발생할 수 있는 계산량은 어느정도 고정되어 있다. 더 많은 참가자가 네트워크에 참여한다고 해도 증가하지 않고, 그 네트워크에 속해 있는 노드들 중 가장 느린 노드에 제한된다. 

TrueBit 은 이 문제를 interactive verification 으로 해결하고자 한다.  It allows more or less arbitrarily complex  computations  to  be  performed  under  the  assumption  that  there  is  at  least  one  honest participant.   It  does  not  require  that  participant  to  be  altruistic,  though.   TrueBit  also  includes some drawbacks, especially the drawback that transactions usually take more time to be acceptedand they can also be delayed arbitrarily by an attacker, as long as this attacker has enough financial resources. 이 시스템은 정확도(correctness) 를 활성도(liveness) 보다 우선시 한다. 최소한 1개의 정직한 참가자가 있다면, 부정확한 computation/transaction 을 포함시키는 것은 불가능하다 하지만 공격자는 옳은 computation/transaction 을 포함시키는 것을 임의로 늦출수는 있다.

본 글에서는 Dogecoin-Ethereum bridge 에 특정하여 서술해보려고 한다. Dogecoin-Ethereum bridge는 Dogecoin 블록체인의 블록들이 Ethereum smart contract 에서 검증되어야 한다. 이 검증을 직접적으로 진행하게 되면 비용이 많이 든다. 그래서, 트루빗을 활용한 off-chain 계산을 할 것입니다. 

Having  said  that,  all  analyses  are  equally  applicable  to  bridges  between  blockchains  where  theavailability of blocks can be reasonably assumed.  This means that it can be used to e.g. offload processing volume from the main Ethereum chain to another blockchain (which might even be a proof-of-authority chain) as long as all participants in that chain rightfully assume that block datais available to all potential challengers.

Truebit-light 프로토콜이 정직한 참가자들을 네트워크로 끌어오는 것으로 생각하지 않는다.(의역) 어느 정도 이타적인 참가자가 있는 것으로 가정한다. This includes keeping up with the Ethereum network, paying the gas fee and having a certain amount of money to pay for an initial deposit.

또한 오직 1개의 병렬 작업만을 허용하는 수준으로 시스템을 간소화(simplify) 시킬 것이다. This might be extended to a constant amount of tasks but would also make the presentation here more complex

## page 2

모든 타임아웃들은 lower bound 이고 ethereum 네트워크의 혼잡을 대비하여 확장되어야만 한다. This means that if you want to claim a timeout to take effect, you have to provide a proof that several  previous  blocks  had  enough  capacity  to  include  a  potential  response  by  the  other  party. Since TrueBit never makes a claim that state transitions take effect in a finite amount of time, this is still consistent with the theory presented below.

TrueBit 컨트랙트는 다음과 같은 속성이 있습니다: 

TrueBit 은 fact claiming component 와 verification game 으로 구성되어 있다. In both cases, we fix a mathematical function f : Σn → Σn which can be implemented on a given machine in s elementary computation steps. This is not a big limitation, since f can be an interpreter for another machine, thus allowing to run arbitrary programs, which have a certain finite running time. Limiting the running time is a crucial component, although this limit can be magnitudes higher than what is possible to compute in a single block of the underlying blockchain.

We will start with describing a language that helps us treat the smart contract systems.

**Definition 0.1.** 

A 를 모든 계정의 세트로 T 를 모든 timestamp 의 세트로 설정하자. smart contract C 를 T × A × I (timestamp, sender and input) 으로 부터 input 을 받는 state machine 으로 설정하고 input 에 따라서 S 로부터 상태를 변환하고 O 으로 부터 output 값을 생성해낸다. 여기서, I, O 와 S 는 각 smart contract 타입에 따라 달라진다. smart contract 는 그 상태 변환 함수 (state transition function) C : S × T × A × I → S × O 으로 식별된다. 함수는 부분적 함수이다. (partial function)

즉, machine 은 특정 input 값에 대해서는 거부를 할 수 있다. Smart contract 는 거절되지 않은 이전 input 의 timestamp 보다 높지 않은 timestamp 를 가지고 있는 것들은 모두 거절한다. C 는 또한, iterated state transition function C : (T × A × I)∗ → S × O 으로 식별된다. 묵시적 초기 상태 S0 를 (s0, o0) = C(ε) 으로 가정한다. The iterated state transition function is then defined inductively as C(In, (t, a, i)) = C(s′, t, a, i), where (s′, o′) = C(In).

이를 간단히 표현하자면 다음과 같다.  

C: s =⇒ (i / t,a) s' | o if C(s,t,a,i) = (s' ,o) (식이 조금 이상하네)

a ∈ A 인 플레이어 a 에게는 S → (T ×I)∪{⊥} 라는 전략이 있다. (⊥ 는 플레이어가 아무런 움직임을 일으키지 않았을 때를 의미한다.) 

strategy assignment S : A → (S → (T × I) ∪ {⊥}) 에서 S 에 따른 smart contract C 의 게임 g는  일련의 움직임들이다. i.e. g ∈ (T × A × I)∗ such that C(g) is defined and g 는 g = ε (빈 게임) 또는 g = g′ ·(t,a,i) 인 상황이다. g′ 는 C에서의 게임이다. 

그런데 C 는 C(g′) = (s,o), S(a)(s) = (t,i) 이며 S(a′) = (t′,i′), t′ < t 인데 C(g′ · (t′, a′, i′)) 정의된 a' (a′ ∈ A) 는 없는 상태이다. game 의 길이는 round 의 숫자로 불리운다. 

S:A → (S → (T ×I)∪{⊥}) 에 관한 설명이다. S 에 따른 C 안에 있는 게임 g 중 어느 하나라도 C(g) = (s,o) 를 만족하는 것이 있다면 C ~ → S o 라고 적는다. 

For a single strategy function s:S→T×I∪{⊥} for a player a∈A we write C~ →s o if C~→S o for any S that satisfies S(a) = s. (이것도 화살표가 개판이 되네)

**Theorem 0.2.** 

s 스텝들을 계산해야하는 모든 함수 f 에는 interative game 이 있다. game 에는 a,b 2명의 참가자가 있으며 smart contract 에 의해 진행된다. smart contract G[a, b, ., ., .] 는 아래의 특성을 가진다.

1. it takes at most 1+2 log2 s rounds and at most tG log2 s time (assuming no network congestion) for some intra-round timeout tG

2. 모든 x, y 에게 선수 a 를 위한 전략 s가 존재한다. 
전략 s 는 G[a, b, x, f(x), y] ~->s f(x) 이다.

3. 모든 x, y 에게 선수 b 를 위한 전략 s 가 존재한다. 
전략 s는 G[a,b,x,y,f(x)] ~->s f(x) 이다. 

## page 3

*Proof.* The game will keep the invariant that both players agree on the internal state of the computation at some step l but disagree about the state at step h. Note that we can also work with hashes of internal states, so the data sent in each round is not very large.

Initiall, l = 0 and h = s and the game halves the distance h − l with every second message (round). 

t0 를 게임이 생성된 timestamp 로 설정하자. 초기 상태는 다음과 같을 것이다. 

G [a, b, x, ya, yb, t0] (ε) = (t0, (0, x), (s, ya, yb))

이후의 모든 메세지들은 state 에 있는 timestamp 보다 커야 한다, i.e. we have an implicit requirement that t > tp. Furthermore, we will omit the parameters of G in the following. We will use α for a generic accounts that can be either a or b.

If h − l > 1, we ask both participants to submit what they think is the internal state at step ⌊ h−l ⌋: 2 (식 애매하다)

(여기는 식 1,2,3이 다 개판됨)

h-l = 1 이 되면, smart contract 는 아래의 계산을 실행한다 : 

 f(s,i,p)를 알고리즘의 내부 상태라고 가정해보자. 알고리즘은 i 단계 부터 단일 단계를 실행 후 f 를 계산하며 내부 상태 s 는 보조 증명 데이터 p 를 고려한다.(만약 p 가 잘못되었거나 유효하지 않으면 값은 정의되지 않는다.)

 (식 4 추가)

 게다가, 특정 시점에 t > tp + tG 가 되면, timeout 이 선언된다. 

 (식 5,6)

 최악의 경우에 게임 round 의 숫자를 한번 계산해보자. timeout 은 어느 상태이던지 게임을 바로 끝낼 수 있다. 메세지 타입 (1) 그리고 (2),(3) 모두 h-l 을 반으로 줄인다. 타임아웃을 제하고는, h=l+1 의 상태가 될 때까지 이들이 메세지의 형태로 가능한 유일한 경우들이다. h=l+1 의 시점에서는 오로지 메세지 (4) 만이 가능하다. 

## page 4

만약 timeout 이 없다면, 게임은 1 + 2log2s 만큼의 메세지가 필요할 것이다. 

Note that the timeouts for messages of type (1), (2) and (3) all start at the same time. 이 말은 양쪽 모두 h-l 의 크기 감소를 수행하는데 있어서 tG 만큼의 시간이 걸린다는 것이다. 만약 이것이 tG 보다 오래 걸린다면, 누구든지 게임을 종료시킬 수 있다. 즉, 게임이 걸리는 최대 시간은 tG log2 s 으로 한정된다.(assuming there is an actor who will trigger the timeout).

마침내, 우리는 왜 양쪽 선수 모두가 게임을 끝내는데 있어서 f(x)를 사용하는지에 대해서 논한다. 대칭 형태이기에, 오직 a 선수에 대해서만 논하려고 한다.

Obviously, by responding in time, a can always avert the situation that the game ends with a timeout in a state different from f(x).

만약 game 의 현 state 가 tp, (l, s1), (h, sa, sb) 이라면, 전략은 ⌊ h−l / 2⌋ 단계에서 알고리즘 컴퓨팅 f 의 내부 상태를 포함한 메세지를 보내는 것이다. 그렇게 함으로써, smart contract 는 상태 tp,(l,s1),(l+1,sa,sb) 에서 끝내게 된다. 여기서 s1 은 l 단계에서의 상태이고 sa 는 l+1 단계에서의 상태이다. sb와 sa 는 같지 않고 알고리즘 컴퓨팅 f 는 결정적이기 때문에, b 는 smart contract 를 yb 의 상태로 바꾸기 위해 (4) 형태의 메세지를 사용할 수 없다. 대신 (4) 를 ya = f(x) 의 smart contract output 을 만들기 위해 사용할 수 있다. 

할 일 : 나머지 statement 들을 수식화하고 증명하기.
