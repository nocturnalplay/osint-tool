# Importing required module
import subprocess
import sys
# Using system() method to
# execute shell commands
subprocess.Popen(f'harvest {sys.argv[1]}', shell=True)