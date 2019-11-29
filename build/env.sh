#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
fnsdir="$workspace/src/github.com/fns"
if [ ! -L "$fnsdir/fns" ]; then
    mkdir -p "$fnsdir"
    cd "$fnsdir"
    ln -s ../../../../../. fns
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$fnsdir/fns"
PWD="$fnsdir/fns"

# Launch the arguments with the configured environment.
exec "$@"
