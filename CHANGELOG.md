# [1.6.0](https://github.com/databk/rustdesk-console/compare/1.5.1...1.6.0) (2026-07-15)


### Bug Fixes

* make SMTP username and password fields optional ([#194](https://github.com/databk/rustdesk-console/issues/194)) ([eae647c](https://github.com/databk/rustdesk-console/commit/eae647c464ddc08019c4f7063e7cd425ec7503c1)), closes [#193](https://github.com/databk/rustdesk-console/issues/193)
* **update-check:** fix update check API URL and package.json path issues ([#207](https://github.com/databk/rustdesk-console/issues/207)) ([1d1bc6e](https://github.com/databk/rustdesk-console/commit/1d1bc6e92ec881c88d2c32ab129dba81cf686fad))


### Features

* add nexus module for custom client generation ([#192](https://github.com/databk/rustdesk-console/issues/192)) ([a8c0fcf](https://github.com/databk/rustdesk-console/commit/a8c0fcfede80ee2d5a591e283361540ddfa1edb3))
* add update check module ([#182](https://github.com/databk/rustdesk-console/issues/182)) ([b952d41](https://github.com/databk/rustdesk-console/commit/b952d414071635675c328ce22cbbb1cf4e8ed64a))



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



