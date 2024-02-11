import { ethers } from 'ethers'
import * as dotenv from 'dotenv'

dotenv.config({ path: '../.env' })

const INFURA_ID = process.env.INFURA_ID
const provider = new ethers.JsonRpcProvider(`https://kovan.infura.io/v3/${INFURA_ID}`)

const account1 = '' // Your account address 1
const account2 = '' // Your account address 2

const privateKey1 = '' // Private key of account 1
const wallet = new ethers.Wallet(privateKey1, provider)

const main = async () => {
    const senderBalanceBefore = await provider.getBalance(account1)
    const recieverBalanceBefore = await provider.getBalance(account2)

    console.log(`\nSender balance before: ${ethers.formatEther(senderBalanceBefore)}`)
    console.log(`reciever balance before: ${ethers.formatEther(recieverBalanceBefore)}\n`)

    const tx = await wallet.sendTransaction({
        to: account2,
        value: ethers.parseEther("0.025")
    })

    await tx.wait()
    console.log(tx)

    const senderBalanceAfter = await provider.getBalance(account1)
    const recieverBalanceAfter = await provider.getBalance(account2)

    console.log(`\nSender balance after: ${ethers.formatEther(senderBalanceAfter)}`)
    console.log(`reciever balance after: ${ethers.formatEther(recieverBalanceAfter)}\n`)
}

main()
