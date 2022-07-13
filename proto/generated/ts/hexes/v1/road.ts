/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface Road {
  type: Road_Type;
  /** Amount of move points required to move from adjacent tile. */
  baseCost: number;
  /** `base_cost` * 1.41 (rounded) */
  diagonalCost: number;
}

export enum Road_Type {
  TYPE_UNSPECIFIED = 0,
  TYPE_DIRT = 1,
  TYPE_GRAVEL = 2,
  TYPE_COBBLESTONE = 3,
  TYPE_FAVORABLE_WINDS = 4,
  UNRECOGNIZED = -1,
}

export function road_TypeFromJSON(object: any): Road_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return Road_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_DIRT":
      return Road_Type.TYPE_DIRT;
    case 2:
    case "TYPE_GRAVEL":
      return Road_Type.TYPE_GRAVEL;
    case 3:
    case "TYPE_COBBLESTONE":
      return Road_Type.TYPE_COBBLESTONE;
    case 4:
    case "TYPE_FAVORABLE_WINDS":
      return Road_Type.TYPE_FAVORABLE_WINDS;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Road_Type.UNRECOGNIZED;
  }
}

export function road_TypeToJSON(object: Road_Type): string {
  switch (object) {
    case Road_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case Road_Type.TYPE_DIRT:
      return "TYPE_DIRT";
    case Road_Type.TYPE_GRAVEL:
      return "TYPE_GRAVEL";
    case Road_Type.TYPE_COBBLESTONE:
      return "TYPE_COBBLESTONE";
    case Road_Type.TYPE_FAVORABLE_WINDS:
      return "TYPE_FAVORABLE_WINDS";
    case Road_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseRoad(): Road {
  return { type: 0, baseCost: 0, diagonalCost: 0 };
}

export const Road = {
  encode(message: Road, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.baseCost !== 0) {
      writer.uint32(16).uint32(message.baseCost);
    }
    if (message.diagonalCost !== 0) {
      writer.uint32(24).uint32(message.diagonalCost);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Road {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRoad();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.baseCost = reader.uint32();
          break;
        case 3:
          message.diagonalCost = reader.uint32();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Road {
    return {
      type: isSet(object.type) ? road_TypeFromJSON(object.type) : 0,
      baseCost: isSet(object.baseCost) ? Number(object.baseCost) : 0,
      diagonalCost: isSet(object.diagonalCost)
        ? Number(object.diagonalCost)
        : 0,
    };
  },

  toJSON(message: Road): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = road_TypeToJSON(message.type));
    message.baseCost !== undefined &&
      (obj.baseCost = Math.round(message.baseCost));
    message.diagonalCost !== undefined &&
      (obj.diagonalCost = Math.round(message.diagonalCost));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Road>, I>>(object: I): Road {
    const message = createBaseRoad();
    message.type = object.type ?? 0;
    message.baseCost = object.baseCost ?? 0;
    message.diagonalCost = object.diagonalCost ?? 0;
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
