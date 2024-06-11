import os
import platform
import subprocess

def build_go_lib():
    system = platform.system()
    if system == "Windows":
        lib_name = "addons.dll"
        subprocess.run(["go", "build", "-o", lib_name, "addons/addons.go"])
    elif system == "Linux":
        lib_name = "addons.so"
        subprocess.run(["go", "build", "-o", lib_name, "-buildmode=c-shared", "addons/addons.go"])
    else:
        raise ValueError("Unsupported OS")

    return lib_name

if __name__ == "__main__":
    build_go_lib()