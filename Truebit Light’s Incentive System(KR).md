#Truebit Light's Incentive System - Work in Progress - 

Christian Reitwiebner
chris@ethereum.org

블록체인은 탈중앙화된 상태에서 실행되는 프로그램들이 정확하게 실행되는 것을 합리적으로 확신시켜준다. 거래가 승인되었다는 것을 보장하지 않는다, 또한 승인된 거래가 지속적으로 승인된 상태로 남는다는 것 역시 보장하지 않는다. 하지만 시간이 흐름에 따라서 상태가 바뀔 확률은 지속적으로 감소한다. 

블록체인의 가장 주된 문제점은 그들의 확장성이다: 시간당 발생할 수 있는 계산량은 어느정도 고정되어 있다. 더 많은 참가자가 네트워크에 참여한다고 해도 증가하지 않고, 그 네트워크에 속해 있는 노드들 중 가장 느린 노드에 제한된다. 

TrueBit 은 이 문제를 interactive verification 으로 해결하고자 한다.  It allows more or less arbitrarily complex  computations  to  be  performed  under  the  assumption  that  there  is  at  least  one  honest participant.   It  does  not  require  that  participant  to  be  altruistic,  though.   TrueBit  also  includes some drawbacks, especially the drawback that transactions usually take more time to be acceptedand they can also be delayed arbitrarily by an attacker, as long as this attacker has enough financial resources. 이 시스템은 정확도(correctness) 를 활성도(liveness) 보다 우선시 한다. 최소한 1개의 정직한 참가자가 있다면, 부정확한 computation/transaction 을 포함시키는 것은 불가능하다 하지만 공격자는 옳은 computation/transaction 을 포함시키는 것을 임의로 늦출수는 있다.

본 글에서는 Dogecoin-Ethereum bridge 에 특정하여 서술해보려고 한다. Dogecoin-Ethereum bridge는 Dogecoin 블록체인의 블록들이 Ethereum smart contract 에서 검증되어야 한다. 이 검증을 직접적으로 진행하게 되면 비용이 많이 든다. 그래서, 트루빗을 활용한 off-chain 계산을 할 것입니다. 

Having  said  that,  all  analyses  are  equally  applicable  to  bridges  between  blockchains  where  theavailability of blocks can be reasonably assumed.  This means that it can be used to e.g. offload processing volume from the main Ethereum chain to another blockchain (which might even be a proof-of-authority chain) as long as all participants in that chain rightfully assume that block datais available to all potential challengers.

We call TrueBit-light the protocol that does not make an effort to bring honest participants to the network.  We assume that an honest participant is present who is altruistic to a certain degree.  This includes keeping up with the Ethereum network, paying the gas fee and having a certain amountof money to pay for an initial deposit.

We also simplify the system to a degree to only allow one parallel task.  This might be extended to a constant amount of tasks but would also make the presentation here more complex

모든 타임아웃들은 lower bound 이고 ethereum 네트워크의 혼잡을 대비하여 확장되어야만 한다. This means that if you want to claim a timeout to take effect, you have to provide a proof that several  previous  blocks  had  enough  capacity  to  include  a  potential  response  by  the  other  party. Since TrueBit never makes a claim that state transitions take effect in a finite amount of time, this is still consistent with the theory presented below.

TrueBit 컨트랙트는 다음과 같은 속성이 있습니다: 

TrueBit 은 fact claiming component 와 verification game 으로 구성되어 있다. In both cases, we fix a mathematical function f : Σn → Σn which can be implemented on a given machine in s elementary computation steps. This is not a big limitation, since f can be an interpreter for another machine, thus allowing to run arbitrary programs, which have a certain finite running time. Limiting the running time is a crucial component, although this limit can be magnitudes higher than what is possible to compute in a single block of the underlying blockchain.

We will start with describing a language that helps us treat the smart contract systems.

**Definition 0.1.** A 를 모든 계정의 세트로 T 를 모든 timestamp 의 세트로 설정하자. smart contract C 를 T × A × I (timestamp, sender and input) 으로 부터 input 을 받는 state machine 으로 설정하고 input 에 따라서 S 로부터 상태를 변환하고 O 으로 부터 output 값을 생성해낸다. 여기서, I, O 와 S 는 각 smart contract 타입에 따라 달라진다. smart contract 는 그 상태 변환 함수 (state transition function) C : S × T × A × I → S × O 으로 식별된다. 함수는 부분적 함수이다. (partial function)즉, machine 은 특정 input 값에 대해서는 거부를 할 수 있다. Smart contract 는 거절되지 않은 이전 input 의 timestamp 보다 높지 않은 timestamp 를 가지고 있는 것들은 모두 거절한다. C 는 또한, iterated state transition function C : (T × A × I)∗ → S × O 으로 식별된다. 묵시적 초기 상태 S0 를 (s0, o0) = C(ε) 으로 가정한다. The iterated state transition function is then defined inductively as C(In, (t, a, i)) = C(s′, t, a, i), where (s′, o′) = C(In).

