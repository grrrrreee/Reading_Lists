# Mimblewimble 

## Andrew Poelstra

## 2016-10-06(commit e9f45ec)

### Abstract

tom elvis 이름을 사용하는 익명의 사람이 비트코인 조사 채널 IRC에 서명을 하여 Tor hidden service[Jed16] 호스트된 문서를 삭제하였다 Mimblewimble이라는 제목의 이 문서는 비트코인과 근본적으로 다른 거래 구조를 가진 블록체인을 묘사했다. 단절적 합병 와 거래중단, 비밀거래 그리고 새유저들이 코인의이력 증명을 요구하지 않는 현재 chainstate의 완벽한 증명 을 지원한다 불행하게도 그 논문은 주요 아이디어들을 간결하게 적혀 있지만 그것은 보안의 대한 논쟁도 없고 심지어 실수도 하나 있다. 이 논문의 목적은 원래의 아이디어를 구체화 시키는 것이고, 저자가 개발한 확장 개선사항을 추가하는 것이다 특히 Mimblewimbled은 비트코인 기록이 있는 체인의 모든거래 을 기록하는데 15기가 데이터밖에 안 든다 ( 범위증명이 포함되고 100기가 가 필요한 UTXO set은 포함시키지 않는다) Jedusor는 이것을 어떻게 줄일 것인가에 대한 문제를 열어뒀다 우리는 이것을 해결하고, 15Gb를 1메가바이트 이하로 줄이기 위해 작업 증명 블록체인 압축에 대한 연구와 결합한다

### contents

### 1. Introduction 



#### 1.1 Overview and Goals



#### 1.2 Trust Model



### 2. Preliminaries 

#### 2.1 Cryptographic Primitives 



##### 2.1.1 Standard Primitives 



##### 2.1.2 Sinking Signatures



#### 2.2 Mimblewimble Primitives



### 3. The Mimblewimble Payment System

#### 3.1 Fundamental Theorem



#### 3.2 The Blockchain



#### 3.3 Consensus




##### 3.3.1 Block Headers



##### 3.3.2 Proven and Expected Work



##### 3.3.3 Sinking Signatures and Compact Chains 




### 4. Extensions and Future Research



### 5. Acknowledgements



### References 