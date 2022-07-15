import { Color } from '../../proto/hexes/v1/color'
import { Coordinate } from '../../proto/hexes/v1/coordinate'
import { colorString } from '../../utils/color'

import HorseIcon from '../../assets/horse.svg'
import styles from './index.module.css'


interface P {
    color: Color,
    coordinate: Coordinate,
}

export const GeneralView = (props: P) => {
    if (props.coordinate === undefined) {
        return null
    }

    return (
        <div
            style={{
                position: 'relative',
                'grid-column': props.coordinate.x + 1,
                'grid-row': props.coordinate.y + 1,
            }}
        >
            <HorseIcon class={ styles.general + ' ' + styles[colorString(props.color)] }/>
        </div>
    )
}
