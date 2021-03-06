// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateHello } from "./types/blog/tx";
import { MsgUpdateHello } from "./types/blog/tx";
import { MsgDeleteHello } from "./types/blog/tx";


const types = [
  ["/alice.blog.blog.MsgCreateHello", MsgCreateHello],
  ["/alice.blog.blog.MsgUpdateHello", MsgUpdateHello],
  ["/alice.blog.blog.MsgDeleteHello", MsgDeleteHello],
  
];
export const MissingWalletError = new Error("wallet is required");

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgCreateHello: (data: MsgCreateHello): EncodeObject => ({ typeUrl: "/alice.blog.blog.MsgCreateHello", value: data }),
    msgUpdateHello: (data: MsgUpdateHello): EncodeObject => ({ typeUrl: "/alice.blog.blog.MsgUpdateHello", value: data }),
    msgDeleteHello: (data: MsgDeleteHello): EncodeObject => ({ typeUrl: "/alice.blog.blog.MsgDeleteHello", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
