import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "alice.blog.blog";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgCreateHello {
    creator: string;
    title: string;
    body: string;
}
export interface MsgCreateHelloResponse {
    id: number;
}
export interface MsgUpdateHello {
    creator: string;
    id: number;
    title: string;
    body: string;
}
export interface MsgUpdateHelloResponse {
}
export interface MsgDeleteHello {
    creator: string;
    id: number;
}
export interface MsgDeleteHelloResponse {
}
export declare const MsgCreateHello: {
    encode(message: MsgCreateHello, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHello;
    fromJSON(object: any): MsgCreateHello;
    toJSON(message: MsgCreateHello): unknown;
    fromPartial(object: DeepPartial<MsgCreateHello>): MsgCreateHello;
};
export declare const MsgCreateHelloResponse: {
    encode(message: MsgCreateHelloResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgCreateHelloResponse;
    fromJSON(object: any): MsgCreateHelloResponse;
    toJSON(message: MsgCreateHelloResponse): unknown;
    fromPartial(object: DeepPartial<MsgCreateHelloResponse>): MsgCreateHelloResponse;
};
export declare const MsgUpdateHello: {
    encode(message: MsgUpdateHello, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateHello;
    fromJSON(object: any): MsgUpdateHello;
    toJSON(message: MsgUpdateHello): unknown;
    fromPartial(object: DeepPartial<MsgUpdateHello>): MsgUpdateHello;
};
export declare const MsgUpdateHelloResponse: {
    encode(_: MsgUpdateHelloResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgUpdateHelloResponse;
    fromJSON(_: any): MsgUpdateHelloResponse;
    toJSON(_: MsgUpdateHelloResponse): unknown;
    fromPartial(_: DeepPartial<MsgUpdateHelloResponse>): MsgUpdateHelloResponse;
};
export declare const MsgDeleteHello: {
    encode(message: MsgDeleteHello, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteHello;
    fromJSON(object: any): MsgDeleteHello;
    toJSON(message: MsgDeleteHello): unknown;
    fromPartial(object: DeepPartial<MsgDeleteHello>): MsgDeleteHello;
};
export declare const MsgDeleteHelloResponse: {
    encode(_: MsgDeleteHelloResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgDeleteHelloResponse;
    fromJSON(_: any): MsgDeleteHelloResponse;
    toJSON(_: MsgDeleteHelloResponse): unknown;
    fromPartial(_: DeepPartial<MsgDeleteHelloResponse>): MsgDeleteHelloResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    CreateHello(request: MsgCreateHello): Promise<MsgCreateHelloResponse>;
    UpdateHello(request: MsgUpdateHello): Promise<MsgUpdateHelloResponse>;
    DeleteHello(request: MsgDeleteHello): Promise<MsgDeleteHelloResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    CreateHello(request: MsgCreateHello): Promise<MsgCreateHelloResponse>;
    UpdateHello(request: MsgUpdateHello): Promise<MsgUpdateHelloResponse>;
    DeleteHello(request: MsgDeleteHello): Promise<MsgDeleteHelloResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
