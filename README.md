# Parallel Distributed Processing in AWS
AWSサービスにおける並列分散バッチ処理の検証

## 検証サービス
1. Step Functions（Lambda）
2. Step Functions（ECS）
3. AWS Batch（Fargate）

## 検証内容
`usersテーブル`上の各ユーザーに対して`num`数分のTweetデータを作成する。<br>
そして、そのTweetデータを`tweetsテーブル`に格納する。<br>

#### # usersテーブル
|  id  |  name  |  num  |
| ---- | ---- | ---- |
|  1  |  akio  |  50  |
|  2  |  miki  |  50  |

#### # tweetsテーブル
|  twid  |  tw_user_id  | text |  created_at  |
| ---- | ---- | ---- | ---- |
|  27463  |  1  |  ほにゃらら  |  2022-06-30 12:57:25  |
|  59349  |  2  |  hogehoge  |  2022-06-30 12:57:25  |

## 1. Step Functions（Lambda）
### ◇ 構成
![stepfunctions_graph](StepFunctions-Lambda/stepfunctions_graph.png)

### ◇ 実行時間
#### # 検証状況
* users：500
* num：50（一律で設定）
#### # 検証結果
* 実行時間：18.4s
* tweets：25,000

#### # 考察
単体のworkerだと1sほどで完了するため実行に時間がかかっている。<br>
⇒ 各workerの起動時刻にラグがある<br>
⇒ タスクの受け渡しやLambdaの起動にオーバーヘッドがある

とはいえ単一で実行する時と比べて圧倒的に速い

## 2. Step Functions（ECS）
上記の `Step Functions（Lambda）` においてECSに変更するだけのため作成は割愛

## 3. AWS Batch（Fargate）
### ◇ 構成

### ◇ 実行時間
