import axios from 'axios'

export default class PlaylistService {

    static URL = "http://localhost:8091/api/v1/account-service/playlists"

    static getJWT() {
        let user = JSON.parse(localStorage.getItem("user"))
        return user.token
    }

    static async getAccountPlaylistCollection() {
        let result;
        let token = PlaylistService.getJWT()
        try {
            result = await axios.get(PlaylistService.URL, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            console.log(result)
        } catch (err) {
            console.log(err)
        }
        return result != undefined ? result.data : null
    }

    static async createAccountPlaylist(name) {
        let result;
        let token = PlaylistService.getJWT()
        let body = {
            name:name
        }
        try {
            result = await axios.post(PlaylistService.URL, body,{
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            })
            console.log(result)
        } catch (err) {
            console.log(err)
        }
        return result != undefined ? result.data : null
    }
}