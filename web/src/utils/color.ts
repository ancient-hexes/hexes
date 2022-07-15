import { Color } from '../proto/hexes/v1/color'

export const colorString = (color: Color): string => {
    switch (color) {
        case Color.COLOR_RED:
            return 'red'
        case Color.COLOR_BLUE:
            return 'blue'
        case Color.COLOR_TAN:
            return 'tan'
        case Color.COLOR_GREEN:
            return 'green'
        case Color.COLOR_ORANGE:
            return 'orange'
        case Color.COLOR_PURPLE:
            return 'purple'
        case Color.COLOR_TEAL:
            return 'teal'
        default:
            return ''
    }
}
