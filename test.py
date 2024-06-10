import ctypes

lib = ctypes.CDLL('./addons/addons.dll')  # Or hello.so if on Linux.
hello = lib.hello

hello()