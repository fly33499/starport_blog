/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Hello } from "../blog/hello";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "alice.blog.blog";

/** this line is used by starport scaffolding # 3 */
export interface QueryGetHelloRequest {
  id: number;
}

export interface QueryGetHelloResponse {
  Hello: Hello | undefined;
}

export interface QueryAllHelloRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllHelloResponse {
  Hello: Hello[];
  pagination: PageResponse | undefined;
}

const baseQueryGetHelloRequest: object = { id: 0 };

export const QueryGetHelloRequest = {
  encode(
    message: QueryGetHelloRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetHelloRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetHelloRequest } as QueryGetHelloRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHelloRequest {
    const message = { ...baseQueryGetHelloRequest } as QueryGetHelloRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetHelloRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetHelloRequest>): QueryGetHelloRequest {
    const message = { ...baseQueryGetHelloRequest } as QueryGetHelloRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetHelloResponse: object = {};

export const QueryGetHelloResponse = {
  encode(
    message: QueryGetHelloResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.Hello !== undefined) {
      Hello.encode(message.Hello, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetHelloResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetHelloResponse } as QueryGetHelloResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Hello = Hello.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetHelloResponse {
    const message = { ...baseQueryGetHelloResponse } as QueryGetHelloResponse;
    if (object.Hello !== undefined && object.Hello !== null) {
      message.Hello = Hello.fromJSON(object.Hello);
    } else {
      message.Hello = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetHelloResponse): unknown {
    const obj: any = {};
    message.Hello !== undefined &&
      (obj.Hello = message.Hello ? Hello.toJSON(message.Hello) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetHelloResponse>
  ): QueryGetHelloResponse {
    const message = { ...baseQueryGetHelloResponse } as QueryGetHelloResponse;
    if (object.Hello !== undefined && object.Hello !== null) {
      message.Hello = Hello.fromPartial(object.Hello);
    } else {
      message.Hello = undefined;
    }
    return message;
  },
};

const baseQueryAllHelloRequest: object = {};

export const QueryAllHelloRequest = {
  encode(
    message: QueryAllHelloRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllHelloRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllHelloRequest } as QueryAllHelloRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHelloRequest {
    const message = { ...baseQueryAllHelloRequest } as QueryAllHelloRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllHelloRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllHelloRequest>): QueryAllHelloRequest {
    const message = { ...baseQueryAllHelloRequest } as QueryAllHelloRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllHelloResponse: object = {};

export const QueryAllHelloResponse = {
  encode(
    message: QueryAllHelloResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.Hello) {
      Hello.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllHelloResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllHelloResponse } as QueryAllHelloResponse;
    message.Hello = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.Hello.push(Hello.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllHelloResponse {
    const message = { ...baseQueryAllHelloResponse } as QueryAllHelloResponse;
    message.Hello = [];
    if (object.Hello !== undefined && object.Hello !== null) {
      for (const e of object.Hello) {
        message.Hello.push(Hello.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllHelloResponse): unknown {
    const obj: any = {};
    if (message.Hello) {
      obj.Hello = message.Hello.map((e) => (e ? Hello.toJSON(e) : undefined));
    } else {
      obj.Hello = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllHelloResponse>
  ): QueryAllHelloResponse {
    const message = { ...baseQueryAllHelloResponse } as QueryAllHelloResponse;
    message.Hello = [];
    if (object.Hello !== undefined && object.Hello !== null) {
      for (const e of object.Hello) {
        message.Hello.push(Hello.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a hello by id. */
  Hello(request: QueryGetHelloRequest): Promise<QueryGetHelloResponse>;
  /** Queries a list of hello items. */
  HelloAll(request: QueryAllHelloRequest): Promise<QueryAllHelloResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Hello(request: QueryGetHelloRequest): Promise<QueryGetHelloResponse> {
    const data = QueryGetHelloRequest.encode(request).finish();
    const promise = this.rpc.request("alice.blog.blog.Query", "Hello", data);
    return promise.then((data) =>
      QueryGetHelloResponse.decode(new Reader(data))
    );
  }

  HelloAll(request: QueryAllHelloRequest): Promise<QueryAllHelloResponse> {
    const data = QueryAllHelloRequest.encode(request).finish();
    const promise = this.rpc.request("alice.blog.blog.Query", "HelloAll", data);
    return promise.then((data) =>
      QueryAllHelloResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
