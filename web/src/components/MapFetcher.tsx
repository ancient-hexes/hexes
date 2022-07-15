import { grpc } from '@improbable-eng/grpc-web'
import { createChannel, createClient } from 'nice-grpc-web'
import { onMount } from 'solid-js'

import { Event } from '../proto/hexes/v1/event'
import { EventAPIDefinition } from '../proto/hexes/v1/event_api'
import { PayloadChat } from '../proto/hexes/v1/event_payload'


async function* connectTest(): AsyncIterable<Event> {
    for (let i = 0; i < 10; i++) {
        yield Event.fromPartial({
            chat: PayloadChat.fromPartial({
                text: `[web] ${i}`,
            }),
        })

        // Sleeps for 1 second.
        await new Promise(resolve => setTimeout(resolve, 1000))
    }
}

export const MapFetcher = () => {
    onMount(async () => {
        const channel = createChannel('http://localhost:8080', grpc.WebsocketTransport())
        const client = createClient(EventAPIDefinition, channel)

        for await (const event of client.connect(connectTest() as any)) {
            console.info('received server event: %s', event.chat?.text)
        }
    })

    return (
        <h1>TBA</h1>
    )
}
