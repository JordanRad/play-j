import axios from 'axios'

export default class AccountService {

    static URL = "http://localhost:8091/api/v1/account-service/accounts"


    static async login(email, password) {
        const body = {
            email: email,
            password: password
        }
        let result;
        try {
            result = await axios.post(AccountService.URL + `/login`, body)
        } catch (err) {
            console.log(err)
        }

        return result != undefined ? result.data : null
    }
}