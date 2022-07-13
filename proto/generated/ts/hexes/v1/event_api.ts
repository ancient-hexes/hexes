/* eslint-disable */
export const protobufPackage = "hexes.v1";

export interface EventAPI {}

export class EventAPIClientImpl implements EventAPI {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}
