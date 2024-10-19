# Learn LLD and HLD

This repository is created for the action of learning lld and hld.

# Low Level design

## Solid Principles 🏗️

Solid principles help us to write better code.

1. S - Single responsibility principle 🎯
    
    A class should have only 1 reason to change
    
2. O - Open/Closed principle 🚪
    
    Open for extension but closed for modification.
    
3. L - Liskov substitution principle 🔄
    
    If class B is subtype of class A, then we should be able to replace object of A with B without breaking the behaviour of the program.
    
    Subclass should extend the capability of parent class not narrow it down.
    
4. I - Interface segmented Principle 🧩
    
    Interfaces should be such, that client should not implement unnecessary functions they do not need.
    
5. D - Dependency inversion principle 🔀
    
    Class should depend on interfaces rather than concrete classes.
    
    Always create an object of interface, never create object of concrete class in parent class (which we are implementing).

Advantages of following these principles:

- Avoid duplicate code 🚫
- easy to maintain 🛠️
- easy to understand 🧠
- flexible software 🤸
- reduce complexity 🧵


## How to add .gitignore file

If you want to ignore any generated files like a.out, you can use the .gitignore file.

1. Create a file name '.gitignore'.
2. Add `*.out` inside file.

After this git will be not tracking all .out files inside your repository.

To ignore a folder add `folder_name/` in gitignore. If some file is already tracked by git and you can untrack if using `git rm -r --cached [filename]` without deleting it from local system.