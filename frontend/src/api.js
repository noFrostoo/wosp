import { default as Axios } from 'axios';
import { API_URL, API_SUFFIX } from "./config";

const axios = Axios.create({
  validateStatus: function (status) {
    return status >= 200 && status < 300;
  },
  baseURL: API_URL
});

function authHeaders(token) {
    return {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    };
  }

export const api = {
    async logIn(username, password) {
      let data = {
        "user" : {
            "username": username,
            "password": password
        }
      }
  
      return axios.post(`${API_SUFFIX}/login/`, data);
    },
  
    async getMe(token) {
      return axios.get(`${API_SUFFIX}/user/me`, authHeaders(token));
    },
  
    async updateMe(token, username, password) {
        let data = {
            "user" : {
                "username": username,
                "password": password
            }
            }
      return axios.put(`${API_SUFFIX}/api/users/me`, data, authHeaders(token));
    },

    async create_todo(token, todo) {
        let data = {
            "todo" : {
                "user_id": todo.user_id,
                "title": todo.title,
                "description": todo.description,
                "done": todo.done,
            }
          }
          return axios.post(`${API_SUFFIX}/todo/`, data, authHeaders(token));
    },

    async edit_todo(token, todo) {
        let data = {
            "todo" : {
                "id": todo.id,
                "user_id": todo.user_id,
                "title": todo.title,
                "description": todo.description,
                "done": todo.done,
            }
          }
          return axios.put(`${API_SUFFIX}/todo/${todo.id}`, data, authHeaders(token));
    },

    async get_todo(token, id) {
        return axios.post(`${API_SUFFIX}/todo/${id}`, authHeaders(token));
    },

    async delete_todo(token, id) {
        return axios.delete(`${API_SUFFIX}/todo/${id}`, authHeaders(token));
    }

}