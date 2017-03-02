1. Set up layout
2. Fill out routes/practice_app.proto
3. Generate pb:
  `protoc -I routes/ routes/practice_app.proto --go_out=plugins=grpc:routes`
4. Add to .env:
  PORT=:50051
  ADDRESS=localhost:50051
5. Set up server file
6. Set up client file
7. Run:
  `go run server/server.go`
8. In another terminal tab run:
  `go run client/client.go` (can add a number as a command line argument, like `go run client/client.go 15`)
9. Should return:
  (without CLA)
  2017/03/02 16:30:28 AddOne: 3
  2017/03/02 16:30:28 Square: 4
  (with CLA of 15)
  2017/03/02 16:31:08 AddOne: 16
  2017/03/02 16:31:08 Square: 225
