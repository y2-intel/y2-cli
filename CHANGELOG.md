# Changelog

## 0.4.0 (2026-02-10)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/y2-intel/y2-cli/compare/v0.3.0...v0.4.0)

### ⚠ BREAKING CHANGES

* add support for passing files as parameters

### Features

* add readme documentation for passing files as arguments ([c44fe53](https://github.com/y2-intel/y2-cli/commit/c44fe53db83f1eb97530050f38c5ed4c125ab585))
* add support for passing files as parameters ([b8f244f](https://github.com/y2-intel/y2-cli/commit/b8f244fd59d88fed15153a3e0ad4d8c872b5b929))
* **api:** manual updates ([e0a6899](https://github.com/y2-intel/y2-cli/commit/e0a6899f76fd6bcf4f484d21322d301366262c43))
* **client:** provide file completions when using file embed syntax ([715608d](https://github.com/y2-intel/y2-cli/commit/715608d0cf4f317766ef88e60e02ba616e61ba28))
* **cli:** improve shell completions for namespaced commands and flags ([5894389](https://github.com/y2-intel/y2-cli/commit/589438991b07e9b02d674ff52c654ff3130cfd09))
* enable suggestion for mistyped commands and flags ([8380198](https://github.com/y2-intel/y2-cli/commit/838019879730d0fb80863ad0d28d5be55113fdbd))


### Bug Fixes

* avoid consuming request bodies when printing redacted outputs ([43b2396](https://github.com/y2-intel/y2-cli/commit/43b2396c532f2e6b2aaccf45022cd81b4fd77482))
* **client:** do not use pager for short paginated responses ([acd270e](https://github.com/y2-intel/y2-cli/commit/acd270ea0856f3b2dd6589afac3356cf911a83c2))
* fix for file uploads to octet stream and form encoding endpoints ([071e3da](https://github.com/y2-intel/y2-cli/commit/071e3da7c0cdb0e6d43c3f6aa4d3425265fb5d53))
* fix for paginated output not writing to pager correctly ([7e414b7](https://github.com/y2-intel/y2-cli/commit/7e414b7a7ea407c7475e59b30c71e2c45767d3dd))
* fix for when terminal width is not available ([ec407aa](https://github.com/y2-intel/y2-cli/commit/ec407aa954a9340897884dd59e41507bb965e1d1))
* fix terminal height issues causing test failures ([57183a3](https://github.com/y2-intel/y2-cli/commit/57183a39720e63b50de6c6cf6ecfcc8fd17756fd))
* flag defaults ([35b1dba](https://github.com/y2-intel/y2-cli/commit/35b1dbacd938641fb25790477f7aea6608dbab5e))
* overly broad redaction of Authorization ([3fccc7f](https://github.com/y2-intel/y2-cli/commit/3fccc7f2a6d85b37a10058802cc161d8783f2b2d))
* prevent flag duplication ([6b3655d](https://github.com/y2-intel/y2-cli/commit/6b3655dea46ce38ee18ea132fe43265239140851))
* use RawJSON for iterated values instead of re-marshalling ([50ef392](https://github.com/y2-intel/y2-cli/commit/50ef392b447c372d36c44e5adbb5e49b3cc29c4e))


### Chores

* add build step to ci ([0bcb727](https://github.com/y2-intel/y2-cli/commit/0bcb7272eb5703a363b1ec8587bf43a3c8d46305))
* **internal:** codegen related update ([b7d11fb](https://github.com/y2-intel/y2-cli/commit/b7d11fb07bc9798196b617a6ef62e31cde1f7532))
* **internal:** codegen related update ([65a9fa3](https://github.com/y2-intel/y2-cli/commit/65a9fa3154822643ad3c0a8264701d3d41e5b907))
* **internal:** codegen related update ([c1b88a8](https://github.com/y2-intel/y2-cli/commit/c1b88a8f92a73570ed8780dec32db3ebcfcad58b))
* **internal:** codegen related update ([64e86f1](https://github.com/y2-intel/y2-cli/commit/64e86f1681ad150e82e0f036888e0f3a8d18d4be))
* **internal:** update `actions/checkout` version ([7da50c8](https://github.com/y2-intel/y2-cli/commit/7da50c841da88d81278173fb409745b7bb8bdccf))
* update documentation in readme ([f64a9c4](https://github.com/y2-intel/y2-cli/commit/f64a9c44849f47cfa7d287c96ecad84bc03b85e9))
* update Go SDK version ([ffa2972](https://github.com/y2-intel/y2-cli/commit/ffa29721400fc0af15f9a609a3e6fddf3a97686a))
* update internal comment ([3fccc7f](https://github.com/y2-intel/y2-cli/commit/3fccc7f2a6d85b37a10058802cc161d8783f2b2d))
* updated README.md with more flag information ([8c2e8e3](https://github.com/y2-intel/y2-cli/commit/8c2e8e3eb9607b3e16ea51ba47c3f1c4e9c46d64))

## 0.3.0 (2026-01-12)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/y2-intel/y2-cli/compare/v0.2.0...v0.3.0)

### Features

* added support for --foo.baz inner field flags ([4b3c5d4](https://github.com/y2-intel/y2-cli/commit/4b3c5d43a7f18f527627e3d0ba7a8a9a2727546f))
* enable CI tests ([7d6d8de](https://github.com/y2-intel/y2-cli/commit/7d6d8dee7fdb46a23fa3b008481ec119c11e308b))
* improved behavior for exploring paginated/streamed endpoints ([48543cc](https://github.com/y2-intel/y2-cli/commit/48543ccd557cd40a4276dfb96c5c3b7f1e827f02))


### Bug Fixes

* check required arguments ([4bb1f8f](https://github.com/y2-intel/y2-cli/commit/4bb1f8ff36c9ffe850a321eb36e4c4b6d0727b5c))
* fixed placeholders for date/time arguments ([7aa41a9](https://github.com/y2-intel/y2-cli/commit/7aa41a9852e226832a00ca17e09c0ea61dd784d9))


### Chores

* **internal:** codegen related update ([3d98cfa](https://github.com/y2-intel/y2-cli/commit/3d98cfa19a6d0a4aaf100e50a8bdb6fe22599188))
* **internal:** codegen related update ([ff9767b](https://github.com/y2-intel/y2-cli/commit/ff9767bcab51ae4b093ff163c940b6d46cd09d20))

## 0.2.0 (2025-12-31)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/y2-intel/y2-cli/compare/v0.1.0...v0.2.0)

### Features

* **api:** api update ([68ec209](https://github.com/y2-intel/y2-cli/commit/68ec209e9103bd3da4bdaf3ca691dcadd365f33f))

## 0.1.0 (2025-12-28)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/y2-intel/y2-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([a59b091](https://github.com/y2-intel/y2-cli/commit/a59b0914db8b98b6e9ccdd3a8ac716babb716e99))


### Chores

* update SDK settings ([ed540b0](https://github.com/y2-intel/y2-cli/commit/ed540b09465356e36db090d0b6377fa35fdf925b))
* update SDK settings ([909e543](https://github.com/y2-intel/y2-cli/commit/909e5436f6de09d5a94cf2933669cc23a622c755))
* update SDK settings ([58f5f48](https://github.com/y2-intel/y2-cli/commit/58f5f482aee906166a8656dc08a34470b1cd0322))
* update SDK settings ([901dc0a](https://github.com/y2-intel/y2-cli/commit/901dc0a285e652c6949401bff5fd5d1d3dabcea2))
