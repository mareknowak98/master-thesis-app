import jwt_decode from "jwt-decode";

const TOKEN_KEY = 'user-token'
const ID_TOKEN_KEY = 'id-user-token'

const TokenService = {
  getToken() {
    return localStorage.getItem(TOKEN_KEY)
  },
  decodeToken(token) {
    return jwt_decode(token);
  },
  setToken(token) {
    return localStorage.setItem(TOKEN_KEY, token);
  },
  deleteToken() {
    localStorage.removeItem(TOKEN_KEY)
  },
  deleteIdToken() {
    localStorage.removeItem(ID_TOKEN_KEY)
  },
  getIdToken() {
    return localStorage.getItem(ID_TOKEN_KEY)
  },
  setIdToken(token) {
    return localStorage.setItem(ID_TOKEN_KEY, token);
  },
}

export { TokenService }