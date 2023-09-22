@echo off
set /p message="Enter commit message: "
set /p tag="Enter tag: "

git add .
git commit -m "%message%"
git tag %tag%
git push --tags