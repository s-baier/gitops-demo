# Contributing

To enable effective and efficient collaboration, we would like to establish some contributing standards.

## File & Directory Naming Conventions

To avoid issues dealing with handling of case sensitivity by various operating systems, please adhere to the following ground rules:

* Keep file and directory names within this repository all small caps.
* Use dashes instead of whitespace for separation to ease the use of tab completion.

This might be over-ruled by language / framework specific standards.

## Merge Request Mechanism

The main branch is protected.
Therefore, create a branch for your changeset, and create a merge request to merge (squash commits) your proposed changes into main.
All merge requests require the approval of at least 4 eyes.
The person initiating the merge request, is responsible to execute the merge after approval.

## Git Commit

The cardinal rule for creating good commits is to ensure there is only one "logical change" per commit. There are many reasons why this is an important rule:

The smaller the amount of code being changed, the quicker & easier it is to review & identify potential flaws.
If a change is found to be flawed later, it may be necessary to revert the broken commit.
This is much easier to do if there are not other unrelated code changes entangled with the original commit.
When troubleshooting problems using Git's bisect capability, small well defined changes will aid in isolating exactly where the code problem was introduced.
When browsing history using Git annotate/blame, small well defined changes also aid in isolating exactly where & why a piece of code came from.

### Git Commit Messages

* Commit messages must have a subject line and may have body copy. These must be separated by a blank line.
* The subject line must not exceed 50 characters.
* The subject line should be capitalized and must not end in a period.
* The subject line must be written in imperative mood (Fix, not Fixed / Fixes etc.).
* The body copy must be wrapped at 72 columns.
* The body copy must only contain explanations as to what and why, never how. The latter belongs in documentation and implementation.

### Information in Git Messages

* Describe *why* a change is being made.
* How does it address the issue?
* What effects does the patch have?
* Do not assume the reviewer understands what the original problem was.
* Do not assume the reviewer has access to external web services/site.
* Do not assume the code is self-evident/self-documenting.
* Read the commit message to see if it hints at improved code structure.
* The first commit line is the most important.
* Describe any limitations of the current code.
* Do not include patch set-specific comments.

### Subject Line Standard Terminology

| First Word | Meaning                                              |
| ---------- | ---------------------------------------------------- |
| Add        | Create a capability e.g. feature, test, dependency.  |
| Cut        | Remove a capability e.g. feature, test, dependency.  |
| Fix        | Fix an issue e.g. bug, typo, accident, misstatement. |
| Bump       | Increase the version of something e.g. dependency.   |
| Make       | Change the build process, or tooling, or infra.      |
| Start      | Begin doing something; e.g. create a feature flag.   |
| Stop       | End doing something; e.g. remove a feature flag.     |
| Refactor   | A code change that MUST be just a refactoring.       |
| Reformat   | Refactor of formatting, e.g. omit whitespace.        |
| Optimize   | Refactor of performance, e.g. speed up code.         |
| Document   | Refactor of documentation, e.g. help files.          |

### Example Git Commit Message

```log
Refactor libvirt get_cpu_info method to use config APIs

The get_cpu_info method in the libvirt driver currently uses
XPath queries to extract information from the capabilities
XML document. Switch this over to use the new config class
LibvirtConfigCaps. Also provide a test case to validate
the data being returned.

DocImpact
Closes-Bug: #1003373
Implements: blueprint libvirt-xml-cpu-model
Change-Id: I4946a16d27f712ae2adf8441ce78e6c0bb0bb657
```
