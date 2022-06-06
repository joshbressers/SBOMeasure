# SBOMeasure
Tooling to test and measure SBOM generators

Let’s create some tooling that allows various SBOM tools to be run through
a set of tests to show what they do and don’t pick up correctly. Some of
these tests will be simple. Some will be purposely bad intended to trick
the scanners.

The intent isn’t to declare one scanner the best. It’s to give the tool
authors a concrete list of bugs and features to work on.

We will expand the testing in ways that nudge the SBOM creators to add new
and interesting features to their tools.

# Goals

- No human interaction. The tools should be fully automated. SBOM tooling
  that requires manual input will not be considered
- Tool should output CycloneDX or SPDX
- Test should run scanner only against a pre-determined set of data. Do not
  try to scan Java artifacts with an npm scanner

# Non-Goals

- There will not be a "winner". The intent isn't to show one scanner is
  better or worse than another scanner. The intent is to help scanners
  identify and fill gaps


# Testcases
We want to collect reasonable examples of difficult to scan artifacts to
point SBOM scanners at.

When scanning docker images it will be reasonable to assume the local
system contains a functioning instance of docker.

Example data should be placed in the data directory.

Not every tool will be able to scan every type of testcase (for example a
scanner that only scans one language, or doesn't scan containers). We will
need a way to denote which example datasets are appropriate for each
scanner.

# Scanners

What are the various SBOM scanners and how will we run them?

Ideally every scanner should get its own directory in this system that will
allow for scripts to be written that can install and run the scanner.

## Scanners to investigate

- [Dependency-Check](https://owasp.org/www-project-dependency-check/)
- [FOSSology](https://github.com/fossology/fossology)
- [npm-spdx](https://github.com/swinslow/npm-spdx)
- [OSS Review Toolkit](https://github.com/oss-review-toolkit/ort)
- [Quartermaster](https://github.com/QMSTR/qmstr)
- [reuse](https://git.fsfe.org/reuse/tool)
- [Scanode toolkit](https://github.com/nexB/scancode-toolkit)
- [Syft](https://github.com/anchore/syft/)
- [Tern](https://github.com/tern-tools/tern)

# Output parsing

We will need tooling that can parse the output of the scanner and
understand if the results line up with expectations. For example if a
scanner is expecing a certain version of log4j to be present, we should
specifically check for that.

Ideally we will not need a parser for every scanner as they should be
outputting in standard formats. A scanner that does not output SPDX or
CycloneDX that is well formed should be considered a fail.

We will need the ability to save the output data in a manner than can be
easily displayed. The sort of graph that is shown here would be friendly
and nice
https://www.rezilion.com/blog/log4j-blindspots-what-your-scanner-is-still-missing/

# TODO
Here is an incomplete list of things that need to be done. Feel free to
jump in and contribute. PRs are most welcome!

- Identify SBOM scanners. There are a lot of these
- Document how to install and run an SBOM sanner
- Create testcases (we will need a lot of these)
- SPDX and CycloneDX parsing applications (Josh has started a go spdx
  parser on the josh-dev branch)
- Setup a github actions for the repo. It will be a mono repo and will need
  a nice way to automatically verify PRs
