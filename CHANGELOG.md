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



# [1.2.0](https://github.com/databk/rustdesk-console/compare/1.1.0...1.2.0) (2026-05-17)


### Bug Fixes

* **dashboard:** resolve device ID to numeric id from peers table ([#33](https://github.com/databk/rustdesk-console/issues/33)) ([54bd9d1](https://github.com/databk/rustdesk-console/commit/54bd9d183913701d93388356e3f3b0bfbc36c7da))
* **dashboard:** resolve device name to hostname from sysinfos table ([#34](https://github.com/databk/rustdesk-console/issues/34)) ([0d81f31](https://github.com/databk/rustdesk-console/commit/0d81f312faf50690d1594a712339e80bd75fa3f4))
* handle JSON parsing for file audit data in statistics ([#32](https://github.com/databk/rustdesk-console/issues/32)) ([84fd24c](https://github.com/databk/rustdesk-console/commit/84fd24c8f1c45fc29555b19e2416daf2de3ee589))
* resolve eslint errors across multiple modules ([#40](https://github.com/databk/rustdesk-console/issues/40)) ([37802d1](https://github.com/databk/rustdesk-console/commit/37802d13223c25aa1069aa53561c17e308d23c56))
* **throttler:** resolve rate limiting double counting issue ([#35](https://github.com/databk/rustdesk-console/issues/35)) ([931d209](https://github.com/databk/rustdesk-console/commit/931d209ab413ef1ad748a98062edab8dcc96a9d9))


### Features

* add dashboard API for statistics and analytics ([#31](https://github.com/databk/rustdesk-console/issues/31)) ([e0d5a07](https://github.com/databk/rustdesk-console/commit/e0d5a07e1301e65f1105cbf11606d606903f1871))
* add device_group_guid parameter to peer query ([#37](https://github.com/databk/rustdesk-console/issues/37)) ([d685cb8](https://github.com/databk/rustdesk-console/commit/d685cb85a2746bc7702cb1d7cd5f25644179c969))
* **audit:** enhance file audit query with advanced filters ([#38](https://github.com/databk/rustdesk-console/issues/38)) ([ecf3845](https://github.com/databk/rustdesk-console/commit/ecf38455a539e14ebdfc2dd23cb6b9cf84b68e41))
* **security:** Implement JTI Token Blacklist Mechanism ([#36](https://github.com/databk/rustdesk-console/issues/36)) ([16e3b6f](https://github.com/databk/rustdesk-console/commit/16e3b6f8ef069180581511de4438431bd1558914))
* **settings:** add SMTP configuration management API with generic settings table ([#24](https://github.com/databk/rustdesk-console/issues/24)) ([c5e3c67](https://github.com/databk/rustdesk-console/commit/c5e3c671a659fb2626b2ef17737de717d40a660b))



# [1.1.0](https://github.com/databk/rustdesk-console/compare/1.0.0...1.1.0) (2026-04-29)


### Bug Fixes

* **ab:** support single tag parameter in peers query ([#25](https://github.com/databk/rustdesk-console/issues/25)) ([446c7bd](https://github.com/databk/rustdesk-console/commit/446c7bd40e1f0b1be413e5da774d1b0c7514cda5))
* **ci:** remove v prefix from docker image tags ([#15](https://github.com/databk/rustdesk-console/issues/15)) ([18464d6](https://github.com/databk/rustdesk-console/commit/18464d66d8a014fa5070afecae4740833b9585c9))
* **ci:** scope changelog to changes since last release ([#28](https://github.com/databk/rustdesk-console/issues/28)) ([812a66e](https://github.com/databk/rustdesk-console/commit/812a66e83b07d27607849a2c9dc5fa42fefe4c6b))
* restore device_name field and update user field in peer service response ([#21](https://github.com/databk/rustdesk-console/issues/21)) ([345a2b6](https://github.com/databk/rustdesk-console/commit/345a2b6db681179fe9ebbf3fa1834b67d019814b))


### Features

* **ab:** add tag, id, alias filtering for peers endpoint ([#23](https://github.com/databk/rustdesk-console/issues/23)) ([86d858e](https://github.com/databk/rustdesk-console/commit/86d858e4e2ac94a99063987eb98c124edc3b772f))
* **peers:** add pagination and multi-field filtering to peers API ([#22](https://github.com/databk/rustdesk-console/issues/22)) ([00ef135](https://github.com/databk/rustdesk-console/commit/00ef13577a4709296a1d721d176d7e47da227d87))
* **peers:** add status column to peers table for device enable/disable ([#18](https://github.com/databk/rustdesk-console/issues/18)) ([50950ce](https://github.com/databk/rustdesk-console/commit/50950ce724cc5b7112011733999d21081e22033f))
* **workflow:** update release workflow to use manual trigger and conventional changelog ([#29](https://github.com/databk/rustdesk-console/issues/29)) ([afd3f72](https://github.com/databk/rustdesk-console/commit/afd3f728fcb41c7676e1c4d200451041cd607330))



# [1.0.0](https://github.com/databk/rustdesk-console/compare/b30cdb26ff86c0693428612a98564e0eb106e492...1.0.0) (2026-04-19)


### Bug Fixes

* **ci:** correct release workflow title, version and contributors ([#13](https://github.com/databk/rustdesk-console/issues/13)) ([288da33](https://github.com/databk/rustdesk-console/commit/288da33600f83c6cedccce84fcf1774b398ca31d))
* **ci:** remove npm test step from release workflow ([#12](https://github.com/databk/rustdesk-console/issues/12)) ([6122a4c](https://github.com/databk/rustdesk-console/commit/6122a4ca3eb7afe758cbddc7b3d7f6123a6449fe))
* **ci:** use package.json version for changelog and skip auto-bump ([#14](https://github.com/databk/rustdesk-console/issues/14)) ([57aa5cd](https://github.com/databk/rustdesk-console/commit/57aa5cd297cbce9fe682ec21d11ed947fc5e5f67))
* **docker:** correct HEALTHCHECK endpoint ([#9](https://github.com/databk/rustdesk-console/issues/9)) ([7e89853](https://github.com/databk/rustdesk-console/commit/7e8985344bb5a9d4b631abb9a0db7b41a292e7de))
* resolve all eslint errors and warnings ([#11](https://github.com/databk/rustdesk-console/issues/11)) ([ded5cf3](https://github.com/databk/rustdesk-console/commit/ded5cf3c6a0b8cb277eab1e83e9f3788ebfd06a2))


### Features

* add GitHub Actions release workflow ([#7](https://github.com/databk/rustdesk-console/issues/7)) ([7825096](https://github.com/databk/rustdesk-console/commit/782509694901023a788aaf409606e4854c8834c6))
* **address-book:** add address book module and controller ([09cb17e](https://github.com/databk/rustdesk-console/commit/09cb17e06a534b1b2bed84f341a6ec4f46afe56f))
* **address-book:** add database entities for address book ([d7be57d](https://github.com/databk/rustdesk-console/commit/d7be57db1932a59c5e998532fcbfe44e41f2cf1c))
* **address-book:** add DTOs for peers, queries, rules, and tags ([8887995](https://github.com/databk/rustdesk-console/commit/888799531ae97253e1ab599f624f8795817be147))
* **address-book:** add main address book service ([2682e1b](https://github.com/databk/rustdesk-console/commit/2682e1be5e6ceafe4b7fbf1012f578e4582724d2))
* **address-book:** add specialized services for peers, tags, rules, permissions, and legacy support ([86137a6](https://github.com/databk/rustdesk-console/commit/86137a6831a870b124558c794b0fd1d6b15427af))
* **audit:** add audit module with controller and service ([6f1cb03](https://github.com/databk/rustdesk-console/commit/6f1cb0327af005b3497c1964c631811ba7f49517))
* **audit:** add DTOs and entities for alarm, connection, and file audits ([6977e72](https://github.com/databk/rustdesk-console/commit/6977e7220f43db2f68ac603f23a1e9ae30156810))
* **auth:** add auth guards, decorators, and JWT strategy ([865cc33](https://github.com/databk/rustdesk-console/commit/865cc332559ef525f24224f1cdef185ea01cb197))
* **auth:** add authentication controller ([89503a6](https://github.com/databk/rustdesk-console/commit/89503a664d86cd77de544932aa84ddbb05f7bebc))
* **auth:** add authentication DTOs ([9a9c68d](https://github.com/databk/rustdesk-console/commit/9a9c68d03d7173f48eea14ec1a3ae536b8af1205))
* **auth:** add authentication module ([3fb8096](https://github.com/databk/rustdesk-console/commit/3fb80969e4f97b8b8ed33212cee98dd54b44201f))
* **auth:** add device tracking service and exports ([032bb5d](https://github.com/databk/rustdesk-console/commit/032bb5d6552c4d724b5368e6837e2d5c19a7bf52))
* **auth:** add email verification service ([a85b9c7](https://github.com/databk/rustdesk-console/commit/a85b9c764c09bf7288915aa5fffc0b3dbba09ee6))
* **auth:** add email verification session entity ([bbb1504](https://github.com/databk/rustdesk-console/commit/bbb1504cc37d402d5d72ed26956df00f67efe520))
* **auth:** add main authentication service ([9b17e40](https://github.com/databk/rustdesk-console/commit/9b17e40e773ed144188af2671129f6cd066040b1))
* **auth:** add token management service ([2718240](https://github.com/databk/rustdesk-console/commit/2718240595cc6cfe9a66ba2a4188b7931eff8f91))
* **auth:** add two-factor authentication service ([641dfba](https://github.com/databk/rustdesk-console/commit/641dfba0eade4afc6a10c53165ef23779648a682))
* **common:** add common module index exports ([73e9f11](https://github.com/databk/rustdesk-console/commit/73e9f11eb140800bb7ae124dd90fb0798cee3303))
* **core:** add application entry point ([b30cdb2](https://github.com/databk/rustdesk-console/commit/b30cdb26ff86c0693428612a98564e0eb106e492))
* **core:** add root application module ([e6f6c06](https://github.com/databk/rustdesk-console/commit/e6f6c0606b23725be47994e753641cce5d706c0c))
* **database:** add database initialization service ([f071e3e](https://github.com/databk/rustdesk-console/commit/f071e3e66112364a9962e5753cc37fddfe229011))
* **database:** add database module ([d2bd012](https://github.com/databk/rustdesk-console/commit/d2bd012627b1273710f6e293db232944acd49562))
* **decorators:** add @CurrentUser() decorator ([7d4b130](https://github.com/databk/rustdesk-console/commit/7d4b130e3880af3a6afe3161ff3be3364950eb98))
* **decorators:** add @Public() decorator ([c66cc7c](https://github.com/databk/rustdesk-console/commit/c66cc7cc7dcce45c8c4b7409f9a2eed7f5672663))
* **decorators:** add decorators index exports ([f6422c3](https://github.com/databk/rustdesk-console/commit/f6422c32d10f4a91a3f14035933117c22523ab7e))
* **device-group:** add device group module with controller and service ([0983c7b](https://github.com/databk/rustdesk-console/commit/0983c7b61fcc06190e274d1f9a3c14978a54e443))
* **device-group:** add DTOs for device groups, devices, peers, and users ([1a82557](https://github.com/databk/rustdesk-console/commit/1a82557ffa5c5f35c62220e1a47b224f5408ab53))
* **device-group:** add entities and peer service ([aea5936](https://github.com/databk/rustdesk-console/commit/aea5936e975f2699e6422b556faf75be248e39e4))
* **email:** add email service module with Handlebars templates ([27be75b](https://github.com/databk/rustdesk-console/commit/27be75b5e584509e36c7ded55b0369448a8b59e6))
* **entities:** add entities index exports ([117aab0](https://github.com/databk/rustdesk-console/commit/117aab06c440fe03d176a4773630b2c3644f8b5b))
* **entities:** add Peer entity ([9e06107](https://github.com/databk/rustdesk-console/commit/9e06107636286f5783ca57cdcc9b7656265012e1))
* **entities:** add SysInfo entity ([349c327](https://github.com/databk/rustdesk-console/commit/349c327913ac8e3e3ac939ae367f043d3d0a8c39))
* **guards:** add admin role guard ([6469e47](https://github.com/databk/rustdesk-console/commit/6469e475af63bcf424138776688237098e9bfc7b))
* **guards:** add device throttling guard ([ee1bbee](https://github.com/databk/rustdesk-console/commit/ee1bbee2b60199d1d8c626627afa82760e435ed0))
* **guards:** add guards index exports ([fae530f](https://github.com/databk/rustdesk-console/commit/fae530f9171649c6a8522d92650b57244b037377))
* **guards:** add JWT authentication guard ([c2d7b22](https://github.com/databk/rustdesk-console/commit/c2d7b22b7bb424c798a606afbcdc4cd4e4f73592))
* **heartbeat:** add heartbeat monitoring module ([d0cbb75](https://github.com/databk/rustdesk-console/commit/d0cbb75e39f625be82388e0061bde269785fe556))
* **oidc:** add OpenID Connect authentication module ([f9bbef7](https://github.com/databk/rustdesk-console/commit/f9bbef751c719bbfdabec01e97032703f4c1a369))
* **services:** add services index exports ([9dd663d](https://github.com/databk/rustdesk-console/commit/9dd663db1832072a633d1115ab2c1c578e47692b))
* **services:** add token service ([6e926f8](https://github.com/databk/rustdesk-console/commit/6e926f874dffa69033dfd802bafadf61276ec3b9))
* **sysinfo:** add system information module ([a7bf6f9](https://github.com/databk/rustdesk-console/commit/a7bf6f9773d03495b2f2fb416791cd7ccdf5eadc))
* **user:** add user entities and service ([74d1744](https://github.com/databk/rustdesk-console/commit/74d1744a173ce80a531c5baa4f75472cb18042b8))
* **user:** add user module with controller ([073aa43](https://github.com/databk/rustdesk-console/commit/073aa43b64c7b80f344cef5f84787151d11c1e19))



