# spdx-parse

This is meant to be the initial cut of a go application to parse an spdx
file

We will need the ability to pass in a spdx file along with some sort of
expected content. The expected content list should be absolute. Anything
found not in the content list should be considered an error.

A bad incomplete example:

spdx-parse -p log4j:2.17.2 -i scanner.json

Then it can return what was or wasn't found. This output format should be
something simple

# TODO
- tests
- some sort of build system
