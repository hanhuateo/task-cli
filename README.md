# task-cli

- ensure that go 1.26.6 or later is installed

1. clone the repository into your desired folder
2. open terminal
3. cd into the cloned repository
4. run `go install`
5. for mac users, run `echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.zshrc && source ~/.zshrc` in terminal
6. re-open the terminal that you are using
7. start using by typing `task-cli`

e.g. of how to use 
```go
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

note: wherever you use the task-cli command, it will create a .json file, so ideally you could simply use it in wherever the default folder location is for your terminal
