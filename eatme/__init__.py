#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
from functools import partial

from plumbum import local, colors, cli
from plumbum.commands.processes import ProcessExecutionError

__author__ = 'Taras Drapalyuk <taras@drapalyuk.com>'
__date__ = '02.11.2015'
__version__ = '0.0.5'


def get_repos(start_path='.'):
    for file_name in os.listdir(start_path) + [start_path]:
        if os.path.isdir(file_name) and not file_name.startswith('..'):
            if os.path.exists(os.path.join(file_name, '.hg')):
                yield file_name


def run_for_all_repos(func, start_path='.'):
    for repo_path in get_repos(start_path):
        func(path=repo_path)


def pull_update(path, branch=None, clean=False):
    """hg pull + hg update"""
    hg = local['hg']
    hg = hg['--repository', path]
    hg_pull = hg['pull']
    hg_update = hg['update']

    if clean is True:
        hg_update = hg_update['--clean']

    if branch:
        hg_update = hg_update['--rev', branch]

    with colors.green:
        print(path)

    with colors.yellow:
        print(hg_pull)

    print(hg_pull())

    with colors.yellow:
        print(hg_update)

    # Игнорируем ошибки:
    # 255 - unknown revision
    try:
        print(hg_update(retcode=[0]))
    except ProcessExecutionError as e:
        with colors.red:
            print(e.stderr)


def push(path, branch=None, new_branch=True):
    """hg push"""
    hg = local['hg']
    hg = hg['--repository', path]

    hg_push = hg['push']

    if branch:
        hg_push = hg_push['--branch', branch]

    if new_branch is True:
        hg_push = hg_push['--new-branch']

    with colors.green:
        print(path)

    with colors.yellow:
        print(hg_push)

    print(hg_push())


def status(path, branch=None, new_branch=True):
    """hg status"""
    hg = local['hg']
    hg = hg['--repository', path]

    hg_status = hg['status']

    with colors.green:
        print(path)

    with colors.yellow:
        print(hg_status)

    print(hg_status())


class EatMe(cli.Application):
    PROGNAME = 'eatme'
    VERSION = __version__
    verbose = cli.Flag(["v", "verbose"], help="enable additional output")

    def main(self, *args):
        if args:
            print "Unknown command %r" % (args[0],)
            return 1  # error exit code
        if not self.nested_command:  # will be ``None`` if no sub-command follows
            print "No command given"
            return 1  # error exit code


@EatMe.subcommand("push")
class Push(cli.Application):
    new_branch = cli.Flag(["new-branch"], help="hg push --new-branch")
    branch = None

    @cli.switch(["-b", "--branch"], help="hg update --rev BRANCH")
    def set_branch(self, branch):
        self.branch = branch

    def main(self):
        hg_push = partial(push, branch=self.branch, new_branch=self.new_branch)
        run_for_all_repos(hg_push)


@EatMe.subcommand("update")
class Update(cli.Application):
    clean = cli.Flag(["C", "clean"], help="hg update --clean")
    branch = None

    @cli.switch(["-b", "--branch"], help="hg update --rev BRANCH")
    def set_branch(self, branch):
        self.branch = branch

    def main(self):
        hg_pull_update = partial(pull_update, branch=self.branch, clean=self.clean)
        run_for_all_repos(hg_pull_update)


@EatMe.subcommand("status")
class Status(cli.Application):
    def main(self):
        run_for_all_repos(status)


def run_cli():
    import argparse

    parser = argparse.ArgumentParser()
    parser.add_argument('-b', '--branch', help=u'branch name')
    parser.add_argument('-C', '--clean', action='store_true', default=False,
                        help=u'discard local changes (no backup). Can only be used with branch option (-b)')
    parser.add_argument('--push', action='store_true', default=False,
                        help=u'hg push')
    parser.add_argument('--update', action='store_true', default=False,
                        help=u'hg pull + hg update')

    parser.add_argument('--status', action='store_true', default=False,
                        help=u'hg status')

    args = parser.parse_args()

    if args.update:
        hg_pull_update = partial(pull_update, branch=args.branch, clean=args.clean)
        run_for_all_repos(hg_pull_update)

    if args.push:
        hg_push = partial(push, branch=args.branch, clean=args.clean)
        run_for_all_repos(hg_push)

    if args.status:
        run_for_all_repos(status)


if __name__ == '__main__':
    EatMe.run()
    # run_cli()