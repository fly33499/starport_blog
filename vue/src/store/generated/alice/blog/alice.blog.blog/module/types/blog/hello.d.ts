import { Writer, Reader } from "protobufjs/minimal";
export declare const protobufPackage = "alice.blog.blog";
export interface Hello {
    creator: string;
    id: number;
    title: string;
    body: string;
}
export declare const Hello: {
    encode(message: Hello, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): Hello;
    fromJSON(object: any): Hello;
    toJSON(message: Hello): unknown;
    fromPartial(object: DeepPartial<Hello>): Hello;
};
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
