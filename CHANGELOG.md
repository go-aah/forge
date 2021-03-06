**v0.8.0** 2018-12-12 Jeevanandam M <jeeva@myjeeva.com>
	* Handling external and native fs more effectively

**v0.7.0** 2018-12-07 Jeevanandam M <jeeva@myjeeva.com>
	* Removed vfs module dependency
	* Added FileSystem interface to extending possibility  
...

**v0.5** 2018-06-20 Jeevanandam M <jeeva@myjeeva.com>
	* Virtual FileSystem capability added `aahframework.org/vfs`

**v0.4.4** 2017-03-22 Jeevanandam M <jeeva@myjeeva.com>
	* Upstream Refresh
	* File closer enhanced
	* Relative subdirectory support added for `include`

**v0.4 to v0.4.3** 2017-03-02 Jeevanandam M <jeeva@myjeeva.com>
* Package versioning tested and released

**v0.3.2** 2017-03-02 Jeevanandam M <jeeva@myjeeva.com>
* Package name updated to aahframework.org/*
* Travis build config update
* Resolving relative path for `include` tag to parent conf directory
* Fix line no and EOF check for scanner

**v0.3.1** 2016-09-02 Jeevanandam M <jeeva@myjeeva.com> (Upstream refresh from v0.2.1)
* Add new `func (*Section) Merge` to merge multiple `Section`s together.

**v0.3.0**	2016-07-02	Jeevanandam M <jeeva@myjeeva.com>
* Forked `github.com/brettlangdon/forge` into `go-aah` org, please refer [here](https://github.com/brettlangdon/forge/issues/13#issuecomment-229080913) why?
* Merged PR [brettlangdon/forge#19](https://github.com/brettlangdon/forge/pull/19/files) from upstream and made required changes
* Version bumped for merging new feature environment variable support
* I'm continuing version number to honor his efforts made nice library `forge` which is similar to `HOCON` syntax but forge direction is different :)

-----

2016-02-26  Brett Langdon  <me@brett.is>

	* v0.2.0: Fix bug with parsing of lists without ending semicolon

	Note: We bumped to `v0.2.0` since we were not doing a great job of following
	semver in previous releases, this version changes aims to correct that mistake.

2015-10-02  Brett Langdon  <me@brett.is>

	* v0.1.7: Allow empty string values

2015-08-09  Brett Langdon  <me@brett.is>

	* v0.1.6: Allow integers in identifiers

2015-07-15  Brett Langdon  <me@brett.is>

	* v0.1.5: Add support for lists

2015-07-14  Brett Langdon  <me@brett.is>

	* v0.1.4: Make semicolons optional

2015-07-14  Brett Langdon  <me@brett.is>

	* v0.1.3: Allow single quoted strings

2015-07-04  Brett Langdon  <me@brett.is>

	* v0.1.2: Allow escaped double quotes and backslashes in strings

2015-06-28  Brett Langdon  <me@brett.is>

	* v0.1.1: Add `Section.Keys` method

2015-06-22  Brett Langdon  <me@brett.is>

	* v0.1.0: Initial release
