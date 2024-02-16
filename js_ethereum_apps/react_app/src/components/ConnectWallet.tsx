import React from "react";
import {Account} from "./Account";
import {useAccount} from "wagmi";
import {ConnectButton} from "./ConnectButton";

export function ConnectWallet() {
    const { isConnected } = useAccount()
    if (isConnected) return <Account />
    return <ConnectButton />
}
