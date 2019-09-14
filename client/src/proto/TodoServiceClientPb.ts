/**
 * @fileoverview gRPC-Web generated client stub for todo
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


import * as grpcWeb from 'grpc-web';

import {
  addTodoParams,
  completeResponse,
  completeTodoParams,
  deleteResponse,
  deleteTodoParams,
  getTodoParams,
  todoObject,
  todoResponse} from './todo_pb';

export class todoServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoaddTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    todoObject,
    (request: addTodoParams) => {
      return request.serializeBinary();
    },
    todoObject.deserializeBinary
  );

  addTodo(
    request: addTodoParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: todoObject) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/todo.todoService/addTodo',
      request,
      metadata || {},
      this.methodInfoaddTodo,
      callback);
  }

  methodInfodeleteTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    deleteResponse,
    (request: deleteTodoParams) => {
      return request.serializeBinary();
    },
    deleteResponse.deserializeBinary
  );

  deleteTodo(
    request: deleteTodoParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: deleteResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/todo.todoService/deleteTodo',
      request,
      metadata || {},
      this.methodInfodeleteTodo,
      callback);
  }

  methodInfocompleteTodo = new grpcWeb.AbstractClientBase.MethodInfo(
    completeResponse,
    (request: completeTodoParams) => {
      return request.serializeBinary();
    },
    completeResponse.deserializeBinary
  );

  completeTodo(
    request: completeTodoParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: completeResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/todo.todoService/completeTodo',
      request,
      metadata || {},
      this.methodInfocompleteTodo,
      callback);
  }

  methodInfogetTodos = new grpcWeb.AbstractClientBase.MethodInfo(
    todoResponse,
    (request: getTodoParams) => {
      return request.serializeBinary();
    },
    todoResponse.deserializeBinary
  );

  getTodos(
    request: getTodoParams,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: todoResponse) => void) {
    return this.client_.rpcCall(
      this.hostname_ +
        '/todo.todoService/getTodos',
      request,
      metadata || {},
      this.methodInfogetTodos,
      callback);
  }

}

