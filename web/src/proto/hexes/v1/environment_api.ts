/* eslint-disable */
import { Environment_ListRequest, Environment_ListResponse } from "./env.js";

export const protobufPackage = "hexes.v1";

export type EnvironmentAPIDefinition = typeof EnvironmentAPIDefinition;
export const EnvironmentAPIDefinition = {
  name: "EnvironmentAPI",
  fullName: "hexes.v1.EnvironmentAPI",
  methods: {
    list: {
      name: "List",
      requestType: Environment_ListRequest,
      requestStream: false,
      responseType: Environment_ListResponse,
      responseStream: true,
      options: {},
    },
  },
} as const;
