/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "cdbo.cdnode.coinmaster";

/** Params defines the parameters for the module. */
export interface Params {
  minters: string;
  denoms: string;
}

const baseParams: object = { minters: "", denoms: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.minters !== "") {
      writer.uint32(10).string(message.minters);
    }
    if (message.denoms !== "") {
      writer.uint32(18).string(message.denoms);
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
          message.minters = reader.string();
          break;
        case 2:
          message.denoms = reader.string();
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
    if (object.minters !== undefined && object.minters !== null) {
      message.minters = String(object.minters);
    } else {
      message.minters = "";
    }
    if (object.denoms !== undefined && object.denoms !== null) {
      message.denoms = String(object.denoms);
    } else {
      message.denoms = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.minters !== undefined && (obj.minters = message.minters);
    message.denoms !== undefined && (obj.denoms = message.denoms);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.minters !== undefined && object.minters !== null) {
      message.minters = object.minters;
    } else {
      message.minters = "";
    }
    if (object.denoms !== undefined && object.denoms !== null) {
      message.denoms = object.denoms;
    } else {
      message.denoms = "";
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
