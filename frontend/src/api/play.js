import { removePrefix } from './utils'
import { baseURL } from '@/utils/constants'
import store from '@/store'

async function playerAction (data) {

    const res = await fetch(`${baseURL}/api/airplay`, {
        method: 'POST',
        headers: {
            'X-Auth': store.state.jwt,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })

    if (res.status === 404) {
        return 404
    }else {
        return res
    }
}
export function airplay (url, type) {

    let host = window.location.host

    url = host+"/api/raw"+removePrefix(url)+"?auth="+store.state.jwt+"&inline=true"

    if(url.substr(0,7).toLowerCase() !== "http://" || url.substr(0,8).toLowerCase() !== "https://"){
        url = "http://" + url;
    }

    let content = {
        "media_url":url,
        "starting_position":0.0,
        "play_timeout":0,
        "type":type
    }

    return playerAction(content)

}