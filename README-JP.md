# Go - ビジネスロジック層におけるトランザクション

このプロジェクトはサービスレイヤーにトランザクションを実装します。

MVC パターンのアーキテクチャでは、リポジトリがすべてのデータベース関連のタスクを含み、サービス/Usecase がビジネスロジックを含みます。

## Go でのトランザクション

あるプロジェクトのエンドポイントが、多くのデータベーステーブルにエフェクトを生成するとします。そのような場合、これらのデータベースアクションの 1 つが失敗すると、対応するデータベーステーブルには他のデータが保存されません。そこで、トランザクションは他のテーブルがデータベースに保存されないようにするテクニックを提供します。

## ゴームトランザクション

人気のある ORM ライブラリである[Gorm](https://gorm.io/docs/transactions.html)は、このトランザクション機能を提供しています。

Gorm トランザクションの基本例を以下に示します：

```bash
// トランザクションを開始する
tx := db.Begin()

// トランザクションの中でいくつかのデータベース操作を行う (ここからは 'db' ではなく 'tx' を使用する)
tx.Create(...)

// ...

// エラーの場合はトランザクションをロールバックする
tx.Rollback()

// あるいはトランザクションをコミットする
tx.Commit()

```

#### 参照

- [Gorm における手動トランザクション](https://gorm.io/docs/transactions.html#Control-the-transaction-manually)

これは最も基本的なトランザクションで、リポジトリ層のメソッド内部で使用できます。

例えば、商品を購入する際、商品価格がユーザー残高から差し引かれ、商品在庫も商品テーブルから差し引かれます。これは商品を購入する際の単純なビジネスロジックです。

この 2 つのテーブルのトランザクションを、リポジトリ内の 1 つの `purchaseProduct()` メソッドで行えば、[SOLID の原則](https://s8sg.medium.com/solid-principle-in-go-e1a624290346) の単一責任の法則を排除することができます。
また、テストケースの作成も難しくなります。このプロセスはビジネスで必要なロジックなので、この `purchaseProduct()` はサービスレイヤーに置くべきです。このメソッドの内部では、データベースに関連する部分を実行するために、2 つの異なるリポジトリメソッド `reduceStockAmount()` と `reduceBalance()` を呼び出します。

## サービスレベルのトランザクション

サービスレベルにはデータベースロジックが含まれておらず、サービスメソッド内で別々のリポジトリメソッドが呼び出されるため、最初のセクションの例で見たようなトランザクションをサービス内で設定することはできません。このため、ミドルウェアを作成する必要があります。

### transaction.go

このファイルには、サービス層レベルでトランザクションを管理するための UoW (Unit of Work) ミドルウェアを作成するためのコードが含まれています。

## 参照

- [英語で書かれた README | README file in English](README.md)

## 著者

- [@IshmamAbir](https://www.github.com/IshmamAbir)

## 🔗 リンク

[![ポートフォリオ](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://linktr.ee/ishmam_abir)

[![Linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ishmam-abir/)

[![Facebook](https://img.shields.io/badge/facebook-1DA1F2?style=for-the-badge&logo=facebook&logoColor=white)](https://facebook.com/ishmam.abir)

## 📝 ライセンス

[MIT](https://github.com/IshmamAbir/Go-Service_Level_Transaction/blob/main/LICENSE)
