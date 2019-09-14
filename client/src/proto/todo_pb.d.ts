import * as jspb from "google-protobuf"

export class getTodoParams extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): getTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: getTodoParams): getTodoParams.AsObject;
  static serializeBinaryToWriter(message: getTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): getTodoParams;
  static deserializeBinaryFromReader(message: getTodoParams, reader: jspb.BinaryReader): getTodoParams;
}

export namespace getTodoParams {
  export type AsObject = {
  }
}

export class addTodoParams extends jspb.Message {
  getTask(): string;
  setTask(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): addTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: addTodoParams): addTodoParams.AsObject;
  static serializeBinaryToWriter(message: addTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): addTodoParams;
  static deserializeBinaryFromReader(message: addTodoParams, reader: jspb.BinaryReader): addTodoParams;
}

export namespace addTodoParams {
  export type AsObject = {
    task: string,
  }
}

export class completeTodoParams extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): completeTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: completeTodoParams): completeTodoParams.AsObject;
  static serializeBinaryToWriter(message: completeTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): completeTodoParams;
  static deserializeBinaryFromReader(message: completeTodoParams, reader: jspb.BinaryReader): completeTodoParams;
}

export namespace completeTodoParams {
  export type AsObject = {
    id: string,
  }
}

export class deleteTodoParams extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): deleteTodoParams.AsObject;
  static toObject(includeInstance: boolean, msg: deleteTodoParams): deleteTodoParams.AsObject;
  static serializeBinaryToWriter(message: deleteTodoParams, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): deleteTodoParams;
  static deserializeBinaryFromReader(message: deleteTodoParams, reader: jspb.BinaryReader): deleteTodoParams;
}

export namespace deleteTodoParams {
  export type AsObject = {
    id: string,
  }
}

export class todoObject extends jspb.Message {
  getId(): string;
  setId(value: string): void;

  getCompleted(): boolean;
  setCompleted(value: boolean): void;

  getTask(): string;
  setTask(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): todoObject.AsObject;
  static toObject(includeInstance: boolean, msg: todoObject): todoObject.AsObject;
  static serializeBinaryToWriter(message: todoObject, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): todoObject;
  static deserializeBinaryFromReader(message: todoObject, reader: jspb.BinaryReader): todoObject;
}

export namespace todoObject {
  export type AsObject = {
    id: string,
    completed: boolean,
    task: string,
  }
}

export class todoResponse extends jspb.Message {
  getTodosList(): Array<todoObject>;
  setTodosList(value: Array<todoObject>): void;
  clearTodosList(): void;
  addTodos(value?: todoObject, index?: number): todoObject;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): todoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: todoResponse): todoResponse.AsObject;
  static serializeBinaryToWriter(message: todoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): todoResponse;
  static deserializeBinaryFromReader(message: todoResponse, reader: jspb.BinaryReader): todoResponse;
}

export namespace todoResponse {
  export type AsObject = {
    todosList: Array<todoObject.AsObject>,
  }
}

export class deleteResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): deleteResponse.AsObject;
  static toObject(includeInstance: boolean, msg: deleteResponse): deleteResponse.AsObject;
  static serializeBinaryToWriter(message: deleteResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): deleteResponse;
  static deserializeBinaryFromReader(message: deleteResponse, reader: jspb.BinaryReader): deleteResponse;
}

export namespace deleteResponse {
  export type AsObject = {
    message: string,
  }
}

export class completeResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): completeResponse.AsObject;
  static toObject(includeInstance: boolean, msg: completeResponse): completeResponse.AsObject;
  static serializeBinaryToWriter(message: completeResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): completeResponse;
  static deserializeBinaryFromReader(message: completeResponse, reader: jspb.BinaryReader): completeResponse;
}

export namespace completeResponse {
  export type AsObject = {
    message: string,
  }
}

