import { Terrain_Type } from '../proto/hexes/v1/terrain'

export const terrainString = (terrain?: Terrain_Type): string => {
    switch (terrain) {
        case Terrain_Type.TYPE_UNSPECIFIED:
            return 'Unspecified'
        case Terrain_Type.TYPE_GRASS:
            return 'Grass'
        case Terrain_Type.TYPE_HIGHLANDS:
            return 'Highlands'
        case Terrain_Type.TYPE_DIRT:
            return 'Dirt'
        case Terrain_Type.TYPE_LAVA:
            return 'Lava'
        case Terrain_Type.TYPE_SUBTERRANEAN:
            return 'Subterranean'
        case Terrain_Type.TYPE_ROCK:
            return 'Rock'
        case Terrain_Type.TYPE_ROUGH:
            return 'Rough'
        case Terrain_Type.TYPE_WASTELAND:
            return 'Wasteland'
        case Terrain_Type.TYPE_SAND:
            return 'Sand'
        case Terrain_Type.TYPE_SNOW:
            return 'Snow'
        case Terrain_Type.TYPE_SWAMP:
            return 'Swamp'
        case Terrain_Type.TYPE_WATER:
            return 'Water'
    }
    return '?'
}
