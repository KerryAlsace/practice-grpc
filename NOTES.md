1. Set up layout
2. Fill out routes/practice_app.proto
3. Generate pb:
  `protoc -I routes/ routes/practice_app.proto --go_out=plugins=grpc:routes`
4. Add to .env:
  PORT=:50051
  ADDRESS=localhost:50051
5. Set up server file
6. Set up client file
