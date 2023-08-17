# Go - Transaction in business Logic Layer

This project implements transaction in the service layer of the project.

In a MVC pattern architecture, Repository contains all the database related task and Service/Usecase contains the business logic.

## Transaction in Go

Suppose an endpoint of a project create effects in more then one database tables. In that case, if one of these database action fails to perform, other datas should also not be saved in their respective database tables. So there is a mechanism provided by Transaction that rollback other tables to save in the database.

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

If we use transaction for this two table from one single `purchaseProduct()` method inside repository, it will eliminate the Single Responsibility of [SOLID Principle](https://s8sg.medium.com/solid-principle-in-go-e1a624290346).
Also, writting test case would be difficult. As the process is a business required logic, this `purchaseProduct()` should be in Service layer. Inside this method, we will call two different repository method `reduceStockAmount()` and `reduceBalance()`.

## Transaction in Service level

As service level dont contain database logic and separate repository method is called inside the service method, it is unable to set the transaction in service that we have seen in the example of first section. We have to create a middleware for this.

### Transaction Middleware

For creating Transaction in business layer, a middleware `DBTransactionMiddleware()` method is created where we pass the handler function along with the gorm db initializer which handles the transaction part and `Commit()` the complete process. But if the handler returns anything rather than a `http.StatusOk` or `http.StatusCreated`, it call the
`Rollback()` method and takes back all the entry.

### `WithTx()` method

This method assign the same `*gorm.DB` initializer to all the repositories where the service method will make an effect. call this method before calling the `purchaseProduct()` . for example

```bash
err := userUsecase.WithTx(txHandle).PurchaseProduct(orderRequest);
```

or we can call it separately one after another.

```bash
txHandleErr := userUsecase.WithTx(txHandle)
purchaseErr := userUsecase.PurchaseProduct(orderRequest);
```

### Calling Middleware

Call the transaction middleware in the handleFunc.

```bash
router.HandleFunc("/order-product", transaction.DBTransactionMiddleware(gorm.Db, <PurchaseProduct handler>)).Methods("POST")
```

## Authors

- [@IshmamAbir](https://www.github.com/IshmamAbir)

## 🔗 Links

[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://linktr.ee/ishmam_abir)

[![Linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ishmam-abir/)

[![Facebook](https://img.shields.io/badge/facebook-1DA1F2?style=for-the-badge&logo=facebook&logoColor=white)](https://facebook.com/ishmam.abir)

## License

[MIT](https://github.com/IshmamAbir/Go-Service_Level_Transaction/blob/main/LICENSE)
