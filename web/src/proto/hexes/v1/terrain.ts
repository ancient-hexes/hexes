/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface Terrain {
  type: Terrain_Type;
  /** A movement penalty for a given terrain without a road. */
  movementPenalty: number;
}

export enum Terrain_Type {
  TYPE_UNSPECIFIED = 0,
  TYPE_GRASS = 1,
  TYPE_HIGHLANDS = 2,
  TYPE_DIRT = 3,
  TYPE_LAVA = 4,
  TYPE_SUBTERRANEAN = 5,
  TYPE_ROCK = 6,
  TYPE_ROUGH = 7,
  TYPE_WASTELAND = 8,
  TYPE_SAND = 9,
  TYPE_SNOW = 10,
  TYPE_SWAMP = 11,
  TYPE_WATER = 12,
  UNRECOGNIZED = -1,
}

export function terrain_TypeFromJSON(object: any): Terrain_Type {
  switch (object) {
    case 0:
    case "TYPE_UNSPECIFIED":
      return Terrain_Type.TYPE_UNSPECIFIED;
    case 1:
    case "TYPE_GRASS":
      return Terrain_Type.TYPE_GRASS;
    case 2:
    case "TYPE_HIGHLANDS":
      return Terrain_Type.TYPE_HIGHLANDS;
    case 3:
    case "TYPE_DIRT":
      return Terrain_Type.TYPE_DIRT;
    case 4:
    case "TYPE_LAVA":
      return Terrain_Type.TYPE_LAVA;
    case 5:
    case "TYPE_SUBTERRANEAN":
      return Terrain_Type.TYPE_SUBTERRANEAN;
    case 6:
    case "TYPE_ROCK":
      return Terrain_Type.TYPE_ROCK;
    case 7:
    case "TYPE_ROUGH":
      return Terrain_Type.TYPE_ROUGH;
    case 8:
    case "TYPE_WASTELAND":
      return Terrain_Type.TYPE_WASTELAND;
    case 9:
    case "TYPE_SAND":
      return Terrain_Type.TYPE_SAND;
    case 10:
    case "TYPE_SNOW":
      return Terrain_Type.TYPE_SNOW;
    case 11:
    case "TYPE_SWAMP":
      return Terrain_Type.TYPE_SWAMP;
    case 12:
    case "TYPE_WATER":
      return Terrain_Type.TYPE_WATER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Terrain_Type.UNRECOGNIZED;
  }
}

export function terrain_TypeToJSON(object: Terrain_Type): string {
  switch (object) {
    case Terrain_Type.TYPE_UNSPECIFIED:
      return "TYPE_UNSPECIFIED";
    case Terrain_Type.TYPE_GRASS:
      return "TYPE_GRASS";
    case Terrain_Type.TYPE_HIGHLANDS:
      return "TYPE_HIGHLANDS";
    case Terrain_Type.TYPE_DIRT:
      return "TYPE_DIRT";
    case Terrain_Type.TYPE_LAVA:
      return "TYPE_LAVA";
    case Terrain_Type.TYPE_SUBTERRANEAN:
      return "TYPE_SUBTERRANEAN";
    case Terrain_Type.TYPE_ROCK:
      return "TYPE_ROCK";
    case Terrain_Type.TYPE_ROUGH:
      return "TYPE_ROUGH";
    case Terrain_Type.TYPE_WASTELAND:
      return "TYPE_WASTELAND";
    case Terrain_Type.TYPE_SAND:
      return "TYPE_SAND";
    case Terrain_Type.TYPE_SNOW:
      return "TYPE_SNOW";
    case Terrain_Type.TYPE_SWAMP:
      return "TYPE_SWAMP";
    case Terrain_Type.TYPE_WATER:
      return "TYPE_WATER";
    case Terrain_Type.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

function createBaseTerrain(): Terrain {
  return { type: 0, movementPenalty: 0 };
}

export const Terrain = {
  encode(
    message: Terrain,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.type !== 0) {
      writer.uint32(8).int32(message.type);
    }
    if (message.movementPenalty !== 0) {
      writer.uint32(17).double(message.movementPenalty);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Terrain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTerrain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.int32() as any;
          break;
        case 2:
          message.movementPenalty = reader.double();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Terrain {
    return {
      type: isSet(object.type) ? terrain_TypeFromJSON(object.type) : 0,
      movementPenalty: isSet(object.movementPenalty)
        ? Number(object.movementPenalty)
        : 0,
    };
  },

  toJSON(message: Terrain): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = terrain_TypeToJSON(message.type));
    message.movementPenalty !== undefined &&
      (obj.movementPenalty = message.movementPenalty);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Terrain>, I>>(object: I): Terrain {
    const message = createBaseTerrain();
    message.type = object.type ?? 0;
    message.movementPenalty = object.movementPenalty ?? 0;
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
