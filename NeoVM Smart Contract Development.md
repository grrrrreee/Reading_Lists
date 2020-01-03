## *NOTE: Ignore the broken links and placeholders for images.*

NeoVM 과 SmartX 를 이용한 스마트 컨트랙트 개발-- 기본편

1.  **작업 전 확인사항**

실제 개발 과정을 진행하기 전에, 필요한 툴들을 모두 갖추었는지 확인해야합니다. 

아래에 온톨로지 플랫폼을 활용한 스마트 컨트랙트 개발 과정에서 사용되는 필수적인 툴들이 나열되어 있습니다. 

-   SmartX -- 온톨로지 온라인 스마트컨트랙트 IDE 및 디버거. 

-   Cyano wallet -- 구글 크롬 확장 프로그램

-   Explorer -- 전반적인 블록체인 활동과 거래내역을 추적할 수 있는 웹 기반 공개 툴 

2.  **환경 설정**

개발 환경설정 과정은 복잡하고 시간이 많이 소요되는 작업입니다. 그래서 이 과정을 최대한 부드럽고 편하게 하여 많은 사람들이 어려움을 느끼지 않게 하기 위해 노력했습니다. 이 과정에서 사용될 툴들은 모두 웹기반이고 실용적이며 독자 여러분 모두 바로 사용할 수 있습니다. 

먼저, 편의성과 실제 네트워크에서 작동 확인이라는 2가지 목표를 달성하기 위해서 스마트 컨트랙트를 개발 및 테스팅을 테스트넷에서 진행합니다. 프라이빗 체인은 사용하지 않습니다. 물론, 프라이빗 체인은 로컬 환경에서도 Ontology의 Punica suite를 사용하여 쉽게 설정될 수 있습니다. Punica suite는 프라이빗 넷에서도 스마트 컨트랙트를 배포하고 테스팅 할 수 있는 개발 툴 모음집입니다. 하지만, 해당 과정에서는 사용되지 않습니다.


