/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Hello } from "../blog/hello";
export const protobufPackage = "alice.blog.blog";
const baseGenesisState = { helloCount: 0 };
export const GenesisState = {
    encode(message, writer = Writer.create()) {
        for (const v of message.helloList) {
            Hello.encode(v, writer.uint32(10).fork()).ldelim();
        }
        if (message.helloCount !== 0) {
            writer.uint32(16).uint64(message.helloCount);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseGenesisState };
        message.helloList = [];
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.helloList.push(Hello.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.helloCount = longToNumber(reader.uint64());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseGenesisState };
        message.helloList = [];
        if (object.helloList !== undefined && object.helloList !== null) {
            for (const e of object.helloList) {
                message.helloList.push(Hello.fromJSON(e));
            }
        }
        if (object.helloCount !== undefined && object.helloCount !== null) {
            message.helloCount = Number(object.helloCount);
        }
        else {
            message.helloCount = 0;
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        if (message.helloList) {
            obj.helloList = message.helloList.map((e) => e ? Hello.toJSON(e) : undefined);
        }
        else {
            obj.helloList = [];
        }
        message.helloCount !== undefined && (obj.helloCount = message.helloCount);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseGenesisState };
        message.helloList = [];
        if (object.helloList !== undefined && object.helloList !== null) {
            for (const e of object.helloList) {
                message.helloList.push(Hello.fromPartial(e));
            }
        }
        if (object.helloCount !== undefined && object.helloCount !== null) {
            message.helloCount = object.helloCount;
        }
        else {
            message.helloCount = 0;
        }
        return message;
    },
};
var globalThis = (() => {
    if (typeof globalThis !== "undefined")
        return globalThis;
    if (typeof self !== "undefined")
        return self;
    if (typeof window !== "undefined")
        return window;
    if (typeof global !== "undefined")
        return global;
    throw "Unable to locate global object";
})();
function longToNumber(long) {
    if (long.gt(Number.MAX_SAFE_INTEGER)) {
        throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
    }
    return long.toNumber();
}
if (util.Long !== Long) {
    util.Long = Long;
    configure();
}
