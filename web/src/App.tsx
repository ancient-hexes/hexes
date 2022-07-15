import type { Component } from 'solid-js'
import { MapFetcher } from './components/MapFetcher'

import { MapView } from './components/MapView'
import { generateRandomMap } from './components/MapView/sample'
import { StatusBar } from './components/StatusBar'


const App: Component = () => {
    return (
        <>
            <MapFetcher/>
            <MapView map={ generateRandomMap(64, 64) }/>
            <StatusBar/>
        </>
    )
}

export default App
