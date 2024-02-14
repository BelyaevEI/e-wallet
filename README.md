# E-Wallet
REST API for application that implements a payment system transaction processing system

#### Structure:
3 Domain layers:

- Service layer
- Repository layer
- Database layer

## API:

### POST /api/v1/wallet

Creating new wallet 

##### Example response: 
```
{
	"id": "12345",
	"balance": "100.00"
} 
```
##### Response status:
```
200 - create successfully
400 - error in the request
500 - internal error
```

### POST /api/v1/wallet/{walletid}/send

Transferring funds from one wallet to another. The wallet ID is in the request path.

##### Example Input: 
```
{
	"to": "123456",
	"amount": "10.00"
} 
```

##### Response status:
```
200 - transfer successfully
404 - the outgoing wallet was not found
400 - the target wallet has not been found or the required amount is not available on the outgoing account
500 - internal error
```

### GET /api/v1/wallet/{walletid}/history

Getting the history of incoming and outgoing transactions. The wallet ID is in the request path.

##### Example response: 
```
{
	"time": "2024-02-14T15:35:00Z",
	"from": "12345",
  "to": "123456",
  "amount": "10.00"

} 
```
##### Response status:
```
200 - getting successfully
404 - wallet was not found
500 - internal error
```

### GET /api/v1/wallet/{walletid}

Getting the current wallet status. The wallet ID is in the request path.

##### Example Response: 
```
{
  "id": "12345",
  "balance": "123.00"
} 
```

## Requirements
- go 1.21
- docker & docker-compose
