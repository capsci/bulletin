# bulletin
Bulletin helps you quickly generate release notes for the git project you specify.

Based on patterns found in git commit messages, group commits into following buckets:
* Features : New features added
* Fixes : Bug fixes
* Enhancements : Change in behavior of existing
* Misc : Everything else

#### Format
```
Section Name
* Commit short message (commit-id) <author>
```

## Phase 1 - Bare Bones Basic
* [ ] Group commits into buckets based on emojis
  * [ ] Check for reference [Gimoji](https://gist.github.com/parmentf/035de27d6ed1dce0b36a)
* [ ] Support sections Features, Fixes, Enhancements, Misc
* [ ] Select commit range based on tags

## Phase 2 - Bare Bones Extended
* [ ] Group commits into buckets based on emojis or text
* [ ] Select commit range based on tags or commit ids

##
* [ ] Add necessary tests
* [ ] Add/update README.md/ CONTRIBUTING.md
* [ ] Create usable release


## Phase 3 - Customization
* [ ] Add feature to generate TODO list (i.e. no sections from above, but list of commits)
* [ ] Allow ordering/grouping of commits (commit dates/authors/etc)
* [ ] Allow text replacement
  * [ ] replace text to another text
  * [ ] replace text to link

## Phase 4 - YAML
* [ ] Provide YAML configuration
* [ ] Check if YAML configuration is valid
* [ ] Allow specifying fields to include in release notes
* [ ] Allow specifying field to group grouping is based on 
* [ ] Group commits into buckets based pattern/emojis specified in specified in YAML
* [ ] Allow user specified sections

## Phase 5 - a11y
* [ ] Provide templates for release notes
* [ ] specify branch
* [ ] `git clone --bare` if repository is provided
