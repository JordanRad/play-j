import axios from 'axios'

export default class PlayerService {

    static URL = "http://localhost:8092/api/v1/player-service/tracks"

    static getJWT() {
        let user = JSON.parse(localStorage.getItem("user"))
        return user.token
    }

    static async getTrackById(id) {
        let result;
        let token = PlayerService.getJWT()
        try {
            result = await axios.get(`${PlayerService.URL}/${id}`, {
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