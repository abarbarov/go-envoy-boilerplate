Golang server behind envoy proxy that communicates with a react client (typescript) via grpc (demo)

#### Client
1. Create react app with typescript support
```bash
npx create-react-app client --typescript
```
2. create client/proto folder
3. 
```bash
yarn add google-protobuf grpc grpc-web
```
4. generate client proto libraries
```
protoc -I todo/ todo/todo.proto --js_out=import_style=commonjs:client/src/proto --grpc-web_out=import_style=typescript,mode=grpcwebtext:client/src/proto
```
5. add /* eslint-disable */ to generated *.js files to make tslint happy

#### Server
6. Generate go models
```bash
protoc -I todo/ todo/todo.proto --go_out=plugins=grpc:todo
```

#### Envoy proxy
7. build envoy docker image:
```bash
docker build -t abarbarov/go-envoy-boilerplate:v1 .
```

#### Run everything
8. run go server
```bash
go run server.go
```
9. run envoy proxy
```bash
docker run -d -p 9090:9090 abarbarov/go-envoy-boilerplate:v1
```
10. run client app
```bash
yarn start
```


