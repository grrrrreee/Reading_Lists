## *NOTE: Ignore the broken links and placeholders for images.*

Smart contract development using NeoVM and SmartX -- The Basics

1.  **Prerequisites**

실제 개발 과정을 진행하기 전에, 필요한 툴들을 가지고 있는지 확인해야합니다. 

아래의 목록이 온톨로지 플랫폼을 활용한 스마트 컨트랙트 개발 과정에서 사용되는 필수적인 툴들입니다. 

-   SmartX -- 온톨로지 온라인 스마트컨트랙트 IDE 및 디버거. 

-   Cyano wallet -- 구글 크롬 확장 프로그램

-   Explorer -- 전반적인 블록체인 활동과 거래내역을 추적할 수 있는 웹 기반 공개 툴 

2.  **환경 설정**

개발 환경설정 과정은 복잡하고 시간이 많이 소요되는 작업일 수 있습니다. 그래서 이 과정은 최대한 swift 하고 편하게 하여 많은 사람들이 어려움을 느끼지 않게 하기 위해 노력했습니다. 이 과정에서 사용될 툴들은 모두 웹기반이고 실용적이며 바로 사용할 수 있습니다. 

먼저, 편의를 극대화 시키기 위해서, 우리는 스마트 컨트랙트를 개발 및 테스팅을 테스트넷에서 진행할 예정입니다. 그리고 이를 통해 프라이빗 체인의 필요성을 없앨 계획입니다. (thereby eliminating the need of a private chain.) A private blockchain network architecture can be set up just as easily on your local environment using Ontology's Punica suite, a set of development tools that allows for smart contract deployment and testing on the private net, but that is beyond the scope of this tutorial.

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
: 테스트넷에서 스마트 컨트랙트 배포 및 테스팅하는 과정은 메인넷의 ONT/ONG의 잔고없이도 진행할 수 있습니다. But, to pay the gas cast of deploying a contract on the test net, you will still need a nominal testnet ONG balance. The gas cost is calculated by taking the product of gas price and gas limit(gas price \* gas limit). Test tokens have been made available for free by Ontology and can be applied for by following

