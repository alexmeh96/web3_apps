import {createConfig, http} from "wagmi";
import {base, mainnet, optimism, sepolia} from "wagmi/chains";
import { injected, safe, walletConnect } from 'wagmi/connectors'
import {createWeb3Modal} from "@web3modal/wagmi/react";

const projectId = process.env.PROJECT_ID || ''

const metadata = {
    name: 'Web3Modal',
    description: 'Web3Modal Example',
    url: 'https://web3modal.com', // origin must match your domain & subdomain
    icons: ['https://avatars.githubusercontent.com/u/37784886']
}

export const config = createConfig({
    chains: [mainnet, optimism, base],
    connectors: [
        // injected(),
        walletConnect({ projectId, metadata, showQrModal: false }),
        // metaMask(),
        // safe(),
    ],
    transports: {
        [mainnet.id]: http(),
        [optimism.id]: http(),
        [base.id]: http(),
    },
})

createWeb3Modal({
    wagmiConfig: config,
    projectId: projectId,
    enableAnalytics: false
})
