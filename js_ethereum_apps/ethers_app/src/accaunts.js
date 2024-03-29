import { ethers } from 'ethers'
import * as dotenv from 'dotenv'

dotenv.config({ path: '../.env' })

const INFURA_ID = process.env.INFURA_ID

const provider = new ethers.JsonRpcProvider(`https://sepolia.infura.io/v3/${INFURA_ID}`)

const address = '0xcD00aC6C1d67ebD195BC296B65a29E8196dF2DAb'

const main = async () => {
    const balance = await provider.getBalance(address)
    console.log(`${ethers.formatEther(balance)} ETH`)
}

main()
