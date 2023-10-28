const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  publicPath: '/admin/',
  outputDir: 'dist',
  assetsDir: 'static',
})
