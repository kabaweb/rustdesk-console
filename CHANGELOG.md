## [1.5.1](https://github.com/databk/rustdesk-console/compare/1.5.0...1.5.1) (2026-06-26)


### Bug Fixes

* map os field to standardized platform constants in /ab/peers response ([#178](https://github.com/databk/rustdesk-console/issues/178)) ([c1b0d67](https://github.com/databk/rustdesk-console/commit/c1b0d676717d7d93631d9d0358380d63d328775d)), closes [#175](https://github.com/databk/rustdesk-console/issues/175)
* merge saved LDAP config when testing connection ([#180](https://github.com/databk/rustdesk-console/issues/180)) ([7390ca5](https://github.com/databk/rustdesk-console/commit/7390ca54a364c7f91933239ac72aeed6b68774d8))
* specify varchar type for User.email column ([#181](https://github.com/databk/rustdesk-console/issues/181)) ([924bcda](https://github.com/databk/rustdesk-console/commit/924bcdaf6834b7face280e931db0070878767069))
* store null instead of empty string for user email to avoid unique constraint violation ([#176](https://github.com/databk/rustdesk-console/issues/176)) ([cff06b7](https://github.com/databk/rustdesk-console/commit/cff06b77fbf009eab533730935f542ee8bc83395)), closes [#173](https://github.com/databk/rustdesk-console/issues/173)



# [1.5.0](https://github.com/databk/rustdesk-console/compare/1.4.1...1.5.0) (2026-06-14)


### Bug Fixes

* align alarm audit query interface with file/connection audit ([#138](https://github.com/databk/rustdesk-console/issues/138)) ([825af6a](https://github.com/databk/rustdesk-console/commit/825af6a4dddb4fc02d1f37353d616dfb90072a99))


### Features

* add audit log auto-cleanup with configurable retention ([#143](https://github.com/databk/rustdesk-console/issues/143)) ([e83545a](https://github.com/databk/rustdesk-console/commit/e83545a0e13f4c097a2d24ba52a4b3a4fc52a143))
* add LDAP authentication support ([#148](https://github.com/databk/rustdesk-console/issues/148)) ([efa90ec](https://github.com/databk/rustdesk-console/commit/efa90ecbf3766c7825f1a2f2a0f0d91814afb410)), closes [#135](https://github.com/databk/rustdesk-console/issues/135)



## [1.4.1](https://github.com/databk/rustdesk-console/compare/1.4.0...1.4.1) (2026-06-07)


### Bug Fixes

* **docker:** use login-options endpoint for health check ([#133](https://github.com/databk/rustdesk-console/issues/133)) ([aab6d92](https://github.com/databk/rustdesk-console/commit/aab6d92a687c1ccdf0581b1a502601e9abdb3557))



# [1.4.0](https://github.com/databk/rustdesk-console/compare/1.3.0...1.4.0) (2026-06-06)


### Bug Fixes

* add validation decorators to StrategyQueryDto ([#108](https://github.com/databk/rustdesk-console/issues/108)) ([0e64942](https://github.com/databk/rustdesk-console/commit/0e6494276b979e7e4ee381c11d266de24ba63158))
* correct is_admin query parameter handling in admin users API ([#122](https://github.com/databk/rustdesk-console/issues/122)) ([ea8f3da](https://github.com/databk/rustdesk-console/commit/ea8f3da73678dac7642e287363c5e08abbf488ec))
* return full API path for avatar field ([#117](https://github.com/databk/rustdesk-console/issues/117)) ([af98ebf](https://github.com/databk/rustdesk-console/commit/af98ebfb6e823074529beefd35d5d3b61f1013bd))
* unify login response type field to email_check for client compatibility ([#118](https://github.com/databk/rustdesk-console/issues/118)) ([9fa36fc](https://github.com/databk/rustdesk-console/commit/9fa36fc6a79f890b1b5c361ee6100d16a90bbc3c))
* **user:** specify varchar type for avatar column ([#115](https://github.com/databk/rustdesk-console/issues/115)) ([dd7fb10](https://github.com/databk/rustdesk-console/commit/dd7fb10f89ada59baa79be17ad79a314988d0ab8))


### Features

* add admin users API for management-side user queries ([#119](https://github.com/databk/rustdesk-console/issues/119)) ([92682d1](https://github.com/databk/rustdesk-console/commit/92682d1d668652fbcf3ee9c8355b9bea8778d63e))
* add change password API for current user ([#114](https://github.com/databk/rustdesk-console/issues/114)) ([ae6c2f6](https://github.com/databk/rustdesk-console/commit/ae6c2f6e1d993c2adc43a6fe682f111e2c6b03f3))
* add strategy delivery via heartbeat ([#100](https://github.com/databk/rustdesk-console/issues/100)) ([015b489](https://github.com/databk/rustdesk-console/commit/015b48962a1f428924f24314a29cc9929f967694))
* add user avatar upload and management ([#111](https://github.com/databk/rustdesk-console/issues/111)) ([58af9fa](https://github.com/databk/rustdesk-console/commit/58af9fa3cab151338d5e4db2ca5e52a3507acfed))
* **auth:** implement complete 2FA user setup flow and remove tfa_url ([#110](https://github.com/databk/rustdesk-console/issues/110)) ([8eabe87](https://github.com/databk/rustdesk-console/commit/8eabe87acc6446bf767cbc9a23c7224bdd0e447f))
* **strategy:** add assignments query API ([#113](https://github.com/databk/rustdesk-console/issues/113)) ([cb50e41](https://github.com/databk/rustdesk-console/commit/cb50e41bb14430105833b845c54dfa00558dace9))
* update default admin account creation logic ([#120](https://github.com/databk/rustdesk-console/issues/120)) ([6899302](https://github.com/databk/rustdesk-console/commit/68993029699c1e1891a30aa6b5547b5e32e681e6))



# [1.3.0](https://github.com/databk/rustdesk-console/compare/1.2.0...1.3.0) (2026-05-30)


### Bug Fixes

* **audit:** simplify note handling for client requests ([#77](https://github.com/databk/rustdesk-console/issues/77)) ([34c8d23](https://github.com/databk/rustdesk-console/commit/34c8d23713d300a150be459c6be954fd3c8d9eac))
* **build:** correct OIDC templates path and add to build assets ([#95](https://github.com/databk/rustdesk-console/issues/95)) ([56cb794](https://github.com/databk/rustdesk-console/commit/56cb794b3a2d5c04ebe11d4a85b8e850f0a4d463))
* change file audit query param from peerId to deviceId ([#68](https://github.com/databk/rustdesk-console/issues/68)) ([7598dd9](https://github.com/databk/rustdesk-console/commit/7598dd923386117609664d1ecaa11adf687e4f10))
* correct login response type field for TFA check ([#69](https://github.com/databk/rustdesk-console/issues/69)) ([64e3a0f](https://github.com/databk/rustdesk-console/commit/64e3a0f6af4c5aa97fb45b77163d87dbd6e4d518))
* **oidc:** change login options response format to oidc/{name} ([#91](https://github.com/databk/rustdesk-console/issues/91)) ([e09d0e2](https://github.com/databk/rustdesk-console/commit/e09d0e246bd59e6d782f157755c1edd75ebc85f7))
* **oidc:** make id, uuid and deviceInfo optional for web frontend login ([#97](https://github.com/databk/rustdesk-console/issues/97)) ([8c59758](https://github.com/databk/rustdesk-console/commit/8c597589bef5bcf6d662ad5393fe9b0cf89ea39d))
* use dedicated lastHeartbeat field for device online status ([#67](https://github.com/databk/rustdesk-console/issues/67)) ([cfe3e28](https://github.com/databk/rustdesk-console/commit/cfe3e28ba356beb85ebaf9aa126b9ea662f4f17f))
* use tag_name instead of tag for action-gh-release ([#42](https://github.com/databk/rustdesk-console/issues/42)) ([bd62cf8](https://github.com/databk/rustdesk-console/commit/bd62cf8e40b345928c3af0e884027addef73ed85))


### Features

* add force disconnect connection via heartbeat response ([#79](https://github.com/databk/rustdesk-console/issues/79)) ([82fa275](https://github.com/databk/rustdesk-console/commit/82fa275c1b80586b1b5134704fb399d1d90bac17))
* **audit:** add admin endpoint to update connection audit note ([#78](https://github.com/databk/rustdesk-console/issues/78)) ([981ff0f](https://github.com/databk/rustdesk-console/commit/981ff0f975292ac8b8e234f4491c733e2009e901))
* **audit:** add ConnType enum and use -1 for unestablished connections ([#75](https://github.com/databk/rustdesk-console/issues/75)) ([ae55042](https://github.com/databk/rustdesk-console/commit/ae5504296ef1fe46b00cb2357c4a5da3ea857bd5))
* **audit:** add note field to connection audit ([#76](https://github.com/databk/rustdesk-console/issues/76)) ([114db9e](https://github.com/databk/rustdesk-console/commit/114db9e1cc269fecb8a38f198a6d4f667a7e3883))
* **device:** add batch device status update API ([#56](https://github.com/databk/rustdesk-console/issues/56)) ([3c674a1](https://github.com/databk/rustdesk-console/commit/3c674a1c1a60763c0dc25dcd7685479556de1cee))
* implement OIDC login with Authorization Code Flow + PKCE ([#70](https://github.com/databk/rustdesk-console/issues/70)) ([6218647](https://github.com/databk/rustdesk-console/commit/6218647973ff8c1100519080c9c4ef46d92ff648)), closes [#71](https://github.com/databk/rustdesk-console/issues/71) [#73](https://github.com/databk/rustdesk-console/issues/73) [#88](https://github.com/databk/rustdesk-console/issues/88)
* **oidc:** add admin CRUD API for OIDC provider management ([#92](https://github.com/databk/rustdesk-console/issues/92)) ([13901a4](https://github.com/databk/rustdesk-console/commit/13901a4046eaf8a31149ec3e81ceb48fcaec97c9))
* **oidc:** add dual protocol support for OIDC and OAuth2 providers ([#94](https://github.com/databk/rustdesk-console/issues/94)) ([d636603](https://github.com/databk/rustdesk-console/commit/d6366038c3f46a82d2f836277b5ef783744dca32)), closes [#5](https://github.com/databk/rustdesk-console/issues/5)
* **oidc:** add token storage logic for web login callback ([#99](https://github.com/databk/rustdesk-console/issues/99)) ([778680b](https://github.com/databk/rustdesk-console/commit/778680b013eec5cde00a3a5b078364ce56d32825))
* **oidc:** add web frontend login support with cookie-based authentication ([#96](https://github.com/databk/rustdesk-console/issues/96)) ([ebe01d5](https://github.com/databk/rustdesk-console/commit/ebe01d57a9dfb79a4a34db4af7de60f70bcd4d3a))



