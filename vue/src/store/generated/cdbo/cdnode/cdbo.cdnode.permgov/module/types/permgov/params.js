/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";
export const protobufPackage = "cdbo.cdnode.permgov";
const baseParams = { governor: "" };
export const Params = {
    encode(message, writer = Writer.create()) {
        if (message.governor !== "") {
            writer.uint32(10).string(message.governor);
        }
        return writer;
    },
    decode(input, length) {
        const reader = input instanceof Uint8Array ? new Reader(input) : input;
        let end = length === undefined ? reader.len : reader.pos + length;
        const message = { ...baseParams };
        while (reader.pos < end) {
            const tag = reader.uint32();
            switch (tag >>> 3) {
                case 1:
                    message.governor = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
            }
        }
        return message;
    },
    fromJSON(object) {
        const message = { ...baseParams };
        if (object.governor !== undefined && object.governor !== null) {
            message.governor = String(object.governor);
        }
        else {
            message.governor = "";
        }
        return message;
    },
    toJSON(message) {
        const obj = {};
        message.governor !== undefined && (obj.governor = message.governor);
        return obj;
    },
    fromPartial(object) {
        const message = { ...baseParams };
        if (object.governor !== undefined && object.governor !== null) {
            message.governor = object.governor;
        }
        else {
            message.governor = "";
        }
        return message;
    },
};