이를 간단히 표현하자면 다음과 같다.  C: s =⇒ (i / t,a) s' | o if C(s,t,a,i) = (s' ,o) (식이 조금 이상하네)

a ∈ A 인 플레이어 a 에게는 S → (T ×I)∪{⊥} 라는 전략이 있다. (⊥ 는 플레이어가 아무런 움직임을 일으키지 않았을 때를 의미한다.) strategy assignment S : A → (S → (T × I) ∪ {⊥}) 에서 S 에 따른 smart contract C 의 게임 g는  일련의 움직임들이다. i.e. g ∈ (T × A × I)∗ such that C(g) is defined and g 는 g = ε (빈 게임) 또는 g = g′ ·(t,a,i) 인 상황이다. g′ 는 C에서의 게임이다. 그런데 C 는 C(g′) = (s,o), S(a)(s) = (t,i) 이며 S(a′) = (t′,i′), t′ < t 인데 C(g′ · (t′, a′, i′)) 정의된 a' (a′ ∈ A) 는 없는 상태이다. game 의 길이는 round 의 숫자로 불리운다. 

S:A → (S → (T ×I)∪{⊥}) 에 관한 설명이다. S 에 따른 C 안에 있는 게임 g 중 어느 하나라도 C(g) = (s,o) 를 만족하는 것이 있다면 C ~ → S o 라고 적는다. For a single strategy function s:S→T×I∪{⊥} for a player a∈A we write C~ →s o if C~→S o for any S that satisfies S(a) = s. (이것도 화살표가 개판이 되네)

**Theorem 0.2.** For any function f taking s steps to compute, there is an interactive game with two participants a and b implemented by a smart contract G[a, b, ·, ·, ·] with the following properties:

1. it takes at most 1+2 log2 s rounds and at most tG log2 s time (assuming no network congestion) for some intra-round timeout tG

2. for any x and y, there is always a strategy s for player a such that G[a,b,x,f(x),y] 􏰁s f(x) (화살표)

3. for any x and y, there is always a strategy s for player b such that G[a,b,x,y,f(x)] 􏰁s f(x) (화살표)

*Proof.* The game will keep the invariant that both players agree on the internal state of the computation at some step l but disagree about the state at step h. Note that we can also work with hashes of internal states, so the data sent in each round is not very large.

Initiall, l = 0 and h = s and the game halves the distance h − l with every second message (round). 

t0 를 게임이 생성된 timestamp 로 설정하자. 초기 상태는 다음과 같을 것이다. 

G [a, b, x, ya, yb, t0] (ε) = (t0, (0, x), (s, ya, yb))

이후의 모든 메세지들은 state 에 있는 timestamp 보다 커야 한다, i.e. we have an implicit requirement that t > tp. Furthermore, we will omit the parameters of G in the following. We will use α for a generic accounts that can be either a or b.

If h − l > 1, we ask both participants to submit what they think is the internal state at step ⌊ h−l ⌋: 2 (식 애매하다)

(여기는 식 1,2,3이 다 개판됨)

If h − l = 1, the smart contract can actually perform the computation:

 Let f(s,i,p) be the internal state of the algorithm that computes f after running a single step starting from step number i and internal state s taking into account auxiliary proof data p (the value is undefined if p is malformed or invalid).

 (식 4 추가)

 Furthermore, at a certain time t > tp + tG, a timeout can be claimed:

 (식 5,6)

 Let us now analyze the number of rounds of the game in the worst case. Note that timeouts (i.e. mesages of type (5) or (6)) can directly end the game from any state. A message of type (1) followed by either (2) or (3) reduce h − l roughly by half. Apart from timeouts, these are the only messages possible until h = l + 1. At that point, only message (4) is possible.

만약 timeout 이 없다면, 게임은 1 + 2log2s 만큼의 메세지가 필요할 것이라는 것을 의미한다. 

Note that the timeouts for messages of type (1), (2) and (3) all start at the same time. This mean that both parties have tG time to perform the magnitude reduction of h − l. If this takes longer than tG, anyone can step in and end the game. This means that the game will take at most tG log2 s time (assuming there is an actor who will trigger the timeout).

Finally, we argue why both players have a strategy to end the game with f(x). Due to symmetry, we only argue for player a.

Obviously, by responding in time, a can always avert the situation that the game ends with a timeout in a state different from f(x).

If the current state of the game is tp, (l, s1), (h, sa, sb), the strategy is to send a message that contains the internal state of the algorithm computing f at step ⌊ h−l / 2⌋. In doing so, the smart contract will
end up with a state tp,(l,s1),(l+1,sa,sb) where s1 is the state at step l and sa is the state at step
l + 1. Since sb ̸= sa and the algorithm computing f is deterministic, b cannot use a message of type (4) to turn the smart contract into state yb. Instead, a uses (4) to make the smart contract output ya = f(x).

TODO: Formulate rest of statements and prove them.
