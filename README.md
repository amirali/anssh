# Ansible inventory file ssh utility
Problem: I couldn't memorize all of the ansible inventory hosts and groups.  
Solution: So I wrote this script to help me.

## Installation
```bash
go install github.com/amirali/anssh@latest
```

## Usage
You can use this to select your host from the inventory file using an interactive shell and then ssh into it.
```bash
anssh -user=root -identity=/path/to/identity/file -inv=first_inventory.ini -inv=second_inventory.ini ...
```
