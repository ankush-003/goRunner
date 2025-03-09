# Docker-from-scratch 🐳

## How to create a child process ?

- `fork()` system call is used to create a new process. It creates a new process by duplicating the calling process. The new process, called the child, is an exact copy of the calling process, called the parent, except for a few values that are different. [Read more](https://cs.brown.edu/courses/csci0300/2022/notes/l17.html)
- golang uses `os/exec` package to create a child process. [Read more](https://golang.org/pkg/os/exec/)
- `os/exec` package provides a way to create a child process and execute a command in that process. 

Steps Involved in Execution:

1. Create a Command Struct
   - exec.Command(name string, arg ...string) initializes a command with the given program name and arguments.

2. Start the Process
   - cmd.Start() spawns a new process but does not wait for it to complete.
   - The process is run asynchronously.

3. Wait for the Process to Finish
   - cmd.Wait() waits for the process to complete.
   - This is needed when using cmd.Start().

4. Get Output
   - cmd.Output() runs the command and returns its standard output.
   - cmd.CombinedOutput() returns both stdout and stderr.

## How do we isolate the child process ?

- namespace is a feature of the Linux kernel that partitions kernel resources such that one set of processes sees one set of resources while another set of processes sees a different set of resources.
- The Linux kernel provides six types of namespaces: 
  - Mount (mnt)
  - Process ID (pid)
  - Network (net)
  - Inter-process communication (ipc)
  - UTS (uts includes hostname and NIS domain name)
  - User (user)
  - cgroup (control group means resource limitation)
