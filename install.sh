#!/bin/bash
BIN_NAME=syn
INSTALL_PATH=/usr/local/share/$BIN_NAME
WD=${PWD##*/}
if [[ ! $WD == $BIN_NAME ]]; then
	echo "Please run script from $BIN_NAME folder."
	exit 1
fi
echo "Removing old installation.."
sudo rm -rf $INSTALL_PATH
echo "Copying files.."
if [ "$(uname)" == "Darwin" ]; then
	sudo cp -rf ./dist/macOS $INSTALL_PATH
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
	sudo cp -rf ./dist/linux $INSTALL_PATH
fi
sudo ln -sf $INSTALL_PATH/$BIN_NAME /usr/local/bin/$BIN_NAME
sudo chmod +x $INSTALL_PATH/$BIN_NAME

echo "Done! Installed at /usr/local/share/$BIN_NAME"
echo "Use this command to get started: "
echo "> $BIN_NAME -h"
