# Go - Transaction in business Logic Layer

This project implements transaction in the service layer.

In a MVC pattern architecture, Repository contains all the database related task and Service/Usecase contains the business logic.

## Transaction in Go

Assume a project's endpoint generates effects in many database tables. In that instance, if one of these database actions fails, no other data should be saved in the corresponding database tables. So Transaction provides a technique that prevents other tables from being saved in the database.

## Gorm Transactions

[Gorm](https://gorm.io/docs/transactions.html), a popular ORM library provide this transaction feature.

Basic example of Gorm transaction is:

```bash
// begin a transaction
tx := db.Begin()

// do some database operations in the transaction (use 'tx' from this point, not 'db')
tx.Create(...)

// ...

// rollback the transaction in case of error
tx.Rollback()

// Or commit the transaction
tx.Commit()
```

#### Reference

- [Manual Transaction in Gorm](https://gorm.io/docs/transactions.html#Control-the-transaction-manually)

This is the most basic level transaction that can be used inside of a method of Repository layer.

For example:
while purchasing a product, product price will be deducted from the user balance, product stock will also be deducted from product table. This is the simple business logic of purchasing an item.

If we use transaction for this two table from one single `purchaseProduct()` method inside repository, it will eliminate the Single Responsibility law of [SOLID Principle](https://s8sg.medium.com/solid-principle-in-go-e1a624290346).
Also, writting test case would be difficult. As the process is a business required logic, this `purchaseProduct()` should be in Service layer. Inside this method, we will call two different repository method `reduceStockAmount()` and `reduceBalance()` to do the database related part.

## Transaction in Service level

As service level dont contain database logic and separate repository methods are called inside the service method, it is unable to set the transaction in service that we have seen in the example of first section. We have to create a middleware for this.

### transaction.go

This file contains the code for creating a UoW(unit of work) middleware for managing the transaction in the service layer level.

## Reference

- [Instructions in Japanese | 日本語で README](README-JP.md)

## Authors

- [@IshmamAbir](https://www.github.com/IshmamAbir)

## 🔗 Links

[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://linktr.ee/ishmam_abir)

[![Linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ishmam-abir/)

[![Facebook](https://img.shields.io/badge/facebook-1DA1F2?style=for-the-badge&logo=facebook&logoColor=white)](https://facebook.com/ishmam.abir)

## 📝 License

[MIT](https://github.com/IshmamAbir/Go-Service_Level_Transaction/blob/main/LICENSE)
