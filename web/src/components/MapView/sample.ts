import { Map } from '../../proto/hexes/v1/map'
import { Terrain_Type } from '../../proto/hexes/v1/terrain'
import { Tile } from '../../proto/hexes/v1/tile'


export const generateRandomMap = (height: number, width: number): Map => {
    const map = Map.fromPartial({
        height,
        width,
    })
    // const terrain: Terrain_Type = Math.round(Math.random() * 100 % 1) + 1
    const terrain = Terrain_Type.TYPE_SAND
    for (let x = 0; x < map.width; x++) {
        for (let y = 0; y < map.height; y++) {
            map.tiles.push(Tile.fromPartial({
                terrain,
                coordinate: { x, y },
            }))
        }
    }
    return map
}
