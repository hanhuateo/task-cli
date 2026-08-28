# task-cli

- ensure that go 1.26.6 or later is installed

- [project url](https://github.com/hanhuateo/task-cli)

1. clone the repository into your desired folder
2. open terminal
3. cd into the cloned repository
4. run `go install`
5. for mac users, run `echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc && source ~/.zshrc` in terminal
6. for linux users, run `echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.bashrc && source ~/.bashrc` in bash
7. for windows users, run `setx PATH "%PATH%;%USERPROFILE%\go\bin"` in command prompt
8. re-open the terminal that you are using
9. start using by typing `task-cli` 

note: wherever you use the task-cli command, it will create a .json file, so ideally you could simply use it in wherever the default folder location is for your terminal
