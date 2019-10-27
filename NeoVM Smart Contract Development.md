## *NOTE: Ignore the broken links and placeholders for images.*

NeoVM 과 SmartX 를 이용한 스마트 컨트랙트 개발-- 기본편

1.  **작업 전 확인사항**

실제 개발 과정을 진행하기 전에, 필요한 툴들을 가지고 있는지 확인해야합니다. 

아래의 목록이 온톨로지 플랫폼을 활용한 스마트 컨트랙트 개발 과정에서 사용되는 필수적인 툴들입니다. 

-   SmartX -- 온톨로지 온라인 스마트컨트랙트 IDE 및 디버거. 

-   Cyano wallet -- 구글 크롬 확장 프로그램

-   Explorer -- 전반적인 블록체인 활동과 거래내역을 추적할 수 있는 웹 기반 공개 툴 

2.  **환경 설정**

개발 환경설정 과정은 복잡하고 시간이 많이 소요되는 작업일 수 있습니다. 그래서 이 과정은 최대한 swift 하고 편하게 하여 많은 사람들이 어려움을 느끼지 않게 하기 위해 노력했습니다. 이 과정에서 사용될 툴들은 모두 웹기반이고 실용적이며 바로 사용할 수 있습니다. 

먼저, 편의를 극대화 시키기 위해서, 우리는 스마트 컨트랙트를 개발 및 테스팅을 테스트넷에서 진행할 예정입니다. 그리고 이를 통해 프라이빗 체인의 필요성을 없앨 계획입니다. (thereby eliminating the need of a private chain.) 프라이빗 블록체인 네트워크 구조는 로컬 환경에서도 Ontology의 Punica suite를 사용하여 쉽게 설정될 수 있습니다. Punica suite는 프라이빗 넷에서도 스마트 컨트랙트를 배포하고 테스팅 할 수 있는 개발 툴 모음집입니다. 하지만, 해당 과정에서는 사용되지 않습니다.

