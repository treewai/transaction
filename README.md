# Transaction

a REST API service that provides CRUD operations for a simple resource called ”Transaction”

## Usage
- Download and run
```
git clone https://github.com/treewai/transaction`

cd transaction

go build

./transaction
```

- Auth
```
curl -X POST "http://localhost:8080/auth" -H "Content-type: application/json" -d '{"username": "user1", "password": "password1"}'

{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODM1Njk0MjUsInVzZXJuYW1lIjoidXNlcjEifQ._BiGFnv2wCY4uounyCM0LaifnkWXARZUB45vDdQQM-s",
    "username": "user1"
}
```

- Add Transaction
```
curl -X POST "http://localhost:8080/transactions" -H "Content-type: application/json" -H "Authorization: Bearer {token}" -d '{"account": "ABC", "type": "Withdrawal", "amount": 100.0, "user": "user1"}'

{
    "id": 1,
    "account": "ABC",
    "type": "Withdrawal",
    "amount": 100,
    "date": "2023-05-08T01:17:43.7598293+07:00",
    "user":"user1"
}
```

- Get All Transactions
```
curl -X GET "http://localhost:8080/transactions" -H "Content-type: application/json" -H "Authorization: Bearer {token}"

[
    {
        "id": 1,
        "account": "ABC",
        "type": "Withdrawal",
        "amount": 100,
        "date": "2023-05-08T01:17:43.7598293+07:00",
        "user":"user1"
    }
]
```

- Get Transaction by ID
```
curl -X GET "http://localhost:8080/transactions/1" -H "Content-type: application/json" -H "Authorization: Bearer {token}"

{
    "id": 1,
    "account": "ABC",
    "type": "Withdrawal",
    "amount": 100,
    "date": "2023-05-08T01:17:43.7598293+07:00",
    "user":"user1"
}
```

- Update Transaction
```
curl -X PUT "http://localhost:8080/transactions/1" -H "Content-type: application/json" -H "Authorization: Bearer {token}" -d '{"account": "DEF", "type": "Deposit", "amount": 200.0, "user": "user2"}'

{
    "id": 1,
    "account": "DEF",
    "type": "Deposit",
    "amount": 200,
    "date": "2023-05-08T01:17:43.7598293+07:00",
    "user":"user2"
}
```

- Update Transaction
```
curl -X DELETE "http://localhost:8080/transactions/1" -H "Content-type: application/json" -H "Authorization: Bearer {token}"

{
    "id": 1,
    "account": "DEF",
    "type": "Deposit",
    "amount": 200,
    "date": "2023-05-08T01:17:43.7598293+07:00",
    "user":"user2"
}
```