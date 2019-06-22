<template>
    <button @click="show" :aria-label="$t('buttons.play')" :title="$t('buttons.play')" class="action" id="play-button">
        <i class="material-icons">play_arrow</i>
        <span>{{ $t('buttons.player') }}</span>
    </button>
</template>

<script>

import { airplay } from '../../api/play'
import {mapGetters, mapState} from 'vuex'
    export default {
        name: 'play-button',
        computed: {
            ...mapState(['req', 'selected']),
            ...mapGetters(['isListing', 'selectedCount']),

        },
        methods: {
            show: function () {
                let url = this.req.items[this.selected[0]].url
                let type = this.req.items[this.selected[0]].type
                let res = airplay(url, type)
                let vm = this
                res.then(function (result) {
                    if (result === 404){
                        vm.$store.commit('showHover', 'play')
                    }
                })
            }
        }
    }
</script>