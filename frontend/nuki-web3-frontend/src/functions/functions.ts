
import { getContract } from '../utils/utils';
import { ethers } from 'ethers';

/* 
---~~~=*%$}>    addVisitor    <{$%&=~~~---

Inputs:  visitor: address

*/


export async function addVisitor(visitor: string) {
  try {
    const { ethereum } = window;
    const contract = await getContract(ethereum);
    await contract.addVisitor(visitor);
    
  }
  catch (error) {
    console.log(error);
    return "addVisitor failed";
  }
}
/* 
---~~~=*%$}>    isSenderAllowedToEnter    <{$%&=~~~---

Inputs:  home: address

*/


export async function isSenderAllowedToEnter(home: string) {
  try {
    const { ethereum } = window;
    const contract = await getContract(ethereum);
    const data = await contract.isSenderAllowedToEnter(home);
    return data;
  }
  catch (error) {
    console.log(error);
    return "isSenderAllowedToEnter failed";
  }
}
/* 
---~~~=*%$}>    removeVisitor    <{$%&=~~~---

Inputs:  visitor: address

*/


export async function removeVisitor(visitor: string) {
  try {
    const { ethereum } = window;
    const contract = await getContract(ethereum);
    await contract.removeVisitor(visitor);
    
  }
  catch (error) {
    console.log(error);
    return "removeVisitor failed";
  }
}
/* 
---~~~=*%$}>    name    <{$%&=~~~---

Inputs:  

*/


export async function name() {
  try {
    const { ethereum } = window;
    const contract = await getContract(ethereum);
    const data = await contract.name();
    return data;
  }
  catch (error) {
    console.log(error);
    return "name failed";
  }
}