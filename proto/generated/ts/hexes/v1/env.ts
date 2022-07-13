/* eslint-disable */
import { Timestamp } from "../../google/protobuf/timestamp.js";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface Environment {
  id: Environment_ID | undefined;
  meta: Environment_Meta | undefined;
}

export enum Environment_Type {
  TYPE_UNSPECIFIED = 0,
  TYPE_LOBBY = 1,
  TYPE_SESSION = 2,
  UNRECOGNIZED = -1,
}

export function environment_TypeFromJSON(object: any): Environment_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return Environment_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_LOBBY":
      return Environment_Type.TYPE_LOBBY;
    case 2:
    case "TYPE_SESSION":
      return Environment_Type.TYPE_SESSION;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Environment_Type.UNRECOGNIZED;
  }
}

export function environment_TypeToJSON(object: Environment_Type): string {
  switch (object) {
    case Environment_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case Environment_Type.TYPE_LOBBY:
      return "TYPE_LOBBY";
    case Environment_Type.TYPE_SESSION:
      return "TYPE_SESSION";
    case Environment_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Environment_ID {
  type: Environment_Type;
  key: string;
}

export interface Environment_Meta {
  /**
   * Time of environment creation.
   * Not set for LOBBY-environments.
   */
  createdAt: Date | undefined;
}

export interface Environment_ListRequest {
  batchSize: number;
}

export interface Environment_ListResponse {
  items: Environment[];
}

function createBaseEnvironment(): Environment {
  return { id: undefined, meta: undefined };
}

export const Environment = {
  encode(
    message: Environment,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== undefined) {
      Environment_ID.encode(message.id, writer.uint32(10).fork()).ldelim();
    }
    if (message.meta !== undefined) {
      Environment_Meta.encode(message.meta, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Environment {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnvironment();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = Environment_ID.decode(reader, reader.uint32());
          break;
        case 2:
          message.meta = Environment_Meta.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Environment {
    return {
      id: isSet(object.id) ? Environment_ID.fromJSON(object.id) : undefined,
      meta: isSet(object.meta)
        ? Environment_Meta.fromJSON(object.meta)
        : undefined,
    };
  },

  toJSON(message: Environment): unknown {
    const obj: any = {};
    message.id !== undefined &&
      (obj.id = message.id ? Environment_ID.toJSON(message.id) : undefined);
    message.meta !== undefined &&
      (obj.meta = message.meta
        ? Environment_Meta.toJSON(message.meta)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Environment>, I>>(
    object: I
  ): Environment {
    const message = createBaseEnvironment();
    message.id =
      object.id !== undefined && object.id !== null
        ? Environment_ID.fromPartial(object.id)
        : undefined;
    message.meta =
      object.meta !== undefined && object.meta !== null
        ? Environment_Meta.fromPartial(object.meta)
        : undefined;
    return message;
  },
};

function createBaseEnvironment_ID(): Environment_ID {
  return { type: 0, key: "" };
}

export const Environment_ID = {
  encode(
    message: Environment_ID,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Environment_ID {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnvironment_ID();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.key = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Environment_ID {
    return {
      type: isSet(object.type) ? environment_TypeFromJSON(object.type) : 0,
      key: isSet(object.key) ? String(object.key) : "",
    };
  },

  toJSON(message: Environment_ID): unknown {
    const obj: any = {};
    message.type !== undefined &&
      (obj.type = environment_TypeToJSON(message.type));
    message.key !== undefined && (obj.key = message.key);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Environment_ID>, I>>(
    object: I
  ): Environment_ID {
    const message = createBaseEnvironment_ID();
    message.type = object.type ?? 0;
    message.key = object.key ?? "";
    return message;
  },
};

function createBaseEnvironment_Meta(): Environment_Meta {
  return { createdAt: undefined };
}

export const Environment_Meta = {
  encode(
    message: Environment_Meta,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.createdAt !== undefined) {
      Timestamp.encode(
        toTimestamp(message.createdAt),
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Environment_Meta {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnvironment_Meta();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.createdAt = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Environment_Meta {
    return {
      createdAt: isSet(object.createdAt)
        ? fromJsonTimestamp(object.createdAt)
        : undefined,
    };
  },

  toJSON(message: Environment_Meta): unknown {
    const obj: any = {};
    message.createdAt !== undefined &&
      (obj.createdAt = message.createdAt.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Environment_Meta>, I>>(
    object: I
  ): Environment_Meta {
    const message = createBaseEnvironment_Meta();
    message.createdAt = object.createdAt ?? undefined;
    return message;
  },
};

function createBaseEnvironment_ListRequest(): Environment_ListRequest {
  return { batchSize: 0 };
}

export const Environment_ListRequest = {
  encode(
    message: Environment_ListRequest,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.batchSize !== 0) {
      writer.uint32(8).uint32(message.batchSize);
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): Environment_ListRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnvironment_ListRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.batchSize = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Environment_ListRequest {
    return {
      batchSize: isSet(object.batchSize) ? Number(object.batchSize) : 0,
    };
  },

  toJSON(message: Environment_ListRequest): unknown {
    const obj: any = {};
    message.batchSize !== undefined &&
      (obj.batchSize = Math.round(message.batchSize));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Environment_ListRequest>, I>>(
    object: I
  ): Environment_ListRequest {
    const message = createBaseEnvironment_ListRequest();
    message.batchSize = object.batchSize ?? 0;
    return message;
  },
};

function createBaseEnvironment_ListResponse(): Environment_ListResponse {
  return { items: [] };
}

export const Environment_ListResponse = {
  encode(
    message: Environment_ListResponse,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    for (const v of message.items) {
      Environment.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: _m0.Reader | Uint8Array,
    length?: number
  ): Environment_ListResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnvironment_ListResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.items.push(Environment.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Environment_ListResponse {
    return {
      items: Array.isArray(object?.items)
        ? object.items.map((e: any) => Environment.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Environment_ListResponse): unknown {
    const obj: any = {};
    if (message.items) {
      obj.items = message.items.map((e) =>
        e ? Environment.toJSON(e) : undefined
      );
    } else {
      obj.items = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Environment_ListResponse>, I>>(
    object: I
  ): Environment_ListResponse {
    const message = createBaseEnvironment_ListResponse();
    message.items = object.items?.map((e) => Environment.fromPartial(e)) || [];
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
