@echo off
set /p message="Enter commit message: "
set /p tag="Enter tag: "

git add .
git commit -m "%message%"


git tag v%tag%
git push origin main --tags
