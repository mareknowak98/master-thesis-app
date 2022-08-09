import jwt_decode from "jwt-decode";

const TOKEN_KEY = 'user-token'

const TokenService = {
  getToken() {
    return localStorage.getItem(TOKEN_KEY)
  },
  decodeToken(token) {
    return jwt_decode(token);
  },
  setToken(token) {
    return localStorage.setItem(TOKEN_KEY, token);
  }
}

export { TokenService }