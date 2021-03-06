/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface PayloadChat {
  text: string;
  alliesOnly: boolean;
}

function createBasePayloadChat(): PayloadChat {
  return { text: "", alliesOnly: false };
}

export const PayloadChat = {
  encode(
    message: PayloadChat,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.text !== "") {
      writer.uint32(10).string(message.text);
    }
    if (message.alliesOnly === true) {
      writer.uint32(16).bool(message.alliesOnly);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PayloadChat {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePayloadChat();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.text = reader.string();
          break;
        case 2:
          message.alliesOnly = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PayloadChat {
    return {
      text: isSet(object.text) ? String(object.text) : "",
      alliesOnly: isSet(object.alliesOnly) ? Boolean(object.alliesOnly) : false,
    };
  },

  toJSON(message: PayloadChat): unknown {
    const obj: any = {};
    message.text !== undefined && (obj.text = message.text);
    message.alliesOnly !== undefined && (obj.alliesOnly = message.alliesOnly);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PayloadChat>, I>>(
    object: I
  ): PayloadChat {
    const message = createBasePayloadChat();
    message.text = object.text ?? "";
    message.alliesOnly = object.alliesOnly ?? false;
    return message;
  },
};

type Builtin =
  | Date
  | Function
  | Uint8Array
  | string
  | number
  | boolean
  | undefined;

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & Record<
        Exclude<keyof I, KeysOfUnion<P>>,
        never
      >;

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
