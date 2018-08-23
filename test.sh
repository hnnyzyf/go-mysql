#!/bin/bash

#get the path of test file
SHELL_FOLDER=$(dirname $(readlink -f "$0"))

echo "====Workspace Is "$SHELL_FOLDER"===="

echo "==============Testing==============="

 function read_dir(){
    for file in `ls $1`       
    do
        if [ -d $1"/"$file ] 
        then
            read_dir $1"/"$file
        else
            if !([[ $file =~ "_test.go" ]]);then
                continue
            else
               echo Test:$1
               cd $1
               echo `go test -v *.go`
               break
            fi
        fi
    done
}   


for folder in `ls $SHELL_FOLDER`   
do
    if [ -d $SHELL_FOLDER"/"$folder ] && !([[ $folder =~ "vendor" ]])
    then
        read_dir $SHELL_FOLDER"/"$folder
    fi
done




