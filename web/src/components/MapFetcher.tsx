import { createChannel, createClient } from 'nice-grpc-web'
import { onMount } from 'solid-js'
import { Environment_ListRequest } from '../proto/hexes/v1/env'
import { EnvironmentAPIDefinition } from '../proto/hexes/v1/environment_api'


export const MapFetcher = () => {
    onMount(async () => {
        const channel = createChannel('http://localhost:8080')
        console.info('channel', channel)

        const client = createClient(EnvironmentAPIDefinition, channel)
        const request = Environment_ListRequest.fromPartial({})
        const response = await client.list(request as any)
        console.info('got response:', response)
        for await (const item of response) {
            console.info('=>', item)
        }
    })

    return (
        <h1>TBA</h1>
    )
}
