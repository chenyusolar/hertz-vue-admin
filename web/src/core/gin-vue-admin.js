/*
 * gin-vue-admin web框架组
 *
 * */
// 加载网站配置文件夹
import { register } from './global'
import packageInfo from '../../package.json'

export default {
  install: (app) => {
    register(app)
    console.log(`
       欢迎使用 Hertz-Vue-Admin
       当前版本:v${packageInfo.version}
       项目地址：https://github.com/chenyusolar/hertz-vue-admin
       插件市场:https://plugin.hertz-vue-admin.com
       
       默认自动化文档地址:http://127.0.0.1:${import.meta.env.VITE_SERVER_PORT}/swagger/index.html
       默认前端文件运行地址:http://127.0.0.1:${import.meta.env.VITE_CLI_PORT}
       --------------------------------------版权声明--------------------------------------
    `)
  }
}