우리에게 제일 먼저 필요한 것은 웹 브라우저입니다. 그 중에서도 구글 크롬을 추천합니다. 과정에서 사용할 Cyano 지갑이 구글 크롬 플러그인이기 때문입니다.
Download link -
[[https://www.google.com/chrome/]{custom-style="Hyperlink"}](https://www.google.com/chrome/)

다음으로, Cyano 지갑 플러그인을 브라우저에 설치합니다. Cyano 지갑은 크롬 스토어에서 검색하여 찾을 수 있습니다. 아래의 링크에 들어가십시오.  -
[[https://chrome.google.com/webstore/detail/cyano-wallet/dkdedlpgdmmkkfjabffeganieamfklkm]{custom-style="Hyperlink"}](https://chrome.google.com/webstore/detail/cyano-wallet/dkdedlpgdmmkkfjabffeganieamfklkm)

![](media/image1.jpg){width="3.7666666666666666in"
height="6.191666666666666in"}

설치가 완료되고 난 후에 이제 처음으로 지갑을 실행해보겠습니다. 지갑을 실행하면 로그인하라는 메세지가 표시될 것입니다. 

만약 이미 Cyano 계정을 가지고 있다면, 개인키나 연상기호 단어를 사용하여 로그인 하십시오. 

만약 아직 계정이 없다면, 새로운 계정을 등록하는 작업을 진행해주십시오. 당신의 개인키와 연상기호 단어는 반드시 안전한 곳에 보관해야 합니다.

로그인을 한 후에는 지갑에서 **.dat** 파일을 export 합니다. 오른쪽 상단의 톱니바퀴 모양을 클릭하면 설정 화면으로 넘어가게 됩니다. 이 파일은 안전하게 저장되어야 합니다. 나중에 이 파일은 SmartX와 함께 사용될 것입니다.

다음으로,네트워크 세팅을 TEST-NET으로 변경하여 주십시오. 우리의 결과물을 테스트 넷에서 배포하고 테스팅할 것입니다.(?)

![](media/image2.jpg){width="3.75in" height="6.216666666666667in"}

<<<<<<< HEAD
: 테스트넷에서 스마트 컨트랙트 배포 및 테스팅하는 과정은 메인넷의 ONT/ONG의 잔고없이도 진행할 수 있습니다. 하지만, 테스트 넷의 ONG 잔고는 필요합니다. 필요한 가스 총 비용은 가스 값과 가스 limit으로 결정됩니다. (가스 값 \* 가스 limit) 테스트 토큰은 무료로 제공될 수 있으며 아래와 같이 신청할 수 있습니다.

given link:
[[https://developer.ont.io/applyONG]{custom-style="Hyperlink"}](https://developer.ont.io/applyONG)

![](media/image3.jpeg){width="5.356871172353456in" height="2.6287871828521436in"}

![](media/image4.jpeg){width="7.242424540682415in" height="3.5619203849518812in"}

우리가 사용할 IDE는 브라우저 기반의 개발 환경이자 Python, C\#, and JavaScript(coming soon) 를 지원하는 Ontology의 Smart X입니다. 해당 과정에서 소개할 개발 과정에서는 Python이 사용될 것입니다.

NeoVM 은 SmartX에서 Python으로 작성된 프로그램을 실행하는 실행 엔진의 역할을 합니다. SmartX 코어는 모든 Ontology의 API들을 통합하여, 블록체인 관련 행위를 실행하는데 있어서 필요한 관련 API를 import 하여 직접적으로 사용할 수 있도록하게 합니다. 이 부분은 해당 과정에서 추후에 더 자세하게 이야기 해보겠습니다.

![](media/image5.jpeg){width="7.268055555555556in" height="3.609027777777778in"}

탐색기는 온라인 웹기반 툴로 거래를 추적하고 블록체인 활동 관련 이벤트들의 진행상황 및 결과를 확인할 수 있게 합니다. 테스트넷과 메인넷 2개 모두의 거래 해시, ONT ID, 컨트랙트 주소 그리고 블록 높이 등에 대한 정보를 얻을 수 있습니다. 

3.  **Launching SmartX**

알림 : 스마트 컨트랙트 테스팅을 하려면 Cyano 계정은 필요하지만, ONT/ONG 잔고는 필요하지 않습니다.

![](media/image4.jpeg){width="7.242424540682415in" height="3.5619203849518812in"}

![](media/image6.jpeg){width="7.268055555555556in" height="3.5527777777777776in"}

로그인을 하게되면 프로젝트를 선택하라고 하는 메세지를 보게 될 것입니다. 만약 이미 존재하는 프로젝트가 없다면 새로운 프로젝트를 생성하십시오. 

![](media/image7.jpg){width="2.2229440069991253in" height="2.3939391951006126in"}

새로운 프로젝트를 만드려고 할 때, 작업에 사용할 프로그래밍 언어를 선택할 수 있습니다. 해당 과정에서 제시할 예시나 사용할 샘플 코드는 모두 Python으로 작성되어 있습니다. 해당 과정은 개발 언어 자체보다는 NeoVM을 이용하여 스마트 컨트랙트 개발하는 방법에 더 집중하여 작성되었습니다.

![](media/image8.jpeg){width="7.268055555555556in" height="3.5375in"}

파이썬을 선택하고 다음 메뉴로 진행하십시오. 스마트 컨트랙트의 예시들과 템플릿들이 담겨있는 목록이 나타날 것입니다. 아무 예시를 선택하여 샘플 코드를 실행하거나 이를 기반으로 수정하고 추가하여 나만의 스마트 컨트랙트 논리구조를 구성할 수 있습니다.(?) 해당 과정에서는 OEP-4 형태를 살펴볼 예정입니다.

4.  **프로그램 논리구조 작성 시작**

![](media/image9.jpeg){width="7.268055555555556in" height="3.597916666666667in"}

SmartX가 설정되고 작동되기 시작하면 그 화면은 다음과 같을 것입니다. 

OEP-4 템플릿을 선택하였다면 이미 코드는 쓰여져 있을 것입니다. 물론, 원한다면 해당 코드를 직접 수정하고 원하는대로 구현하고자 하는 바를 나타낼 수도 있습니다. 하지만, 해당 과정은 최대한 간단한 과정을 보여주는 것이 목표이기에 우리는 제공된 코드를 사용하겠습니다.

다음 단계로 넘어가기 전에, 해당 과정에서 사용될 코드를 살펴보고 전체적인 구조를 한번 파악해보십시오.

![](media/image10.jpg){width="7.258333333333334in" height="1.525in"}

이 섹션으 코드에서 정의된 변수들은 프로토콜 그 자체와 이 기능들을 제어할 특징점들을 정의합니다.

"NAME" 그리고 "SYMBOL" 와 같은 변수는 토큰의 식별자로 사용됩니다. 

"FACTOR" is the base 10 value that dictates the precision of amounts that can be transferred. 예를 들어, 값이 100으로 설정되면, 최대 2레벨의 정밀도를 시스템에서 지원합니다. 

"DECIMALS" stores the multiplier value for access convenience.

"OWNER" 는 Base58 주소를 저장하고 있습니다. 해당 주소는 토큰에 대한 전반적인 권한과 필요에 따라 이를 분산시킬 수 있는 권한을 가진 계정 혹은 주체의 주소입니다.

"TOTAL\_AMOUNT" 는 현재 존재하는 토큰의 총량을 나타냅니다. 항상 같은 값을 반환합니다.

"BALANCE\_PREFIX" is an access modifier that is used with account addresses for authentication purposes.

"APPROVE\_PREFIX" serves the same purpose, but for the approve operation wherein the owner can authenticate another account to use tokens.

"SUPPLY\_KEY" correlates directly to the total amount of tokens and is used for any operations that may be carried with the total tokens figure, since the value isn't directly accessible.

![](media/image11.jpg){width="1.3in" height="0.19166666666666668in"}

This line immediately stands out in the upper section of the code.

GetContext() 는 스마트 컨트랙트와 블록체인의 다리 역할을 하는 함수입니다. 해당함수는 Storage API의 일부인 GET과 PUT를 호출하여 체인에 혹은 체인으로부터 정보를 fetching 하거나 transmitting 할 때 사용됩니다. 

We will go through the relevant APIs as we come across the respective functions, and the full set of available APIs in a later section of the tutorial.

![](media/image12.jpg){width="5.083333333333333in"
height="0.8083333333333333in"}

해당 과정은 SmartX IDE를 사용하고 있기 때문에, 과정에서 사용되는 모든 API들과 함수들은 프로그래밍을 시작할 때 importing하여 바로 접근 및 사용할 수 있습니다. 과정에서 우리가 import 할 함수들은 Runtime API에서 GetContext(), Get(), Put(), and Delete() from the Storage API, Notify(), CheckWitness() Base58ToAddress()과 빌트인 함수(built-in function) concat() 입니다. 

알림: In later versions, built-in functions don't need to be imported and can be called directly.

이제 Main() 함수를 한번 살펴봅시다 
function.![](media/image13.jpg){width="5.189394138232721in" height="4.5815594925634295in"}

Main() 함수는 2개의 argument를 취합니다. *operation* 과 *args* 입니다. The *operation* argument is based on the operation to be performed and dictates the function to the executed. The *args* argument helps passing the important information that a function needs to carry out further execution, for example account addresses or input data.

Here, clearly there are 11 different functions that can be called depending upon the argument that is passed in Main(). SmartX는 이 인수들을 우측 하단에 있는 "Options"르 이용하여 전달합니다.

-   **init()** **:** init() 함수는 프로그램 논리 구조의 시작점을 설정합니다. 해당 함수는 initializes the definition variables declared at the top based on the values provided. Thus, this is the first function that needs to be executed post deployment.

-   **name() :** 해당 함수는 토큰에 설정된 이름을 반환합니다. 이번 경우에는 "My Token" 입니다. 

-   **symbol() :** 해당 함수는 토큰의 symbol을 반환합니다. 이번 경우에는 "MYT"입니다.

-   **decimals() :** 유효한 토큰 값을 정확하게 소수점까지 반환합니다, 이 경우에는 8자리입니다.

-   **totalSupply() :** Returns the total number of tokens assigned while initializing. Denotes the fixed number of tokens allocated for circulation. (Uses SUPPLY\_KEY to fetch the value from the chain, stored earlier during initialization)

-   **balanceOf(acct) :** Fetches the corresponding token balance of the account that identifies with the Base58 address passed as argument to the function.

-   **transfer(from\_acc, to\_acc, amount) :** 해당 함수는 인수로 받은 값만큼을 from\_acc 주소에서 to\_acc 주소로 전송합니다.

-   **transferMulti(args) :** The parameter here is an array that contains the same information, i.e., the sender's address, receiver's address, and the amount to be sent, in that sequence at the respective indices. It can be iterated for multiples transactions by passing the respective account addresses and the corresponding amount.

-   **transferFrom(spender, from\_acc, to\_acc, amount) :** 송신자는 \_acc address에서 해당하는 양을 받아서 \_acc address에 전송을 합니다.

-   **approve(owner, spender, amount) :** The owner authorizes the spender to use a certain amount of tokens from their own account. 여기서 소유자와 소비자는 모두 Base58 주소와 소비자가 허가받은 양을 인수로 사용합니다.

-   **allowance(owner, spender) :** This function can be used to the amount that the owner account has authorized spender to use. 이 경우에는 2개의 인수 모두 Base58 주소입니다.

위에 언급된 함수들은 2가지 종류로 나누어집니다. - 접근형 함수와 유틸리티형입니다. Let us consider the flow of control as these functions are called.

Access functions are primarily used to fetch data post contract deployment. Functions such as name(), symbol(), totalSupply() and balanceOf(acc) allow us to achieve this by using get() function from the Storage API which fetches relevant data from the chain. Let us look at how it is implemented in program logic.

**name() function definition**

![](media/image14.jpg){width="2.1212117235345582in" height="0.8444630358705162in"}

This is how a simple access function can be defined. The name() 함수는 인수가 필요하지 않습니다. 비록 name(), symbol() 그리고 decimals()가 명시적으로 get() 함수를 호출하지는 않지만, 스마트 컨트랙트가 배포된 후에, 모든 정보는 체인에서 fetch 된다. 

**balanceOf(acc) function definition**

![](media/image15.jpg){width="3.25in" height="1.252205818022747in"}

balanceOf() 함수는 1개의 인수만을 받습니다. 계정을 나타내는 Base58 주소 입니다. 해당 인수에 대해서 길이 검사를 하고 유효하지 않다면 예외 처리를 하는 유효성 검사도 있습니다. 만약에 주소가 유효하다면 get() 함수가 2개의 인수들과 함께 호출됩니다. 2개의 인수는 BALANCE\_PREFIX을 접두로 가진 계정 주소와 context 입니다.

The context allows for data reference on the chain to fetch the account balance value, while the prefixed account address ensures authenticated access. Next, get() returns this data to Main() where it is output to the log window using the notify() function. The totalSupply() function works in a similar fashion.

알림：잔액 및 approve prefixes 은 ASCII 형식의 16진수 값이며 당신의 프로그램 논리 구조에 맞게 변경될 수 있습니다.

샘플 코드에 있는 유틸리티형 함수를 살펴보겠습니다. 

유틸리티형 함수는 기능들을 더욱 풍부하게 하고 온체인 데이터들을 수정할 수 있게하며, 이느 실행될 거래와 작업들의 기반이 됩니다. 해당 과정에서 사용되는 OEP-4 토큰 논리구조는 다양한 사용사례들과 시나리오들을 기능형태로 묘사합니다. This is to exhibit just how versatile smart contracts are in nature and the different kinds of business logic that they can be used to generate.

**transfer(from\_acc, to\_acc, amount)**

![](media/image16.jpg){width="4.95in" height="3.6666666666666665in"}

The transfer() 함수는 토큰을 한 계정에서 다른 계정으로 전송하는 거래의 가장 기본적인 특징을 실행합니다. 해당 함수는 3개의 인수를 받습니다. 보내는 사람의 주소, 받는 사람의 주소 그리고 거래량입니다.

해당 함수는 간단한 길이 검사로 검증을 진행하지만, 원한다면 더 복잡한 논리구조를 개발 및 구현할 수 있습니다. 

Next, the BALANCE\_PREFIX is concatenated to the sender's account address, and balance is retrieved by making a get() call using this address. A quick comparison is made in the next step where the sender account's balance is compared with the amount to be transferred. All three scenarios have been defined clearly.

만약에 잔고가 거래량보다 적다면, 거래는 실패하게 되고 통제권은 바로 Main() 으로 넘어가게 된다.

If the amount equates to the balance amount exactly, the balance of sender account is set to 0 by calling the delete() method using the sender's prefixed address. This is practically equivalent to using the put() method to manually assign the value 0 to sender's account but using put() method in this case might give rise to security vulnerabilities.

만약 잔고가 거래량보다 높다면, the amount is deducted from the balance by making a put() call and updating the sender accounts balance with the deducted value.

다음으로, BALANCE\_PREFIX를 접두로 둔 받는 사람의 주소는 받는 사람의 계좌잔고에 거래량을 더하는 과정에서 사용됩니다. 

마지막으로, 해당 거래는 RegisterAction() 함수를 사용하여 체인에 전송되고 원장에 기록됩니다.

TransferEvent is the alias that RegisterAction() uses here. RegisterAction() is a method of the Action API and it takes four arguments that are transferred to the chain in order to record transaction details. The transaction hash and certain other details are output in the logs section.

 **transferMulti()** 도 비슷한 방법으로 작동합니다. transferMulti()가 실행되면 transfer()를 호출하기에 기본적인 논리구조는 거의 비슷합니다, 하지만 차이점이 있다면 여러 개의 거래가 동시에 처리된다는 점입니다. 인수는 nested array 1개를 필요로 합니다.

![](media/image17.jpg){width="5.2551924759405075in" height="1.7727274715660541in"}

The sub-array elements are processed in sets of three such that the first and second elements still represent the sender's and receiver's addresses, and the third element represents the transfer amount. The sub arrays are iterated till there are no more elements left in the args\[\] array.

만약에 sub-array가 3개의 인수를 포함하지 않거나 이전 거래가 특정 이유로 인하여 실패하였다면 예외처리가 됩니다. 예외처리가 되면 통제권은 loop에서 나와 Main()으로 넘어갑니다.

**approve(owner, spender, amount)**

![](media/image18.jpg){width="5.208333333333333in" height="3.1166666666666667in"}

The approve function implements another complex logic wherein an account, namely the "spender" is given the permission to utilize a certain amount in tokens from another account, namely the "owner".

The function first carries out address validation in terms of length, and user authentication for the "owner" who is about to perform the approval.

Next, the amount selected for approval is compared to the available balance in the "owner" account.

만약 잔고에 충분한 양이 있지 않다면 과정은 종료되고 통제권으 Main()에 넘어가게 됩니다. 

만약 잔고에 충분하 양이 있다면, APPROVAL\_PREFIX르 접두로 가진 "소유자" 주소와 "소비자" 주소를 연결한 키가 형성됩니다. It is then added to the ledger using the put() method by passing the context, the above generated key, and the approval amount.

거래 이벤ㅌ는 ApprovalEvent()방법으로 기록되고 NeoVM엔진이 반환한 거래 해시값이 포함된 결과값은 로그 섹션에 나타나게 됩니다. 

**transferFrom(spender, from\_acc, to\_acc, amount)**

![](media/image19.jpg){width="6.85in" height="5.2in"}

transferFrom() 함수는 더 복잡한 논리구조를 실행하고 and carries out a task that may prove to be useful for certain applications.

This function allows a third party, namely the spender, to utilize a certain amount in tokens that are provided from an account that does not designate to their own credentials, basically implementing the same logic as that of the approve() function. It takes four arguments, which are three Byte58 addresses and one transfer amount.

먼저, 함수는 주소의 유효성 검사를 실행합니다. 후에 소비자에게 Runtime API의 일부인 CheckWitness()함수를 사용하여 이 거래를 실행할 권한이 있는지 확인합니다.  

Next, the balance of the "from" account is fetched and cross-checked with the transaction amount to ensure the account has enough balance. The process to fetch the balance remains the same.

The spender's address is then prefixed with the APPROVE\_PREFIX and the approved amount from is fetched using the prefixed address.

거래는 그 후에 실행됩니다. 만약 거래 금액이 승인된 금액을 초과한다면, 거래는 중지되고(aborted) 통제권은 Main()함수에게 돌아가게 됩니다. 

만약에 승인된 금액이 정확히 양과 일치한다면, 거래량은 "보낸 사람"의 계좌 잔고에서 차감됩니다. 

만약 승인된 금액이 거래량을 초과한다면, 그 차액이 계산되어 추후에 참조할 수 있도록 원장에 기록됩니다. 그리고 그 거래량은 "보낸 사람" 계좌 잔고에서 차감됩니다.

거래량은 put()함수를 사용하여 "받는" 사람의 계정으로 전송됩니다. 후에 해당 이벤트는 기록되고 거래 해시와 함께 결과는 IDE의 로그 섹션에 나타납니다.

Another function that implements a similar logic has be defined as **allowance(owner, spender)** which facilitates querying the amount of allowance that has been allocated to the "spender" account from the "owner" account.

![](media/image20.jpg){width="4.633333333333334in" height="1.1666666666666667in"}

Practically speaking, this function cam be classified as an access function too in the sense that it returns the allowance value. But it also performs a get() query to fetch this result from the chain.

A key generated by concatenating the prefixed owner address and the spender address is passed to the get() method along with the context to fetch the required allowance value, which is then returned to Main(). 이 값은 확인할 수 있으며 다른 작업에 사용될 수도 있습니다.

5.  **배포와 테스팅**

![](media/image21.jpg){width="5.175in" height="1.5916666666666666in"}

논리 구조 구현이 완료되고 나면, 이제는 스마트 컨트랙트를 실행(compile)할 차례입니다. 

스마트 컨트랙트를 실행하고 나면, 그 결과는 IDE 화면에 다음과 같이 보입니다 - 

![](media/image22.jpeg){width="2.5737970253718285in" height="3.25in"}

compile 탭에 보이는, AVM 바이트 코드는 컴파일 후에 생성된 중 간결과물입니다. NEOVM 은 이 AVM 코드를 처리하여 계약을 실행합니다.

The opcode indicates the stack status line by line; an advanced debugging tool.

ABI stores the parameter information for all the functions and the contract hash itself. The contract ABI can be saved for future retrieval.

![](media/image23.jpg){width="7.268055555555556in" height="0.47727252843394574in"}

The logs section displays the compilers response which includes everything from debugging results to the information that the VM returns.

컴파일 결과 나타나는 오류가 모두 수정되거나 해결된 후에, 컴파일을 성공하면 이제 해당 컨트랙트는 배포될 수 있다.

![](media/image24.jpg){width="5.091666666666667in" height="6.133333333333334in"}

Here, we fill in the relevant information regarding the contract and proceed. A confirmation window pops up where you can enter the gas price and gas limit.

**알림: 가스 한계는 정확히 소수점 9자리까지 포함됩니다. 그렇기에, 10^9^ 가스 unit은 1 ONG 토큰과 일치한다고 볼 수 있습니다. 표현할 수 있는 값 중 가장 작은 값은 0.000000001 입니다.**

**The wallet automatically sets a suitable limit based on the complexity of the code being compiled and run. But you always have the option to set a limit yourselves. Ensure that limit is higher than the cost, otherwise the contract may fail to deploy or invoke.**

![](media/image25.jpg){width="1.9874278215223098in" height="3.113636264216973in"}

성공적으로 배포가 되면, 거래 해시값은 로그 섹션과 오른쪽 하단 결과창에서 확인할 수 있습니다.

![](media/image23.jpg){width="7.268055555555556in" height="1.6145833333333333in"}

다음으로, 이제는 컨트랙트를 실행하려고 합니다. 

Before executing other functions, we must first initialize the wallet using init() so as to ensure that it has enough balance to carry out transactions.

We can then choose the function that we want to execute from the drop-down menu.

At this point we have two options, pre-run and run. We can choose to directly run the contract, and the engine would run the AVM code generated before. Pre-run is an option which can be chosen to test run the contract to check if it runs as expected. There is no gas cost associated with pre-running a contract.

We can select the function that we want to execute from the drop-down menu, select the data type of the value to be passed and pass the argument in the blank field.

![](media/image26.jpg){width="5.066666666666666in" height="5.416666666666667in"}

![](media/image27.jpg){width="7.268055555555556in" height="0.5909087926509187in"}

All the results that are displayed in the logs section are in hexadecimal format. The tool option in the right-side pane provides several different conversion tools that are at the user's disposal which can be used to perform data conversions and a few other functions.

![](media/image28.jpg){width="4.983333333333333in" height="6.258333333333334in"}

거래가 실행될 때마다, 거래 해시값이 결과로 반환되고 해당 거래 해시값은 결과를 추적하는데 사용된다. Let us try another transaction, this time with actual tokens being transferred.

![](media/image29.jpg){width="5.108333333333333in" height="5.408333333333333in"}

Once you run the contract, you will be prompted to confirm the transaction and enter the gas price and gas limit that will used to calculate the gas which the contract consumes in its execution process.
(ONG)

가스 가격과 가스 리밋을 입력합니다. 거래가 성공적으로 전달이 되면 거래 해시갑이 로그 섹션에 나타날 것입니다.

![](media/image30.jpg){width="7.268055555555556in" height="0.5875in"}

To see the current balance of the account that we transferred the tokens to, we can pre-run the balanceOf() function. We pre-run it because we want to see the output of the function. Running it directly would execute it, but we will not be able to see the value returned.

![](media/image31.jpg){width="7.268055555555556in" height="0.7659722222222223in"}

Clearly, 50 units of our sample token have been transferred to the receiver's address. The displayed value 32 is the hexadecimal value.

![](media/image32.jpg){width="4.925in" height="1.6583333333333334in"}

이 16진수 형태 값(hex value)을 tools 섹션에 제공된 변환기(converter)를 통하여 decimal 값으로 변환할 수 있습니다. 

거래 해시 값은 나중에 탐색기에서 거래 세부사항을 보는데에 사용될 수 있습니다. 

![](media/image33.jpeg){width="7.268055555555556in" height="3.454861111111111in"}

The last section in the right-hand pane, Restful, is the API that is used to communicate with the chain to fetch useful information regarding the chain's status, such as the current block height or the smartcode for an event. Smartcode 는 사실상 스마트 컨트랙트 정보를 JSON기반으로 표현한 것이다. JSON 은 쿼리를 기반으로 생성되고 반환되며 이때 함수들은 Restful API로 호출된다.

![](media/image34.jpg){width="4.25in" height="5.99959864391951in"}

The transaction hash can be used to fetch the smartcode for the respective transaction.

The Restful API can be accessed from SDKs too, but SDK integration is another tutorial all by itself.

![](media/image35.jpg){width="3.2727274715660544in" height="3.0673534558180227in"}
