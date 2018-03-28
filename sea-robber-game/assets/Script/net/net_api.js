var http = require('hhttp')
var config = require('net_config')

// getRankAll 获取战力排行
// 只缓存前10000名的 
// params 
// id 用户ID
// page 目标页数
// size 每页面记录数

module.exports.getRankAll = function (user_id, page, size) {
  let url = `${config.baseURL}/rank/?page=${page}&size=${size}` 
  return http.get(url, { headers: { "Authorization": user_id }})
    .then(resp => {
      return Promise.resolve(JSON.parse(resp))
    })
}

// postCreateUser 创建用户
// display_name 用户名称，用于展示显示，可留空
// battle_score 战斗力
// battle_ship_id 战舰ID 对应本地图片

// response
// user { uuid: 用户唯一ID }
module.exports.postCreateUser = function(display_name, battle_score, battle_ship_id) {
  let url = config.baseURL + '/user/'
  return http.post(url, {display_name, battle_score, battle_ship_id})
    .then(resp => {
      return Promise.resolve(JSON.parse(resp).user)
    })
}

// postUpdateUser 修改用户信息
// display_name 用户名称，用于展示显示，可留空
// battle_score 战斗力
// battle_ship_id 战舰ID 对应本地图片
// uuid 用户唯一标示

module.exports.postUpdateUser = function(display_name, battle_score, battle_ship_id, uuid) {
  let url = config.baseURL + '/user/' + uuid
  return http.post(url, {display_name, battle_score, battle_ship_id})
    .then(resp => {
      return Promise.resolve(JSON.parse(resp).user)
    })
}