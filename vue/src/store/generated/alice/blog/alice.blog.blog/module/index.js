// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateHello } from "./types/blog/tx";
import { MsgDeleteHello } from "./types/blog/tx";
import { MsgUpdateHello } from "./types/blog/tx";
const types = [
    ["/alice.blog.blog.MsgCreateHello", MsgCreateHello],
    ["/alice.blog.blog.MsgDeleteHello", MsgDeleteHello],
    ["/alice.blog.blog.MsgUpdateHello", MsgUpdateHello],
];
export const MissingWalletError = new Error("wallet is required");
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw MissingWalletError;
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee, memo } = { fee: defaultFee, memo: "" }) => client.signAndBroadcast(address, msgs, fee, memo),
        msgCreateHello: (data) => ({ typeUrl: "/alice.blog.blog.MsgCreateHello", value: data }),
        msgDeleteHello: (data) => ({ typeUrl: "/alice.blog.blog.MsgDeleteHello", value: data }),
        msgUpdateHello: (data) => ({ typeUrl: "/alice.blog.blog.MsgUpdateHello", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
