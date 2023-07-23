import { ethers } from "hardhat";

async function main() {
  const nuki = await ethers.deployContract("Nuki");

  await nuki.waitForDeployment();

  console.log(
    `Nuki deployed to ${nuki.target}`
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});