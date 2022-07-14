/* eslint-disable */
export const protobufPackage = "hexes.v1";

export enum Color {
  COLOR_UNSPECIFIED = 0,
  COLOR_RED = 1,
  COLOR_BLUE = 2,
  COLOR_TAN = 3,
  COLOR_GREEN = 4,
  COLOR_ORANGE = 5,
  COLOR_PURPLE = 6,
  COLOR_TEAL = 7,
  COLOR_PINK = 8,
  UNRECOGNIZED = -1,
}

export function colorFromJSON(object: any): Color {
  switch (object) {
    case 0:
    case "COLOR_UNSPECIFIED":
      return Color.COLOR_UNSPECIFIED;
    case 1:
    case "COLOR_RED":
      return Color.COLOR_RED;
    case 2:
    case "COLOR_BLUE":
      return Color.COLOR_BLUE;
    case 3:
    case "COLOR_TAN":
      return Color.COLOR_TAN;
    case 4:
    case "COLOR_GREEN":
      return Color.COLOR_GREEN;
    case 5:
    case "COLOR_ORANGE":
      return Color.COLOR_ORANGE;
    case 6:
    case "COLOR_PURPLE":
      return Color.COLOR_PURPLE;
    case 7:
    case "COLOR_TEAL":
      return Color.COLOR_TEAL;
    case 8:
    case "COLOR_PINK":
      return Color.COLOR_PINK;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Color.UNRECOGNIZED;
  }
}

export function colorToJSON(object: Color): string {
  switch (object) {
    case Color.COLOR_UNSPECIFIED:
      return "COLOR_UNSPECIFIED";
    case Color.COLOR_RED:
      return "COLOR_RED";
    case Color.COLOR_BLUE:
      return "COLOR_BLUE";
    case Color.COLOR_TAN:
      return "COLOR_TAN";
    case Color.COLOR_GREEN:
      return "COLOR_GREEN";
    case Color.COLOR_ORANGE:
      return "COLOR_ORANGE";
    case Color.COLOR_PURPLE:
      return "COLOR_PURPLE";
    case Color.COLOR_TEAL:
      return "COLOR_TEAL";
    case Color.COLOR_PINK:
      return "COLOR_PINK";
    case Color.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
