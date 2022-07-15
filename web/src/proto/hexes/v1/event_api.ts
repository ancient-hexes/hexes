/* eslint-disable */
import { Event } from "./event.js";

export const protobufPackage = "hexes.v1";

export type EventAPIDefinition = typeof EventAPIDefinition;
export const EventAPIDefinition = {
  name: "EventAPI",
  fullName: "hexes.v1.EventAPI",
  methods: {
    connect: {
      name: "Connect",
      requestType: Event,
      requestStream: true,
      responseType: Event,
      responseStream: true,
      options: {},
    },
  },
} as const;
