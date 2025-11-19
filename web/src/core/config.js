/**
 * 网站配置文件
 */
import packageInfo from '../../package.json'

const greenText = (text) => `\x1b[32m${text}\x1b[0m`

export const config = {
  appName: 'Gin-Vue-Admin',
  showViteLogo: true,
  keepAliveTabs: false,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(
      greenText(
        `> 欢迎使用Hertz-Vue-Admin，开源地址：https://github.com/chenyusolar/hertz-vue-admin`
      )
    )
    console.log(greenText(`> 当前版本:v${packageInfo.version}`))
   
    console.log('\n')
  }
}

export default config
