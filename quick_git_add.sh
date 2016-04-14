#!/bin/sh
git add *.py
find . -name "*.go" | xargs git add
git add *.sh
git add README.md
git add .gitignore
git status
