# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [0.2.1](https://github.com/NateScarlet/qidian/compare/v0.2.0...v0.2.1) (2021-10-11)


### Bug Fixes

* can not parse empty count ([bb6c3e2](https://github.com/NateScarlet/qidian/commit/bb6c3e20758ba791f76a7da7cb912493aa68f431))
* follow book info page change ([e201b42](https://github.com/NateScarlet/qidian/commit/e201b42daacdab393eaf7698b9e140cf80761935))

## [0.2.0](https://github.com/NateScarlet/qidian/compare/v0.1.9...v0.2.0) (2021-08-19)


### ⚠ BREAKING CHANGES

* change type of (book.RankType).URL from string to url.URL
* use options API for book.CategorySearch
* use options API for book.Rank

### Features

* change type of (book.RankType).URL from string to url.URL ([52f370c](https://github.com/NateScarlet/qidian/commit/52f370c7ae356b5f92e74d624eb80f4703068437))
* use options API for book.CategorySearch ([c4ac056](https://github.com/NateScarlet/qidian/commit/c4ac056b1f646bac62be3171fb6c1cd0e8a9caf0))
* use options API for book.Rank ([f5c7acf](https://github.com/NateScarlet/qidian/commit/f5c7acff850a8907547c862cde0b6c589d946fa3))

### [0.1.9](https://github.com/NateScarlet/qidian/compare/v0.1.8...v0.1.9) (2021-08-19)


### Features

* add Rank.Page field ([bada173](https://github.com/NateScarlet/qidian/commit/bada1738af7c128cbbc41a45d23612eb72bff5a0))

### [0.1.8](https://github.com/NateScarlet/qidian/compare/v0.1.7...v0.1.8) (2021-08-19)


### Bug Fixes

* **deps:** update module github.com/natescarlet/snapshot to v0.6.0 ([709e584](https://github.com/NateScarlet/qidian/commit/709e584ae51bb444d0accc6ce7ec35fa57421834))
* **deps:** update module github.com/puerkitobio/goquery to v1.7.0 ([513c265](https://github.com/NateScarlet/qidian/commit/513c2650a6ac620a81dce350e8af21be0d632223))
* **deps:** update module github.com/puerkitobio/goquery to v1.7.1 ([d3c46f7](https://github.com/NateScarlet/qidian/commit/d3c46f71f3f5d9dff002e5724c3bac409ec77ce3))
* follow site change ([bf9f05e](https://github.com/NateScarlet/qidian/commit/bf9f05e1d9e9c9a99ca5c5c1436d5f3ed44922b6))

### [0.1.7](https://github.com/NateScarlet/qidian/compare/v0.1.6...v0.1.7) (2021-01-13)


### Features

* use fixed timezone ([a268cf0](https://github.com/NateScarlet/qidian/commit/a268cf04eeb981ef11999c971b2cda4bd1b10d69))


### Bug Fixes

* handle book table column name change ([5ab29b0](https://github.com/NateScarlet/qidian/commit/5ab29b0a803254e2b6b736f2a2cda9e166c60c0f))

### [0.1.6](https://github.com/NateScarlet/qidian/compare/v0.1.5...v0.1.6) (2020-10-18)


### Bug Fixes

* error when fetch free book ([7bc141a](https://github.com/NateScarlet/qidian/commit/7bc141a2b287314138806957399e7d55720020aa))

### [0.1.5](https://github.com/NateScarlet/qidian/compare/v0.1.4...v0.1.5) (2020-10-18)


### Features

* add (book.Rank).URL method ([9ea54f8](https://github.com/NateScarlet/qidian/commit/9ea54f84f9f7b1aec5de4203f79baa1449986f28))

### [0.1.4](https://github.com/NateScarlet/qidian/compare/v0.1.3...v0.1.4) (2020-10-18)


### Features

* add (CategorySearch).URL method ([b6a78da](https://github.com/NateScarlet/qidian/commit/b6a78da1d61728825fd7726d60f57f5888d6db0c))

### [0.1.3](https://github.com/NateScarlet/qidian/compare/v0.1.2...v0.1.3) (2020-10-18)


### Features

* improve error message ([7b31beb](https://github.com/NateScarlet/qidian/commit/7b31bebadc8c279572928719155bf5f6406e0b5e))

### [0.1.2](https://github.com/NateScarlet/qidian/compare/v0.1.1...v0.1.2) (2020-10-17)


### Features

* fetch book cover url ([845525d](https://github.com/NateScarlet/qidian/commit/845525d2298c4a8ae56aab3d64add106862508d8))

### [0.1.1](https://github.com/NateScarlet/qidian/compare/v0.1.0...v0.1.1) (2020-10-16)

## 0.1.0 (2020-10-16)


### Features

* allow use custom http.Client ([479d30a](https://github.com/NateScarlet/qidian/commit/479d30a36900567af2f7efb4ea6a85dffedcb67e))
* book use author struct ([333b1b1](https://github.com/NateScarlet/qidian/commit/333b1b1ad8530603cafa82585a6fae3c34abe358))
* category search ([dae0708](https://github.com/NateScarlet/qidian/commit/dae0708887389d9bab70f7df109266d18da8a3d5))
* category search novel by tag ([bbae36b](https://github.com/NateScarlet/qidian/commit/bbae36bc2e146fa41d30057df5bca74610578002))
* fetch author data ([6d26219](https://github.com/NateScarlet/qidian/commit/6d26219528d16b3d8e7afeea5362892a96fb5d34))
* fetch book info ([566ab55](https://github.com/NateScarlet/qidian/commit/566ab55083d8b080df8c6bb8137363410aff765e))
* font deobfuscate ([464ec96](https://github.com/NateScarlet/qidian/commit/464ec96e7b1c948c5be898592e628cc2d5abe82c))
* generate category.go ([78c1f4f](https://github.com/NateScarlet/qidian/commit/78c1f4ff5b95658fc94e550fd0c272b4865e0a1f))
* implement book search ([4411299](https://github.com/NateScarlet/qidian/commit/4411299b4c9d294411bcf3cee612363047dde5f1))
* implement rank fetch ([128d8f9](https://github.com/NateScarlet/qidian/commit/128d8f9b922196dfcdcaeb9f29b32d0ed942040f))
* parse book bookmark count ([6c7338a](https://github.com/NateScarlet/qidian/commit/6c7338a3c4d115d845e379a3926663fdfc5c78f1))
* parse book char count ([455aa0a](https://github.com/NateScarlet/qidian/commit/455aa0ad3b5f2c33cd64e8ff58c2e5f95229607d))
* parse book finished ([9e6140d](https://github.com/NateScarlet/qidian/commit/9e6140d0c33b14ce3d4967b328b72a5131aab7ce))
* parse book last updated ([b4f2a09](https://github.com/NateScarlet/qidian/commit/b4f2a094ef5d9ab5a6acf0d924b86392e58aacfa))
* parse book recommend count ([eb8a746](https://github.com/NateScarlet/qidian/commit/eb8a74689b5197ec06e5f77150543faea1e864ed))