개발을 하는데 있어서 가장 먼저 필요한 것은 웹 브라우저입니다. 그 중에서도 구글 크롬을 추천합니다. 과정에서 사용할 Cyano 지갑이 구글 크롬 플러그인이기 때문입니다.
Download link -
[[https://www.google.com/chrome/]{custom-style="Hyperlink"}](https://www.google.com/chrome/)

다음으로, Cyano 지갑 플러그인을 브라우저에 설치합니다. Cyano 지갑은 크롬 스토어에서 검색하여 찾을 수 있습니다. 아래의 링크에 들어가십시오.  -
[[https://chrome.google.com/webstore/detail/cyano-wallet/dkdedlpgdmmkkfjabffeganieamfklkm]{custom-style="Hyperlink"}](https://chrome.google.com/webstore/detail/cyano-wallet/dkdedlpgdmmkkfjabffeganieamfklkm)

![](media/image1.jpg){width="3.7666666666666666in"
height="6.191666666666666in"}

설치가 완료되었다면 이제 지갑을 실행해보겠습니다. 지갑을 실행하면 로그인하라는 메세지가 표시될 것입니다. 

만약 이미 Cyano 계정을 가지고 있다면, 개인키나 연상기호 단어를 사용하여 로그인 하십시오. 

만약 아직 계정이 없다면, 새로운 계정을 등록하는 작업을 진행해주십시오. 당신의 개인키와 연상기호 단어는 반드시 안전한 곳에 보관되어야 합니다.

로그인을 한 후, 지갑에서 **.dat** 파일을 내보내기(export) 합니다. 오른쪽 상단의 톱니바퀴 모양을 클릭하면 설정 화면이 나오는데 그곳에서 진행하면 됩니다. 이 파일 역시 안전하게 저장되어야 합니다. 나중에 이 파일은 SmartX와 함께 사용될 것입니다.

다음으로, 네트워크 세팅을 TEST-NET으로 변경하여 주십시오. 앞서 언급하였듯이, 이 과정에서는 테스트 넷에 컨트랙트를 배포하고 테스팅합니다.

![](media/image2.jpg){width="3.75in" height="6.216666666666667in"}


알림 : 테스트넷에서 스마트 컨트랙트 배포 및 테스팅하는 과정에서는 메인넷의 ONT/ONG의 잔고 즉, 실제 ONT/ONG 없이도 진행할 수 있습니다. 하지만, 테스트 넷의 ONG 잔고 즉, 테스트 넷의 ONT/ONG는 필요합니다. 필요한 가스 총 비용은 가스 값과 가스 한계값(limit)으로 결정됩니다. (가스 값 \* 가스 한계값) 테스트 토큰은 무료로 제공될 수 있으며 아래와 같이 신청할 수 있습니다.

참고 링크:
[[https://developer.ont.io/applyONG]{custom-style="Hyperlink"}](https://developer.ont.io/applyONG)

![](media/image3.jpeg){width="5.356871172353456in" height="2.6287871828521436in"}

![](media/image4.jpeg){width="7.242424540682415in" height="3.5619203849518812in"}

<<<<<<< HEAD
우리가 사용할 IDE는 브라우저 기반의 개발 환경이자 Python, C\# 그리고 JavaScript(곧 지원)를 지원하는 Ontology의 Smart X입니다. 해당 과정에서는 Python이 사용될 것입니다.

NeoVM 은 SmartX에서 Python으로 작성된 프로그램을 실행하는 실행 엔진의 역할을 합니다. SmartX 코어는 모든 Ontology의 API들을 통합하여, 블록체인 관련 행위를 실행하는데 있어서 필요한 관련 API를 불러오기(import)하여 직접 사용할 수 있게 지원합니다. 이 부분은 중후반부에 더 자세하게 이야기 해보겠습니다.

![](media/image5.jpeg){width="7.268055555555556in" height="3.609027777777778in"}

탐색기는 온라인 웹기반 툴로 거래를 추적하고 블록체인 행위 관련 이벤트들의 진행상황 및 결과를 확인할 수 있게 지원합니다. 테스트넷과 메인넷 2개 모두의 거래 해시, ONT ID, 컨트랙트 주소 그리고 블록 높이 등에 대한 정보를 얻을 수 있습니다. 

3.  **Launching SmartX**

알림 : 스마트 컨트랙트 테스팅을 하려면 Cyano 계정만 필요하며, ONT/ONG는 필요하지 않습니다.

![](media/image4.jpeg){width="7.242424540682415in" height="3.5619203849518812in"}

![](media/image6.jpeg){width="7.268055555555556in" height="3.5527777777777776in"}

로그인을 하게되면 프로젝트를 선택하라는 메세지를 보게 될 것입니다. 만약 진행중인 프로젝트가 없다면 새로운 프로젝트를 생성하십시오. 

![](media/image7.jpg){width="2.2229440069991253in" height="2.3939391951006126in"}

새로운 프로젝트를 생성할 때, 작업에 사용할 프로그래밍 언어를 선택할 수 있습니다. 해당 과정에서 제시할 예시나 사용할 샘플 코드는 모두 Python으로 작성되어 있습니다. 해당 과정은 개발 언어 자체보다는 NeoVM을 이용하여 스마트 컨트랙트 개발하는 방법에 더 집중하여 작성되었습니다.

![](media/image8.jpeg){width="7.268055555555556in" height="3.5375in"}

파이썬을 선택하고 다음 메뉴로 진행하십시오. 스마트 컨트랙트의 예시들과 템플릿들이 담겨있는 목록이 나타날 것입니다. 원하는 템플릿을 선택하여 샘플 코드를 바로 실행하거나 이를 기반으로 수정하고 추가하여 나만의 스마트 컨트랙트를 만들수 있습니다. 해당 과정에서는 OEP-4 템플릿으로 진행할 예정입니다.

4.  **프로그램 논리구조 작성 시작**

![](media/image9.jpeg){width="7.268055555555556in" height="3.597916666666667in"}

SmartX가 정상적으로 작동하기 시작하면 아래와 같은 화면이 보일 것입니다. 

OEP-4 템플릿을 선택하면 이미 코드가 일정부분 쓰여져 있을 것입니다. 물론, 원한다면 해당 코드를 직접 수정하고 원하는대로 구현하고자 하는 바를 나타낼 수도 있습니다. 하지만, 이 과정에서는 최대한 간단하게 보여주는 것이 목표이기에 제공된 코드를 그대로 사용하겠습니다.

다음 단계로 넘어가기 전에, 이 과정에서 사용될 코드를 살펴보고 전체적인 구조를 한번 파악해보는 것도 나쁘지 않을 것 같습니다. 

![](media/image10.jpg){width="7.258333333333334in" height="1.525in"}

이 섹션은 코드에서 정의된 변수들은 프로토콜 그 자체와 이 기능들을 제어할 특징점들을 정의합니다.(?)

"NAME" 그리고 "SYMBOL" 와 같은 변수는 토큰의 식별자로 사용됩니다. 

"FACTOR" 은 전송량의 정확도를 나타내는 값으로 기본값은 10으로 설정되어있습니다. 예를 들어, 값이 100으로 설정되면, 최대 2 수준의 정밀도를 시스템에서 지원합니다. 

"DECIMALS" 은 접근(access) 편의성을 위한 승수 값을 저장하고 있습니다. 

"OWNER" 는 Base58 주소를 저장하고 있습니다. 해당 주소는 토큰에 대한 전반적인 권한과 필요에 따라 이를 분산시킬 수 있는 권한을 가진 계정 혹은 주체의 주소입니다.

"TOTAL\_AMOUNT" 는 현재 존재하는 토큰의 총량을 나타냅니다. 항상 같은 값을 반환합니다.

"BALANCE\_PREFIX" 는 인증 목적으로 사용되는 access modifier 입니다.

"APPROVE\_PREFIX" 는 같은 목적으로 사용됩니다. 소유자(owner) 다른 계정으로 하여금 토큰을 사용할 수 있게 하는 기능이 있습니다. 이때 타 계정 소유자가 허가를 받았는지 여부를 확인하는 인증작업에 사용됩니다.

"SUPPLY\_KEY" 는 토큰 전체량과 직접적으로 관련이 있습니다. 값에 직접적으로 접근할 수 없기에 해당 값을 이용하는 모든 작동들에 이 함수가 이용됩니다.

![](media/image11.jpg){width="1.3in" height="0.19166666666666668in"}

이 행은 코드의 상단 부분에 바로 드러납니다.

GetContext() 는 스마트 컨트랙트와 블록체인의 다리 역할을 하는 함수입니다. 해당함수는 Storage API의 일부인 GET과 PUT을 호출하여 체인에 혹은 체인으로부터 정보를 전송하거나 받아올 때 사용됩니다. 

과정에서 사용될 관련 API들을 하나씩 살펴보면서 진행할 예정입니다. 또한, 후반부에는 사용가능한 전체 API를 간단하게 알아보겠습니다.

![](media/image12.jpg){width="5.083333333333333in"
height="0.8083333333333333in"}

독자들도 SmartX IDE를 사용하여 개발과정을 함께하고 계신다면 과정에서 사용되는 모든 API들과 함수들은 프로그래밍을 시작할 때 불러오기(import)하여 바로 접근 및 사용할 수 있습니다. 과정에서 우리가 불러올 함수들은 Runtime API에서 GetContext(), Get(), Put(), and Delete() from the Storage API, Notify(), CheckWitness() Base58ToAddress()과 기본 내재함수(built-in function) concat() 입니다. 

알림: 기본 내재함수는 이후 버전에서는 불러올 필요없이 바로 사용할 수 있습니다.

이제 Main() 함수를 한번 살펴봅시다 
function.![](media/image13.jpg){width="5.189394138232721in" height="4.5815594925634295in"}

Main() 함수는 2개의 인수(argument)를 취합니다. *operation* 과 *args* 입니다. *operation* 인수는 수행할 조작들을 기반으로하여 실행될 함수들을 지시합니다. *args* 인수는 함수가 추가적인 실행을 하는데 있어서 필요한 중요한 정보(예를 들면 계정 주소나 입력값 데이터 등)를 전달하는데 도움을 주는 역할을 합니다.

Main()에 전달된 인수에 따라서 호출될 수 있는 함수는 11가지가 있습니다. SmartX는 이 인수들을 우측 하단에 있는 "Options"을 이용하여 전달합니다.

-   **init()** **:** 프로그램 논리 구조의 시작점을 설정합니다. 해당 함수는 주어진 값을 기반으로 위에서 언급된 변수들의 초기값 설정을 합니다. 그렇기에, 해당 함수는 배포 후에 가장 먼저 실행되어야 하는 함수입니다.

-   **name() :** 토큰에 설정된 이름을 반환합니다. 이 경우에는 "My Token" 입니다. 

-   **symbol() :** 토큰의 symbol을 반환합니다. 이 경우에는 "MYT"입니다.

-   **decimals() :** 유효한 토큰 값을 정확하게 소수점까지 반환합니다, 이 경우에는 8자리입니다.

-   **totalSupply() :** 초기값 설정 때 배정된 토큰 총량을 반환합니다. 순환에 할당된 고정된 숫자를 반환합니다. (SUPPLY\_KEY를 사용하여 초기값 설정을 하는 과정에서 이전에 저장된 값들을 체인에서 불러옵니다.)

-   **balanceOf(acct) :** 함수에 인수로 전달된 Base58 주소의 계정 잔고 금액을 불러옵니다. 

-   **transfer(from\_acc, to\_acc, amount) :** 해당 함수는 인수로 받은 값만큼을 from\_acc 주소에서 to\_acc 주소로 전송합니다.

-   **transferMulti(args) :** 여기에서 인수는 다음과 같은 정보를 담은 배열(array)입니다 : 보내는 사람의 주소, 받는 사람의 주소 그리고 거래량. 각 계정주소와 그에 해당하는 금액을 전달하여 다수의 거래에서 반복적으로 실행될 수 있습니다.

-   **transferFrom(spender, from\_acc, to\_acc, amount) :** 송신자는 \_acc address에서 해당하는 양을 받아서 \_acc address에 전송을 합니다.

-   **approve(owner, spender, amount) :** 소유자는 소비자에게 자신의 계정에 있는 특정량의 토큰을 사용할 수 있는 권한을 부여합니다. 여기서 소유자와 소비자는 모두 Base58 주소이고 이 둘과 소비자가 허가받은 양 총 3개를 인수로 사용합니다.

-   **allowance(owner, spender) :** 소유자 계정이 허가한 소비자가 허가된 양에 한해서 사용될 수 있습니다. 이 경우에는 2개의 인수를 필요로 하는데 2개 모두 Base58 주소입니다.

위에 언급된 함수들은 2가지 종류로 나누어집니다. - 접근(access)형 함수와 유틸리티(utility)형입니다. 이 함수들이 호출되었을 때의 플로우를 한번 생각해봅시다. 

접근형 함수들은 주로 컨트랙트 배포 후에 데이터를 가져오는 역할을 합니다. name(), symbol(), totalSupply() 그리고 balanceOf(acc) 와 같은 함수들은 정보들을 체인으로부터 정보를 취하는 기능들을 모아놓은 Storage API의 get() 함수를 사용하여 수집합니다. 실제 프로그램 논리구조에서는 어떤식으로 적용되는지 한번 살펴봅시다.

**name() function definition**

![](media/image14.jpg){width="2.1212117235345582in" height="0.8444630358705162in"}

가장 간단한 접근형 함수입니다. name() 함수는 인수가 필요하지 않습니다. 비록 name(), symbol() 그리고 decimals()가 명시적으로 get() 함수를 호출하지는 않지만, 스마트 컨트랙트가 배포된 후에 모든 정보는 체인에서 불러옵니다. 

**balanceOf(acc) function definition**

![](media/image15.jpg){width="3.25in" height="1.252205818022747in"}

balanceOf() 함수는 1개의 인수만을 받습니다. 계정을 나타내는 Base58 주소 입니다. 유효성 검사는 해당 인수에 대해서 길이 검사를 하고 유효하지 않다면 예외 처리하는 형식으로 진행합니다. 만약에 주소가 유효하다면 get() 함수가 2개의 인수들과 함께 호출됩니다. 2개의 인수는 BALANCE\_PREFIX을 접두로 가진 계정 주소와 context 입니다.

context는 데이터 참조가 계정의 잔고를 가져오도록 지원하며, 접두를 붙인 계정 주소에 대한 접근 권한을 허용합니다. 다음으로, get()은 해당 데이터를 Main()으로 반환하고 notify() 함수를 이용하여 로그 윈도우창에 결과값을 출력합니다. totalSupply() 함수도 비슷한 방식으로 작동합니다.

알림：잔액 및 허용 접두부(approve prefixe)는 ASCII 형식의 16진수 값이며 당신의 프로그램 논리 구조에 맞게 변경될 수 있습니다.

샘플 코드에 있는 유틸리티형 함수를 살펴보겠습니다. 

유틸리티형 함수는 기능들을 더욱 풍부하게 하고 온체인 데이터 수정을 지원하며, 실행될 거래와 작업들의 기반이 됩니다. 해당 과정에서 사용되는 OEP-4 토큰 논리구조는 다양한 사용사례들과 시나리오들을 묘사합니다. 이는 스마트 컨트랙트가 얼마나 많은 분야에 사용될 수 있고 다양한 방법으로 응용될 수 있는지와 우리가 생성할 수많은 종류의 비즈니스 로직에 사용될 수 있다는 것을 보여줍니다.

**transfer(from\_acc, to\_acc, amount)**

![](media/image16.jpg){width="4.95in" height="3.6666666666666665in"}

transfer() 함수는 토큰을 한 계정에서 다른 계정으로 전송하는 거래의 가장 기본적인 특징을 실행합니다. 3개의 인수를 받습니다. 보내는 사람의 주소, 받는 사람의 주소 그리고 거래량입니다.

유효성 검사는 간단한 길이 검사로 검증하지만, 원한다면 더 복잡한 논리구조를 개발 및 구현할 수 있습니다. 

다음으로 BALANCE\_PREFIX 가 보내는 사람의 계정 주소와 합쳐집니다. 그리고 get()은 계정 주소를 사용하여 잔액정보를 가져옵니다. 다음으로 계정의 잔고와 전송될 금액을 비교합니다. 아래에 3가지 시나리오 모두 명확하게 정의되어 있습니다.

만약에 잔고가 거래량보다 적다면, 거래는 실패하게 되고 통제권은 바로 Main() 으로 넘어가게 됩니다. 

만약에 잔고와 거래량이 정확히 일치한다면 보내는 사람의 계정 잔고는 0으로 설정됩니다. 이때 보내는 사람의 주소의 접두를 이용하여 delete() 방법이 호출됩니다. 실질적으로 put() 방법을 사용하여 보내는 사람의 계좌 잔고를 0으로 만드는 것과 같은 결과를 내기는 하지만 put()을 사용한다면 보안상 취약점이 나타날 수 있어서 delete() 방법을 사용하였습니다. 

만약 잔고가 거래량보다 높다면, put()을 통해서 해당량만큼 잔고에서 차감하고 차감후 그 결과량으로 보내는 사람의 잔고를 업데이트합니다.

다음으로, BALANCE\_PREFIX를 접두로 둔 받는 사람의 주소는 받는 사람의 계좌잔고에 거래량을 더하는 과정에서 사용됩니다. 

마지막으로, 해당 거래는 RegisterAction() 함수를 사용하여 체인에 전송되고 원장에 기록됩니다.

TransferEvent 는 해당과정에서 RegisterAction()이 사용하는 별칭이라고 생각하면 됩니다. RegisterAction()은 Acition API에 속해있으며 거래의 세부사항을 기록하고 체인에 전송되는 4개의 인수를 취합니다. 거래 해시값과 특정 몇가지 다른 결과값들은 로그 섹션에서 확인할 수 있습니다.

 **transferMulti()** 도 비슷한 방법으로 작동합니다. transferMulti()가 실행되면 transfer()를 호출하기에 기본적인 논리구조는 transfer()와 거의 비슷합니다. 하지만 차이점이 있다면 여러 개의 거래가 동시에 처리된다는 점입니다. 인수는 중첩 배열(nested array) 1개를 필요로 합니다.

![](media/image17.jpg){width="5.2551924759405075in" height="1.7727274715660541in"}

sub-array 요소는 3개의 세트로 구성됩니다. 첫번째와 두번째는 보내는 사람과 받는 사람의 주소이고, 세번째는 거래량입니다. sub array는 args\[\]배열에 더 이상 요소가 남지 않을때까지 반복 실행됩니다.

만약에 sub-array가 3개의 인수를 포함하지 않거나 이전 거래가 특정 이유로 인하여 실패하였다면 예외처리가 됩니다. 예외처리가 되면 통제권은 loop에서 나와 Main()으로 넘어갑니다.

**approve(owner, spender, amount)**

![](media/image18.jpg){width="5.208333333333333in" height="3.1166666666666667in"}

이 함수는 또다른 복잡한 논리구조를 구현한 것입니다. "소비자(spender)"가 "소유자(owner)"에게 특정량만큼의 토큰을 사용할 수 있는 허가를 받는 과정입니다.

해당 함수는 2가지 검사를 진행합니다. 길이로 유효성 검사를 진행하고, 승인을 실행할 수 있는 “소유자”여부로 사용자 검사를 진행합니다.

다음으로, 승인된 금액과 “소유자” 계정의 잔고를 비교합니다.

만약 잔고에 충분한 양이 있지 않다면 종료되고 통제권은 Main()에 넘어가게 됩니다. 

만약 잔고에 충분한 양이 있다면, APPROVAL\_PREFIX를 접두로 가진 "소유자" 주소와 "소비자" 주소를 연결한 키(key)가 형성됩니다. 다음으로 put()을 이용하여 context, 위에서 언급된 키 그리고 승인 금액이 원장에 전달 및 기록됩니다.

거래 이벤트는 ApprovalEvent()방법으로 기록되고 NeoVM엔진이 반환한 거래 해시값이 포함된 결과값은 로그 섹션에 나타나게 됩니다. 

**transferFrom(spender, from\_acc, to\_acc, amount)**

![](media/image19.jpg){width="6.85in" height="5.2in"}

transferFrom() 함수는 더 복잡한 논리구조를 실행합니다. 이는 특정 어플리케이션에서 유용하게 사용될 수 있을 것입니다.

이 함수에서는 spender라고 불리우는 제3자를 허용합니다. 이 기능을 사용하면 spender가 자격증명을 하도록 선정되지 않은 계정의 일정량의 토큰을 사용할 수 있게 합니다. 원칙적으로 approve() 함수와 같은 논리구조를 사용한다고 보면 될 것입니다. 4개의 인수를 받습니다. 3개는 Byte58 주소이고 1개는 전송량입니다.

먼저, 함수는 주소의 유효성 검사를 실행합니다. 후에 소비자에게 Runtime API의 CheckWitness()함수를 사용하여 이 거래를 실행할 권한이 있는지 확인합니다.  

다음으로, "보낸 사람"의 계정 잔고를 가져와 거래량과 비교한 후에 충분한 잔고가 있는지를 확인합니다. 잔고 정보를 가져오는 과정과 똑같습니다. 

spender의 주소는 APPROVE\_PREFIX를 접두로 붙고 승인된 금액은 접두사 주소를 사용(참조)하여 가져옵니다. 

거래는 그 후에 실행됩니다. 만약 거래 금액이 승인된 금액을 초과한다면, 거래는 중지되고 통제권은 Main()함수에게 돌아가게 됩니다. 

만약에 승인된 금액이 정확히 양과 일치한다면, 거래량은 "보낸 사람"의 계좌 잔고에서 차감됩니다. 

만약 승인된 금액이 거래량을 초과한다면, 그 차액이 계산되어 추후에 참조할 수 있도록 원장에 기록됩니다. 그리고 그 거래량은 "보낸 사람" 계좌 잔고에서 차감됩니다.

거래량은 put()함수를 사용하여 "받는" 사람의 계정으로 전송됩니다. 후에 해당 이벤트는 기록되고 거래 해시와 함께 결과는 IDE의 로그 섹션에 나타납니다.

이와 비슷한 논리 구조를 구현한 또다른 함수는 바로 **allowance(owner, spender)** 입니다. 이 함수는 "소유자" 계정에서 "spender" 계정에 할당한 허용량을 쉽게 불러올 수 있게 지원합니다.

![](media/image20.jpg){width="4.633333333333334in" height="1.1666666666666667in"}

굳이 말하자면, 해당 함수는 허용량 값을 받아온다는 점에서 접근형 함수로 분류될 수 있습니다. 그러나 get() 쿼리를 행하여 체인으로 부터 결과를 가져온다는 점에서 약간의 차이가 있다고 할 수 있습니다. 

접두를 가진 소유자의 주소와 소비자의 주소를 합쳐서 형성된 key는 get() 방법으로 context와 함께 전달되어 필요한 허용량 값을 가져오고, Main()으로 반환됩니다. 이 값은 확인할 수 있으며 다른 작업에 사용될 수도 있습니다.

5.  **배포와 테스팅**

![](media/image21.jpg){width="5.175in" height="1.5916666666666666in"}

논리 구조 구현이 완료되고 나면, 이제는 스마트 컨트랙트를 실행(compile)할 차례입니다. 

스마트 컨트랙트를 실행하고 나면, 그 결과는 IDE 화면에 다음과 같이 보입니다 - 

![](media/image22.jpeg){width="2.5737970253718285in" height="3.25in"}

compile 탭에 보이는, AVM 바이트 코드는 컴파일 후에 생성된 중간결과물입니다. NEOVM은 이 AVM 코드를 처리하여 계약을 실행합니다.

opcode는 스택 상황을 줄단위로 지시합니다 ; 발전된 형태의 디버깅 툴입니다. 

ABI는 모든 함수의 매개변수 및 컨트랙트 해시 그 자체를 저장합니다. 컨트랙트 ABI는 추후에 검색을 하는데 사용됩니다. 

![](media/image23.jpg){width="7.268055555555556in" height="0.47727252843394574in"}

로그 섹션은 디버깅 결과부터 VM이 반환하는 결과까지 다양한 컴파일러 결과를 보여줍니다.

컴파일 결과 나타나는 오류가 모두 수정되거나 해결된 후에, 컴파일을 성공하면 이제 해당 컨트랙트는 배포될 수 있습니다.

![](media/image24.jpg){width="5.091666666666667in" height="6.133333333333334in"}

여기서, 컨트랙트에 대한 정보를 입력합니다. 가스 값과 가스 한계값을 입력할 수 있는 창이 나타납니다.

**알림: 가스 한계는 정확히 소수점 9자리까지 포함됩니다. 그렇기에, 10^9^ 가스 unit은 1 ONG 토큰과 일치한다고 볼 수 있습니다. 표현할 수 있는 값 중 가장 작은 값은 0.000000001 입니다.**

**지갑은 컴파일하고 실행될 코드의 복잡성에 기반하여 적절한 한계값을 자동으로 설정합니다. 하지만 그 한계값을 사용자가 스스로 결정할수도 있습니다. 한계값이 비용보다 높은지 확인하십시오. 그렇지 않다면 컨트랙트의 배포 및 호출 행위는 수행되지 않을 것입니다.**

![](media/image25.jpg){width="1.9874278215223098in" height="3.113636264216973in"}

성공적으로 배포가 되면, 거래 해시값은 로그 섹션과 오른쪽 하단 결과창에서 확인할 수 있습니다.

![](media/image23.jpg){width="7.268055555555556in" height="1.6145833333333333in"}

다음으로, 이제는 컨트랙트를 실행하려고 합니다. 

다른 함수들을 실행하기 전에, 먼저 지갑의 초기값 설정을 해야합니다. 이때 init()함수를 사용하며 거래를 실행하는데 있어서 충분한 양의 잔고가 있는지를 확인합니다.  

후에 drop-down 메뉴에서 실행할 함수를 선택할 수 있습니다. 

이 부분에서 우리는 2가지의 선택권이 있습니다. 사전 실행(pre-run) 혹은 실행(run) 입니다. 바로 컨트랙트 코드를 실행시키고, 엔진을 AVM코드가 생성되기 전에 실행시킬 수 있습니다. 사전실행은 컨트랙트 코드를 테스팅할 때 즉, 해당 코드가 잘 작동하는지를 확인하는데에 사용됩니다. 컨트랙트 코드를 사전 실행할 때에는 가스 값이 요구되지 않습니다.

drop-down 메뉴에서 실행시키고 싶은 함수를 선택할 수 있습니다, 전달될 값의 데이터 형태와 공란에 들어갈 인수를 전달할 수 있습니다. 

![](media/image26.jpg){width="5.066666666666666in" height="5.416666666666667in"}

![](media/image27.jpg){width="7.268055555555556in" height="0.5909087926509187in"}

로그 섹션에 나타나는 모든 결과는 16진수 형태입니다. 오른쪽 창의 도구(tool) 옵션은 데이터 변환 및 몇가지 기능을 수행하는데 사용되는 다양한 변환 도구를 제공합니다.

![](media/image28.jpg){width="4.983333333333333in" height="6.258333333333334in"}

거래가 실행될 때마다, 거래 해시값이 결과로 반환되고 해당 거래 해시값은 결과를 추적하는데 사용됩니다. 이제 다른 거래도 한번 실행해보려고 합니다. 이번에는 토큰을 실제로 전송해보도록 하겠습니다. 

![](media/image29.jpg){width="5.108333333333333in" height="5.408333333333333in"}

컨트랙트를 실행시키면, 거래를 확인하라는 메세지가 표시됩니다. 계약을 실행하는데 있어서 사용될 가스 가격 및 가스 한도를 입력하라는 메세지가 표시됩니다. (ONG)

가스 가격과 가스 한도를 입력합니다. 거래가 성공적으로 실행이 되면 거래 해시갑이 로그 섹션에 나타날 것입니다.

![](media/image30.jpg){width="7.268055555555556in" height="0.5875in"}

balanceOf() 함수를 사전 실행하면, 토큰을 받은 계정의 현재 잔고를 확인할 수 있습니다. 사전 실행을 하여 balanceOf()의 결과값을 한번 살펴봅시다. 직접 실행을 한다면 실행은 되지만, 값을 반환하지는 않습니다.

![](media/image31.jpg){width="7.268055555555556in" height="0.7659722222222223in"}

10진수 기준 50개의 샘플 토큰이 받는 사람의 주소로 전송되었습니다. 표시된 32라는 값은 16진수 값입니다. 

![](media/image32.jpg){width="4.925in" height="1.6583333333333334in"}

이 16진수 형태 값은 tools 섹션에 제공된 변환기(converter)를 통하여 10진수 값으로 변환할 수 있습니다. 

거래 해시 값은 나중에 탐색기에서 거래 세부사항을 보는데에 사용될 수 있습니다. 

![](media/image33.jpeg){width="7.268055555555556in" height="3.454861111111111in"}

오른쪽 창의 마지막 섹션은 Restful한 API입니다. API는 현재 블록높이 이벤트에 대한 Smartcode와 유용한 체인의 상태 정보를 불러오는 행위 등 통신하는데에 사용됩니다. Smartcode 는 사실상 스마트 컨트랙트 정보를 JSON기반으로 표현한 것입니다. JSON은 쿼리를 기반으로 생성되고 반환되며 이때 함수들은 Restful API로 호출됩니다.

![](media/image34.jpg){width="4.25in" height="5.99959864391951in"}

거래 해시값은 거래 각각의 Smartcode를 가져오는데에 사용됩니다. 

Restful API는 SDK를 통해서도 접근 및 사용할 수 있습니다. SDK 통합 방법 관련 튜토리얼은 다른 글에서도 확인하실 수 있습니다.

![](media/image35.jpg){width="3.2727274715660544in" height="3.0673534558180227in"}