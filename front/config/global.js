import axios from "axios";

var header = ""

const instance = axios.create({
  host: 'http://127.0.0.1:8080',
  timeout: 1000,
  headers: header
});
