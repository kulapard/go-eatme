#!/usr/bin/env python
# -*- coding: utf-8 -*-
import re

try:
    from setuptools import setup
except ImportError:
    from distutils.core import setup
import os

with open('eatme/__init__.py', 'r') as fd:
    PACKAGE_VERSION = re.search(r'^__version__\s*=\s*[\'"]([^\'"]*)[\'"]',
                                fd.read(), re.MULTILINE).group(1)

if not PACKAGE_VERSION:
    raise RuntimeError('Cannot find version information')

this_dir = os.path.dirname(__file__)
readme_filename = os.path.join(this_dir, 'README.rst')
requirements_filename = os.path.join(this_dir, 'requirements.txt')

with open(readme_filename) as f:
    PACKAGE_LONG_DESCRIPTION = f.read()

with open(requirements_filename) as f:
    # Игнорируем комментарии,
    # пустые строки и пакеты из репозиториев (-e)
    PACKAGE_INSTALL_REQUIRES = [
        line.strip('\n') for line in f
        if not line.startswith('#') and not line.startswith('-e') and line.strip()
        ]

setup(
    name='EatMe',
    version=PACKAGE_VERSION,
    author='Taras Drapalyuk',
    author_email='taras@drapalyuk.com',
    description='EatMe Utility',
    url='https://bitbucket.org/KulaPard/eatme',
    long_description=PACKAGE_LONG_DESCRIPTION,
    packages=[
        'eatme',
    ],
    entry_points={
        'console_scripts': [
            'eatme = eatme:EatMe.run',
        ],
    },
    install_requires=PACKAGE_INSTALL_REQUIRES,
)
