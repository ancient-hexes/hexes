/* eslint-disable */
import {
  Terrain_Type,
  terrain_TypeFromJSON,
  terrain_TypeToJSON,
} from "./terrain.js";
import { Road_Type, road_TypeFromJSON, road_TypeToJSON } from "./road.js";
import { Coordinate } from "./coordinate.js";
import { Color, colorFromJSON, colorToJSON } from "./color.js";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface Tile {
  /** A coordinate relative to map's top left corner. */
  coordinate: Coordinate | undefined;
  /** Reference to `Terrain.type`. */
  terrain: Terrain_Type;
  /** Reference to `Road.type`. */
  road: Road_Type;
  /** A list of players who can view & inspect the tile. */
  visibleTo: Color[];
}

function createBaseTile(): Tile {
  return { coordinate: undefined, terrain: 0, road: 0, visibleTo: [] };
}

export const Tile = {
  encode(message: Tile, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.coordinate !== undefined) {
      Coordinate.encode(message.coordinate, writer.uint32(10).fork()).ldelim();
    }
    if (message.terrain !== 0) {
      writer.uint32(16).int32(message.terrain);
    }
    if (message.road !== 0) {
      writer.uint32(24).int32(message.road);
    }
    writer.uint32(34).fork();
    for (const v of message.visibleTo) {
      writer.int32(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Tile {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTile();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coordinate = Coordinate.decode(reader, reader.uint32());
          break;
        case 2:
          message.terrain = reader.int32() as any;
          break;
        case 3:
          message.road = reader.int32() as any;
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.visibleTo.push(reader.int32() as any);
            }
          } else {
            message.visibleTo.push(reader.int32() as any);
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Tile {
    return {
      coordinate: isSet(object.coordinate)
        ? Coordinate.fromJSON(object.coordinate)
        : undefined,
      terrain: isSet(object.terrain) ? terrain_TypeFromJSON(object.terrain) : 0,
      road: isSet(object.road) ? road_TypeFromJSON(object.road) : 0,
      visibleTo: Array.isArray(object?.visibleTo)
        ? object.visibleTo.map((e: any) => colorFromJSON(e))
        : [],
    };
  },

  toJSON(message: Tile): unknown {
    const obj: any = {};
    message.coordinate !== undefined &&
      (obj.coordinate = message.coordinate
        ? Coordinate.toJSON(message.coordinate)
        : undefined);
    message.terrain !== undefined &&
      (obj.terrain = terrain_TypeToJSON(message.terrain));
    message.road !== undefined && (obj.road = road_TypeToJSON(message.road));
    if (message.visibleTo) {
      obj.visibleTo = message.visibleTo.map((e) => colorToJSON(e));
    } else {
      obj.visibleTo = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Tile>, I>>(object: I): Tile {
    const message = createBaseTile();
    message.coordinate =
      object.coordinate !== undefined && object.coordinate !== null
        ? Coordinate.fromPartial(object.coordinate)
        : undefined;
    message.terrain = object.terrain ?? 0;
    message.road = object.road ?? 0;
    message.visibleTo = object.visibleTo?.map((e) => e) || [];
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
