## How to scan the latest pairs added to Sushiswap ?

출처 : https://medium.com/coinmonks/how-to-scan-the-latest-pairs-added-to-sushiswap-4e5fee7880e2

Update: for those interested on having the latest pairs for Uniswap -> go Here

### What is Sushiswap ?

SushiSwap은 이더리움에서 작동하는 소프트웨어로 이를 작동하게 하기 위해 암호 자산을 사고 파는 네트워크 사용자들에게 인센티브를 주려고한다.  Similar to platforms like Uniswap and Balancer, SushiSwap uses a collection of liquidity pools to achieve this goal. Users first lock up assets into smart contracts, and traders then buy and sell cryptocurrencies from those pools, swapping out one token for another.

수많이 생성되는 DeFi 플랫폼 중 하나인 SushiSwap은 중앙화된 운영주체 없이 암호자산을 거래할 수 있게 해준다.

Sushiswap 소프트웨어와 관련된 결정은 그 토종 암호화폐 SUSHI를 보유한 보유자에 의해 결정된다. 누구든 소유자는 새로운 안건을 낼 수 있고 다른 이들이 제안한 안건에 대해서 투표할 수 있다.

image

서로 다른 사용자간 각기 다른 암호 자산을 사고 파는 것을 가능케하여 기존의 거래소처럼 작동하게 하는 것이 가장 핵심 기능이다. 중앙화된 주체의 지원에 의해서 이루어지는 것이 아니라, SushiSwap은 스마트 컨트랙트에 의해서 거래가 실행된다. 사용자는 암호 자산을 다른 거래자들이 접근할 수 있는 소프트웨어에 lock 해놓는다.

lock 되어있는 자산을 거래하는 사람은 거래할 때 일정량의 수수료를 지불한다. 그 수수료는 유동성을 공급한 이들에게 pool에 기여한 그 비율에 따라 나뉘어진다.

이게 Sushiswap의 탈중앙화 거래소의 모습이다: 

image

암호자산 데이터 분석 이용자들은 Sushiswap이나 Uniswap에서 가장 마지막으로 거래된 토큰을 찾는 방법에 대해서 관심을 갖게 되었다. 주로 거래 분석용 툴로 사용되었을 것이다.

그래서, 필자가 구글 시트 템플릿을 하나 만들었다. 이 템플릿으로 2개의 탈중앙 거래소에서 거래 가능한 새로운 코인들을 확인할 수 있다. 

ACCESS LIVE TEMPLATE SHEET HERE

The sheet returns all new tradable pairs on Sushiswap and Uniswap, giving constraints on the Number of Days the pair has been active, the Volume ($), the Liquidity ($), and the number of Transactions.

image

In order to get Sushiswap’s analytics I used The Graph which is an indexing protocol for querying networks like Ethereum and IPFS. 누구나 usbgraph라는 Open API를 사용할 수 있고, 만들거나 배포할 수 있다. 이를 통해서 데이터를 더 쉽게 접근할 수 있게 만들 수 있다.

image

SUSHISWAP FUNCTION IN GOOGLE SHEETS:
코인이 활성화된 일수, 규모($), liquidity($) 그리고 거래 숫자를 기반으로 Sushiswap에서 새롭게 거래될 수 있는 pair를 반환한다.

image

예를 들어, 새로운 Sushiswap pair를 보고 싶다면 아래와 같은 조건을 만족하면 된다 :  
- 지난 4일 안에 pool이 열려야 한다. 
- 하루 거래량이 $5,000보다는 커야 한다. 
- Liquidity는 $10,000보다는 커야 한다. 
그리고 런칭이 된 이후에 200개 이상의 거래가 이뤄졌어야 한다. 

식은 다음과 같이 작성하면 된다: =SUSHISWAP(4,5000,10000,200)

@param {days} the number of Days since the pair is active
@param {volume} the minimum Volume ($)
@param {liquidity} the minimum Liquidity ($)
@param {tx_count} the number of Transactions existant since creation

* @return a table (see GIF above)with all new tradable pairs on Uniswap and their number of Days since Active, the Volume ($), the Liquidity ($), the number of Transactions.

### More indicators for scanning ?

There are plenty more functionalities that can be added through the TheGraph API. Don’t hesitate to have a look at all available end points like:

image 

- totalSupply
- untrackedVolumeUSD
-  liquidityProviderCount

If you are interested in getting some help in integrating more personalized indicators, DM me.

### Conclusion

An easy accessible way to get the latest pairs on Sushiswap using Google Sheets which can be used for data analytics, and as a trading tools for screening of new market participants.