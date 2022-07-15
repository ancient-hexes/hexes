import { createSignal } from 'solid-js'
import { Tile } from '../proto/hexes/v1/tile'

const [getHoveredTile, setHoveredTile] = createSignal<Tile | undefined>(undefined)

export const cursor = {
    hoveredTile: {
        get: getHoveredTile,
        set: setHoveredTile,
    }
}
