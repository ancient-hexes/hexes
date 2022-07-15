import type { Component } from 'solid-js'

import { MapView } from './components/MapView'
import { generateRandomMap } from './components/MapView/sample'
import { StatusBar } from './components/StatusBar'


const App: Component = () => {
    return (
        <>
            <MapView map={ generateRandomMap(64, 64) }/>
            <StatusBar/>
        </>
    )
}

export default App
