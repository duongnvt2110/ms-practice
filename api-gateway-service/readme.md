# Setup gRPC
- https://grpc-ecosystem.github.io/grpc-gateway/docs/tutorials/introduction/
# Generate stubs 
 protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   ./proto/api_gateway/api_gateway.proto


 protoc -I ./proto \
  --go_out ./proto --go_opt paths=source_relative \
  --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   ./proto/api_gateway/api_gateway.proto



## Sample gRPC gateway 
```
  // lis, err := net.Listen("tcp", ":3000")
  // if err != nil {
  // 	log.Fatalln("Failed to listen:", err)
  // }
  // s := grpc.NewServer()
  // api_gateway.RegisterGreeterServer(s, &server{})
  // log.Println("Serving gRPC on 0.0.0.0:3000")
  // go func() {
  // 	log.Fatalln(s.Serve(lis))
  // }()

  // conn, err := grpc.NewClient(
  // 	":3000",
  // 	grpc.WithTransportCredentials(insecure.NewCredentials()),
  // )
  // if err != nil {
  // 	log.Fatalln("Failed to dial server:", err)
  // }
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{
    grpc.WithInsecure(),
    grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
  }
  webAEndPoints := "webA-1:8081"
  err := api_gateway.RegisterWebAHandlerFromEndpoint(context.Background(), mux, webAEndPoints, opts)
  if err != nil {
    log.Fatalln("Failed to register gateway:", err)
  }
  webBEndPoints := "webB-1:8084"
  err = api_gateway.RegisterWebBHandlerFromEndpoint(context.Background(), mux, webBEndPoints, opts)
  if err != nil {
    log.Fatalln("Failed to register gateway:", err)
  }
  gwServer := &http.Server{
    Addr:    ":3000",
    Handler: mux,
  }

  log.Println("Serving test gRPC-Gateway on http://0.0.0.0:3000")
  log.Fatalln(gwServer.ListenAndServe())
```