given link:
[[https://developer.ont.io/applyONG]{custom-style="Hyperlink"}](https://developer.ont.io/applyONG)

![](media/image3.jpeg){width="5.356871172353456in"
height="2.6287871828521436in"}

![](media/image4.jpeg){width="7.242424540682415in"
height="3.5619203849518812in"}

The IDE that we will be using is Ontology's SmartX, a browser-based development environment that supports Python, C\#, and JavaScript(coming soon). We are going to take an in-depth look at the development process using Python, as far as this tutorial is concerned.

NeoVM serves as the execution engine for the programs written using Python in SmartX. SmartX core integrates all of Ontology's APIs, and so all the different functions which allow us to perform blockchain related tasks can be used directly by importing the relevant API, which we will be discussing later in this tutorial.

![](media/image5.jpeg){width="7.268055555555556in" height="3.609027777777778in"}

탐색기는 온라인 웹기반 툴로 거래를 추적하고 블록체인 활동 관련 이벤트들의 진행상황 및 결과를 확인할 수 있게 합니다. 테스트넷과 메인넷 2개 모두의 거래 해시, ONT ID, 컨트랙트 주소 그리고 블록 높이 등에 대한 정보를 얻을 수 있습니다. 

3.  **Launching SmartX**

주의: Testing smart contracts requires a Cyano account but does not need ONT/ONG balance.

![](media/image4.jpeg){width="7.242424540682415in" height="3.5619203849518812in"}

![](media/image6.jpeg){width="7.268055555555556in" height="3.5527777777777776in"}

Once logged in you will be prompted to select a project. Create a new one if there are no existing projects.

![](media/image7.jpg){width="2.2229440069991253in" height="2.3939391951006126in"}

Upon selecting the new project option, you can choose a programming language to work with. In this tutorial, since our focus is on developing smart contracts with NeoVM, we will be illustrating examples and sample code using Python.

![](media/image8.jpeg){width="7.268055555555556in" height="3.5375in"}

파이썬을 선택하고 다음 메뉴로 진행하십시오. Here, a list populated with a set of examples and templates of smart contracts will be displayed. You can choose any example and go through the sample code or use it as a base to give structure to your own smart contract logic. For this tutorial we will be looking at the OEP-4 illustration.

4.  **프로그램 논리구조 작성 시작**

![](media/image9.jpeg){width="7.268055555555556in" height="3.597916666666667in"}

SmartX가 설정되고 작동되기 시작하면 그 화면은 다음과 같을 것입니다. 

Since we have selected the OEP-4 template, the code is already present in the editor area. You may choose to edit this code as you please based on the logic that you're trying to implement. But for the sake of simplicity and staying within the scope of this tutorial we will use the code as it is to ensure uniformity.

It is worth taking some time to take a look at the code and the overall structure of the program that we'll be working with.

![](media/image10.jpg){width="7.258333333333334in" height="1.525in"}

The variables declared in this section of the code define the protocol itself and the specifics that will govern its functionality.

"NAME" 그리고 "SYMBOL" 와 같은 변수는 토큰의 식별자로 사용됩니다. 

"FACTOR" is the base 10 value that dictates the precision of amounts that can be transferred. For example, if the value is set to 100, transfer amounts with values up to two levels of precision are supported by the system. 

"DECIMALS" stores the multiplier value for access convenience.

"OWNER" stores the Base58 address of the entity or account that holds the authority over the totality of tokens and can choose to distribute them as and when needed.

"TOTAL\_AMOUNT" 는 현재 존재하는 토큰의 총량을 stores the total number of tokens that exist. Always a fixed number.

"BALANCE\_PREFIX" is an access modifier that is used with account addresses for authentication purposes. "APPROVE\_PREFIX" serves the same purpose, but for the approve operation wherein the owner can authenticate another account to use tokens.

"SUPPLY\_KEY" correlates directly to the total amount of tokens and is used for any operations that may be carried with the total tokens figure, since the value isn't directly accessible.

![](media/image11.jpg){width="1.3in" height="0.19166666666666668in"}

This line immediately stands out in the upper section of the code.

GetContext() 는 스마트 컨트랙트와 블록체인의 다리 역할을 하는 함수입니다. 해당함수는 Storage API의 일부인 GET과 PUT를 호출하여 체인에 혹은 체인으로부터 정보를 fetching 하거나 transmitting 할 때 사용됩니다. 

We will go through the relevant APIs as we come across the respective functions, and the full set of available APIs in a later section of the tutorial.

![](media/image12.jpg){width="5.083333333333333in"
height="0.8083333333333333in"}

해당 과정은 SmartX IDE를 사용하고 있기 때문에, 과정에서 사용되는 모든 API들과 함수들은 프로그래밍을 시작할 때 importing하여 바로 접근 및 사용할 수 있습니다. 과정에서 우리가 import 할 함수들은 Runtime API에서 GetContext(), Get(), Put(), and Delete() from the Storage API, Notify(), CheckWitness() Base58ToAddress() 이며 빌트인 함수(built-in function) concat() 입니다. 

알림: In later versions, built-in functions don't need to be imported and can be called directly.

Let's take a look at the Main()
function.![](media/image13.jpg){width="5.189394138232721in" height="4.5815594925634295in"}

The Main() function takes two arguments, *operation* and *args*. The *operation* argument is based on the operation to be performed and dictates the function to the executed. The *args* argument helps passing the important information that a function needs to carry out further execution, for example account addresses or input data.

Here, clearly there are 11 different functions that can be called depending upon the argument that is passed in Main(). SmartX passes these arguments using the "Options" pane on the bottom right.

-   **init()** **:** init() 함수는 프로그램 논리 구조의 시작점을 설정합니다. 해당 함수는 initializes the definition variables declared at the top based on the values provided. Thus, this is the first function that needs to be executed post deployment.

-   **name() :** 해당 함수는 토큰에 설정된 이름을 반환합니다. 이번 경우에는 "My Token" 입니다. 

-   **symbol() :** 해당 함수는 토큰의 symbol을 반환합니다. 이번 경우에는 "MYT"입니다.

-   **decimals() :** 유효한 토큰 값의 정확도를 지정하는 소수점을 반환합니다, 이 경우에는 8자리입니다.

-   **totalSupply() :** Returns the total number of tokens assigned while initializing. Denotes the fixed number of tokens allocated for circulation. (Uses SUPPLY\_KEY to fetch the value from the chain, stored earlier during initialization)

-   **balanceOf(acct) :** Fetches the corresponding token balance of the account that identifies with the Base58 address passed as argument to the function.

-   **transfer(from\_acc, to\_acc, amount) :** Transfers the equivalent token value passed as the amount argument to the to\_acc address from the from\_acc address.

-   **transferMulti(args) :** The parameter here is an array that contains the same information, i.e., the sender's address, receiver's address, and the amount to be sent, in that sequence at the respective indices. It can be iterated for multiples transactions by passing the respective account addresses and the corresponding amount.

-   **transferFrom(spender, from\_acc, to\_acc, amount) :** The spender here takes a certain amount of tokens from the from\_acc address, and transfers them to the to\_acc address.

-   **approve(owner, spender, amount) :** The owner authorizes the spender to use a certain amount of tokens from their own account. Both the owner and spender arguments here are Base58 addresses and the amount specifies the amount that the spender is authorized to spend.

-   **allowance(owner, spender) :** This function can be used to the amount that the owner account has authorized spender to use. Both arguments in this case are Base58 addresses.

위에 언급된 함수들은 2가지 종류로 나누어집니다. - 접근형 함수와 유틸리티형입니다. Let us consider the flow of control as these functions are called.

Access functions are primarily used to fetch data post contract deployment. Functions such as name(), symbol(), totalSupply() and balanceOf(acc) allow us to achieve this by using get() function from the Storage API which fetches relevant data from the chain. Let us look at how it is implemented in program logic.

**name() function definition**

![](media/image14.jpg){width="2.1212117235345582in" height="0.8444630358705162in"}

This is how a simple access function can be defined. The name() function takes no arguments. Even though functions such as name(), symbol() and decimals() do not explicitly call the get() function, after the contract is deployed, all the data is fetched from the chain.

**balanceOf(acc) function definition**

![](media/image15.jpg){width="3.25in" height="1.252205818022747in"}

The balanceOf() function takes one argument, a Base58 address which denotes an account. There is a validity check in place that verifies the length of the address and raises an exception if the address is invalid. If the address is valid the get() function is called with two arguments, the account address prefixed with BALANCE\_PREFIX, and the context.

The context allows for data reference on the chain to fetch the account balance value, while the prefixed account address ensures authenticated access. Next, get() returns this data to Main() where it is output to the log window using the notify() function. The totalSupply() function works in a similar fashion.

알림：잔액 및 approve prefixes 은 ASCII 형식의 16진수 값이며 당신의 프로그램 논리 구조에 맞게 변경될 수 있습니다.

Let us look at the utilities in the sample code.

Utilities are methods that have richer functionality and help modifying the on-chain data, which is the basis for transactions and tasks that may be carried out. The OEP-4 token logic that the code template is using illustrates various use cases and scenarios in the form of functionality. This is to exhibit just how versatile smart contracts are in nature and the different kinds of business logic that they can be used to generate.

**transfer(from\_acc, to\_acc, amount)**

![](media/image16.jpg){width="4.95in" height="3.6666666666666665in"}

The transfer() function implements the most fundamental transaction feature, transferring tokens from one account to another. It takes three arguments, the sender's address, the receiver's address, and the amount to be transferred.

해당 함수는 간단한 길이 검사로 검증을 진행하지만, 원한다면 더 복잡한 논리구조를 개발 및 구현할 수 있습니다. 

Next, the BALANCE\_PREFIX is concatenated to the sender's account address, and balance is retrieved by making a get() call using this address. A quick comparison is made in the next step where the sender account's balance is compared with the amount to be transferred. All three scenarios have been defined clearly.

만약에 잔고가 거래량보다 적다면, 거래는 실패하게 되고 통제권은 바로 Main() 으로 넘어가게 된다.

If the amount equates to the balance amount exactly, the balance of sender account is set to 0 by calling the delete() method using the sender's prefixed address. This is practically equivalent to using the put() method to manually assign the value 0 to sender's account but using put() method in this case might give rise to security vulnerabilities.

만약 잔고가 거래량보다 높다면, the amount is deducted from the balance by making a put() call and updating the sender accounts balance with the deducted value.

Next, the receiver's address is prefixed with the BALANCE\_PREFIX, and the prefixed address is used to add the transfer amount to the receiver's account.

마지막으로, 해당 거래는  Finally, this transaction event is sent to the chain using the RegisterAction() method for recording in a ledger.

TransferEvent is the alias that RegisterAction() uses here. RegisterAction() is a method of the Action API and it takes four arguments that are transferred to the chain in order to record transaction details. The transaction hash and certain other details are output in the logs section.

 **transferMulti()** 도 비슷한 방법으로 작동합니다. transferMulti()가 실행되면 transfer()를 호출하기에 기본적인 논리구조는 거의 비슷합니다,  but it allows for multiple transfers to take place simultaneously. It takes one argument which is a nested array.

![](media/image17.jpg){width="5.2551924759405075in"
height="1.7727274715660541in"}

The sub-array elements are processed in sets of three such that the first and second elements still represent the sender's and receiver's addresses, and the third element represents the transfer amount. The sub arrays are iterated till there are no more elements left in the args\[\] array.

Exception is thrown in case the sub-array does not contain exactly three elements, or the previous transaction fails for some reason, and the control comes out of the loop and goes back to Main()

**approve(owner, spender, amount)**

![](media/image18.jpg){width="5.208333333333333in"
height="3.1166666666666667in"}

The approve function implements another complex logic wherein an account, namely the "spender" is given the permission to utilize a certain amount in tokens from another account, namely the "owner".

The function first carries out address validation in terms of length, and user authentication for the "owner" who is about to perform the approval.

Next, the amount selected for approval is compared to the available balance in the "owner" account.

If the account does not have enough balance the process is terminated, and the control returns to Main().

If the account has enough balance, a key is generated by concatenating the "owner" address prefixed with APPROVAL\_PREFIX and the "spender" address. It is then added to the ledger using the put() method by passing the context, the above generated key, and the approval amount.

The transaction event is then recorded using the ApprovalEvent() method, and the result containing the transaction hash returned by the NeoVM engine is displayed in the logs section.

**transferFrom(spender, from\_acc, to\_acc, amount)**

![](media/image19.jpg){width="6.85in" height="5.2in"}

transferFrom() 함수는 더 복잡한 논리구조를 실행하고 and carries out a task that may prove to be useful for certain applications.

This function allows a third party, namely the spender, to utilize a certain amount in tokens that are provided from an account that does not designate to their own credentials, basically implementing the same logic as that of the approve() function. It takes four arguments, which are three Byte58 addresses and one transfer amount.

First, the function carries out the conventional address validation. Then it verifies whether the spender has the authorization to carry out this transaction using the CheckWitness() function which is a part of the Runtime API.

Next, the balance of the "from" account is fetched and cross-checked with the transaction amount to ensure the account has enough balance. The process to fetch the balance remains the same.

The spender's address is then prefixed with the APPROVE\_PREFIX and the approved amount from is fetched using the prefixed address.

The transaction comes next. 만약 거래 금액이 승인된 금액을 초과한다면, 거래는 중지되고(aborted) 통제권은 Main()함수에게 돌아가게 됩니다. 

만약에 승인된 금액이 정확히 양과 일치한다면, 거래량은 "보낸 사람"의 계좌 잔고에서 차감됩니다. 

만약 승인된 금액이 거래량을 초과한다면, 그 차액이 계산되어 추후에 참조할 수 있도록 원장에 기록됩니다. 그리고 그 거래량은 "보낸 사람" 계좌 잔고에서 차감됩니다.

The transaction amount is then transferred to the "to" account using the put() function. The event is then recorded and the result with the transaction hash is displayed in the logs section of the IDE.

Another function that implements a similar logic has be defined as **allowance(owner, spender)** which facilitates querying the amount of allowance that has been allocated to the "spender" account from the "owner" account.

![](media/image20.jpg){width="4.633333333333334in"
height="1.1666666666666667in"}

Practically speaking, this function cam be classified as an access function too in the sense that it returns the allowance value. But it also performs a get() query to fetch this result from the chain.

A key generated by concatenating the prefixed owner address and the spender address is passed to the get() method along with the context to fetch the required allowance value, which is then returned to Main(). The value can then be displayed or used to perform other tasks.

5.  **배포와 테스팅**

![](media/image21.jpg){width="5.175in" height="1.5916666666666666in"}

논리 구조 구현이 완료되고 나면, 이제는 스마트 컨트랙트를 컴파일링 할 차례입니다. 

Upon compiling the smart contract, the following results can be seen in the IDE-

![](media/image22.jpeg){width="2.5737970253718285in" height="3.25in"}

In the compile tab, the AVM byte code is the resulting intermediate code produced after compilation. NeoVM processes this AVM code to execute our contract.

The opcode indicates the stack status line by line; an advanced debugging tool.

ABI stores the parameter information for all the functions and the contract hash itself. The contract ABI can be saved for future retrieval.

![](media/image23.jpg){width="7.268055555555556in" height="0.47727252843394574in"}

The logs section displays the compilers response which includes everything from debugging results to the information that the VM returns.

After all the compilation errors have been dealt with and the contract is successfully compiled, we can proceed to deploy it.

![](media/image24.jpg){width="5.091666666666667in" height="6.133333333333334in"}

Here, we fill in the relevant information regarding the contract and proceed. A confirmation window pops up where you can enter the gas price and gas limit.

**알림: 가스 한계는 정확히 소수점 9자리까지 포함됩니다. 그렇기에, 10^9^ 가스 unit은 1 ONG 토큰과 일치한다고 볼 수 있습니다. with the minimum valid value being 0.000000001.**

**The wallet automatically sets a suitable limit based on the complexity of the code being compiled and run. But you always have the option to set a limit yourselves. Ensure that limit is higher than the cost, otherwise the contract may fail to deploy or invoke.**

![](media/image25.jpg){width="1.9874278215223098in" height="3.113636264216973in"}

Once deployed successfully, the transaction hash will be displayed in the logs section and in the result pane in the bottom right.

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

Every time a transaction is executed, a transaction hash will be returned that can be used to track the results. Let us try another transaction, this time with actual tokens being transferred.

![](media/image29.jpg){width="5.108333333333333in" height="5.408333333333333in"}

Once you run the contract, you will be prompted to confirm the transaction and enter the gas price and gas limit that will used to calculate the gas which the contract consumes in its execution process.
(ONG)

가스 가격과 가스 리밋을 입력하고 and provide the confirmation.  After the transaction is carried out successfully the transaction hash will be displayed in the logs section.

![](media/image30.jpg){width="7.268055555555556in" height="0.5875in"}

To see the current balance of the account that we transferred the tokens to, we can pre-run the balanceOf() function. We pre-run it because we want to see the output of the function. Running it directly would execute it, but we will not be able to see the value returned.

![](media/image31.jpg){width="7.268055555555556in" height="0.7659722222222223in"}

Clearly, 50 units of our sample token have been transferred to the receiver's address. The displayed value 32 is the hexadecimal value.

![](media/image32.jpg){width="4.925in" height="1.6583333333333334in"}

이 16진수 형태 값(hex value)을 tools 섹션에 제공된 변환기(converter)를 통하여 decimal 값으로 변환할 수 있습니다. 

거래 해시 값은 나중에 탐색기에서 거래 세부사항을 보는데에 사용될 수 있습니다. 

![](media/image33.jpeg){width="7.268055555555556in" height="3.454861111111111in"}

The last section in the right-hand pane, Restful, is the API that is used to communicate with the chain to fetch useful information regarding the chain's status, such as the current block height or the smartcode for an event. Smartcode is basically a JSON based representation of smart contract information. The JSON is generated and returned based on the query, i.e., the function called in the Restful API.

![](media/image34.jpg){width="4.25in" height="5.99959864391951in"}

The transaction hash can be used to fetch the smartcode for the respective transaction.

The Restful API can be accessed from SDKs too, but SDK integration is another tutorial all by itself.

![](media/image35.jpg){width="3.2727274715660544in" height="3.0673534558180227in"}
