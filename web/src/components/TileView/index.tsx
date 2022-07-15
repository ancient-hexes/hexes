import { JSX } from 'solid-js'
import { tileSize } from '../../const'

import { Terrain_Type } from '../../proto/hexes/v1/terrain'
import { Tile } from '../../proto/hexes/v1/tile'
import state from '../../state'

import styles from './index.module.css'


interface P {
    tile: Tile,
}

const getTheme = (terrain: Terrain_Type): string => {
    switch (terrain) {
        case Terrain_Type.TYPE_GRASS:
            return styles.grass
        case Terrain_Type.TYPE_HIGHLANDS:
            return styles.highlands
        case Terrain_Type.TYPE_DIRT:
            return styles.dirt
        case Terrain_Type.TYPE_LAVA:
            return styles.lava
        case Terrain_Type.TYPE_SUBTERRANEAN:
            return styles.subterranean
        case Terrain_Type.TYPE_ROCK:
            return styles.rock
        case Terrain_Type.TYPE_ROUGH:
            return styles.rough
        case Terrain_Type.TYPE_WASTELAND:
            return styles.wasteland
        case Terrain_Type.TYPE_SAND:
            return styles.sand
        case Terrain_Type.TYPE_SNOW:
            return styles.snow
        case Terrain_Type.TYPE_SWAMP:
            return styles.swamp
        case Terrain_Type.TYPE_WATER:
            return styles.water
    }
    return ''
}

export const TileView = (props: P) => {
    const { coordinate, terrain } = props.tile

    if (coordinate === undefined) {
        throw new Error(`got a Tile without a coordinate: ${props.tile}`)
    }

    const style: JSX.CSSProperties = {
        'grid-column': coordinate.x + 1,
        'grid-row': coordinate.y + 1,
        height: tileSize,
        width: tileSize,
    }

    return (
        <div
            class={ `${styles.tile} ${getTheme(terrain)}` }
            style={ style }
            onMouseEnter={ () => state.cursor.hoveredTile.set(props.tile) }
            onClick={ () => state.cursor.hoveredTile.set(props.tile) }
        />
    )
}
