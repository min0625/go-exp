#!/bin/bash

echo 'eval "export MISE_GO_SET_GOBIN=false"' >>~/.bash_profile
echo 'eval "export MISE_GO_SET_GOBIN=false"' >>~/.bashrc

# note that bash will read from ~/.profile or ~/.bash_profile if the latter exists
# ergo, you may want to check to see which is defined on your system and only append to the existing file
echo 'eval "$(mise activate bash --shims)"' >>~/.bash_profile # this sets up non-interactive sessions
echo 'eval "$(mise activate bash)"' >>~/.bashrc               # this sets up interactive sessions

mise trust --all

mise install
