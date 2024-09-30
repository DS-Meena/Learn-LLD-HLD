# Learn LLD and HLD

This repository is created for the action of learning lld and hld.

## How to add .gitignore file

If you want to ignore any generated files like a.out, you can use the .gitignore file.

1. Create a file name '.gitignore'.
2. Add `*.out` inside file.

After this git will be not tracking all .out files inside your repository.

To ignore a folder add `folder_name/` in gitignore. If some file is already tracked by git and you can untrack if using `git rm -r --cached [filename]` without deleting it from local system.