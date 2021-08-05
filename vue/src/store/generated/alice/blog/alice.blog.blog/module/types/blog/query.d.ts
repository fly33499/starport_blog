import { Reader, Writer } from "protobufjs/minimal";
import { Hello } from "../blog/hello";
import { PageRequest, PageResponse } from "../cosmos/base/query/v1beta1/pagination";
export declare const protobufPackage = "alice.blog.blog";
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
export declare const QueryGetHelloRequest: {
    encode(message: QueryGetHelloRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetHelloRequest;
    fromJSON(object: any): QueryGetHelloRequest;
    toJSON(message: QueryGetHelloRequest): unknown;
    fromPartial(object: DeepPartial<QueryGetHelloRequest>): QueryGetHelloRequest;
};
export declare const QueryGetHelloResponse: {
    encode(message: QueryGetHelloResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryGetHelloResponse;
    fromJSON(object: any): QueryGetHelloResponse;
    toJSON(message: QueryGetHelloResponse): unknown;
    fromPartial(object: DeepPartial<QueryGetHelloResponse>): QueryGetHelloResponse;
};
export declare const QueryAllHelloRequest: {
    encode(message: QueryAllHelloRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllHelloRequest;
    fromJSON(object: any): QueryAllHelloRequest;
    toJSON(message: QueryAllHelloRequest): unknown;
    fromPartial(object: DeepPartial<QueryAllHelloRequest>): QueryAllHelloRequest;
};
export declare const QueryAllHelloResponse: {
    encode(message: QueryAllHelloResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryAllHelloResponse;
    fromJSON(object: any): QueryAllHelloResponse;
    toJSON(message: QueryAllHelloResponse): unknown;
    fromPartial(object: DeepPartial<QueryAllHelloResponse>): QueryAllHelloResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Queries a hello by id. */
    Hello(request: QueryGetHelloRequest): Promise<QueryGetHelloResponse>;
    /** Queries a list of hello items. */
    HelloAll(request: QueryAllHelloRequest): Promise<QueryAllHelloResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Hello(request: QueryGetHelloRequest): Promise<QueryGetHelloResponse>;
    HelloAll(request: QueryAllHelloRequest): Promise<QueryAllHelloResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
