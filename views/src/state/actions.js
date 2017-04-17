// import Vue from 'vue'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:3000'

export const register = ({ commit }, user) => {
  return new Promise((resolve, reject) => {
    axios.post('/api/users', user)
        .then(res => {
          resolve(console.log(res))
        })
        .catch(err => {
          console.error(err)
        })
  })
}
