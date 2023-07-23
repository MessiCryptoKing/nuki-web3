
import { ethers } from 'ethers';
import { MetaMaskInpageProvider } from "@metamask/providers";
import Nuki from '../Nuki.json';

export const contractAddress = '0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9';

export const getContract = async (ethereum: MetaMaskInpageProvider) => {
  const provider = new ethers.BrowserProvider(ethereum as any);
  const signer = await provider.getSigner();
  const connectedContract: ethers.Contract = new ethers.Contract(contractAddress, Nuki.abi, signer);
  return connectedContract;
}
