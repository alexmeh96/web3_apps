import {createConfig, getBalance, getBlock, http} from '@wagmi/core'
import {sepolia} from '@wagmi/core/chains'
import * as dotenv from 'dotenv'

dotenv.config({ path: '../.env' })

const INFURA_ID = process.env.INFURA_ID

const config = createConfig({
    // chains: [mainnet, sepolia],
    chains: [sepolia],
    transports: {
        [sepolia.id]: http(`https://sepolia.infura.io/v3/${INFURA_ID}`),
        // [mainnet.id]: http(`https://mainnet.infura.io/v3/${INFURA_ID}`),
    },
})

const main = async () => {

    const blockNumber = await getBlock(config)

    const balance = await getBalance(config, {
        address: '0xcD00aC6C1d67ebD195BC296B65a29E8196dF2DAb',
    })

    console.log(`blockNumber: ${blockNumber.number}`)
    console.log(`balance: ${balance.value}`)
}

main()
