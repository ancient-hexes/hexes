import { Show } from 'solid-js'
import { Tile } from '../../proto/hexes/v1/tile'

import state from '../../state'
import { terrainString } from '../../utils/terrain'

import styles from './index.module.css'

export const StatusBar = () => {
    const { hoveredTile } = state.cursor

    return (
        <div class={ styles.statusBar }>
            <Show when={ hoveredTile.get() !== undefined }>
                { formatCoordinates(hoveredTile.get()) }
                { ' ' + terrainString(hoveredTile.get()?.terrain) }
            </Show>
        </div>
    )
}

const formatCoordinates = (tile?: Tile): string => {
    if (tile === undefined) {
        return ''
    }

    const x = String(tile.coordinate?.x || 0)
    const y = String(tile.coordinate?.y || 0)
    const z = String(tile.coordinate?.z || 0)
    return `(${x}, ${y}, ${z})`
}
