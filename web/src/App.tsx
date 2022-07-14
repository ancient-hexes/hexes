import type { Component } from 'solid-js'

import logo from './logo.svg'
import styles from './App.module.css'

import { MapFetcher } from './components/MapFetcher'

const App: Component = () => {
    return (
        <div class={ styles.App }>
            <header class={ styles.header }>
                <img src={ logo } class={ styles.logo } alt="logo"/>
            </header>
            <MapFetcher/>
        </div>
    )
}

export default App
