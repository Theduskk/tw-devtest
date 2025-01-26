Requires Go 1.23.4 (Requires 1.20+); Docker Engine 27.4.0; (Tested on these, may work on older versions)

# Go Only
go run main.go

The below is noisy as it queries the public endpoint (TODO Mock it)
go test -v ./...

# Docker
go build; docker build --tag devtask0.1 .; docker run --publish 3000:3000 devtask0.1;

# Example curl commands
#Test address `0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2`
```
curl -X POST 'http://127.0.0.1:3000/subscribe?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2'
curl -X POST 'http://127.0.0.1:3000/unsubscribe?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2'
curl -X GET 'http://127.0.0.1:3000/getSubscribedAccounts'
curl -X GET 'http://127.0.0.1:3000/getCurrentBlock?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2'
```
Both of these dump a large amount of data as theres no min/max filter on the transactions
```
curl -X GET 'http://127.0.0.1:3000/getTransactions?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2'
curl -X GET 'http://127.0.0.1:3000/getAllTransactions'
```