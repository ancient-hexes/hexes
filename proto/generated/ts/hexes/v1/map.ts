/* eslint-disable */
import * as Long from "long";
import { Tile } from "./tile.js";
import { Coordinate } from "./coordinate.js";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface Map {
  generation: Map_GenerationSpec | undefined;
  tiles: Tile[];
  objects: Map_Object[];
}

export interface Map_Object {
  /** A reference to `Map.Object.Type.id`. */
  type: string;
  /** A coordinate of bottom left tile. */
  coordinate: Coordinate | undefined;
}

export interface Map_Object_Type {
  /** e.g. `Base/CreatureBanks/DragonFlyHive` */
  id: string;
  /**
   * Object size in tiles.
   * e.g. an object with height = 2, width = 3 would look like this:
   *  |   |   |   |   |   |
   *  |   | x | x | x |   |
   *  |   | x | x | x |   |
   *  |   |   |   |   |   |
   *
   * The minimum value for both is 1.
   */
  height: number;
  width: number;
  /**
   * Coordinates relative to object's bottom left corner.
   * Describe a complex shapes.
   * For example, here's Dragon Utopia (height = 3, width = 6):
   *  |   |   |   |   |   |   |   |   |
   *  |   |   |   | x | x | x | x |   |
   *  |   | x | x | x | x | x |   |   |
   *  |   |   |   | x | x | x |   |   |
   *  |   |   |   |   |   |   |   |   |
   *
   *  [
   *      [0, 1],
   *      [1, 1],
   *      [2, 0], [2, 1], [2, 2],
   *      [3, 0], [3, 1], [3, 2],
   *      [4, 0], [4, 1], [4, 2],
   *      [5, 2],
   *  ]
   *
   * If not set, objects takes all tiles within its dimensions.
   */
  shape: Coordinate[];
  /**
   * A coordinate relative to object's bottom left corner.
   * e.g. for object with height = 2, width = 3 an entrance could be:
   *  – [0, 0] // bottom left
   *  – [1, 0] // bottom center
   *  – [2, 0] // bottom right
   *  – [0, 1] // top left
   *  – [1, 1] // top center
   *  – [2, 1] // top right
   *
   * For example, here's an object with [1, 0] entrance:
   *  |   |   |   |   |   |
   *  |   | x | x | x |   |
   *  |   | x | E | x |   |
   *  |   |   |   |   |   |
   *
   * Entrance is optional.
   */
  entrance: Coordinate | undefined;
}

/** Describes how the map was generated. */
export interface Map_GenerationSpec {
  /** Random seed, only set for randomly generated maps. */
  seed: number;
  /** Filename of the template used for generation. */
  template: string;
  /** Template file hash (MD5). */
  templateHash: string;
}

function createBaseMap(): Map {
  return { generation: undefined, tiles: [], objects: [] };
}

