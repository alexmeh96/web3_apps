import {createPublicClient, http} from 'viem'
import {mainnet} from "viem/chains";

const INFURA_ID = ''


const client = createPublicClient({
    chain: mainnet,
    transport: http(`https://sepolia.infura.io/v3/${INFURA_ID}`)
})


const main = async () => {
    const blockNumber = await client.getBlockNumber()

    const balance = await client.getBalance({
        address: '0xcD00aC6C1d67ebD195BC296B65a29E8196dF2DAb',
    })

    console.log(`blockNumber = ${blockNumber}`)
    console.log(`balance = ${balance}`)
}

main()

