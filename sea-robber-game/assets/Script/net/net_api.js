var http = require('hhttp')
var config = require('net_config')

// getRankAll 获取按战力排行的 前100名玩家信息
module.exports.getRankAll = function () {
  let url = config.baseURL + '/rank/'
  return http.get(url)
    .then(resp => {
      return Promise.resolve(JSON.parse(resp).rank)
    })
}

module.exports.postCreateUser = function(display_name, battle_score, battle_ship_id) {
  let url = config.baseURL + '/user/'
  return http.post(url, {display_name, battle_score, battle_ship_id})
    .then(resp => {
      return Promise.resolve(JSON.parse(resp).user)
    })
}

module.exports.postUpdateUser = function(display_name, battle_score, battle_ship_id, uuid) {
  let url = config.baseURL + '/user/' + uuid
  return http.post(url, {display_name, battle_score, battle_ship_id})
    .then(resp => {
      return Promise.resolve(JSON.parse(resp).user)
    })
}