export const Map = {
  encode(message: Map, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.generation !== undefined) {
      Map_GenerationSpec.encode(
        message.generation,
        writer.uint32(10).fork()
      ).ldelim();
    }
    for (const v of message.tiles) {
      Tile.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.objects) {
      Map_Object.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Map {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMap();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.generation = Map_GenerationSpec.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.tiles.push(Tile.decode(reader, reader.uint32()));
          break;
        case 3:
          message.objects.push(Map_Object.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map {
    return {
      generation: isSet(object.generation)
        ? Map_GenerationSpec.fromJSON(object.generation)
        : undefined,
      tiles: Array.isArray(object?.tiles)
        ? object.tiles.map((e: any) => Tile.fromJSON(e))
        : [],
      objects: Array.isArray(object?.objects)
        ? object.objects.map((e: any) => Map_Object.fromJSON(e))
        : [],
    };
  },

  toJSON(message: Map): unknown {
    const obj: any = {};
    message.generation !== undefined &&
      (obj.generation = message.generation
        ? Map_GenerationSpec.toJSON(message.generation)
        : undefined);
    if (message.tiles) {
      obj.tiles = message.tiles.map((e) => (e ? Tile.toJSON(e) : undefined));
    } else {
      obj.tiles = [];
    }
    if (message.objects) {
      obj.objects = message.objects.map((e) =>
        e ? Map_Object.toJSON(e) : undefined
      );
    } else {
      obj.objects = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Map>, I>>(object: I): Map {
    const message = createBaseMap();
    message.generation =
      object.generation !== undefined && object.generation !== null
        ? Map_GenerationSpec.fromPartial(object.generation)
        : undefined;
    message.tiles = object.tiles?.map((e) => Tile.fromPartial(e)) || [];
    message.objects =
      object.objects?.map((e) => Map_Object.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMap_Object(): Map_Object {
  return { type: "", coordinate: undefined };
}

export const Map_Object = {
  encode(
    message: Map_Object,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.type !== "") {
      writer.uint32(10).string(message.type);
    }
    if (message.coordinate !== undefined) {
      Coordinate.encode(message.coordinate, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Map_Object {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMap_Object();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.type = reader.string();
          break;
        case 2:
          message.coordinate = Coordinate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map_Object {
    return {
      type: isSet(object.type) ? String(object.type) : "",
      coordinate: isSet(object.coordinate)
        ? Coordinate.fromJSON(object.coordinate)
        : undefined,
    };
  },

  toJSON(message: Map_Object): unknown {
    const obj: any = {};
    message.type !== undefined && (obj.type = message.type);
    message.coordinate !== undefined &&
      (obj.coordinate = message.coordinate
        ? Coordinate.toJSON(message.coordinate)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Map_Object>, I>>(
    object: I
  ): Map_Object {
    const message = createBaseMap_Object();
    message.type = object.type ?? "";
    message.coordinate =
      object.coordinate !== undefined && object.coordinate !== null
        ? Coordinate.fromPartial(object.coordinate)
        : undefined;
    return message;
  },
};

function createBaseMap_Object_Type(): Map_Object_Type {
  return { id: "", height: 0, width: 0, shape: [], entrance: undefined };
}

export const Map_Object_Type = {
  encode(
    message: Map_Object_Type,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.height !== 0) {
      writer.uint32(16).uint32(message.height);
    }
    if (message.width !== 0) {
      writer.uint32(24).uint32(message.width);
    }
    for (const v of message.shape) {
      Coordinate.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.entrance !== undefined) {
      Coordinate.encode(message.entrance, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Map_Object_Type {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMap_Object_Type();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.height = reader.uint32();
          break;
        case 3:
          message.width = reader.uint32();
          break;
        case 4:
          message.shape.push(Coordinate.decode(reader, reader.uint32()));
          break;
        case 5:
          message.entrance = Coordinate.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map_Object_Type {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      height: isSet(object.height) ? Number(object.height) : 0,
      width: isSet(object.width) ? Number(object.width) : 0,
      shape: Array.isArray(object?.shape)
        ? object.shape.map((e: any) => Coordinate.fromJSON(e))
        : [],
      entrance: isSet(object.entrance)
        ? Coordinate.fromJSON(object.entrance)
        : undefined,
    };
  },

  toJSON(message: Map_Object_Type): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.height !== undefined && (obj.height = Math.round(message.height));
    message.width !== undefined && (obj.width = Math.round(message.width));
    if (message.shape) {
      obj.shape = message.shape.map((e) =>
        e ? Coordinate.toJSON(e) : undefined
      );
    } else {
      obj.shape = [];
    }
    message.entrance !== undefined &&
      (obj.entrance = message.entrance
        ? Coordinate.toJSON(message.entrance)
        : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Map_Object_Type>, I>>(
    object: I
  ): Map_Object_Type {
    const message = createBaseMap_Object_Type();
    message.id = object.id ?? "";
    message.height = object.height ?? 0;
    message.width = object.width ?? 0;
    message.shape = object.shape?.map((e) => Coordinate.fromPartial(e)) || [];
    message.entrance =
      object.entrance !== undefined && object.entrance !== null
        ? Coordinate.fromPartial(object.entrance)
        : undefined;
    return message;
  },
};

function createBaseMap_GenerationSpec(): Map_GenerationSpec {
  return { seed: 0, template: "", templateHash: "" };
}

export const Map_GenerationSpec = {
  encode(
    message: Map_GenerationSpec,
    writer: _m0.Writer = _m0.Writer.create()
  ): _m0.Writer {
    if (message.seed !== 0) {
      writer.uint32(8).uint64(message.seed);
    }
    if (message.template !== "") {
      writer.uint32(18).string(message.template);
    }
    if (message.templateHash !== "") {
      writer.uint32(26).string(message.templateHash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Map_GenerationSpec {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMap_GenerationSpec();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seed = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.template = reader.string();
          break;
        case 3:
          message.templateHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Map_GenerationSpec {
    return {
      seed: isSet(object.seed) ? Number(object.seed) : 0,
      template: isSet(object.template) ? String(object.template) : "",
      templateHash: isSet(object.templateHash)
        ? String(object.templateHash)
        : "",
    };
  },

  toJSON(message: Map_GenerationSpec): unknown {
    const obj: any = {};
    message.seed !== undefined && (obj.seed = Math.round(message.seed));
    message.template !== undefined && (obj.template = message.template);
    message.templateHash !== undefined &&
      (obj.templateHash = message.templateHash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Map_GenerationSpec>, I>>(
    object: I
  ): Map_GenerationSpec {
    const message = createBaseMap_GenerationSpec();
    message.seed = object.seed ?? 0;
    message.template = object.template ?? "";
    message.templateHash = object.templateHash ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

// If you get a compile-error about 'Constructor<Long> and ... have no overlap',
// add '--ts_proto_opt=esModuleInterop=true' as a flag when calling 'protoc'.
if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
