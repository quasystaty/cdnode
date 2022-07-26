/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cdbo.cdnode.permgov";

/** Params defines the parameters for the module. */
export interface Params {
  governor: string;
}

const baseParams: object = { governor: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.governor !== "") {
      writer.uint32(10).string(message.governor);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
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

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.governor !== undefined && object.governor !== null) {
      message.governor = String(object.governor);
    } else {
      message.governor = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.governor !== undefined && (obj.governor = message.governor);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.governor !== undefined && object.governor !== null) {
      message.governor = object.governor;
    } else {
      message.governor = "";
    }
    return message;
  },
};

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
