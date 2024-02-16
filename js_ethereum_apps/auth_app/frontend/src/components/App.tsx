import {SiweMessage} from "siwe";
import {provider, signer} from "../index";

export const App = () => {

    const init = async () => {
        try {
            const res = await fetch('http://localhost:8085/api/validate', {
                credentials: 'include',
            })

            console.log(await res.text())
        } catch (err) {

            console.log(err)
        }
    }
    //
    const connectWallet = async () => {
        try {
            const accounts = await provider
                .send('eth_requestAccounts', [])
                .catch(() => console.log('user rejected request'))
            if (accounts[0]) {
                console.log(accounts[0])
            }
        } catch (e) {
            console.error(e)
        }
    }

    const signin = async () => {
        try {
            // Get nonce
            const res = await fetch('http://localhost:8085/api/nonce', {
                credentials: 'include',
            })

            console.log(res)

            // Create message
            const messageRaw = new SiweMessage({
                domain: window.location.host,
                address: await signer.getAddress(),
                statement: 'Sign in with Ethereum to the app.',
                uri: window.location.origin,
                version: '1',
                chainId: 1,
                nonce: await res.text(),
            })

            const message = messageRaw.prepareMessage()

            // Get signature
            const signature = await signer.signMessage(message)

            // Send to server
            const res2 = await fetch('http://localhost:8085/api/verify', {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({ message, signature }),
                credentials: 'include',
            })

            console.log(res2)
        } catch (err) {
            console.error(err)
        }
    }


    return (
        <div>
            <button onClick={connectWallet}>connectWallet</button>
            <button onClick={init}>init</button>
            <button onClick={signin}>signin</button>
        </div>
    );
};
