/* eslint-disable */
import { Observable } from "rxjs";
import { Environment_ListResponse, Environment_ListRequest } from "./env.js";
import { map } from "rxjs/operators";
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "hexes.v1";

export interface EnvironmentAPI {
  List(request: Environment_ListRequest): Observable<Environment_ListResponse>;
}

export class EnvironmentAPIClientImpl implements EnvironmentAPI {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.List = this.List.bind(this);
  }
  List(request: Environment_ListRequest): Observable<Environment_ListResponse> {
    const data = Environment_ListRequest.encode(request).finish();
    const result = this.rpc.serverStreamingRequest(
      "hexes.v1.EnvironmentAPI",
      "List",
      data
    );
    return result.pipe(
      map((data) => Environment_ListResponse.decode(new _m0.Reader(data)))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
  clientStreamingRequest(
    service: string,
    method: string,
    data: Observable<Uint8Array>
  ): Promise<Uint8Array>;
  serverStreamingRequest(
    service: string,
    method: string,
    data: Uint8Array
  ): Observable<Uint8Array>;
  bidirectionalStreamingRequest(
    service: string,
    method: string,
    data: Observable<Uint8Array>
  ): Observable<Uint8Array>;
}
