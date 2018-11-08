#!/bin/sh

if [ ! -d .git ];then
    echo "Execute scripts/install-git-hooks in the top-level directory."
    exit 1
fi

ln -sf ../../scripts/commit-msg.hook .git/hooks/commit-msg || exit 1
chmod u+x .git/hooks/commit-msg

touch .git/hooks/applied || exit

echo
echo "Git commit hooks are installed successfuly."
