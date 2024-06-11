from setuptools import setup, find_packages
import os
import platform
from setuptools.command.build_ext import build_ext as _build_ext
import subprocess
import shutil

class build_ext(_build_ext):
    def run(self):

        subprocess.run(["python", "web1337/build_addons.py"], check=True)
        
        system = platform.system()
        lib_name = ""
        if system == "Windows":
            lib_name = "addons.dll"
        elif system == "Linux":
            lib_name = "addons.so"
        else:
            raise ValueError("Unsupported OS")

        if lib_name:
            shutil.copy(lib_name, os.path.join(self.build_lib, 'web1337', lib_name))
        
        super().run()


system = platform.system()
if system == "Windows":
    ext_name = "web1337.addons"
    ext_files = ["addons.dll"]
elif system == "Linux":
    ext_name = "web1337.addons"
    ext_files = ["addons.so"]
else:
    raise ValueError("Unsupported OS")



setup(
    name='web1337',
    version='0.1.0',
    description='Python SDK to work with KLYNTAR',
    long_description=open('README.md').read(),
    long_description_content_type='text/markdown',
    author='KLY Foundation',
    author_email='hello@klyntar.org',
    url='https://github.com/KLYN74R/Web1337Python',
    packages=find_packages(),
    package_data={'web1337': ext_files},
    include_package_data=True,
    cmdclass={'build_ext': build_ext},
    classifiers=[
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.9',
        'License :: OSI Approved :: MIT License',
        'Operating System :: OS Independent',
    ],
    python_requires='>=3.6',
    install_requires=[],
    ext_modules=[],
)
