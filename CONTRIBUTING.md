# How to contribute

## Submitting changes

Please submit changes by creating a
[GitHub Pull-Request](https://github.com/simplesurance/bunny-go/pulls).

- A Pull-Request (PR) should contain only one isolated change. Do not mix multiple
  logical changes in one Pull-Request, e.g. a bugfix, a new feature, an
  unrelated vendor package upgrade and some formatting changes.
  Small Pull-Request are easier to review, can be faster processed and merged.
  They also make it simple to revert specific changes that introduced a bug.
  Big Pull-Request take longer to get reviewed and merged and risk that a
  perfect change is rejected because of another unrelated change in the same PR.
- Add a description about what your Pull-Request changes and why. If you have
  written descriptive commit messages, they can be reused or referenced as
  Pull-Request description.

### Commit Messages

- Create a commit per logical change and document it in the commit message.
  The popular commit message template from
  [Tim Pope](https://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html):
  can be followed:

```
Capitalized, short (50 chars or less) summary

More detailed explanatory text, if necessary.  Wrap it to about 72
characters or so.  In some contexts, the first line is treated as the
subject of an email and the rest of the text as the body.  The blank
line separating the summary from the body is critical (unless you omit
the body entirely); tools like rebase can get confused if you run the
two together.

Write your commit message in the imperative: "Fix bug" and not "Fixed bug"
or "Fixes bug."  This convention matches up with commit messages generated
by commands like git merge and git revert.

Further paragraphs come after blank lines.

- Bullet points are okay, too

- Typically a hyphen or asterisk is used for the bullet, followed by a
  single space, with blank lines in between, but conventions vary here

- Use a hanging indent
```

