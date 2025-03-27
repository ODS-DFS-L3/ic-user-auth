## 概要・目的
本リポジトリでは、ウラノス・エコシステムにおける技術参照文書
「ウラノス・エコシステム・データスペーシズリファレンスアーキテクチャモデル（ODS-RAM）」の参照実装のうち、アイデンティティレイヤ（L3）に規定されたアイデンティティコンポーネント※について公開する。

ODS-RAM、ODS-FDCの詳細については[こちら](https://www.ipa.go.jp/digital/architecture/reports/ouranos-ecosystem-dataspaces-ram-white-paper.html)を参照すること

※アイデンティティコンポーネントはODSでは仕様検討中であり本リポジトリ群では2024年度デジタルライフライン全国総合整備計画におけるデータ連携実証の中で使用したユーザ認証システムを提供する。

## 基本概念
「ウラノス・エコシステム（Ouranos Ecosystem）」とは、デジタル化をイネーブラーとして、企業と企業をつなぐビジネス・デジタルの協調領域を整備し、利⽤可能とすることにより、産業界全体として新たな連携による価値を⽣み出すエコシステムである。

ウラノス・エコシステムでは産業界がデータスペースの社会実装を早急に進めるためのサービスライフサイクルに焦点をおいたアーキテクチャモデルである「ウラノス・エコシステム・データスペーシズリファレンスアーキテクチャモデル（ODS-RAM）」を公開しており、その中でデータスペースコネクタの参照実装としてODS Flex Dataspace Connector（ODS-FDC）およびアイデンティティコンポーネントの参照実装としてODS Identity Component（ODS-IC）をOSS公開する。

[ODS-RAM](https://www.ipa.go.jp/digital/architecture/reports/ouranos-ecosystem-dataspaces-ram-white-paper.html)より、ODS-FDCのシステムアーキテクチャを抜粋して示す。

![ODS-FDCシステムアーキテクチャ](./docs/architecture/picture/ODS-FDC.png)

ODS-RAM40頁参照


# 環境構築手順
本リポジトリ内のユースケースを体験するための環境構築の手順を示す。

## ライブラリ・開発フレームワーク・ビルドツール・各種提供形態のパッケージ
本リポジトリのシステムが動作するための動作環境の一式を事前にインストールする。
- go関係のソフトウェアをインストール
  - golang
  - golangci-lint
  - mockery
  - goreturns
- マイグレーションツール（golang-migrate）のインストール
- makeをインストール
- docker, docker-composeをインストール

### 動作確認済み実行環境
|Name                                        |Version |Notes|
|:-------------------------------------------|:-------|:----|
|golang                                      |1.22||
|[golangci-lint](https://golangci-lint.run/usage/install/)|1.56.2||
|[mockery](https://vektra.github.io/mockery/)|v2.42.3||
|[goreturns](https://github.com/sqs/goreturns)|-|go install|
|[golang-migrate](https://github.com/golang-migrate/migrate)|-||
|make                                        |GNU make 3.81||
|docker                                      |-||
|docker-compose                              |-||

アプリケーションで利用しているOSSやMWについては[開発者向けドキュメント](/docs/architecture/architecture.md)を参照

## ビルド・起動手順
本リポジトリでは、A社の事業者を作成し、ユーザの認証を実施する。

以下、`path/to/authenticator-backend` はローカル環境に `git clone` されたユーザ認証システムへのパスを表す。

コマンドについては、Windows環境（PowerShell）でのコマンド例を記載する。

### 1.必要な外部システムの起動

以下の外部システムが必要のため、起動する
- PostgreSQL : データベースシステム
- Keycloak : 認証・認可サービス

```shell
# Postgresの起動およびKeycloakの起動
cd <path/to/authenticator-backend>
docker-compose up -d
```

### 2. 認証データのセットアップ

1. データベースのスキーマ定義を反映し、A社のデータをDBに必要な情報を作成
```shell
cd <path/to/authenticator-backend>
$env:POSTGRESQL_URL='postgres://dhuser:passw0rd@localhost:5432/dhlocal?sslmode=disable'
migrate -path setup/migrations_tutorials -database $env:POSTGRESQL_URL up
```
補足. Postgresが導入済の環境で、導入されているPostgresを利用したい場合は下記を実施
```shell
# データベースのスキーマ定義を反映
cd <path/to/authenticator-backend>
$env:POSTGRESQL_URL='postgres://dhuser:passw0rd@localhost:5432/dhlocal?sslmode=disable'
migrate -path setup/migrations -database $env:POSTGRESQL_URL up
```
```shell
# A社のデータをDBに必要な情報を作成
cd <path/to/authenticator-backend>
# 下記コマンドが実行できなかった場合、GitBashやWSL等のツール導入を行い下記コマンドを実行すること
./setup/setup_seeds.sh
```

[確認]
- DBに接続し、operatorsテーブル(事業者に関する情報を保存するためのテーブル)のカラム数を確認

2. A社の認証情報をKeycloakに作成
※ 事業者情報を初期値から変更する場合は```/cmd/add_tutorial_user/data/seed.csv```を編集する。
```shell
cd <path/to/authenticator-backend>
# ローカルのKeyclockに事業者情報を追加
make keycloak-add-tutorial
```

[確認]
- http://localhost:4000 にアクセスし、ユーザ名：admin、パスワード：passwordでログインを行った後、KeycloakのUsersにデータが入っていることを確認

3. 各種情報の追加・変更
- 事業者識別子(内部)の追加
  - 左メニューから『Realm settings』を押下し、『User profile』タブを開き、『Create attribute』を押下
  - 「Attribute[Name]」に'operator_id'、「Display name」に'${operator_id}'を入力し、『Create』を押下
  - 以上の操作により、ユーザーのGeneralにoperator_idが追加される。
  - 左メニューから『Users』を押下し、「Username」が「oem_a@example.com」リンクを押下
  - 「operator_id」に'b39e6248-c888-56ca-d9d0-89de1b1adc8e'を入力し、『Save』を押下
  - 以上の操作により、対象ユーザーの「operator_id」に事業者識別子(内部)が追加される
  - 左メニューから『Clients』を押下し、「Client ID」が「UserAuthenticationSystem」リンクを押下
  - 『Client Scopes』タブを開き、「UserAuthenticationSystem-dedicated」リンクを押下
  - 『Configure a new mapper』を押下し、『User Attribute』を押下
  - 「Name」「Token Claim Name」に'operator_id'を入力、「User Attribute」のプルダウンから'operator_id'を選択し、『Save』を押下

- クライアントの設定変更
  - 左メニューから『Clients』を押下し、「Client ID」が「UserAuthenticationSystem」リンクを押下
  - 『Settings』タブから「Capability config」において以下の設定をOnまたはチェックを入れ、『Save』を押下
    - Client authentication
    - Authorization
    - Standard flow
    - Direct access grants
    - Service accounts roles
    - OAuth 2.0 Device Authorization Grant

- アクセストークン有効時間の変更
  - 左メニューから『Realm settings』を押下し、『Tokens』タブを開き、「Access Token Lifespan」の値を'15'(Minutes)に変更し、『Save』を押下  
   ※「Access Token Lifespan」については任意時間となる

### 3. ユーザ認証システムのビルドと起動

1. ビルド手順

```shell
cd <path/to/authenticator-backend>
go build main.go
docker build -t authenticator-backend .
```

2. 起動手順

```shell
docker run -v ${PWD}/config/:/app/config/ -td -i --network docker.internal --env-file config/local.env -p 8081:8081 --name authenticator-backend authenticator-backend
```

# 参考実装
環境準備が整ったら、下記の業務フローのチュートリアルを実行できる。

- [ユーザ主体認証による事業者認証](/docs/tutorials/OperatorIdentification.md)

# ドキュメント
開発者向けのドキュメントは[こちら](/docs/architecture/architecture.md)を参照

# ライセンス
- 本リポジトリはMITライセンスで提供されています。
- ソースコードおよび関連ドキュメントの著作権は株式会社NTTデータグループ、株式会社NTTデータに帰属します。

# 免責事項
- 本リポジトリの内容は予告なく変更・削除する可能性があります。
- 本リポジトリの利用により生じた損失及び損害等について、いかなる責任も負わないものとします。
