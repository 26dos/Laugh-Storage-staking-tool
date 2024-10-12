
# README

## 1. Compilation

### Homebrew

We recommend that macOS users use Homebrew to install each of the necessary packages.

Use the command `brew install` to install the following packages:

```sh
brew install go jq pkg-config hwloc coreutils
```

Next, clone the `dc-pay-channel-cli` repository and build the executables.

### Rust

Rustup is an installer for the systems programming language Rust. Run the installer and follow the onscreen prompts. The default installation option should be chosen unless you are familiar with customization:

```sh
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

### Build and Install `dc-pay-channel-cli`

The installation instructions are different depending on which CPU is in your Mac:

- **M1-based CPUs**
- **Intel CPUs**

#### M1-based CPUs

Clone the repository:

```sh
git clone https://github.com/26dos/dc-pay-channel-cli.git
cd dc-pay-channel-cli.git/
```

Set environment variables:

```sh
export LIBRARY_PATH=/opt/homebrew/lib
export FFI_BUILD_FROM_SOURCE=1
export PATH="$(brew --prefix coreutils)/libexec/gnubin:/usr/local/bin:$PATH"
```

Build `dc-pay-channel-cli`:

```sh
make
```

#### Intel CPUs

> ❗ These instructions are for installing Lotus on an Intel Mac. If you have an M1-based CPU, use the M1-based CPU instructions above.

Clone the repository:

```sh
git clone https://github.com/26dos/dc-pay-channel-cli.git
cd dc-pay-channel-cli/
```

Build and install `dc-pay-channel-cli`:

```sh
make
```

## 2. Usage

1. **Create an F4 Wallet and Send FIL for Staking**

   ```sh
   ./dc-pay-channel-cli new delegated
   ```

2. **Stake FIL**

   Based on the requested `datacap`, staking requirements are as follows:
    - Public datasets require 20 days of staking.
    - Private datasets require 40 days of staking.
    - Staking 10T requires 1 FIL.
    - FIL must be claimed after the staking period ends.

   ```sh
   ./dc-pay-channel-cli stake --data-set-type 0 --dataCap 10 --from f410fhjllu2s5xvujvpvjgeznzqniu3fd3qkts6he6ia
   ```

3. **View Staking Information**

   ```sh
   ./dc-pay-channel-cli getAllStakeInfo --from f410fhjllu2s5xvujvpvjgeznzqniu3fd3qkts6he6ia
   ```

4. **Claim Staked FIL After Expiry**

   ```sh
   ./dc-pay-channel-cli unlock --stakeId 0 --from f410fhjllu2s5xvujvpvjgeznzqniu3fd3qkts6he6ia
   ```
