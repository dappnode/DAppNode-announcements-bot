package params

var RegistryAbi string = `[{"constant":true,"inputs":[{"name":"_semanticVersion","type":"uint16[3]"}],"name":"getBySemanticVersion","outputs":[{"name":"semanticVersion","type":"uint16[3]"},{"name":"contractAddress","type":"address"},{"name":"contentURI","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"EVMSCRIPT_REGISTRY_APP_ID","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_newSemanticVersion","type":"uint16[3]"},{"name":"_contractAddress","type":"address"},{"name":"_contentURI","type":"bytes"}],"name":"newVersion","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_versionId","type":"uint256"}],"name":"getByVersionId","outputs":[{"name":"semanticVersion","type":"uint16[3]"},{"name":"contractAddress","type":"address"},{"name":"contentURI","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"appId","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getInitializationBlock","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_contractAddress","type":"address"}],"name":"getLatestForContractAddress","outputs":[{"name":"semanticVersion","type":"uint16[3]"},{"name":"contractAddress","type":"address"},{"name":"contentURI","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"EVMSCRIPT_REGISTRY_APP","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_sender","type":"address"},{"name":"_role","type":"bytes32"},{"name":"params","type":"uint256[]"}],"name":"canPerform","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_oldVersion","type":"uint16[3]"},{"name":"_newVersion","type":"uint16[3]"}],"name":"isValidBump","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"pure","type":"function"},{"constant":true,"inputs":[],"name":"CREATE_VERSION_ROLE","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getLatest","outputs":[{"name":"semanticVersion","type":"uint16[3]"},{"name":"contractAddress","type":"address"},{"name":"contentURI","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getVersionsCount","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"kernel","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_script","type":"bytes"}],"name":"getExecutor","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"versionId","type":"uint256"},{"indexed":false,"name":"semanticVersion","type":"uint16[3]"}],"name":"NewVersion","type":"event"}]`

var RepositoryAbi string = `[{"constant":true,"inputs":[],"name":"REPO_APP_NAME","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"APM_APP_NAME","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"ENS_SUB_APP_NAME","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"registrar","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_name","type":"string"},{"name":"_dev","type":"address"},{"name":"_initialSemanticVersion","type":"uint16[3]"},{"name":"_contractAddress","type":"address"},{"name":"_contentURI","type":"bytes"}],"name":"newRepoWithVersion","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"CREATE_REPO_ROLE","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"EVMSCRIPT_REGISTRY_APP_ID","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"appId","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"getInitializationBlock","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"EVMSCRIPT_REGISTRY_APP","outputs":[{"name":"","type":"bytes32"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_sender","type":"address"},{"name":"_role","type":"bytes32"},{"name":"params","type":"uint256[]"}],"name":"canPerform","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_name","type":"string"},{"name":"_dev","type":"address"}],"name":"newRepo","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_registrar","type":"address"}],"name":"initialize","outputs":[],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"kernel","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_script","type":"bytes"}],"name":"getExecutor","outputs":[{"name":"","type":"address"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":false,"name":"id","type":"bytes32"},{"indexed":false,"name":"name","type":"string"},{"indexed":false,"name":"repo","type":"address"}],"name":"NewRepo","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"proxy","type":"address"}],"name":"NewAppProxy","type":"event"}]`