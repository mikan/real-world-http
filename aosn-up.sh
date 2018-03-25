#!/bin/bash
git push
cd ../../aosn/real-world-http
git pull
git submodule update
cd mikan
git pull origin master
cd ..
git add mikan
git commit -m ":tangerine: update mikan"
git push
cd ../../mikan/real-world-http