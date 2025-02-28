# task-tracker
Task tracker is a simple command line interface (CLI) used to 
track what you need to do, what you have done, and what you are currently working on.\
Solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) project on [roadmap.sh](https://roadmap.sh)

## Features
- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

## Building from source
Ensure the GO SDK is installed
1. Clone the repository
   ```bash
   git clone git@github.com:hayohtee/task-tracker.git
   ```
3. Change into the project directory
   ```bash
   cd task-tracker
   ```
4. Compile
   ```bash
   go build -o task-cli ./cmd/task
   ```

## Usage
```bash
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
