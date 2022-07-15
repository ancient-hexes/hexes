import { tileSize } from '../../const'
import { Color } from '../../proto/hexes/v1/color'
import { Map } from '../../proto/hexes/v1/map'
import { GeneralView } from '../GeneralView'
import { TileView } from '../TileView'

import styles from './index.module.css'


interface P {
    map: Map,
}

export const MapView = (props: P) => {
    const height = `${tileSize * props.map.height}px`
    const width = `${tileSize * props.map.width}px`

    const style = {
        height,
        width,
        'grid-template-rows': `repeat(${props.map.height}, ${tileSize}px)`,
        'grid-template-columns': `repeat(${props.map.width}, ${tileSize}px)`,
    }

    return (
        <div class={ styles.mapView } style={ style }>
            { props.map.tiles.map(tile => <TileView tile={ tile }/>) }

            <GeneralView color={ Color.COLOR_RED } coordinate={{ x: 0, y: 1, z: 0 }}/>
            <GeneralView color={ Color.COLOR_BLUE } coordinate={{ x: 8, y: 12, z: 0 }}/>
        </div>
    )
}
