# Getting Started 
1.Installing the CLI <br>
2.Activating mise or adding its shims to PATH <br>
3.Adding tools to mise <br>

# Quick start 

## 1. Install mise CLI
First we need to download the mise CLI.<br>
This directory is simply a suggestion. mise can be installed anywhere.
```shell
curl https://mise.jdx.dev/install.sh | sh
~/.local/bin/mise --version
mise 2024.x.x
```
## 2.Activate mise 
Make sure you restart your shell session after modifying your rc file in order for it to take effect. <br>
```shell
echo 'eval "$(~/.local/bin/mise activate bash)"' >> ~/.bashrc
```
then use this command 
```shell
export MISE_GLOBAL_CONFIG_FILE=./config.toml

```
for update mise
```shell
mise self-update
```