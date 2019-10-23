#!/bin/bash

# move /workspace
cd /workspace/src_g

# git add all
git add *

# git commit -m "time"
date1=`date -d '9 hours' '+%y/%m/%d %R'`
echo "$date1"
git commit -m "$date1"

# git push
git push
