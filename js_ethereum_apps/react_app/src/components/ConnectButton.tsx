import * as React from 'react'
import {Connector, useConnect} from 'wagmi'
import {useWeb3Modal} from '@web3modal/wagmi/react'

export function ConnectButton() {
    // const {connectors, connect} = useConnect()

    const {open} = useWeb3Modal()

    return (
        <div>
            {/*{*/}
            {/*    connectors.map((connector) => (*/}
            {/*        <button key={connector.uid} onClick={() => connect({connector})}>*/}
            {/*            {connector.name}*/}
            {/*        </button>*/}
            {/*    ))*/}
            {/*}*/}
            <button onClick={() => open()}>Open Connect Modal</button>
        </div>
    )
}
