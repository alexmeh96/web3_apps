import {createRoot} from "react-dom/client";
import {App} from "./App";
import {QueryClient} from '@tanstack/react-query'

export const queryClient = new QueryClient()

const root = document.getElementById('root')

if (!root) {
    throw new Error('root not found')
}

const container = createRoot(root)

container.render(<App />)
