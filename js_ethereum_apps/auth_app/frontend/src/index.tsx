import {createRoot} from "react-dom/client";
import {App} from "./components/App";
import { ethers } from 'ethers'

const root = document.getElementById('root')

if(!root) {
    throw new Error('root not found')
}

// @ts-ignore
export const provider = new ethers.providers.Web3Provider(window.ethereum)
export const signer = provider.getSigner()

const container = createRoot(root)

container.render(<App />)
