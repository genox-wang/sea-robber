var { getRankAll, postCreateUser, postUpdateUser } = require('net_api')

cc.Class({
    extends: cc.Component,

    properties: {
        btn_rank_all: cc.Button,
        btn_create_user: cc.Button,
        btn_update_user: cc.Button,

        user_id: 3,
        user_uuid: '722527f8-308f-11e8-84b8-0242ac140003',
        display_name: 'wangji',
        battle_score: 10000,
        battle_ship_id: 1,
    },

    onLoad () {
        this.registerBtn(this.btn_rank_all, () => getRankAll(this.user_id, 3, 10))
        this.registerBtn(this.btn_create_user, () => postCreateUser(this.display_name, this.battle_score, this.battle_ship_id))
        this.registerBtn(this.btn_update_user, () => postUpdateUser(this.display_name, this.battle_score, this.battle_ship_id, this.user_uuid))
    },

    registerBtn (btn, req) {
        btn.node.on('click', function () {
            req().then( resp => {
                console.log(resp)
            })
        })
    }
});
