# Project Roy Coin
![screensh](imgs/coin.jpg)
# 프로젝트 설명
### 목표 : 코인에 사용되는 기술에 대한 이해와 Go lang 공부 
<br>

### 주 사용 언어 : Go (v1.16.6)
<br>

### 구현 하고자 하는 기능
  - 작업증명
  - 채굴
  - 보상
  - P2P 거래
  - 지갑 
<br>
<br>
> > > ### 최종 결과물 :  **로이 코인**의 탄생

<br>
ps. Nomad Corder 노마드 코인 클론 코딩

<br>






------

## To Do


|To Do|작업 명|시작 일|종료 일| 
|:---:|:---:|:---:|:---:| 
|[ x ]|[블록 체인 구조체](https://github.com/abc7468/roycoin/blob/main/blockchain)|2021-07-25|2021-07-25| 
|[ x ]|[서버 사이드 렌더링 웹 사이트](https://github.com/abc7468/roycoin/tree/main/explorer/templates)|2021-07-26| 2021-07-27 |
|[ x ]|[Rest API](https://github.com/abc7468/roycoin/blob/main/rest/rest.go)| 2021-07-27 | 2021-07-29 |
|[ x ]|[CLI](https://github.com/abc7468/roycoin/blob/main/cli/cli.go)| 2021-07-31 | 2021-07-31 |
|[ x ]|[DB](https://github.com/abc7468/roycoin/blob/main/db/db.go)| 2021-08-02 | 2021-08-06 |
|[ x ]|채굴 & 작업증명| 2021-08-07 | 2021-08-09 |
|[ x ]|보상| 2021-08-12 | 2021-08-16 |
|[  ]|지갑| 2021-08-17 | ~ |
|[  ]|P2P 거래| ~ | ~ |


---
## Folders
- explorer
  - templates
    - pages : 라우팅 될 **페이지**
    - partials : 페이지에 로드될 **블락** 

- blockchain
  - block : block의 Struct와 Method
  - chain : blockchain의 struct와 Method

- rest : restAPI

- cli : terminal에서 flag설정을 해주는 cli 작성

- utils : 구현에 도움을 주는 helper 코드

- db : persistance를 위한 함수(Bolt DB)

- wallet : 지갑을 구현하기 위한 Struct와 Method